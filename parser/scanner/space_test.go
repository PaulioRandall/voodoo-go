package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func doTestScanSpace(t *testing.T, in string, exp *token.Token) {
	r := dummyRuner(in)
	tk, _, errTk := scanSpace(r)
	assertToken(t, exp, tk)
	assertToken(t, nil, errTk)
}

func dummySpaceToken(end int, s string) token.Token {
	return dummyToken(0, 0, end, s, token.TT_SPACE)
}

func TestScanSpace_1(t *testing.T) {
	in := ` `
	exp := dummySpaceToken(1, ` `)
	doTestScanSpace(t, in, &exp)
}

func TestScanSpace_2(t *testing.T) {
	in := "\t"
	exp := dummySpaceToken(1, "\t")
	doTestScanSpace(t, in, &exp)
}

func TestScanSpace_3(t *testing.T) {
	in := `   abc`
	exp := dummySpaceToken(3, `   `)
	doTestScanSpace(t, in, &exp)
}

func TestScanSpace_4(t *testing.T) {
	in := "\t\f \n\v\r"
	exp := dummySpaceToken(3, "\t\f ")
	doTestScanSpace(t, in, &exp)
}
