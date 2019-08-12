package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanComment(t *testing.T) {
	runScanTest(t, "comment_test.go", scanComment, scanCommentTests())
}

func scanCommentTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `// 123`,
			Expect:         dummyToken(0, 0, 6, `// 123`, token.COMMENT),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          "// 123\n456",
			Expect:         dummyToken(0, 0, 6, `// 123`, token.COMMENT),
			NextUnreadRune: '\n',
		},
	}
}
