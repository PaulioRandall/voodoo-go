package scanner_new

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanNumber(t *testing.T, in string, exp token.Token) {
	r := dummyRuner(in)
	tk, _ := scanNumber(r)
	assertToken(t, exp, *tk)
}

func dummyNumToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_NUMBER)
}

func dummyNumErrToken(end int) token.Token {
	return dummyToken(0, 0, end, ``, token.TT_ERROR_UPSTREAM)
}

func TestScanNumber_1(t *testing.T) {
	in := `123`
	exp := dummyNumToken(3, `123`)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_2(t *testing.T) {
	in := `123 + 456`
	exp := dummyNumToken(3, `123`)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_3(t *testing.T) {
	in := `123_456`
	exp := dummyNumToken(7, `123_456`)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_4(t *testing.T) {
	in := `123.456`
	exp := dummyNumToken(7, `123.456`)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_5(t *testing.T) {
	in := `123.456_789`
	exp := dummyNumToken(11, `123.456_789`)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_6(t *testing.T) {
	in := `1__2__3__.__4__5__6__`
	exp := dummyNumToken(21, `1__2__3__.__4__5__6__`)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_7(t *testing.T) {
	in := `123.`
	exp := dummyNumErrToken(4)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_8(t *testing.T) {
	in := `123..456`
	exp := dummyNumErrToken(4)
	doTestScanNumber(t, in, exp)
}

func TestScanNumber_9(t *testing.T) {
	in := `123.___`
	exp := dummyNumErrToken(7)
	doTestScanNumber(t, in, exp)
}
