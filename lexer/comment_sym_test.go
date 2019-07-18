package lexer

import (
	"strconv"
	"testing"

	sh "github.com/PaulioRandall/voodoo-go/shared"
	sym "github.com/PaulioRandall/voodoo-go/symbol"
	"github.com/stretchr/testify/assert"
)

func TestCommentSym(t *testing.T) {
	for i, tc := range commentSymTests() {
		t.Log("commentSym() test case: " + strconv.Itoa(i+1))

		itr := sh.NewRuneItr(tc.Input)
		s := commentSym(itr)
		assert.Equal(t, tc.ExpectSym, *s)
	}
}

func commentSymTests() []symTest {
	return []symTest{
		symTest{
			Input:     `//`,
			ExpectSym: sym.Symbol{`//`, 0, 2, 0, sym.COMMENT},
		},
		symTest{
			Input:     `// A comment`,
			ExpectSym: sym.Symbol{`// A comment`, 0, 12, 0, sym.COMMENT},
		},
		symTest{
			Input:     `// Abc // 123 // xyz`,
			ExpectSym: sym.Symbol{`// Abc // 123 // xyz`, 0, 20, 0, sym.COMMENT},
		},
	}
}
