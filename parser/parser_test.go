package parser

import (
	"strconv"
	"testing"

	//"github.com/PaulioRandall/voodoo-go/expr/ctx"
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type parseTest struct {
	TestLine int
	Input    []token.Token
	Expect   ParseTree
	Error    fault.Fault
}

func TestParser(t *testing.T) {
	for _, tc := range makeParseTests() {

		testLine := strconv.Itoa(tc.TestLine)
		t.Log("-> parser_test.go : " + testLine)

		var tree *ParseTree
		var err fault.Fault
		tree, err = Parse(tc.Input)

		if tc.Error != nil {
			assert.Nil(t, tree)
			require.NotNil(t, err)

		} else {
			assert.Nil(t, err)
			require.NotNil(t, tree)
			assert.Equal(t, tc.Expect, tree)
		}

	}
}

func makeParseTests() []parseTest {
	return []parseTest{}
}

/*
func makeParseTests() []parseTest {
	return []parseTest{
		parseTest{
			TestLine: fault.CurrLine(),
			Input: []token.Token{
				token.Token{`x`, 0, 1, token.IDENTIFIER_EXPLICIT},
				token.Token{`<-`, 2, 4, token.ASSIGNMENT},
				token.Token{`1`, 5, 6, token.LITERAL_NUMBER},
			},
		},
	}
}
*/
