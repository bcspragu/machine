// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 45, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 6, 2, 19, 10, 2, 13, 2, 14, 2, 20, 3, 3, 6, 3, 24,
	10, 3, 13, 3, 14, 3, 25, 3, 4, 6, 4, 29, 10, 4, 13, 4, 14, 4, 30, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 40, 10, 8, 13, 8, 14, 8, 41,
	3, 8, 3, 8, 2, 2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 3, 2,
	6, 3, 2, 99, 124, 3, 2, 67, 92, 3, 2, 50, 59, 4, 2, 11, 12, 34, 34, 2,
	48, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 18, 3, 2,
	2, 2, 5, 23, 3, 2, 2, 2, 7, 28, 3, 2, 2, 2, 9, 32, 3, 2, 2, 2, 11, 34,
	3, 2, 2, 2, 13, 36, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 19, 9, 2, 2, 2,
	18, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3,
	2, 2, 2, 21, 4, 3, 2, 2, 2, 22, 24, 9, 3, 2, 2, 23, 22, 3, 2, 2, 2, 24,
	25, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 6, 3, 2, 2,
	2, 27, 29, 9, 4, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 8, 3, 2, 2, 2, 32, 33, 7, 45, 2, 2,
	33, 10, 3, 2, 2, 2, 34, 35, 7, 47, 2, 2, 35, 12, 3, 2, 2, 2, 36, 37, 7,
	61, 2, 2, 37, 14, 3, 2, 2, 2, 38, 40, 9, 5, 2, 2, 39, 38, 3, 2, 2, 2, 40,
	41, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2,
	2, 43, 44, 8, 8, 2, 2, 44, 16, 3, 2, 2, 2, 7, 2, 20, 25, 30, 41, 3, 8,
	2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "'+'", "'-'", "';'",
}

var lexerSymbolicNames = []string{
	"", "IDENTIFIER", "ASSET", "NUMBER", "OP_ADD", "OP_SUB", "SEP", "WHITESPACE",
}

var lexerRuleNames = []string{
	"IDENTIFIER", "ASSET", "NUMBER", "OP_ADD", "OP_SUB", "SEP", "WHITESPACE",
}

type NumScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewNumScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *NumScriptLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNumScriptLexer(input antlr.CharStream) *NumScriptLexer {
	l := new(NumScriptLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "NumScript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumScriptLexer tokens.
const (
	NumScriptLexerIDENTIFIER = 1
	NumScriptLexerASSET      = 2
	NumScriptLexerNUMBER     = 3
	NumScriptLexerOP_ADD     = 4
	NumScriptLexerOP_SUB     = 5
	NumScriptLexerSEP        = 6
	NumScriptLexerWHITESPACE = 7
)