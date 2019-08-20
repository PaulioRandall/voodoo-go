package scanner_new

import (
	"bufio"
	"strings"
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// dummyRuner creates a new Runer from the input string.
func dummyRuner(s string) *Runer {
	sr := strings.NewReader(s)
	br := bufio.NewReader(sr)
	return NewRuner(br)
}

// dummyToken creates a new dummy token.
func dummyToken(line, start, end int, v string, t token.TokenType) token.Token {
	return token.Token{
		Line:  line,
		Start: start,
		End:   end,
		Val:   v,
		Type:  t,
	}
}

// errDummyToken creates a new error dummy token.
func errDummyToken(line, start, end int) token.Token {
	return token.Token{
		Line:  line,
		Start: start,
		End:   end,
		Type:  token.TT_ERROR_UPSTREAM,
	}
}

// readRequireNoErr reads the next rune from the Runer and asserts that no error
// occurred while reading. If an error was returned the test immediately exits.
func readRequireNoErr(t *testing.T, r *Runer) rune {
	ru, err := r.ReadRune()
	require.Nil(t, err)
	return ru
}

// assertToken asserts that the actual token equals the expected token except
// for the error messages.
func assertToken(t *testing.T, exp *token.Token, act *token.Token) {
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
