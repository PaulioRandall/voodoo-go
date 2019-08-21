package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanString(t *testing.T, in string, exp, expErr *token.Token) {
	r := dummyRuner(in)
	tk, _, errTk := scanString(r)
	assertToken(t, exp, tk)
	assertToken(t, expErr, errTk)
}

func dummyStrToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_STRING)
}

func TestScanString_1(t *testing.T) {
	in := `""`
	exp := dummyStrToken(2, `""`)
	doTestScanString(t, in, &exp, nil)
}

func TestScanString_2(t *testing.T) {
	in := `"From hell with love"`
	exp := dummyStrToken(21, `"From hell with love"`)
	doTestScanString(t, in, &exp, nil)
}

func TestScanString_3(t *testing.T) {
	in := `"From hell with love", 123.456`
	exp := dummyStrToken(21, `"From hell with love"`)
	doTestScanString(t, in, &exp, nil)
}

func TestScanString_4(t *testing.T) {
	in := `"Simon: \"Leaders eat last!\""`
	exp := dummyStrToken(30, `"Simon: \"Leaders eat last!\""`)
	doTestScanString(t, in, &exp, nil)
}

func TestScanString_5(t *testing.T) {
	in := `"\\\\\""`
	exp := dummyStrToken(8, `"\\\\\""`)
	doTestScanString(t, in, &exp, nil)
}

func TestScanString_6(t *testing.T) {
	in := `"`
	expErr := errDummyToken(0, 0, 1)
	doTestScanString(t, in, nil, &expErr)
}

func TestScanString_7(t *testing.T) {
	in := `"escaped \"`
	expErr := errDummyToken(0, 0, 11)
	doTestScanString(t, in, nil, &expErr)
}

/*
		tfTest{
			TestLine: fault.CurrLine(),
			Input:    `"escaped \"`,
			Expect:   errDummyToken(0, 0, 11),
		},
	}
}
*/
