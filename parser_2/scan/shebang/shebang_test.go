package shebang

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser_2/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/scantok"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
	"github.com/stretchr/testify/require"
)

func TestScanShebang_1(t *testing.T) {
	in := "#!/bin/voodoo\n\nx:1"

	exp := scantok.New(
		"#!/bin/voodoo",
		0,
		0,
		13,
		token.TT_SHEBANG,
	)

	r := runer.NewByStr(in)
	act, err := ScanShebang(r)

	require.Nil(t, err)
	scantok.AssertEqual(t, exp, act)
}
