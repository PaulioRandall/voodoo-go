package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

func TestScanWord(t *testing.T) {
	runScanTest(t, "word_test.go", scanWord, scanWordTests())
}

func scanWordTests() []scanFuncTest {
	return []scanFuncTest{
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `a`,
			Expect:   symbol.Token{`a`, 0, 1, 0, symbol.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `abc`,
			Expect:   symbol.Token{`abc`, 0, 3, 0, symbol.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `abc_123`,
			Expect:   symbol.Token{`abc_123`, 0, 7, 0, symbol.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `a__________123456789`,
			Expect:   symbol.Token{`a__________123456789`, 0, 20, 0, symbol.IDENTIFIER_EXPLICIT},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `fUnC`,
			Expect:   symbol.Token{`fUnC`, 0, 4, 0, symbol.KEYWORD_FUNC},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `loop`,
			Expect:   symbol.Token{`loop`, 0, 4, 0, symbol.KEYWORD_LOOP},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `when`,
			Expect:   symbol.Token{`when`, 0, 4, 0, symbol.KEYWORD_WHEN},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `end`,
			Expect:   symbol.Token{`end`, 0, 3, 0, symbol.KEYWORD_END},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `key`,
			Expect:   symbol.Token{`key`, 0, 3, 0, symbol.KEYWORD_KEY},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `val`,
			Expect:   symbol.Token{`val`, 0, 3, 0, symbol.KEYWORD_VALUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `true`,
			Expect:   symbol.Token{`true`, 0, 4, 0, symbol.BOOLEAN_TRUE},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    `false`,
			Expect:   symbol.Token{`false`, 0, 5, 0, symbol.BOOLEAN_FALSE},
		},
	}
}
