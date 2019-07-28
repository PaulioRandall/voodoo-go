package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanComment(t *testing.T) {
	runScanTest(t, "comment_test.go", scanComment, scanCommentTests())
}

func scanCommentTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `//`,
			Expect:   token.Token{`//`, 0, 2, 0, token.COMMENT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `// A comment`,
			Expect:   token.Token{`// A comment`, 0, 12, 0, token.COMMENT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `// Abc // 123 // xyz`,
			Expect:   token.Token{`// Abc // 123 // xyz`, 0, 20, 0, token.COMMENT},
		},
	}
}
