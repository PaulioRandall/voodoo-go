package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertTokens asserts that the actual token array matches the expected token
// array.
func AssertTokens(t *testing.T, exp []Token, act []Token) {
	for i, expTk := range exp {
		if !assert.True(t, i < len(act)) {
			break
		}

		actTk := act[i]
		AssertToken(t, &expTk, &actTk)
	}

	assert.Equal(t, len(exp), len(act))
}

// AssertToken asserts that the actual token equals the expected token except
// for the error messages.
func AssertToken(t *testing.T, exp *Token, act *Token) {
	if exp == nil {
		assert.Nil(t, act)
	} else {
		require.NotNil(t, act)
		assert.Equal(t, exp.Val, act.Val)
		assert.Equal(t, exp.Line, act.Line)
		assert.Equal(t, exp.Start, act.Start)
		assert.Equal(t, exp.End, act.End)
		assert.Equal(t, exp.Type, act.Type)
	}
}

// DummyToken creates a new dummy token.
func DummyToken(line, start, end int, v string, t TokenType) Token {
	return Token{
		Line:  line,
		Start: start,
		End:   end,
		Val:   v,
		Type:  t,
	}
}

// PtrDummyToken creates a new pointer to a new dummy token.
func PtrDummyToken(line, start, end int, v string, t TokenType) *Token {
	return &Token{
		Line:  line,
		Start: start,
		End:   end,
		Val:   v,
		Type:  t,
	}
}

// ErrDummyToken creates a new error dummy token.
func ErrDummyToken(line, start, end int) Token {
	return Token{
		Line:  line,
		Start: start,
		End:   end,
		Type:  TT_ERROR_UPSTREAM,
	}
}
