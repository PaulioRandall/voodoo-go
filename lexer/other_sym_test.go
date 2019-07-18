package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestOtherSym(t *testing.T) {
	symErrFuncTest(t, "otherSym", otherSym, otherSymTests())
}

func otherSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `==`,
			ExpectSym: symbol.Symbol{`==`, 0, 2, 0, symbol.EQUAL},
		},
		symTest{
			Input:     `!=`,
			ExpectSym: symbol.Symbol{`!=`, 0, 2, 0, symbol.NOT_EQUAL},
		},
		symTest{
			Input:     `<`,
			ExpectSym: symbol.Symbol{`<`, 0, 1, 0, symbol.LESS_THAN},
		},
		symTest{
			Input:     `<=`,
			ExpectSym: symbol.Symbol{`<=`, 0, 2, 0, symbol.LESS_THAN_OR_EQUAL},
		},
		symTest{
			Input:     `>`,
			ExpectSym: symbol.Symbol{`>`, 0, 1, 0, symbol.GREATER_THAN},
		},
		symTest{
			Input:     `>=`,
			ExpectSym: symbol.Symbol{`>=`, 0, 2, 0, symbol.GREATER_THAN_OR_EQUAL},
		},
		symTest{
			Input:     `||`,
			ExpectSym: symbol.Symbol{`||`, 0, 2, 0, symbol.OR},
		},
		symTest{
			Input:     `&&`,
			ExpectSym: symbol.Symbol{`&&`, 0, 2, 0, symbol.AND},
		},
		symTest{
			Input:     `<-`,
			ExpectSym: symbol.Symbol{`<-`, 0, 2, 0, symbol.ASSIGNMENT},
		},
		symTest{
			Input:     `=>`,
			ExpectSym: symbol.Symbol{`=>`, 0, 2, 0, symbol.IF_TRUE_THEN},
		},
		symTest{
			Input:     `_`,
			ExpectSym: symbol.Symbol{`_`, 0, 1, 0, symbol.VOID},
		},
		symTest{
			Input:     `!`,
			ExpectSym: symbol.Symbol{`!`, 0, 1, 0, symbol.NEGATION},
		},
		symTest{
			Input:     `+`,
			ExpectSym: symbol.Symbol{`+`, 0, 1, 0, symbol.ADD},
		},
		symTest{
			Input:     `-`,
			ExpectSym: symbol.Symbol{`-`, 0, 1, 0, symbol.SUBTRACT},
		},
		symTest{
			Input:     `*`,
			ExpectSym: symbol.Symbol{`*`, 0, 1, 0, symbol.MULTIPLY},
		},
		symTest{
			Input:     `/`,
			ExpectSym: symbol.Symbol{`/`, 0, 1, 0, symbol.DIVIDE},
		},
		symTest{
			Input:     `%`,
			ExpectSym: symbol.Symbol{`%`, 0, 1, 0, symbol.MODULO},
		},
		symTest{
			Input:     `(`,
			ExpectSym: symbol.Symbol{`(`, 0, 1, 0, symbol.CURVED_BRACE_OPEN},
		},
		symTest{
			Input:     `)`,
			ExpectSym: symbol.Symbol{`)`, 0, 1, 0, symbol.CURVED_BRACE_CLOSE},
		},
		symTest{
			Input:     `[`,
			ExpectSym: symbol.Symbol{`[`, 0, 1, 0, symbol.SQUARE_BRACE_OPEN},
		},
		symTest{
			Input:     `]`,
			ExpectSym: symbol.Symbol{`]`, 0, 1, 0, symbol.SQUARE_BRACE_CLOSE},
		},
		symTest{
			Input:     `,`,
			ExpectSym: symbol.Symbol{`,`, 0, 1, 0, symbol.VALUE_SEPARATOR},
		},
		symTest{
			Input:     `:`,
			ExpectSym: symbol.Symbol{`:`, 0, 1, 0, symbol.KEY_VALUE_SEPARATOR},
		},
		symTest{
			Input:     `..`,
			ExpectSym: symbol.Symbol{`..`, 0, 2, 0, symbol.RANGE},
		},
		symTest{
			Input:     `=`,
			ExpectErr: true,
		},
	}
}
