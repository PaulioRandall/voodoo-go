package farm

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/token"
	"github.com/stretchr/testify/assert"
)

func assertFarm(t *testing.T, multiline, mature bool, exp, act *Farm) {
	assert.Equal(t, multiline, act.multiline)
	assert.Equal(t, mature, act.mature)
	token.AssertSliceEqual(t, exp.tokens, act.tokens)
}

func TestFarm_Harvest_1(t *testing.T) {
	tks := []token.Token{
		// x <- 1
		token.Dummy{`x`, token.TT_ID},
		token.Dummy{`<-`, token.TT_ASSIGN},
		token.Dummy{`1`, token.TT_NUMBER},
	}

	exp := &Farm{
		tokens: []token.Token{},
	}

	act := &Farm{
		mature: true,
		tokens: token.Copy(tks),
	}

	out := act.Harvest()
	token.AssertSliceEqual(t, tks, out)
	assertFarm(t, false, false, exp, act)
}

func TestFarm_SaltHarvest_1(t *testing.T) {
	tks := []token.Token{
		// x <- 1
		token.Dummy{`x`, token.TT_ID},
		token.Dummy{`<-`, token.TT_ASSIGN},
		token.Dummy{`1`, token.TT_NUMBER},
	}

	exp := &Farm{}

	act := &Farm{
		tokens: token.Copy(tks),
	}

	out := act.SaltHarvest()
	token.AssertSliceEqual(t, tks, out)
	assertFarm(t, false, false, exp, act)
}
