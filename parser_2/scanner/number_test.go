package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

func doTestScanNumber(t *testing.T, in string, exp *scanTok, expErr ScanError) bool {
	r := runer.NewByStr(in)
	tk, err := scanNumber(r)
	return logicalConjunction(
		AssertScanTokEqual(t, exp, tk),
		AssertScanError(t, expErr, err),
	)
}

func dummyNumToken(end int, text string) *scanTok {
	return &scanTok{
		text: text,
		end:  end,
		kind: token.TT_NUMBER,
	}
}

func dummyNumErr(i int) ScanError {
	return scanErr{
		i: i,
		e: []string{
			`:)`,
		},
	}
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
