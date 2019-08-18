package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanWord(t *testing.T) {
	runScanTest_NEW(t, "word_test.go", scanWord, scanWordTests())
}

func scanWordTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `a`,
			Expect:         dummyToken(0, 0, 1, `a`, token.TT_ID),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `abc_123`,
			Expect:         dummyToken(0, 0, 7, `abc_123`, token.TT_ID),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `a__________123456789`,
			Expect:         dummyToken(0, 0, 20, `a__________123456789`, token.TT_ID),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `abc efg`,
			Expect:         dummyToken(0, 0, 3, `abc`, token.TT_ID),
			NextUnreadRune: ' ',
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `func`,
			Expect:         dummyToken(0, 0, 4, `func`, token.TT_FUNC),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `loop`,
			Expect:         dummyToken(0, 0, 4, `loop`, token.TT_LOOP),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `when`,
			Expect:         dummyToken(0, 0, 4, `when`, token.TT_WHEN),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `done`,
			Expect:         dummyToken(0, 0, 4, `done`, token.TT_DONE),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `true`,
			Expect:         dummyToken(0, 0, 4, `true`, token.TT_TRUE),
			NextUnreadRune: EOF,
		},
		scanFuncTest{
			TestLine:       fault.CurrLine(),
			Input:          `false`,
			Expect:         dummyToken(0, 0, 5, `false`, token.TT_FALSE),
			NextUnreadRune: EOF,
		},
	}
}
