package word

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/scantok"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/require"
)

func doTestScanWord(t *testing.T, in string, exp token.Token) {
	r := runer.NewByStr(in)
	act, e := ScanWord(r)
	require.Nil(t, e, `Unexpected ScanError`)
	scantok.AssertEqual(t, exp, act)
}

func wordDummy(t string, e int, k token.Kind) token.Token {
	return scantok.New(t, 0, 0, e, k)
}

func TestScanWord_1(t *testing.T) {
	exp := wordDummy(`a`, 1, token.TT_ID)
	doTestScanWord(t, `a`, exp)
}

func TestScanWord_2(t *testing.T) {
	exp := wordDummy(`abc_123`, 7, token.TT_ID)
	doTestScanWord(t, `abc_123`, exp)
}

func TestScanWord_3(t *testing.T) {
	exp := wordDummy(`a__________123456789`, 20, token.TT_ID)
	doTestScanWord(t, `a__________123456789`, exp)
}

func TestScanWord_4(t *testing.T) {
	exp := wordDummy(`abc`, 3, token.TT_ID)
	doTestScanWord(t, `abc efg`, exp)
}
