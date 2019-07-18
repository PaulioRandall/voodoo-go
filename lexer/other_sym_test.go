package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestOtherSym(t *testing.T) {
	for i, tc := range otherSymTests() {
		t.Log("otherSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s, err := otherSym(itr)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			if assert.NotNil(t, s) {
				assert.Equal(t, tc.Expects, *s)
			}
		}
	}
}

func otherSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `==`,
			Expects: sym.Symbol{`==`, 0, 2, 0, sym.EQUAL},
		},
		symTest{
			Input:   `!=`,
			Expects: sym.Symbol{`!=`, 0, 2, 0, sym.NOT_EQUAL},
		},
		symTest{
			Input:   `<`,
			Expects: sym.Symbol{`<`, 0, 1, 0, sym.LESS_THAN},
		},
		symTest{
			Input:   `<=`,
			Expects: sym.Symbol{`<=`, 0, 2, 0, sym.LESS_THAN_OR_EQUAL},
		},
		symTest{
			Input:   `>`,
			Expects: sym.Symbol{`>`, 0, 1, 0, sym.GREATER_THAN},
		},
		symTest{
			Input:   `>=`,
			Expects: sym.Symbol{`>=`, 0, 2, 0, sym.GREATER_THAN_OR_EQUAL},
		},
		symTest{
			Input:   `||`,
			Expects: sym.Symbol{`||`, 0, 2, 0, sym.OR},
		},
		symTest{
			Input:   `&&`,
			Expects: sym.Symbol{`&&`, 0, 2, 0, sym.AND},
		},
		symTest{
			Input:   `<-`,
			Expects: sym.Symbol{`<-`, 0, 2, 0, sym.ASSIGNMENT},
		},
		symTest{
			Input:   `=>`,
			Expects: sym.Symbol{`=>`, 0, 2, 0, sym.IF_TRUE_THEN},
		},
		symTest{
			Input:   `_`,
			Expects: sym.Symbol{`_`, 0, 1, 0, sym.VOID},
		},
		symTest{
			Input:   `!`,
			Expects: sym.Symbol{`!`, 0, 1, 0, sym.NEGATION},
		},
		symTest{
			Input:   `+`,
			Expects: sym.Symbol{`+`, 0, 1, 0, sym.ADD},
		},
		symTest{
			Input:   `-`,
			Expects: sym.Symbol{`-`, 0, 1, 0, sym.SUBTRACT},
		},
		symTest{
			Input:   `*`,
			Expects: sym.Symbol{`*`, 0, 1, 0, sym.MULTIPLY},
		},
		symTest{
			Input:   `/`,
			Expects: sym.Symbol{`/`, 0, 1, 0, sym.DIVIDE},
		},
		symTest{
			Input:   `%`,
			Expects: sym.Symbol{`%`, 0, 1, 0, sym.MODULO},
		},
		symTest{
			Input:   `(`,
			Expects: sym.Symbol{`(`, 0, 1, 0, sym.CURVED_BRACE_OPEN},
		},
		symTest{
			Input:   `)`,
			Expects: sym.Symbol{`)`, 0, 1, 0, sym.CURVED_BRACE_CLOSE},
		},
		symTest{
			Input:   `[`,
			Expects: sym.Symbol{`[`, 0, 1, 0, sym.SQUARE_BRACE_OPEN},
		},
		symTest{
			Input:   `]`,
			Expects: sym.Symbol{`]`, 0, 1, 0, sym.SQUARE_BRACE_CLOSE},
		},
		symTest{
			Input:   `,`,
			Expects: sym.Symbol{`,`, 0, 1, 0, sym.VALUE_SEPARATOR},
		},
		symTest{
			Input:   `:`,
			Expects: sym.Symbol{`:`, 0, 1, 0, sym.KEY_VALUE_SEPARATOR},
		},
		symTest{
			Input:   `..`,
			Expects: sym.Symbol{`..`, 0, 2, 0, sym.RANGE},
		},
		symTest{
			Input:     `=`,
			ExpectErr: true,
		},
	}
}
