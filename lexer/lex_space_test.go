package lexer

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestSpaceLex(t *testing.T) {
	lexFuncTest(t, "spaceLex", spaceLex, spaceLexTests())
}

func spaceLexTests() []lexTest {
	return []lexTest{
		lexTest{
			Input:     " ",
			ExpectLex: symbol.Lexeme{" ", 0, 1, 0, symbol.WHITESPACE},
		},
		lexTest{
			Input:     "\t",
			ExpectLex: symbol.Lexeme{"\t", 0, 1, 0, symbol.WHITESPACE},
		},
		lexTest{
			Input:     "\t\n \f \v\r",
			ExpectLex: symbol.Lexeme{"\t\n \f \v\r", 0, 7, 0, symbol.WHITESPACE},
		},
	}
}
