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
			Input:    []rune(`a`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `a`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`abc`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `abc`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`abc_123`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `abc_123`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`a__________123456789`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `a__________123456789`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`abc efg`),
			Output:   []rune(` efg`),
			Expect: token.Token{
				Val:  `abc`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`func`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `func`,
				Type: token.KEYWORD_FUNC,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`loop`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `loop`,
				Type: token.KEYWORD_LOOP,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`when`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `when`,
				Type: token.KEYWORD_WHEN,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`done`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `done`,
				Type: token.KEYWORD_DONE,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`true`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `true`,
				Type: token.BOOLEAN_TRUE,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`false`),
			Output:   []rune{},
			Expect: token.Token{
				Val:  `false`,
				Type: token.BOOLEAN_FALSE,
			},
		},
	}
}
