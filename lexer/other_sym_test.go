package lexer

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOtherSym(t *testing.T) {
	for i, tc := range otherSymTests() {
		t.Log("otherSym() test case: " + strconv.Itoa(i+1))

		itr := NewStrItr(tc.Input)
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
			Expects: Symbol{`=>`, 0, 2, 0, UNDEFINED},
		},
		symTest{
			Input:   `!`,
			Expects: Symbol{`!`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `+`,
			Expects: Symbol{`+`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `-`,
			Expects: Symbol{`-`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `*`,
			Expects: Symbol{`*`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `/`,
			Expects: Symbol{`/`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `%`,
			Expects: Symbol{`%`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `(`,
			Expects: Symbol{`(`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `)`,
			Expects: Symbol{`)`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `[`,
			Expects: Symbol{`[`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `]`,
			Expects: Symbol{`]`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `,`,
			Expects: Symbol{`,`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `:`,
			Expects: Symbol{`:`, 0, 1, 0, UNDEFINED},
		},
		symTest{
			Input:   `..`,
			Expects: Symbol{`..`, 0, 2, 0, UNDEFINED},
		},
		symTest{
			Input:     `=`,
			ExpectErr: true,
		},
	}
}
