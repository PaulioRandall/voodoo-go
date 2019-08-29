package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
	"github.com/stretchr/testify/require"
)

func doTestScanWord(t *testing.T, in string, exp *scanTok) {
	r := runer.NewByStr(in)
	tk, err := scanWord(r)
	require.Nil(t, err, `Unexpected ScanError`)
	AssertScanTokEqual(t, exp, tk)
}

func TestScanWord_1(t *testing.T) {
	exp := &scanTok{
		text: `a`,
		end:  1,
		kind: token.TT_ID,
	}

	doTestScanWord(t, `a`, exp)
}

func TestScanWord_2(t *testing.T) {
	exp := &scanTok{
		text: `abc_123`,
		end:  7,
		kind: token.TT_ID,
	}

	doTestScanWord(t, `abc_123`, exp)
}

func TestScanWord_3(t *testing.T) {
	exp := &scanTok{
		text: `a__________123456789`,
		end:  20,
		kind: token.TT_ID,
	}

	doTestScanWord(t, `a__________123456789`, exp)
}

func TestScanWord_4(t *testing.T) {
	exp := &scanTok{
		text: `abc`,
		end:  3,
		kind: token.TT_ID,
	}

	doTestScanWord(t, `abc efg`, exp)
}
