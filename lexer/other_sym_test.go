package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	"github.com/stretchr/testify/assert"
)

func TestOtherSym(t *testing.T) {
	for i, tc := range otherSymTests() {
		t.Log("otherSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		a, err := otherSym(itr, tc.Line)

		if tc.ExpectErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, tc.Expects, a)
		}
	}
}

func otherSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `==`,
			Expects: Symbol{`==`, 0, 2, 0, EQUAL},
		},
		symTest{
			Input:   `!=`,
			Expects: Symbol{`!=`, 0, 2, 0, NOT_EQUAL},
		},
		symTest{
			Input:   `<`,
			Expects: Symbol{`<`, 0, 1, 0, LESS_THAN},
		},
		symTest{
			Input:   `<=`,
			Expects: Symbol{`<=`, 0, 2, 0, LESS_THAN_OR_EQUAL},
		},
		symTest{
			Input:   `>`,
			Expects: Symbol{`>`, 0, 1, 0, GREATER_THAN},
		},
		symTest{
			Input:   `>=`,
			Expects: Symbol{`>=`, 0, 2, 0, GREATER_THAN_OR_EQUAL},
		},
		symTest{
			Input:   `||`,
			Expects: Symbol{`||`, 0, 2, 0, OR},
		},
		symTest{
			Input:   `&&`,
			Expects: Symbol{`&&`, 0, 2, 0, AND},
		},
		symTest{
			Input:   `<-`,
			Expects: Symbol{`<-`, 0, 2, 0, ASSIGNMENT},
		},
		symTest{
			Input:   `=>`,
			Expects: Symbol{`=>`, 0, 2, 0, IF_TRUE_THEN},
		},
		symTest{
			Input:   `_`,
			Expects: Symbol{`_`, 0, 1, 0, VOID},
		},
		symTest{
			Input:   `!`,
			Expects: Symbol{`!`, 0, 1, 0, NEGATION},
		},
		symTest{
			Input:   `+`,
			Expects: Symbol{`+`, 0, 1, 0, ADD},
		},
		symTest{
			Input:   `-`,
			Expects: Symbol{`-`, 0, 1, 0, SUBTRACT},
		},
		symTest{
			Input:   `*`,
			Expects: Symbol{`*`, 0, 1, 0, MULTIPLY},
		},
		symTest{
			Input:   `/`,
			Expects: Symbol{`/`, 0, 1, 0, DIVIDE},
		},
		symTest{
			Input:   `%`,
			Expects: Symbol{`%`, 0, 1, 0, MODULO},
		},
		symTest{
			Input:   `(`,
			Expects: Symbol{`(`, 0, 1, 0, CIRCLE_BRACE_OPEN},
		},
		symTest{
			Input:   `)`,
			Expects: Symbol{`)`, 0, 1, 0, CIRCLE_BRACE_CLOSE},
		},
		symTest{
			Input:   `[`,
			Expects: Symbol{`[`, 0, 1, 0, SQUARE_BRACE_OPEN},
		},
		symTest{
			Input:   `]`,
			Expects: Symbol{`]`, 0, 1, 0, SQUARE_BRACE_CLOSE},
		},
		symTest{
			Input:   `,`,
			Expects: Symbol{`,`, 0, 1, 0, VALUE_SEPARATOR},
		},
		symTest{
			Input:   `:`,
			Expects: Symbol{`:`, 0, 1, 0, KEY_VALUE_SEPARATOR},
		},
		symTest{
			Input:   `..`,
			Expects: Symbol{`..`, 0, 2, 0, RANGE},
		},
		symTest{
			Input:     `=`,
			ExpectErr: true,
		},
	}
}
