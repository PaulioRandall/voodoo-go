package token

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AssertTokens asserts that the actual token array matches the expected token
// array.
func AssertSlice(t *testing.T, exp []Token, act []Token) {
	for i, expTk := range exp {
		if !assert.True(t, i < len(act), `Token %d missing`, i) {
			break
		}

		actTk := act[i]
		Assert(t, &expTk, &actTk)
	}

	assert.Equal(t, len(exp), len(act), `Expected len(tokens) == %d`, len(exp))
}

// Assert asserts that the actual token equals the expected token except
// for the error messages.
func Assert(t *testing.T, exp *Token, act *Token) (ok bool) {
	if exp == nil {
		ok = assert.Nil(t, act)
	} else {
		require.NotNil(t, act)
		a := assert.Equal(t, exp.Val, act.Val, newErrMsg(exp.Val, act.Val))
		b := assert.Equal(t, exp.Line, act.Line, newErrMsg(exp.Line, act.Line))
		c := assert.Equal(t, exp.Start, act.Start, newErrMsg(exp.Start, act.Start))
		d := assert.Equal(t, exp.End, act.End, newErrMsg(exp.End, act.End))
		e := assert.Equal(t, exp.Kind, act.Kind, newErrMsg(exp.Kind, act.Kind))
		ok = a && b && c && d && e
	}
	return
}

// newErrMsg returns an assertion error message from the input for the token
// assertion functions.
func newErrMsg(exp interface{}, act interface{}) string {
	return fmt.Sprintf(`Token assertion failed: %v == %v`, exp, act)
}

// Dummy creates a new dummy token.
func Dummy(line, start, end int, v string, k Kind) Token {
	return Token{
		Line:  line,
		Start: start,
		End:   end,
		Val:   v,
		Kind:  k,
	}
}

// PtrDummy creates a new dummy token then returns a pointer to it.
func PtrDummy(line, start, end int, v string, k Kind) *Token {
	return &Token{
		Line:  line,
		Start: start,
		End:   end,
		Val:   v,
		Kind:  k,
	}
}

// OfKind creates a new token initialised to the specified token type.
func OfKind(k Kind) Token {
	return Token{
		Kind: k,
	}
}

// OfKindUnique creates a new token initialised to the specified token type and
// a unique value
func OfKindUnique(k Kind) Token {
	return Token{
		Val:  strconv.FormatUint(rand.Uint64(), 10),
		Kind: k,
	}
}
