package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestComment(t *testing.T) {
	lexFuncTest(t, "comment_test.go", commentLex, commentLexTests())
}

func commentLexTests() []lexTest {
	return []lexTest{
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `//`,
			Expect:   symbol.Lexeme{`//`, 0, 2, 0, symbol.COMMENT},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `// A comment`,
			Expect:   symbol.Lexeme{`// A comment`, 0, 12, 0, symbol.COMMENT},
		},
		lexTest{
			TestLine: fault.CurrLine(),
			Input:    `// Abc // 123 // xyz`,
			Expect:   symbol.Lexeme{`// Abc // 123 // xyz`, 0, 20, 0, symbol.COMMENT},
		},
	}
}
