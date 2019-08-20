package scanner_new

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanComment(t *testing.T, in string, exp *token.Token) {
	r := dummyRuner(in)
	tk, _, errTk := scanComment(r)
	assertToken(t, exp, tk)
	assertToken(t, nil, errTk)
}

func dummyCommentToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_COMMENT)
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
