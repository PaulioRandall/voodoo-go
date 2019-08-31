package space

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scan/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
	"github.com/stretchr/testify/require"
)

func doTestScanSpace(t *testing.T, in string, exp token.Token) {
	r := runer.NewByStr(in)
	act, e := ScanSpace(r)
	require.Nil(t, e, `Unexpected ScanError`)
	scantok.AssertEqual(t, exp, act)
}

func dummy(t string, e int, k token.Kind) token.Token {
	return scantok.New(t, 0, 0, e, k)
}

func TestScanSpace_1(t *testing.T) {
	exp := dummy(` `, 1, token.TT_SPACE)
	doTestScanSpace(t, ` `, exp)
}

func TestScanSpace_2(t *testing.T) {
	exp := dummy(" \r\v\f\t", 5, token.TT_SPACE)
	doTestScanSpace(t, " \r\v\f\t", exp)
}

func TestScanSpace_3(t *testing.T) {
	exp := dummy("   ", 3, token.TT_SPACE)
	doTestScanSpace(t, "   \r\n   ", exp)
}

func TestScanSpace_4(t *testing.T) {
	exp := dummy("   ", 3, token.TT_SPACE)
	doTestScanSpace(t, "   \n   ", exp)
}

func TestScanSpace_5(t *testing.T) {
	exp := dummy("   ", 3, token.TT_SPACE)
	doTestScanSpace(t, "   abc   ", exp)
}

func TestScanSpace_6(t *testing.T) {
	exp := dummy("\n", 1, token.TT_NEWLINE)
	doTestScanSpace(t, "\n", exp)
}

func TestScanSpace_7(t *testing.T) {
	exp := dummy("\r\n", 2, token.TT_NEWLINE)
	doTestScanSpace(t, "\r\n", exp)
}

func TestScanSpace_8(t *testing.T) {
	exp := dummy("\n", 1, token.TT_NEWLINE)
	doTestScanSpace(t, "\n   ", exp)
}
