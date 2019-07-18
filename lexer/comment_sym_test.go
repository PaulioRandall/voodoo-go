package lexer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/runer"
	"github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestCommentSym(t *testing.T) {
	for i, tc := range commentSymTests() {
		t.Log("commentSym() test case: " + strconv.Itoa(i+1))

		itr := runer.NewRuneItr(tc.Input)
		s := commentSym(itr)
		assert.Equal(t, tc.ExpectSym, *s)
	}
}

func commentSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `//`,
			ExpectSym: symbol.Symbol{`//`, 0, 2, 0, symbol.COMMENT},
		},
		symTest{
			Input:     `// A comment`,
			ExpectSym: symbol.Symbol{`// A comment`, 0, 12, 0, symbol.COMMENT},
		},
		symTest{
			Input:     `// Abc // 123 // xyz`,
			ExpectSym: symbol.Symbol{`// Abc // 123 // xyz`, 0, 20, 0, symbol.COMMENT},
		},
	}
}
