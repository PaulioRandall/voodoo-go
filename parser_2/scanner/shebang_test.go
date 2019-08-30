package scanner

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
	"github.com/stretchr/testify/require"
)

func TestScanShebang_1(t *testing.T) {
	in := "#!/bin/voodoo\n\nx <- 1"

	exp := &scanTok{
		text: "#!/bin/voodoo",
		end:  13,
		kind: token.TT_SHEBANG,
	}

	r := runer.NewByStr(in)
	act, err := scanShebang(r)

	require.Nil(t, err)
	AssertScanTokEqual(t, exp, act)
}
