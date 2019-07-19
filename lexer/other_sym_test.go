package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestOtherSym(t *testing.T) {
	symErrFuncTest(t, "otherSym", otherSym, otherSymTests())
}

func otherSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `==`,
			ExpectSym: lexeme.Symbol{`==`, 0, 2, 0, lexeme.EQUAL},
		},
		symTest{
			Input:     `!=`,
			ExpectSym: lexeme.Symbol{`!=`, 0, 2, 0, lexeme.NOT_EQUAL},
		},
		symTest{
			Input:     `<`,
			ExpectSym: lexeme.Symbol{`<`, 0, 1, 0, lexeme.LESS_THAN},
		},
		symTest{
			Input:     `<=`,
			ExpectSym: lexeme.Symbol{`<=`, 0, 2, 0, lexeme.LESS_THAN_OR_EQUAL},
		},
		symTest{
			Input:     `>`,
			ExpectSym: lexeme.Symbol{`>`, 0, 1, 0, lexeme.GREATER_THAN},
		},
		symTest{
			Input:     `>=`,
			ExpectSym: lexeme.Symbol{`>=`, 0, 2, 0, lexeme.GREATER_THAN_OR_EQUAL},
		},
		symTest{
			Input:     `||`,
			ExpectSym: lexeme.Symbol{`||`, 0, 2, 0, lexeme.OR},
		},
		symTest{
			Input:     `&&`,
			ExpectSym: lexeme.Symbol{`&&`, 0, 2, 0, lexeme.AND},
		},
		symTest{
			Input:     `<-`,
			ExpectSym: lexeme.Symbol{`<-`, 0, 2, 0, lexeme.ASSIGNMENT},
		},
		symTest{
			Input:     `=>`,
			ExpectSym: lexeme.Symbol{`=>`, 0, 2, 0, lexeme.IF_TRUE_THEN},
		},
		symTest{
			Input:     `_`,
			ExpectSym: lexeme.Symbol{`_`, 0, 1, 0, lexeme.VOID},
		},
		symTest{
			Input:     `!`,
			ExpectSym: lexeme.Symbol{`!`, 0, 1, 0, lexeme.NEGATION},
		},
		symTest{
			Input:     `+`,
			ExpectSym: lexeme.Symbol{`+`, 0, 1, 0, lexeme.ADD},
		},
		symTest{
			Input:     `-`,
			ExpectSym: lexeme.Symbol{`-`, 0, 1, 0, lexeme.SUBTRACT},
		},
		symTest{
			Input:     `*`,
			ExpectSym: lexeme.Symbol{`*`, 0, 1, 0, lexeme.MULTIPLY},
		},
		symTest{
			Input:     `/`,
			ExpectSym: lexeme.Symbol{`/`, 0, 1, 0, lexeme.DIVIDE},
		},
		symTest{
			Input:     `%`,
			ExpectSym: lexeme.Symbol{`%`, 0, 1, 0, lexeme.MODULO},
		},
		symTest{
			Input:     `(`,
			ExpectSym: lexeme.Symbol{`(`, 0, 1, 0, lexeme.CURVED_BRACE_OPEN},
		},
		symTest{
			Input:     `)`,
			ExpectSym: lexeme.Symbol{`)`, 0, 1, 0, lexeme.CURVED_BRACE_CLOSE},
		},
		symTest{
			Input:     `[`,
			ExpectSym: lexeme.Symbol{`[`, 0, 1, 0, lexeme.SQUARE_BRACE_OPEN},
		},
		symTest{
			Input:     `]`,
			ExpectSym: lexeme.Symbol{`]`, 0, 1, 0, lexeme.SQUARE_BRACE_CLOSE},
		},
		symTest{
			Input:     `,`,
			ExpectSym: lexeme.Symbol{`,`, 0, 1, 0, lexeme.VALUE_SEPARATOR},
		},
		symTest{
			Input:     `:`,
			ExpectSym: lexeme.Symbol{`:`, 0, 1, 0, lexeme.KEY_VALUE_SEPARATOR},
		},
		symTest{
			Input:     `..`,
			ExpectSym: lexeme.Symbol{`..`, 0, 2, 0, lexeme.RANGE},
		},
		symTest{
			Input:     `=`,
			ExpectErr: expLexError{0, 1},
		},
	}
}
