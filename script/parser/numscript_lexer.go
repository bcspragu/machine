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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 170,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 6, 20, 126, 10, 20, 13, 20, 14, 20, 127, 3, 21, 6, 21, 131,
	10, 21, 13, 21, 14, 21, 132, 3, 22, 3, 22, 6, 22, 137, 10, 22, 13, 22,
	14, 22, 138, 3, 22, 6, 22, 142, 10, 22, 13, 22, 14, 22, 143, 3, 23, 3,
	23, 6, 23, 148, 10, 23, 13, 23, 14, 23, 149, 3, 23, 6, 23, 153, 10, 23,
	13, 23, 14, 23, 154, 3, 24, 6, 24, 158, 10, 24, 13, 24, 14, 24, 159, 3,
	25, 3, 25, 3, 26, 6, 26, 165, 10, 26, 13, 26, 14, 26, 166, 3, 26, 3, 26,
	2, 2, 27, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20,
	39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 3, 2, 8, 3, 2,
	50, 59, 5, 2, 50, 60, 97, 97, 99, 124, 4, 2, 97, 97, 99, 124, 5, 2, 50,
	59, 97, 97, 99, 124, 4, 2, 49, 59, 67, 92, 5, 2, 11, 12, 15, 15, 34, 34,
	2, 177, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 3, 53, 3, 2, 2, 2, 5, 55, 3,
	2, 2, 2, 7, 60, 3, 2, 2, 2, 9, 66, 3, 2, 2, 2, 11, 71, 3, 2, 2, 2, 13,
	76, 3, 2, 2, 2, 15, 78, 3, 2, 2, 2, 17, 80, 3, 2, 2, 2, 19, 82, 3, 2, 2,
	2, 21, 84, 3, 2, 2, 2, 23, 86, 3, 2, 2, 2, 25, 88, 3, 2, 2, 2, 27, 90,
	3, 2, 2, 2, 29, 92, 3, 2, 2, 2, 31, 94, 3, 2, 2, 2, 33, 102, 3, 2, 2, 2,
	35, 108, 3, 2, 2, 2, 37, 115, 3, 2, 2, 2, 39, 125, 3, 2, 2, 2, 41, 130,
	3, 2, 2, 2, 43, 134, 3, 2, 2, 2, 45, 145, 3, 2, 2, 2, 47, 157, 3, 2, 2,
	2, 49, 161, 3, 2, 2, 2, 51, 164, 3, 2, 2, 2, 53, 54, 7, 46, 2, 2, 54, 4,
	3, 2, 2, 2, 55, 56, 7, 120, 2, 2, 56, 57, 7, 99, 2, 2, 57, 58, 7, 116,
	2, 2, 58, 59, 7, 117, 2, 2, 59, 6, 3, 2, 2, 2, 60, 61, 7, 114, 2, 2, 61,
	62, 7, 116, 2, 2, 62, 63, 7, 107, 2, 2, 63, 64, 7, 112, 2, 2, 64, 65, 7,
	118, 2, 2, 65, 8, 3, 2, 2, 2, 66, 67, 7, 104, 2, 2, 67, 68, 7, 99, 2, 2,
	68, 69, 7, 107, 2, 2, 69, 70, 7, 110, 2, 2, 70, 10, 3, 2, 2, 2, 71, 72,
	7, 117, 2, 2, 72, 73, 7, 103, 2, 2, 73, 74, 7, 112, 2, 2, 74, 75, 7, 102,
	2, 2, 75, 12, 3, 2, 2, 2, 76, 77, 7, 45, 2, 2, 77, 14, 3, 2, 2, 2, 78,
	79, 7, 47, 2, 2, 79, 16, 3, 2, 2, 2, 80, 81, 7, 42, 2, 2, 81, 18, 3, 2,
	2, 2, 82, 83, 7, 43, 2, 2, 83, 20, 3, 2, 2, 2, 84, 85, 7, 93, 2, 2, 85,
	22, 3, 2, 2, 2, 86, 87, 7, 95, 2, 2, 87, 24, 3, 2, 2, 2, 88, 89, 7, 125,
	2, 2, 89, 26, 3, 2, 2, 2, 90, 91, 7, 127, 2, 2, 91, 28, 3, 2, 2, 2, 92,
	93, 7, 63, 2, 2, 93, 30, 3, 2, 2, 2, 94, 95, 7, 99, 2, 2, 95, 96, 7, 101,
	2, 2, 96, 97, 7, 101, 2, 2, 97, 98, 7, 113, 2, 2, 98, 99, 7, 119, 2, 2,
	99, 100, 7, 112, 2, 2, 100, 101, 7, 118, 2, 2, 101, 32, 3, 2, 2, 2, 102,
	103, 7, 99, 2, 2, 103, 104, 7, 117, 2, 2, 104, 105, 7, 117, 2, 2, 105,
	106, 7, 103, 2, 2, 106, 107, 7, 118, 2, 2, 107, 34, 3, 2, 2, 2, 108, 109,
	7, 112, 2, 2, 109, 110, 7, 119, 2, 2, 110, 111, 7, 111, 2, 2, 111, 112,
	7, 100, 2, 2, 112, 113, 7, 103, 2, 2, 113, 114, 7, 116, 2, 2, 114, 36,
	3, 2, 2, 2, 115, 116, 7, 111, 2, 2, 116, 117, 7, 113, 2, 2, 117, 118, 7,
	112, 2, 2, 118, 119, 7, 103, 2, 2, 119, 120, 7, 118, 2, 2, 120, 121, 7,
	99, 2, 2, 121, 122, 7, 116, 2, 2, 122, 123, 7, 123, 2, 2, 123, 38, 3, 2,
	2, 2, 124, 126, 9, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2,
	127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 40, 3, 2, 2, 2, 129, 131,
	9, 3, 2, 2, 130, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 130, 3, 2,
	2, 2, 132, 133, 3, 2, 2, 2, 133, 42, 3, 2, 2, 2, 134, 136, 7, 38, 2, 2,
	135, 137, 9, 4, 2, 2, 136, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138,
	136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 141, 3, 2, 2, 2, 140, 142,
	9, 5, 2, 2, 141, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 141, 3, 2,
	2, 2, 143, 144, 3, 2, 2, 2, 144, 44, 3, 2, 2, 2, 145, 147, 7, 66, 2, 2,
	146, 148, 9, 4, 2, 2, 147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 153,
	9, 3, 2, 2, 152, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 152, 3, 2,
	2, 2, 154, 155, 3, 2, 2, 2, 155, 46, 3, 2, 2, 2, 156, 158, 9, 6, 2, 2,
	157, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159,
	160, 3, 2, 2, 2, 160, 48, 3, 2, 2, 2, 161, 162, 7, 61, 2, 2, 162, 50, 3,
	2, 2, 2, 163, 165, 9, 7, 2, 2, 164, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2,
	2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168,
	169, 8, 26, 2, 2, 169, 52, 3, 2, 2, 2, 11, 2, 127, 132, 138, 143, 149,
	154, 159, 166, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'vars'", "'print'", "'fail'", "'send'", "'+'", "'-'", "'('",
	"')'", "'['", "']'", "'{'", "'}'", "'='", "'account'", "'asset'", "'number'",
	"'monetary'", "", "", "", "", "", "';'",
}

