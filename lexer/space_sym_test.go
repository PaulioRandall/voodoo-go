package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
)

func TestSpaceSym(t *testing.T) {
	lexFuncTest(t, "spaceSym", spaceSym, spaceSymTests())
}

func spaceSymTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     " ",
			ExpectLex: lexeme.Lexeme{" ", 0, 1, 0, lexeme.WHITESPACE},
		},
		lexTest{
			Input:     "\t",
			ExpectLex: lexeme.Lexeme{"\t", 0, 1, 0, lexeme.WHITESPACE},
		},
		lexTest{
			Input:     "\t\n \f \v\r",
			ExpectLex: lexeme.Lexeme{"\t\n \f \v\r", 0, 7, 0, lexeme.WHITESPACE},
		},
	}
}
