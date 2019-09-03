package farm

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
)

func assertFarm(t *testing.T, multiline, mature bool, exp, act *Farm) {
	assert.Equal(t, multiline, act.multiline)
	assert.Equal(t, mature, act.mature)
	token.AssertSliceEqual(t, exp.tokens, act.tokens)
}

func dummy(t string, k token.Kind) token.Token {
	return token.New(t, 0, 0, 0, k)
}

func TestFarm_Harvest_1(t *testing.T) {
	tks := []token.Token{
		// x <- 1
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`1`, token.TT_NUMBER),
	}

	exp := &Farm{}

	act := &Farm{
		mature: true,
		tokens: token.Copy(tks),
	}

	out := act.Harvest()
	token.AssertSliceEqual(t, tks, out)
	assertFarm(t, false, false, exp, act)
}

func TestFarm_FinalHarvest_1(t *testing.T) {
	tks := []token.Token{
		// x <- 1
		dummy(`x`, token.TT_ID),
		dummy(`<-`, token.TT_ASSIGN),
		dummy(`1`, token.TT_NUMBER),
	}

	exp := &Farm{}

	act := &Farm{
		tokens: token.Copy(tks),
	}

	out := act.FinalHarvest()
	token.AssertSliceEqual(t, tks, out)
	assertFarm(t, false, false, exp, act)
}
