package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanComment(t *testing.T, in string, exp *token.Token) {
	r := dummyRuner(in)
	tk, _, errTk := scanComment(r)
	token.AssertToken(t, exp, tk)
	token.AssertToken(t, nil, errTk)
}

func dummyCommentToken(end int, s string) token.Token {
	return token.DummyToken(0, 0, end, s, token.TT_COMMENT)
}

func TestScanComment_1(t *testing.T) {
	in := `// 123`
	exp := dummyCommentToken(6, `// 123`)
	doTestScanComment(t, in, &exp)
}

func TestScanComment_2(t *testing.T) {
	in := "// 123\n456"
	exp := dummyCommentToken(6, `// 123`)
	doTestScanComment(t, in, &exp)
}
