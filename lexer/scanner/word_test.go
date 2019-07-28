package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

func TestScanWord(t *testing.T) {
	runScanTest(t, "word_test.go", scanWord, scanWordTests())
}

func scanWordTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `a`,
			Expect:   token.Token{`a`, 0, 1, 0, token.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `abc`,
			Expect:   token.Token{`abc`, 0, 3, 0, token.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `abc_123`,
			Expect:   token.Token{`abc_123`, 0, 7, 0, token.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `a__________123456789`,
			Expect:   token.Token{`a__________123456789`, 0, 20, 0, token.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `fUnC`,
			Expect:   token.Token{`fUnC`, 0, 4, 0, token.KEYWORD_FUNC},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `loop`,
			Expect:   token.Token{`loop`, 0, 4, 0, token.KEYWORD_LOOP},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `when`,
			Expect:   token.Token{`when`, 0, 4, 0, token.KEYWORD_WHEN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `end`,
			Expect:   token.Token{`end`, 0, 3, 0, token.KEYWORD_END},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `key`,
			Expect:   token.Token{`key`, 0, 3, 0, token.KEYWORD_KEY},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `val`,
			Expect:   token.Token{`val`, 0, 3, 0, token.KEYWORD_VALUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `true`,
			Expect:   token.Token{`true`, 0, 4, 0, token.BOOLEAN_TRUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `false`,
			Expect:   token.Token{`false`, 0, 5, 0, token.BOOLEAN_FALSE},
		},
	}
}
