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
		assert.Equal(t, tc.Expects, *s)
	}
}

func spaceSymTests() []symTest {
	return []symTest{
		symTest{
			Input:   " ",
			Expects: sym.Symbol{" ", 0, 1, 0, sym.WHITESPACE},
		},
		symTest{
			Input:   "\t",
			Expects: sym.Symbol{"\t", 0, 1, 0, sym.WHITESPACE},
		},
		symTest{
			Input:   "\t\n \f \v\r",
			Expects: sym.Symbol{"\t\n \f \v\r", 0, 7, 0, sym.WHITESPACE},
		},
	}
}
