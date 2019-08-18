package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanComment(t *testing.T) {
	runScanTokenTests(t, "comment_test.go", scanComment, scanCommentTests())
}

func dummyCommentToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_COMMENT)
}

func scanCommentTests() []tfTest {
	return []tfTest{
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          `// 123`,
			Expect:         dummyCommentToken(6, `// 123`),
			NextUnreadRune: EOF,
		},
		tfTest{
			TestLine:       fault.CurrLine(),
			Input:          "// 123\n456",
			Expect:         dummyCommentToken(6, `// 123`),
			NextUnreadRune: '\n',
		},
	}
}
