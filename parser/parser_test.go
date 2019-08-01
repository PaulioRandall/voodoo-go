package parser

import (
	"strconv"
	"testing"

	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parseTest struct {
	TestLine int
	Input    []token.Token
	Exes     []Instruction
	Values   []token.Token
	Error    fault.Fault
}

func TestParser(t *testing.T) {
	for _, tc := range makeParseTests() {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> parser_test.go : " + testLine)

		var exes []Instruction
		var values []token.Token
		var err fault.Fault
		exes, values, err = Parse(tc.Input)

		if tc.Error != nil {
			assert.Nil(t, exes)
			assert.Nil(t, values)
			require.NotNil(t, err)

		} else {
			assert.Nil(t, err)

			require.NotNil(t, exes)
			require.NotNil(t, values)

			assert.Equal(t, tc.Exes, exes)
			assert.Equal(t, tc.Values, values)
		}
	}
}

func makeParseTests() []parseTest {
	return []parseTest{
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
			},
			Exes: []Instruction{
				Instruction{
					Token:  token.Token{`<-`, 2, 4, token.ASSIGNMENT},
					Params: 2,
				},
			},
			Values: []token.Token{
				token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
			},
		},
	}
}
