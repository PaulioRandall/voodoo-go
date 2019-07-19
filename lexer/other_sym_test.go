package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestOtherSym(t *testing.T) {
	symErrFuncTest(t, "otherSym", otherSym, otherSymTests())
}

func otherSymTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     `==`,
			ExpectSym: lexeme.Lexeme{`==`, 0, 2, 0, lexeme.EQUAL},
		},
		lexTest{
			Input:     `!=`,
			ExpectSym: lexeme.Lexeme{`!=`, 0, 2, 0, lexeme.NOT_EQUAL},
		},
		lexTest{
			Input:     `<`,
			ExpectSym: lexeme.Lexeme{`<`, 0, 1, 0, lexeme.LESS_THAN},
		},
		lexTest{
			Input:     `<=`,
			ExpectSym: lexeme.Lexeme{`<=`, 0, 2, 0, lexeme.LESS_THAN_OR_EQUAL},
		},
		lexTest{
			Input:     `>`,
			ExpectSym: lexeme.Lexeme{`>`, 0, 1, 0, lexeme.GREATER_THAN},
		},
		lexTest{
			Input:     `>=`,
			ExpectSym: lexeme.Lexeme{`>=`, 0, 2, 0, lexeme.GREATER_THAN_OR_EQUAL},
		},
		lexTest{
			Input:     `||`,
			ExpectSym: lexeme.Lexeme{`||`, 0, 2, 0, lexeme.OR},
		},
		lexTest{
			Input:     `&&`,
			ExpectSym: lexeme.Lexeme{`&&`, 0, 2, 0, lexeme.AND},
		},
		lexTest{
			Input:     `<-`,
			ExpectSym: lexeme.Lexeme{`<-`, 0, 2, 0, lexeme.ASSIGNMENT},
		},
		lexTest{
			Input:     `=>`,
			ExpectSym: lexeme.Lexeme{`=>`, 0, 2, 0, lexeme.IF_TRUE_THEN},
		},
		lexTest{
			Input:     `_`,
			ExpectSym: lexeme.Lexeme{`_`, 0, 1, 0, lexeme.VOID},
		},
		lexTest{
			Input:     `!`,
			ExpectSym: lexeme.Lexeme{`!`, 0, 1, 0, lexeme.NEGATION},
		},
		lexTest{
			Input:     `+`,
			ExpectSym: lexeme.Lexeme{`+`, 0, 1, 0, lexeme.ADD},
		},
		lexTest{
			Input:     `-`,
			ExpectSym: lexeme.Lexeme{`-`, 0, 1, 0, lexeme.SUBTRACT},
		},
		lexTest{
			Input:     `*`,
			ExpectSym: lexeme.Lexeme{`*`, 0, 1, 0, lexeme.MULTIPLY},
		},
		lexTest{
			Input:     `/`,
			ExpectSym: lexeme.Lexeme{`/`, 0, 1, 0, lexeme.DIVIDE},
		},
		lexTest{
			Input:     `%`,
			ExpectSym: lexeme.Lexeme{`%`, 0, 1, 0, lexeme.MODULO},
		},
		lexTest{
			Input:     `(`,
			ExpectSym: lexeme.Lexeme{`(`, 0, 1, 0, lexeme.CURVED_BRACE_OPEN},
		},
		lexTest{
			Input:     `)`,
			ExpectSym: lexeme.Lexeme{`)`, 0, 1, 0, lexeme.CURVED_BRACE_CLOSE},
		},
		lexTest{
			Input:     `[`,
			ExpectSym: lexeme.Lexeme{`[`, 0, 1, 0, lexeme.SQUARE_BRACE_OPEN},
		},
		lexTest{
			Input:     `]`,
			ExpectSym: lexeme.Lexeme{`]`, 0, 1, 0, lexeme.SQUARE_BRACE_CLOSE},
		},
		lexTest{
			Input:     `,`,
			ExpectSym: lexeme.Lexeme{`,`, 0, 1, 0, lexeme.VALUE_SEPARATOR},
		},
		lexTest{
			Input:     `:`,
			ExpectSym: lexeme.Lexeme{`:`, 0, 1, 0, lexeme.KEY_VALUE_SEPARATOR},
		},
		lexTest{
			Input:     `..`,
			ExpectSym: lexeme.Lexeme{`..`, 0, 2, 0, lexeme.RANGE},
		},
		lexTest{
			Input:     `=`,
			ExpectErr: expLexError{0, 1},
		},
	}
}
