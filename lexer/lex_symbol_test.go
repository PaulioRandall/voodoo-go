package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestSymbolLex(t *testing.T) {
	lexErrFuncTest(t, "symbolLex", symbolLex, symbolLexTests())
}

func symbolLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `==`,
			ExpectLex: symbol.Lexeme{`==`, 0, 2, 0, symbol.EQUAL},
		},
		lexTest{
			Input:     `!=`,
			ExpectLex: symbol.Lexeme{`!=`, 0, 2, 0, symbol.NOT_EQUAL},
		},
		lexTest{
			Input:     `<`,
			ExpectLex: symbol.Lexeme{`<`, 0, 1, 0, symbol.LESS_THAN},
		},
		lexTest{
			Input:     `<=`,
			ExpectLex: symbol.Lexeme{`<=`, 0, 2, 0, symbol.LESS_THAN_OR_EQUAL},
		},
		lexTest{
			Input:     `>`,
			ExpectLex: symbol.Lexeme{`>`, 0, 1, 0, symbol.GREATER_THAN},
		},
		lexTest{
			Input:     `>=`,
			ExpectLex: symbol.Lexeme{`>=`, 0, 2, 0, symbol.GREATER_THAN_OR_EQUAL},
		},
		lexTest{
			Input:     `||`,
			ExpectLex: symbol.Lexeme{`||`, 0, 2, 0, symbol.OR},
		},
		lexTest{
			Input:     `&&`,
			ExpectLex: symbol.Lexeme{`&&`, 0, 2, 0, symbol.AND},
		},
		lexTest{
			Input:     `<-`,
			ExpectLex: symbol.Lexeme{`<-`, 0, 2, 0, symbol.ASSIGNMENT},
		},
		lexTest{
			Input:     `=>`,
			ExpectLex: symbol.Lexeme{`=>`, 0, 2, 0, symbol.IF_MATCH_THEN},
		},
		lexTest{
			Input:     `_`,
			ExpectLex: symbol.Lexeme{`_`, 0, 1, 0, symbol.VOID},
		},
		lexTest{
			Input:     `!`,
			ExpectLex: symbol.Lexeme{`!`, 0, 1, 0, symbol.NEGATION},
		},
		lexTest{
			Input:     `+`,
			ExpectLex: symbol.Lexeme{`+`, 0, 1, 0, symbol.ADD},
		},
		lexTest{
			Input:     `-`,
			ExpectLex: symbol.Lexeme{`-`, 0, 1, 0, symbol.SUBTRACT},
		},
		lexTest{
			Input:     `*`,
			ExpectLex: symbol.Lexeme{`*`, 0, 1, 0, symbol.MULTIPLY},
		},
		lexTest{
			Input:     `/`,
			ExpectLex: symbol.Lexeme{`/`, 0, 1, 0, symbol.DIVIDE},
		},
		lexTest{
			Input:     `%`,
			ExpectLex: symbol.Lexeme{`%`, 0, 1, 0, symbol.MODULO},
		},
		lexTest{
			Input:     `(`,
			ExpectLex: symbol.Lexeme{`(`, 0, 1, 0, symbol.CURVED_BRACE_OPEN},
		},
		lexTest{
			Input:     `)`,
			ExpectLex: symbol.Lexeme{`)`, 0, 1, 0, symbol.CURVED_BRACE_CLOSE},
		},
		lexTest{
			Input:     `[`,
			ExpectLex: symbol.Lexeme{`[`, 0, 1, 0, symbol.SQUARE_BRACE_OPEN},
		},
		lexTest{
			Input:     `]`,
			ExpectLex: symbol.Lexeme{`]`, 0, 1, 0, symbol.SQUARE_BRACE_CLOSE},
		},
		lexTest{
			Input:     `,`,
			ExpectLex: symbol.Lexeme{`,`, 0, 1, 0, symbol.VALUE_SEPARATOR},
		},
		lexTest{
			Input:     `:`,
			ExpectLex: symbol.Lexeme{`:`, 0, 1, 0, symbol.KEY_VALUE_SEPARATOR},
		},
		lexTest{
			Input:     `..`,
			ExpectLex: symbol.Lexeme{`..`, 0, 2, 0, symbol.RANGE},
		},
		lexTest{
			Input:     `=`,
			ExpectErr: expLexError{0, 1},
		},
	}
}
