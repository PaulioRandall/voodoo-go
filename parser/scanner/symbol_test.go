package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanSymbol(t *testing.T, in string, exp, expErr *token.Token) {
	r := dummyRuner(in)
	tk, _, errTk := scanSymbol(r)
	token.AssertToken(t, exp, tk)
	token.AssertToken(t, expErr, errTk)
}

func dummySymToken(end int, s string, t token.TokenType) token.Token {
	return token.DummyToken(0, 0, end, s, t)
}

func TestScanSymbol_1(t *testing.T) {
	in := `==`
	exp := dummySymToken(2, `==`, token.TT_CMP_EQ)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_2(t *testing.T) {
	in := `!=`
	exp := dummySymToken(2, `!=`, token.TT_CMP_NOT_EQ)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_3(t *testing.T) {
	in := `<`
	exp := dummySymToken(1, `<`, token.TT_CMP_LT)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_4(t *testing.T) {
	in := `<=`
	exp := dummySymToken(2, `<=`, token.TT_CMP_LT_OR_EQ)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_5(t *testing.T) {
	in := `>`
	exp := dummySymToken(1, `>`, token.TT_CMP_MT)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_6(t *testing.T) {
	in := `>=`
	exp := dummySymToken(2, `>=`, token.TT_CMP_MT_OR_EQ)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_7(t *testing.T) {
	in := `<-`
	exp := dummySymToken(2, `<-`, token.TT_ASSIGN)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_8(t *testing.T) {
	in := `||`
	exp := dummySymToken(2, `||`, token.TT_OR)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_9(t *testing.T) {
	in := `&&`
	exp := dummySymToken(2, `&&`, token.TT_AND)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_10(t *testing.T) {
	in := `!`
	exp := dummySymToken(1, `!`, token.TT_NOT)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_11_1(t *testing.T) {
	in := `=>`
	exp := dummySymToken(2, `=>`, token.TT_MATCH)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_11_2(t *testing.T) {
	in := `=`
	exp := dummySymToken(1, `=`, token.TT_ASSIGN)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_11_3(t *testing.T) {
	in := `:=`
	exp := dummySymToken(2, `:=`, token.TT_ASSIGN)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_12(t *testing.T) {
	in := `_`
	exp := dummySymToken(1, `_`, token.TT_VOID)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_13(t *testing.T) {
	in := `+`
	exp := dummySymToken(1, `+`, token.TT_ADD)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_14(t *testing.T) {
	in := `-`
	exp := dummySymToken(1, `-`, token.TT_SUBTRACT)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_15(t *testing.T) {
	in := `*`
	exp := dummySymToken(1, `*`, token.TT_MULTIPLY)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_16(t *testing.T) {
	in := `/`
	exp := dummySymToken(1, `/`, token.TT_DIVIDE)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_17(t *testing.T) {
	in := `%`
	exp := dummySymToken(1, `%`, token.TT_MODULO)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_18(t *testing.T) {
	in := `(`
	exp := dummySymToken(1, `(`, token.TT_CURVED_OPEN)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_19(t *testing.T) {
	in := `)`
	exp := dummySymToken(1, `)`, token.TT_CURVED_CLOSE)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_20(t *testing.T) {
	in := `[`
	exp := dummySymToken(1, `[`, token.TT_SQUARE_OPEN)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_21(t *testing.T) {
	in := `]`
	exp := dummySymToken(1, `]`, token.TT_SQUARE_CLOSE)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_22(t *testing.T) {
	in := `,`
	exp := dummySymToken(1, `,`, token.TT_VALUE_DELIM)
	doTestScanSymbol(t, in, &exp, nil)
}

func TestScanSymbol_23(t *testing.T) {
	in := `+ 69`
	exp := dummySymToken(1, `+`, token.TT_ADD)
	doTestScanSymbol(t, in, &exp, nil)
}
