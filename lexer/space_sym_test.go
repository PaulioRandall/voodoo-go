package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestSpaceSym(t *testing.T) {
	symFuncTest(t, "spaceSym", spaceSym, spaceSymTests())
}

func spaceSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     " ",
			ExpectSym: symbol.Symbol{" ", 0, 1, 0, symbol.WHITESPACE},
		},
		symTest{
			Input:     "\t",
			ExpectSym: symbol.Symbol{"\t", 0, 1, 0, symbol.WHITESPACE},
		},
		symTest{
			Input:     "\t\n \f \v\r",
			ExpectSym: symbol.Symbol{"\t\n \f \v\r", 0, 7, 0, symbol.WHITESPACE},
		},
	}
}
