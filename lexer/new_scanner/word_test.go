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
			Expect: token.Token{
				Val:  `a`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`abc`),
			Expect: token.Token{
				Val:  `abc`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`abc_123`),
			Expect: token.Token{
				Val:  `abc_123`,
				Type: token.IDENTIFIER_EXPLICIT,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`a__________123456789`),
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
			Input:    []rune(`DO`),
			Expect: token.Token{
				Val:  `DO`,
				Type: token.KEYWORD_FUNC,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`loop`),
			Expect: token.Token{
				Val:  `loop`,
				Type: token.KEYWORD_LOOP,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`when`),
			Expect: token.Token{
				Val:  `when`,
				Type: token.KEYWORD_WHEN,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`done`),
			Expect: token.Token{
				Val:  `done`,
				Type: token.KEYWORD_END,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`true`),
			Expect: token.Token{
				Val:  `true`,
				Type: token.BOOLEAN_TRUE,
			},
		},
		scanFuncTest{
			TestLine: fault.CurrLine(),
			Input:    []rune(`false`),
			Expect: token.Token{
				Val:  `false`,
				Type: token.BOOLEAN_FALSE,
			},
		},
	}
}
