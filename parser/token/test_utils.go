package token

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertTokens asserts that the actual token array matches the expected token
// array.
func AssertTokens(t *testing.T, exp []Token, act []Token) {
	for i, expTk := range exp {
		if !assert.True(t, i < len(act), `Token %d missing`, i) {
			break
		}

		actTk := act[i]
		AssertToken(t, &expTk, &actTk)
	}

	assert.Equal(t, len(exp), len(act), `Expected len(tokens) == %d`, len(exp))
}

// AssertToken asserts that the actual token equals the expected token except
// for the error messages.
func AssertToken(t *testing.T, exp *Token, act *Token) (ok bool) {
	if exp == nil {
		ok = assert.Nil(t, act)
	} else {
		require.NotNil(t, act)
		ok = true
		ok = ok && assert.Equal(t, exp.Val, act.Val, TokenErrorString(act, exp, 0))
		ok = ok && assert.Equal(t, exp.Line, act.Line, TokenErrorString(act, exp, 1))
		ok = ok && assert.Equal(t, exp.Start, act.Start, TokenErrorString(act, exp, 2))
		ok = ok && assert.Equal(t, exp.End, act.End, TokenErrorString(act, exp, 3))
		ok = ok && assert.Equal(t, exp.Type, act.Type, TokenErrorString(act, exp, 4))
	}
	return
}

// TokenErrorString creates a string representation of a failed token assertion.
func TokenErrorString(tk *Token, exp *Token, errField int) string {

	sb := strings.Builder{}
	printErr := func(field int, exp string) {
		if exp != `` && field == errField {
			sb.WriteString("\n  ^------ Expected: ")
			sb.WriteString(exp)
		}
	}

	sb.WriteString("Token{")

	sb.WriteString("\n")
	sb.WriteString("  Val: ")
	sb.WriteString(strconv.QuoteToGraphic(tk.Val))
	printErr(0, strconv.QuoteToGraphic(exp.Val))

	sb.WriteString("\n")
	sb.WriteString("  Line: ")
	sb.WriteString(strconv.Itoa(tk.Line))
	printErr(1, strconv.Itoa(exp.Line))

	sb.WriteString("\n")
	sb.WriteString("  Start: ")
	sb.WriteString(strconv.Itoa(tk.Start))
	printErr(2, strconv.Itoa(exp.Start))

	sb.WriteString("\n")
	sb.WriteString("  End: ")
	sb.WriteString(strconv.Itoa(tk.End))
	printErr(3, strconv.Itoa(exp.End))

	sb.WriteString("\n")
	sb.WriteString("  Type: ")
	sb.WriteString(TokenName(tk.Type))
	printErr(4, TokenName(exp.Type))

	sb.WriteString("\n")
	sb.WriteString("  Errors: [")
	for _, v := range tk.Errors {
		sb.WriteString("\n    ")
		sb.WriteString(strconv.QuoteToGraphic(v))
	}
	sb.WriteString("\n  ]")
	printErr(5, "")

	sb.WriteString("\n}")
	return sb.String()
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

// OfType creates a new token initialised to the specified token type.
func OfType(t TokenType) Token {
	return Token{
		Type: t,
	}
}
