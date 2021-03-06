package scan

import (
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
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

	token.AssertSliceEqual(t, exp, act)
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
		token.New(`#!/bin/bash`, 0, 0, 11, token.TK_SHEBANG),
	}

	doTestNext(t, true, r, exp)
}

func TestNext_2(t *testing.T) {
	r := runer.NewByStr("x <- 1")

	exp := []token.Token{
		token.New(`x`, 0, 0, 1, token.TK_ID),
		token.New(` `, 0, 1, 2, token.TK_SPACE),
		token.New(`<-`, 0, 2, 4, token.TK_ASSIGN),
		token.New(` `, 0, 4, 5, token.TK_SPACE),
		token.New(`1`, 0, 5, 6, token.TK_NUMBER),
	}

	doTestNext(t, false, r, exp)
}

func TestNext_3(t *testing.T) {
	r := runer.NewByStr("x <- 1\ny := 2\r\nz <- 32")

	exp := []token.Token{
		// x <- 1
		token.New(`x`, 0, 0, 1, token.TK_ID),
		token.New(` `, 0, 1, 2, token.TK_SPACE),
		token.New(`<-`, 0, 2, 4, token.TK_ASSIGN),
		token.New(` `, 0, 4, 5, token.TK_SPACE),
		token.New(`1`, 0, 5, 6, token.TK_NUMBER),
		token.New("\n", 0, 6, 7, token.TK_NEWLINE),
		// y := 2
		token.New(`y`, 1, 0, 1, token.TK_ID),
		token.New(` `, 1, 1, 2, token.TK_SPACE),
		token.New(`:=`, 1, 2, 4, token.TK_ASSIGN),
		token.New(` `, 1, 4, 5, token.TK_SPACE),
		token.New(`2`, 1, 5, 6, token.TK_NUMBER),
		token.New("\r\n", 1, 6, 8, token.TK_NEWLINE),
		// z <- 3
		token.New(`z`, 2, 0, 1, token.TK_ID),
		token.New(` `, 2, 1, 2, token.TK_SPACE),
		token.New(`<-`, 2, 2, 4, token.TK_ASSIGN),
		token.New(` `, 2, 4, 5, token.TK_SPACE),
		token.New(`32`, 2, 5, 7, token.TK_NUMBER),
	}

	doTestNext(t, false, r, exp)
}

func TestNext_4(t *testing.T) {
	r := runer.NewByStr("x <- _")

	exp := []token.Token{
		token.New(`x`, 0, 0, 1, token.TK_ID),
		token.New(` `, 0, 1, 2, token.TK_SPACE),
		token.New(`<-`, 0, 2, 4, token.TK_ASSIGN),
		token.New(` `, 0, 4, 5, token.TK_SPACE),
		token.New(`_`, 0, 5, 6, token.TK_VOID),
	}

	doTestNext(t, false, r, exp)
}

func TestNext_5(t *testing.T) {
	r := runer.NewByStr("x,y<-1,_")

	exp := []token.Token{
		token.New(`x`, 0, 0, 1, token.TK_ID),
		token.New(`,`, 0, 1, 2, token.TK_DELIM),
		token.New(`y`, 0, 2, 3, token.TK_ID),
		token.New(`<-`, 0, 3, 5, token.TK_ASSIGN),
		token.New(`1`, 0, 5, 6, token.TK_NUMBER),
		token.New(`,`, 0, 6, 7, token.TK_DELIM),
		token.New(`_`, 0, 7, 8, token.TK_VOID),
	}

	doTestNext(t, false, r, exp)
}
