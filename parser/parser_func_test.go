package parser

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/token"
	"github.com/stretchr/testify/assert"
)

func NewValue(val string, t token.TokenType) Expression {
	return Value{
		Token: Token{
			Val:  val,
			Type: t,
		},
	}
}

func TestSplitOnAssignment_1(t *testing.T) {
	in := []Token{
		Token{`x`, 0, 0, token.IDENTIFIER},
		Token{`<-`, 0, 0, token.ASSIGNMENT},
		Token{`1`, 0, 0, token.LITERAL_NUMBER},
		Token{`2`, 0, 0, token.LITERAL_NUMBER},
	}

	exp_a := []Token{
		Token{`x`, 0, 0, token.IDENTIFIER},
		Token{`<-`, 0, 0, token.ASSIGNMENT},
	}

	exp_out := []Token{
		Token{`1`, 0, 0, token.LITERAL_NUMBER},
		Token{`2`, 0, 0, token.LITERAL_NUMBER},
	}

	a, out := splitOnAssignment(in)
	assert.Equal(t, exp_a, a)
	assert.Equal(t, exp_out, out)
}

func TestSplitOnAssignment_2(t *testing.T) {
	in := []Token{
		Token{`@Spell`, 0, 0, token.SPELL},
		Token{`(`, 0, 0, token.PAREN_CURVY_OPEN},
		Token{`1`, 0, 0, token.LITERAL_NUMBER},
		Token{`(`, 0, 0, token.PAREN_CURVY_CLOSE},
	}

	var exp_a []Token = nil
	exp_out := in

	a, out := splitOnAssignment(in)
	assert.Equal(t, exp_a, a)
	assert.Equal(t, exp_out, out)
}

func TestParseAssignment_1(t *testing.T) {
	in := []Token{
		Token{`x`, 0, 0, token.IDENTIFIER},
		Token{`<-`, 0, 0, token.ASSIGNMENT},
	}

	exp_tk := Token{`<-`, 0, 0, token.ASSIGNMENT}

	exp_out := List{
		Tokens: []Token{
			Token{`x`, 0, 0, token.IDENTIFIER},
		},
	}

	tk, out, err := parseAssignment(in)

	assert.Nil(t, err)
	assert.Equal(t, exp_tk, tk)
	assert.Equal(t, exp_out, out)
}

func TestParseAssignment_2(t *testing.T) {
	in := []Token{
		Token{`x`, 0, 0, token.IDENTIFIER},
		Token{`,`, 0, 0, token.SEPARATOR_VALUE},
		Token{`y`, 0, 0, token.IDENTIFIER},
		Token{`<-`, 0, 0, token.ASSIGNMENT},
	}

	exp_tk := Token{`<-`, 0, 0, token.ASSIGNMENT}

	exp_out := List{
		Tokens: []Token{
			Token{`x`, 0, 0, token.IDENTIFIER},
			Token{`y`, 0, 0, token.IDENTIFIER},
		},
	}

	tk, out, err := parseAssignment(in)

	assert.Nil(t, err)
	assert.Equal(t, exp_tk, tk)
	assert.Equal(t, exp_out, out)
}

func TestSplitOnToken_1(t *testing.T) {

	in := []Token{
		Token{`x`, 0, 0, token.IDENTIFIER},
		Token{`,`, 0, 0, token.SEPARATOR_VALUE},
		Token{`y`, 0, 0, token.IDENTIFIER},
		Token{`<-`, 0, 0, token.ASSIGNMENT},
	}

	exp := [][]Token{
		[]Token{
			Token{`x`, 0, 0, token.IDENTIFIER},
		},
		[]Token{
			Token{`y`, 0, 0, token.IDENTIFIER},
			Token{`<-`, 0, 0, token.ASSIGNMENT},
		},
	}

	var out [][]Token
	out = splitOnToken(in, token.SEPARATOR_VALUE)

	assert.Equal(t, exp, out)
}

func TestParseExpression_1(t *testing.T) {
	in := []Token{
		Token{`x`, 0, 0, token.IDENTIFIER},
	}

	exp_out := NewValue(`x`, token.IDENTIFIER)

	out, err := parseExpression(in)

	assert.Nil(t, err)
	assert.Equal(t, exp_out, out)
}

func TestParseExpression_2(t *testing.T) {
	in := []Token{
		Token{`x`, 0, 0, token.IDENTIFIER},
		Token{`+`, 0, 0, token.CALC_ADD},
		Token{`y`, 0, 0, token.IDENTIFIER},
	}

	exp_out := Operation{
		Left:     NewValue(`x`, token.IDENTIFIER),
		Operator: Token{`+`, 0, 0, token.CALC_ADD},
		Right:    NewValue(`y`, token.IDENTIFIER),
	}

	out, err := parseExpression(in)

	assert.Nil(t, err)
	assert.Equal(t, exp_out, out)
}
