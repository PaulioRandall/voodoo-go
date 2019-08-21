package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanWord(t *testing.T, in string, exp *token.Token) {
	r := dummyRuner(in)
	tk, _, errTk := scanWord(r)
	token.AssertToken(t, exp, tk)
	token.AssertToken(t, nil, errTk)
}

func dummyWordToken(end int, s string, t token.TokenType) token.Token {
	return token.DummyToken(0, 0, end, s, t)
}

func TestScanWord_1(t *testing.T) {
	in := `a`
	exp := dummyWordToken(1, `a`, token.TT_ID)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_2(t *testing.T) {
	in := `abc_123`
	exp := dummyWordToken(7, `abc_123`, token.TT_ID)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_3(t *testing.T) {
	in := `a__________123456789`
	exp := dummyWordToken(20, `a__________123456789`, token.TT_ID)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_4(t *testing.T) {
	in := `abc efg`
	exp := dummyWordToken(3, `abc`, token.TT_ID)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_5(t *testing.T) {
	in := `func`
	exp := dummyWordToken(4, `func`, token.TT_FUNC)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_6(t *testing.T) {
	in := `loop`
	exp := dummyWordToken(4, `loop`, token.TT_LOOP)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_7(t *testing.T) {
	in := `when`
	exp := dummyWordToken(4, `when`, token.TT_WHEN)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_8(t *testing.T) {
	in := `done`
	exp := dummyWordToken(4, `done`, token.TT_DONE)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_9(t *testing.T) {
	in := `true`
	exp := dummyWordToken(4, `true`, token.TT_TRUE)
	doTestScanWord(t, in, &exp)
}

func TestScanWord_10(t *testing.T) {
	in := `false`
	exp := dummyWordToken(5, `false`, token.TT_FALSE)
	doTestScanWord(t, in, &exp)
}
