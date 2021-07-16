package vm

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	ledger "github.com/numary/ledger/core"
	"github.com/numary/machine/core"
	"github.com/numary/machine/vm/program"
)

const (
	EXIT_OK = byte(iota + 1)
	EXIT_FAIL
)

func StdOutPrinter(c chan core.Value) {
	for v := range c {
		fmt.Println("OUT:", v)
	}
}

func NewMachine(p *program.Program) *Machine {
	printc := make(chan core.Value)

	m := Machine{
		Program:    p,
		Constants:  p.Constants,
		print_chan: printc,
		Printer:    StdOutPrinter,
	}

	return &m
}

type Machine struct {
	P          uint
	Program    *program.Program
	Constants  []core.Value // Constants and Variables
	Variables  []core.Value // constitute the resources
	Stack      []core.Value
	Postings   []ledger.Posting             // accumulates postings throughout execution
	Balances   map[string]map[string]uint64 // keeps tracks of balances througout execution
	Printer    func(chan core.Value)
	print_chan chan core.Value
}

func (m *Machine) getResource(addr core.Address) core.Value {
	a := uint16(addr)
	if a < (1 << 15) {
		return m.Constants[a]
	} else {
		a -= (1 << 15)
		return m.Variables[a]
	}
}

func (m *Machine) withdraw(account string, asset string, amount uint64) bool {
	if account == "world" {
		return true
	}
	withdraw_ok := false
	if acc_balance, ok := m.Balances[account]; ok {
		if balance, ok := acc_balance[asset]; ok {
			if balance >= amount {
				acc_balance[asset] -= amount
				withdraw_ok = true
			}
		}
	}
	return withdraw_ok
}

func (m *Machine) credit(account string, asset string, amount uint64) {
	if acc_balance, ok := m.Balances[account]; ok {
		if _, ok := acc_balance[asset]; ok {
			acc_balance[asset] += amount
		} else {
			acc_balance[asset] = amount
		}
	} else {
		m.Balances[account] = map[string]uint64{
			asset: amount,
		}
	}
}

func (m *Machine) tick() (bool, byte) {
	op := m.Program.Instructions[m.P]

	switch op {
	case program.OP_APUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+3]
		v := m.getResource(core.Address(binary.LittleEndian.Uint16(bytes)))
		m.Stack = append(m.Stack, v)
		m.P += 2
	case program.OP_IPUSH:
		bytes := m.Program.Instructions[m.P+1 : m.P+9]
		v := core.Number(binary.LittleEndian.Uint64(bytes))
		m.Stack = append(m.Stack, v)
		m.P += 8
	case program.OP_IADD:
		b := m.popNumber()
		a := m.popNumber()
		m.pushValue(core.Number(a + b))
	case program.OP_ISUB:
		b := m.popNumber()
		a := m.popNumber()
		m.pushValue(core.Number(a - b))
	case program.OP_PRINT:
		a := m.popValue()
		m.print_chan <- a
	case program.OP_FAIL:
		return true, EXIT_FAIL
	case program.OP_SOURCE:
		n := m.popNumber()
		sources := []core.Account{}
		for i := uint64(0); i < n; i++ {
			sources = append(sources, m.popAccount())
		}
		mon := m.popMonetary()
		asset := mon.Asset
		target := mon.Amount

		var n_actual_src uint64
		for _, src := range sources {
			src_funds := m.Balances[string(src)][asset]
			var amt_to_withdraw uint64
			if src_funds > target || src == "world" {
				amt_to_withdraw = target
			} else {
				amt_to_withdraw = src_funds
			}
			target -= amt_to_withdraw
			m.pushValue(core.Monetary{
				Asset:  asset,
				Amount: amt_to_withdraw,
			})
			m.pushValue(src)
			n_actual_src++
			if target == 0 {
				break
			}
		}
		if target != 0 {
			return true, EXIT_FAIL
		}
		m.pushValue(core.Number(n_actual_src))
	case program.OP_ALLOC:
		allotment := m.popAllotment()
		n := m.popNumber()
		source_accounts := make([]core.Account, n)
		source_amounts := make([]uint64, n)
		total_src := uint64(0)
		var asset *string
		// extract accounts from stack while checking the assets correspond
		for i := uint64(0); i < n; i++ {
			source_accounts[i] = m.popAccount()
			mon := m.popMonetary()
			source_amounts[i] = mon.Amount
			// check that the assets correspond
			if asset == nil {
				asset = &mon.Asset
			} else if *asset != mon.Asset {
				return true, EXIT_FAIL
			}
			total_src += mon.Amount
		}
		parts := []uint64{}
		total_allocated := uint64(0)
		// for every part in the allotment, calculate the floored value
		for _, part := range allotment {
			var res big.Int
			res.Mul(part.Num(), new(big.Int).SetUint64(total_src))
			res.Div(&res, part.Denom())
			parts = append(parts, res.Uint64())
			total_allocated += res.Uint64()
		}
		// for every part in the floored values, fetch them from the sources
		first_non_empty_idx := 0
		for _, part := range parts {
			// if the total allocated is less than the target amount, add 1 unit until it isn't
			if total_allocated < uint64(total_src) {
				part += 1
				total_allocated += 1
			}
			n := 0 // number of sources needed to fill this part
			// start at the first non empty source
			for i := first_non_empty_idx; i < len(source_accounts); i++ {
				if part == 0 { // if we finished filling this part
					break
				}
				amt := source_amounts[i] // amount to withdraw from the account
				if amt > part {
					// if we have more than enough to fill, don't give too much to this part
					amt = part
				} else { // if we had to empty the source
					first_non_empty_idx++
				}
				part -= amt
				source_amounts[i] -= amt
				m.pushValue(core.Monetary{Asset: *asset, Amount: amt})
				m.pushValue(source_accounts[i])
				n += 1
			}
			m.pushValue(core.Number(n))
		}
	case program.OP_SEND:
		dest := m.popAccount()
		n := m.popNumber()
		for i := uint64(0); i < n; i++ {
			src := string(m.popAccount())
			mon := m.popMonetary()
			if mon.Amount == 0 {
				continue
			}
			// verify and withdraw funds
			if ok := m.withdraw(src, mon.Asset, mon.Amount); !ok {
				return true, EXIT_FAIL
			}
			m.credit(string(dest), mon.Asset, mon.Amount)
			m.Postings = append(m.Postings, ledger.Posting{
				Source:      string(src),
				Destination: string(dest),
				Asset:       string(mon.Asset),
				Amount:      int64(mon.Amount),
			})
		}
	}

	m.P += 1

	if int(m.P) >= len(m.Program.Instructions) {
		return true, EXIT_OK
	}

	return false, 0
}

