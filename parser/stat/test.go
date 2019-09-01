package stat

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/PaulioRandall/voodoo-go/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqual(t *testing.T, exp, act Statement) bool {
	if exp == nil {
		return assert.Nil(t, act, `Statement == nil`)
	}

	require.NotNil(t, act, `Statement != nil`)

	return utils.LogicalConjunction(
		assert.Equal(t, exp.Kind(), act.Kind(), `Statement.Kind: exp == act`),
		token.AssertSliceEqual(t, exp.Assign(), act.Assign()),
		token.AssertSliceEqual(t, exp.Tokens(), act.Tokens()),
	)
}
