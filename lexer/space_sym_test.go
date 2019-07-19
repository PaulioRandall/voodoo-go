package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestSpaceSym(t *testing.T) {
	symFuncTest(t, "spaceSym", spaceSym, spaceSymTests())
}

func spaceSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     " ",
			ExpectSym: lexeme.Lexeme{" ", 0, 1, 0, lexeme.WHITESPACE},
		},
		symTest{
			Input:     "\t",
			ExpectSym: lexeme.Lexeme{"\t", 0, 1, 0, lexeme.WHITESPACE},
		},
		symTest{
			Input:     "\t\n \f \v\r",
			ExpectSym: lexeme.Lexeme{"\t\n \f \v\r", 0, 7, 0, lexeme.WHITESPACE},
		},
	}
}
