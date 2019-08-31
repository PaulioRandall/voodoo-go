package symbols

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/err"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

func doTestScanSymbol(t *testing.T, in string, exp token.Token, expErr err.ScanError) {
	r := runer.NewByStr(in)
	tk, e := ScanSymbol(r)
	scantok.AssertEqual(t, exp, tk)
	err.AssertEqual(t, expErr, e)
}

func dummySymToken(end int, text string, k token.Kind) token.Token {
	return scantok.New(text, 0, 0, end, k)
}

func TestScanSymbol_1(t *testing.T) {
	in := `:`
	exp := dummySymToken(1, `:`, token.TT_ASSIGN)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_4(t *testing.T) {
	in := `+`
	exp := dummySymToken(1, `+`, token.TT_ADD)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_5(t *testing.T) {
	in := `-`
	exp := dummySymToken(1, `-`, token.TT_SUBTRACT)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_6(t *testing.T) {
	in := `*`
	exp := dummySymToken(1, `*`, token.TT_MULTIPLY)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_7(t *testing.T) {
	in := `/`
	exp := dummySymToken(1, `/`, token.TT_DIVIDE)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_8(t *testing.T) {
	in := `%`
	exp := dummySymToken(1, `%`, token.TT_MODULO)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_9(t *testing.T) {
	in := `+ 69`
	exp := dummySymToken(1, `+`, token.TT_ADD)
	doTestScanSymbol(t, in, exp, nil)
}
