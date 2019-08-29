package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

func doTestScanSymbol(t *testing.T, in string, exp *scanTok, expErr ScanError) {
	r := runer.NewByStr(in)
	tk, err := scanSymbol(r)
	AssertScanTokEqual(t, exp, tk)
	AssertScanError(t, expErr, err)
}

func dummySymToken(end int, text string, k token.Kind) *scanTok {
	return &scanTok{
		text: text,
		end:  end,
		kind: k,
	}
}

func TestScanSymbol_1(t *testing.T) {
	in := `<-`
	exp := dummySymToken(2, `<-`, token.TT_ASSIGN)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_2(t *testing.T) {
	in := `=`
	exp := dummySymToken(1, `=`, token.TT_ASSIGN)
	doTestScanSymbol(t, in, exp, nil)
}

func TestScanSymbol_3(t *testing.T) {
	in := `:=`
	exp := dummySymToken(2, `:=`, token.TT_ASSIGN)
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