var lexerSymbolicNames = []string{
	"", "", "VARS", "PRINT", "FAIL", "SEND", "OP_ADD", "OP_SUB", "LPAREN",
	"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET",
	"TY_NUMBER", "TY_MONETARY", "NUMBER", "IDENTIFIER", "VARIABLE_NAME", "ACCOUNT",
	"ASSET", "SEP", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "VARS", "PRINT", "FAIL", "SEND", "OP_ADD", "OP_SUB", "LPAREN",
	"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET",
	"TY_NUMBER", "TY_MONETARY", "NUMBER", "IDENTIFIER", "VARIABLE_NAME", "ACCOUNT",
	"ASSET", "SEP", "WHITESPACE",
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
	NumScriptLexerT__0          = 1
	NumScriptLexerVARS          = 2
	NumScriptLexerPRINT         = 3
	NumScriptLexerFAIL          = 4
	NumScriptLexerSEND          = 5
	NumScriptLexerOP_ADD        = 6
	NumScriptLexerOP_SUB        = 7
	NumScriptLexerLPAREN        = 8
	NumScriptLexerRPAREN        = 9
	NumScriptLexerLBRACK        = 10
	NumScriptLexerRBRACK        = 11
	NumScriptLexerLBRACE        = 12
	NumScriptLexerRBRACE        = 13
	NumScriptLexerEQ            = 14
	NumScriptLexerTY_ACCOUNT    = 15
	NumScriptLexerTY_ASSET      = 16
	NumScriptLexerTY_NUMBER     = 17
	NumScriptLexerTY_MONETARY   = 18
	NumScriptLexerNUMBER        = 19
	NumScriptLexerIDENTIFIER    = 20
	NumScriptLexerVARIABLE_NAME = 21
	NumScriptLexerACCOUNT       = 22
	NumScriptLexerASSET         = 23
	NumScriptLexerSEP           = 24
	NumScriptLexerWHITESPACE    = 25
)
