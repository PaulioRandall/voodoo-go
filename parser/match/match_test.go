package match

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/stat"
	"github.com/PaulioRandall/voodoo-go/parser/token"

	"github.com/stretchr/testify/require"
)

func TestMatch_1(t *testing.T) {
	in := []token.Token{
		// x <- 1
		token.Dummy{`x`, token.TT_ID},
		token.Dummy{`<-`, token.TT_ASSIGN},
		token.Dummy{`1`, token.TT_NUMBER},
	}

	exp := stat.New(stat.SK_EXPRESSION, in[0:1], in[2:len(in)])
	act, e := Match(in)

	require.Nil(t, e)
	stat.AssertEqual(t, exp, act)
}
