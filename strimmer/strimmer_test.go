package strimmer

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/lexeme"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrim(t *testing.T) {
	for i, tc := range strimTests() {
		t.Log("Strim() test case: " + strconv.Itoa(i+1))
		ts, err := Strim(tc.Input)

		if tc.ExpectErr != nil {
			require.Nil(t, ts)
			assert.NotEmpty(t, err.Error())
			assert.Equal(t, tc.ExpectErr.Line(), err.Line())
			assert.Equal(t, tc.ExpectErr.Col(), err.Col())
		}

		if tc.ExpectToks != nil {
			require.Nil(t, err)
			assert.Equal(t, tc.ExpectToks, ts)
		}
	}
}

type strimTest struct {
	Input      []lexeme.Lexeme
	ExpectToks []Token
	ExpectErr  StrimError
}

type expStrimError struct {
	line int // Line number
	col  int // Column number
}

func (e expStrimError) Error() string {
	// Error messages should be semantically validated
	// so this is not required for testing.
	return ""
}

func (e expStrimError) Line() int {
	return e.line
}

func (e expStrimError) Col() int {
	return e.col
}

func strimTests() []strimTest {
	return []strimTest{
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				lexeme.Lexeme{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`1`, 5, 6, 0, lexeme.NUMBER},
			},
			ExpectToks: []Token{
				Token{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				Token{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				Token{`1`, 5, 6, 0, lexeme.NUMBER},
			},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				lexeme.Lexeme{` `, 1, 2, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				lexeme.Lexeme{` `, 4, 5, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`2`, 5, 6, 0, lexeme.NUMBER},
				lexeme.Lexeme{` `, 6, 7, 0, lexeme.WHITESPACE},
				lexeme.Lexeme{`// 'There's a snake in my boot'`, 7, 38, 0, lexeme.COMMENT},
			},
			ExpectToks: []Token{
				Token{`x`, 0, 1, 0, lexeme.IDENTIFIER},
				Token{`<-`, 2, 4, 0, lexeme.ASSIGNMENT},
				Token{`2`, 5, 6, 0, lexeme.NUMBER},
			},
		},
		strimTest{
			Input: []lexeme.Lexeme{
				lexeme.Lexeme{`// 'There's a snake in my boot'`, 0, 31, 0, lexeme.COMMENT},
			},
			ExpectToks: []Token{},
		},
	}
}
