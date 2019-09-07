package string

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/utils"
)

func doTestScanString(t *testing.T, in string, exp token.Token, expErr perror.Perror) bool {
	r := runer.NewByStr(in)
	act, e := ScanString(r)
	return utils.LogicalConjunction(
		token.AssertEqual(t, exp, act),
		perror.AssertEqual(t, expErr, e),
	)
}

func dummy(end int, text string) token.Token {
	return token.New(text, 0, 0, end, token.TK_STRING)
}

func dummyErr(i int) perror.Perror {
	return perror.New(0, i, []string{`:)`})
}

func TestScanString_1(t *testing.T) {
	in := "`abc`"
	exp := dummy(5, "`abc`")
	doTestScanString(t, in, exp, nil)
}

func TestScanString_2(t *testing.T) {
	in := "``"
	exp := dummy(2, "``")
	doTestScanString(t, in, exp, nil)
}

func TestScanString_3(t *testing.T) {
	in := "````"
	exp := dummy(2, "``")
	doTestScanString(t, in, exp, nil)
}
