package scan

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/scantok"
	"github.com/PaulioRandall/voodoo-go/parser/token"

	"github.com/stretchr/testify/require"
)

func doTestNext(t *testing.T, shebang bool, r *runer.Runer, exp []token.Token) {

	act := []token.Token{}
	if shebang {
		f := ShebangScanner()
		act = doTestScanner(t, r, f, nil, act)
	}

	for f, e := Next(r); f != nil; f, e = Next(r) {
		act = doTestScanner(t, r, f, e, act)
	}

	scantok.AssertSliceEqual(t, exp, act)
}

func doTestScanner(t *testing.T, r *runer.Runer, f TokenScanner, e perror.Perror, act []token.Token) []token.Token {
	require.Nil(t, e, `scan.Next(): Perror == nil`)
	require.NotNil(t, f, `scan.Next(): TokenScanner != nil`)

	tk, e := f(r)
	require.Nil(t, e, `TokenScanner? Perror == nil`)
	require.NotNil(t, f, `TokenScanner? Token != nil`)

	return append(act, tk)
}

func TestNext_1(t *testing.T) {
	r := runer.NewByStr("#!/bin/bash\n")

	exp := []token.Token{
		scantok.New(`#!/bin/bash`, 0, 0, 11, token.TT_SHEBANG),
		scantok.New("\n", 0, 11, 12, token.TT_NEWLINE),
	}

	doTestNext(t, true, r, exp)
}

func TestNext_2(t *testing.T) {
	r := runer.NewByStr("x <- 1")

	exp := []token.Token{
		scantok.New(`x`, 0, 0, 1, token.TT_ID),
		scantok.New(` `, 0, 1, 2, token.TT_SPACE),
		scantok.New(`<-`, 0, 2, 4, token.TT_ASSIGN),
		scantok.New(` `, 0, 4, 5, token.TT_SPACE),
		scantok.New(`1`, 0, 5, 6, token.TT_NUMBER),
	}

	doTestNext(t, false, r, exp)
}

func TestNext_3(t *testing.T) {
	r := runer.NewByStr("x <- 1\ny := 2\r\nz <- 32")

	exp := []token.Token{
		// x <- 1
		scantok.New(`x`, 0, 0, 1, token.TT_ID),
		scantok.New(` `, 0, 1, 2, token.TT_SPACE),
		scantok.New(`<-`, 0, 2, 4, token.TT_ASSIGN),
		scantok.New(` `, 0, 4, 5, token.TT_SPACE),
		scantok.New(`1`, 0, 5, 6, token.TT_NUMBER),
		scantok.New("\n", 0, 6, 7, token.TT_NEWLINE),
		// y := 2
		scantok.New(`y`, 1, 0, 1, token.TT_ID),
		scantok.New(` `, 1, 1, 2, token.TT_SPACE),
		scantok.New(`:=`, 1, 2, 4, token.TT_ASSIGN),
		scantok.New(` `, 1, 4, 5, token.TT_SPACE),
		scantok.New(`2`, 1, 5, 6, token.TT_NUMBER),
		scantok.New("\r\n", 1, 6, 8, token.TT_NEWLINE),
		// z <- 3
		scantok.New(`z`, 2, 0, 1, token.TT_ID),
		scantok.New(` `, 2, 1, 2, token.TT_SPACE),
		scantok.New(`<-`, 2, 2, 4, token.TT_ASSIGN),
		scantok.New(` `, 2, 4, 5, token.TT_SPACE),
		scantok.New(`32`, 2, 5, 7, token.TT_NUMBER),
	}

	doTestNext(t, false, r, exp)
}
