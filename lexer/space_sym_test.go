package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestSpaceSym(t *testing.T) {
	for i, tc := range spaceSymTests() {
		t.Log("spaceSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s := spaceSym(itr)
		assert.Equal(t, tc.ExpectSym, *s)
	}
}

func spaceSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     " ",
			ExpectSym: sym.Symbol{" ", 0, 1, 0, sym.WHITESPACE},
		},
		symTest{
			Input:     "\t",
			ExpectSym: sym.Symbol{"\t", 0, 1, 0, sym.WHITESPACE},
		},
		symTest{
			Input:     "\t\n \f \v\r",
			ExpectSym: sym.Symbol{"\t\n \f \v\r", 0, 7, 0, sym.WHITESPACE},
		},
	}
}
