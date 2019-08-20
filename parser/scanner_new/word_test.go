package scanner_new

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

func TestScanWord(t *testing.T) {
	parserTokenTester(t, "word_test.go", scanWord, scanWordTests())
}

func dummyWordToken(end int, s string, t token.TokenType) token.Token {
	return dummyToken(0, 0, end, s, t)
}

func scanWordTests() []parseTokenTest {
	return []parseTokenTest{
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `a`,
			Expect:   dummyWordToken(1, `a`, token.TT_ID),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `abc_123`,
			Expect:   dummyWordToken(7, `abc_123`, token.TT_ID),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `a__________123456789`,
			Expect:   dummyWordToken(20, `a__________123456789`, token.TT_ID),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `abc efg`,
			Expect:   dummyWordToken(3, `abc`, token.TT_ID),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `func`,
			Expect:   dummyWordToken(4, `func`, token.TT_FUNC),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `loop`,
			Expect:   dummyWordToken(4, `loop`, token.TT_LOOP),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `when`,
			Expect:   dummyWordToken(4, `when`, token.TT_WHEN),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `done`,
			Expect:   dummyWordToken(4, `done`, token.TT_DONE),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `true`,
			Expect:   dummyWordToken(4, `true`, token.TT_TRUE),
		},
		parseTokenTest{
			TestLine: fault.CurrLine(),
			Input:    `false`,
			Expect:   dummyWordToken(5, `false`, token.TT_FALSE),
		},
	}
}
