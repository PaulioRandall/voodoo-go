package number

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/utils"
)

func doTestScanNumber(t *testing.T, in string, exp token.Token, expErr perror.Perror) bool {
	r := runer.NewByStr(in)
	act, e := ScanNumber(r)
	return utils.LogicalConjunction(
		token.AssertEqual(t, exp, act),
		perror.AssertEqual(t, expErr, e),
	)
}

func dummyNumToken(end int, text string) token.Token {
	return token.New(text, 0, 0, end, token.TT_NUMBER)
}

func dummyNumErr(i int) perror.Perror {
	return perror.New(0, i, []string{`:)`})
}

func TestScanNumber_1(t *testing.T) {
	in := `123`
	exp := dummyNumToken(3, `123`)
	doTestScanNumber(t, in, exp, nil)
}

func TestScanNumber_2(t *testing.T) {
	in := `123 + 456`
	exp := dummyNumToken(3, `123`)
	doTestScanNumber(t, in, exp, nil)
}

func TestScanNumber_3(t *testing.T) {
	in := `123.456`
	exp := dummyNumToken(7, `123.456`)
	doTestScanNumber(t, in, exp, nil)
}

func TestScanNumber_4(t *testing.T) {
	in := `123.`
	expErr := dummyNumErr(4)
	doTestScanNumber(t, in, nil, expErr)
}

func TestScanNumber_5(t *testing.T) {
	in := `123..456`
	expErr := dummyNumErr(4)
	doTestScanNumber(t, in, nil, expErr)
}
