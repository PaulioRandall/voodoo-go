package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestNumSym(t *testing.T) {
	for i, tc := range numSymTests() {
		t.Log("numSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s, err := numSym(itr)

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

func numSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   `2`,
			Expects: sym.Symbol{`2`, 0, 1, 0, sym.NUMBER},
		},
		symTest{
			Input:   `123`,
			Expects: sym.Symbol{`123`, 0, 3, 0, sym.NUMBER},
		},
		symTest{
			Input:   `123_456`,
			Expects: sym.Symbol{`123_456`, 0, 7, 0, sym.NUMBER},
		},
		symTest{
			Input:   `123.456`,
			Expects: sym.Symbol{`123.456`, 0, 7, 0, sym.NUMBER},
		},
		symTest{
			Input:   `123.456_789`,
			Expects: sym.Symbol{`123.456_789`, 0, 11, 0, sym.NUMBER},
		},
		symTest{
			Input:   `1__2__3__.__4__5__6__`,
			Expects: sym.Symbol{`1__2__3__.__4__5__6__`, 0, 21, 0, sym.NUMBER},
		},
		symTest{
			Input:   `123..456`,
			Expects: sym.Symbol{`123`, 0, 3, 0, sym.NUMBER},
		},
		symTest{
			Input:     `1_._2_._3`,
			ExpectErr: true,
		},
	}
}
