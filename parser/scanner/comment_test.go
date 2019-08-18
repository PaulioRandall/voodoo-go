package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanComment(t *testing.T) {
	runScanTokenTests(t, "comment_test.go", scanComment, scanCommentTests())
}

func scanCommentTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `// 123`,
			Expect:         dummyToken(0, 0, 6, `// 123`, token.TT_COMMENT),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "// 123\n456",
			Expect:         dummyToken(0, 0, 6, `// 123`, token.TT_COMMENT),
			NextUnreadRune: '\n',
		},
	}
}
