package scan

import (
	"testing"

	//	"github.com/PaulioRandall/voodoo-go/parser/scan/err"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/scantok"
	"github.com/PaulioRandall/voodoo-go/parser/token"

	"github.com/stretchr/testify/require"
)

func doTestScanner_Next(t *testing.T, shebang bool, r *runer.Runer, exp []token.Token) {
	s := New(shebang)

	act := []token.Token{}
	for f, e := s.Next(r); f != nil; f, e = s.Next(r) {
		require.Nil(t, e, `Scanner.Next(): ScanError == nil`)
		require.NotNil(t, f, `Scanner.Next(): TokenScanner != nil`)

		tk, e := f(r)
		require.Nil(t, e, `TokenScanner? ScanError == nil`)
		require.NotNil(t, f, `TokenScanner? Token != nil`)

		act = append(act, tk)
	}

	scantok.AssertSliceEqual(t, exp, act)
}

func TestScanner_Next_1(t *testing.T) {
	r := runer.NewByStr("#!/bin/bash\n")

	exp := []token.Token{
		scantok.New(`#!/bin/bash`, 0, 0, 11, token.TT_SHEBANG),
		scantok.New("\n", 0, 11, 12, token.TT_NEWLINE),
	}

	doTestScanner_Next(t, true, r, exp)
}

func TestScanner_Next_2(t *testing.T) {
	r := runer.NewByStr("x <- 1")

	exp := []token.Token{
		scantok.New(`x`, 0, 0, 1, token.TT_ID),
		scantok.New(` `, 0, 1, 2, token.TT_SPACE),
		scantok.New(`<-`, 0, 2, 4, token.TT_ASSIGN),
		scantok.New(` `, 0, 4, 5, token.TT_SPACE),
		scantok.New(`1`, 0, 5, 6, token.TT_NUMBER),
	}

	doTestScanner_Next(t, false, r, exp)
}

func TestScanner_Next_3(t *testing.T) {
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

	doTestScanner_Next(t, false, r, exp)
}
