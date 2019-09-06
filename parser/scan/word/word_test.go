package word

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/require"
)

func doTestScanWord(t *testing.T, in string, exp token.Token) {
	r := runer.NewByStr(in)
	act, e := ScanWord(r)
	require.Nil(t, e, `Unexpected ScanError`)
	token.AssertEqual(t, exp, act)
}

func wordDummy(e int, t string, k token.Kind) token.Token {
	return token.New(t, 0, 0, e, k)
}

func TestScanWord_1(t *testing.T) {
	exp := wordDummy(1, `a`, token.TT_ID)
	doTestScanWord(t, `a`, exp)
}

func TestScanWord_2(t *testing.T) {
	exp := wordDummy(7, `abc_123`, token.TT_ID)
	doTestScanWord(t, `abc_123`, exp)
}

func TestScanWord_3(t *testing.T) {
	exp := wordDummy(20, `a__________123456789`, token.TT_ID)
	doTestScanWord(t, `a__________123456789`, exp)
}

func TestScanWord_4(t *testing.T) {
	exp := wordDummy(3, `abc`, token.TT_ID)
	doTestScanWord(t, `abc efg`, exp)
}

func TestScanWord_5(t *testing.T) {
	exp := wordDummy(4, `true`, token.TT_BOOL)
	doTestScanWord(t, `true`, exp)
}

func TestScanWord_6(t *testing.T) {
	exp := wordDummy(5, `false`, token.TT_BOOL)
	doTestScanWord(t, `false`, exp)
}