func (m *Machine) Execute() (byte, error) {
	go m.Printer(m.print_chan)
	defer close(m.print_chan)

	if m.Variables == nil {
		return 0, errors.New("variables haven't been initialized")
	} else if m.Balances == nil {
		return 0, errors.New("balances haven't been initialized")
	}

	for {
		finished, exit_code := m.tick()
		if finished {
			return exit_code, nil
		}
	}
}

func (m *Machine) GetNeededBalances() (map[string]map[string]struct{}, error) {
	needed := map[string]map[string]struct{}{}
	for addr, needed_assets := range m.Program.NeededBalances {
		account := m.getResource(addr)
		if account, ok := account.(core.Account); ok {
			needed[string(account)] = map[string]struct{}{}
			for addr := range needed_assets {
				mon := m.getResource(addr)
				if mon, ok := mon.(core.Monetary); ok {
					needed[string(account)][string(mon.Asset)] = struct{}{}
				} else {
					return nil, errors.New("incorrect program")
				}
			}
		} else {
			return nil, errors.New("incorrect program")
		}
	}
	return needed, nil
}

func (m *Machine) SetBalances(balances map[string]map[string]uint64) error {
	// for every account that we need balances of, check if it's there
	for addr, needed_assets := range m.Program.NeededBalances {
		account := m.getResource(addr)
		if account, ok := account.(core.Account); ok {
			if string(account) == "world" {
				continue
			}
			if b, ok := balances[string(account)]; ok {
				// for every asset that we need balances of on that account
				for addr := range needed_assets {
					mon := m.getResource(addr)
					if mon, ok := mon.(core.Monetary); ok {
						if _, ok := b[string(mon.Asset)]; !ok {
							return fmt.Errorf("missing %v balance of %v", mon.Asset, account)
						}
					} else {
						return errors.New("incorrect program")
					}
				}
			} else {
				return fmt.Errorf("missing balances of %v", account)
			}
		} else {
			return errors.New("incorrect program")
		}
	}
	m.Balances = balances
	return nil
}

func (m *Machine) SetVars(vars map[string]core.Value) error {
	v, err := m.Program.ParseVariables(vars)
	if err != nil {
		return err
	}
	m.Variables = v
	return nil
}

func (m *Machine) SetVarsFromJSON(vars map[string]json.RawMessage) error {
	v, err := m.Program.ParseVariablesJSON(vars)
	if err != nil {
		return err
	}
	m.Variables = v
	return nil
}
