package scanner_new

import (
	"bufio"
	"strings"
	"testing"

	"github.com/PaulioRandall/voodoo-go/parser/token"
	"github.com/stretchr/testify/assert"
)

func doTestScanner(t *testing.T, in string, exp []token.Token) {
	r := newRuner(in)
	var act []token.Token

	f, errTk := ScanFirst(r)
	if errTk != nil {
		act = []token.Token{*errTk}
	} else {
		act = scanTestTokens(r, f)
	}

	assertTokens(t, exp, act)
}

func newRuner(s string) *Runer {
	sr := strings.NewReader(s)
	br := bufio.NewReader(sr)
	return NewRuner(br)
}

func scanTestTokens(r *Runer, f ParseToken) []token.Token {
	var tk, errTk *token.Token
	out := []token.Token{}

	for f != nil {
		tk, f, errTk = f(r)
		if errTk != nil {
			out = append(out, *errTk)
			break
		}

		out = append(out, *tk)
	}

	return out
}

func assertTokens(t *testing.T, exp []token.Token, act []token.Token) {
	for i, expTk := range exp {
		if !assert.True(t, i < len(act)) {
			break
		}

		actTk := act[i]
		assertToken(t, &expTk, &actTk)
	}

	assert.Equal(t, len(exp), len(act))
}

func TestScanner_1(t *testing.T) {
	in := `x # 1`
	exp := []token.Token{
		dummyToken(0, 0, 1, `x`, token.TT_ID),
		dummyToken(0, 1, 2, ` `, token.TT_SPACE),
		errDummyToken(0, 0, 2),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_2(t *testing.T) {
	in := `123.456.789`
	exp := []token.Token{
		dummyToken(0, 0, 7, `123.456`, token.TT_NUMBER),
		errDummyToken(0, 0, 7),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_3(t *testing.T) {
	in := `x <- 1`
	exp := []token.Token{
		dummyToken(0, 0, 1, `x`, token.TT_ID),
		dummyToken(0, 1, 2, ` `, token.TT_SPACE),
		dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		dummyToken(0, 4, 5, ` `, token.TT_SPACE),
		dummyToken(0, 5, 6, `1`, token.TT_NUMBER),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_4(t *testing.T) {
	in := `y <- -1.1`
	exp := []token.Token{
		dummyToken(0, 0, 1, `y`, token.TT_ID),
		dummyToken(0, 1, 2, ` `, token.TT_SPACE),
		dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		dummyToken(0, 4, 5, ` `, token.TT_SPACE),
		dummyToken(0, 5, 6, `-`, token.TT_SUBTRACT),
		dummyToken(0, 6, 9, `1.1`, token.TT_NUMBER),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_5(t *testing.T) {
	in := `x <- true`
	exp := []token.Token{
		dummyToken(0, 0, 1, `x`, token.TT_ID),
		dummyToken(0, 1, 2, ` `, token.TT_SPACE),
		dummyToken(0, 2, 4, `<-`, token.TT_ASSIGN),
		dummyToken(0, 4, 5, ` `, token.TT_SPACE),
		dummyToken(0, 5, 9, `true`, token.TT_TRUE),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_6(t *testing.T) {
	in := `@Println("Whelp")`
	exp := []token.Token{
		dummyToken(0, 0, 8, `@Println`, token.TT_SPELL),
		dummyToken(0, 8, 9, `(`, token.TT_CURVY_OPEN),
		dummyToken(0, 9, 16, `"Whelp"`, token.TT_STRING),
		dummyToken(0, 16, 17, `)`, token.TT_CURVY_CLOSE),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_7(t *testing.T) {
	in := "\tresult <- func(a, b) r, err     "
	exp := []token.Token{
		dummyToken(0, 0, 1, "\t", token.TT_SPACE),
		dummyToken(0, 1, 7, `result`, token.TT_ID),
		dummyToken(0, 7, 8, ` `, token.TT_SPACE),
		dummyToken(0, 8, 10, `<-`, token.TT_ASSIGN),
		dummyToken(0, 10, 11, ` `, token.TT_SPACE),
		dummyToken(0, 11, 15, `func`, token.TT_FUNC),
		dummyToken(0, 15, 16, `(`, token.TT_CURVY_OPEN),
		dummyToken(0, 16, 17, `a`, token.TT_ID),
		dummyToken(0, 17, 18, `,`, token.TT_VALUE_DELIM),
		dummyToken(0, 18, 19, ` `, token.TT_SPACE),
		dummyToken(0, 19, 20, `b`, token.TT_ID),
		dummyToken(0, 20, 21, `)`, token.TT_CURVY_CLOSE),
		dummyToken(0, 21, 22, ` `, token.TT_SPACE),
		dummyToken(0, 22, 23, `r`, token.TT_ID),
		dummyToken(0, 23, 24, `,`, token.TT_VALUE_DELIM),
		dummyToken(0, 24, 25, ` `, token.TT_SPACE),
		dummyToken(0, 25, 28, `err`, token.TT_ID),
		dummyToken(0, 28, 33, `     `, token.TT_SPACE),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_8(t *testing.T) {
	in := `alphabet <- ["a", "b", "c"]`
	exp := []token.Token{
		dummyToken(0, 0, 8, `alphabet`, token.TT_ID),
		dummyToken(0, 8, 9, ` `, token.TT_SPACE),
		dummyToken(0, 9, 11, `<-`, token.TT_ASSIGN),
		dummyToken(0, 11, 12, ` `, token.TT_SPACE),
		dummyToken(0, 12, 13, `[`, token.TT_SQUARE_OPEN),
		dummyToken(0, 13, 16, `"a"`, token.TT_STRING),
		dummyToken(0, 16, 17, `,`, token.TT_VALUE_DELIM),
		dummyToken(0, 17, 18, ` `, token.TT_SPACE),
		dummyToken(0, 18, 21, `"b"`, token.TT_STRING),
		dummyToken(0, 21, 22, `,`, token.TT_VALUE_DELIM),
		dummyToken(0, 22, 23, ` `, token.TT_SPACE),
		dummyToken(0, 23, 26, `"c"`, token.TT_STRING),
		dummyToken(0, 26, 27, `]`, token.TT_SQUARE_CLOSE),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_9(t *testing.T) {
	in := `x<-2 // A comment`
	exp := []token.Token{
		dummyToken(0, 0, 1, `x`, token.TT_ID),
		dummyToken(0, 1, 3, `<-`, token.TT_ASSIGN),
		dummyToken(0, 3, 4, `2`, token.TT_NUMBER),
		dummyToken(0, 4, 5, ` `, token.TT_SPACE),
		dummyToken(0, 5, 17, `// A comment`, token.TT_COMMENT),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_10(t *testing.T) {
	in := `isLandscape <- length<height`
	exp := []token.Token{
		dummyToken(0, 0, 11, `isLandscape`, token.TT_ID),
		dummyToken(0, 11, 12, ` `, token.TT_SPACE),
		dummyToken(0, 12, 14, `<-`, token.TT_ASSIGN),
		dummyToken(0, 14, 15, ` `, token.TT_SPACE),
		dummyToken(0, 15, 21, `length`, token.TT_ID),
		dummyToken(0, 21, 22, `<`, token.TT_CMP_LT),
		dummyToken(0, 22, 28, `height`, token.TT_ID),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_11(t *testing.T) {
	in := `x<-3.14*(1-2+3)`
	exp := []token.Token{
		dummyToken(0, 0, 1, `x`, token.TT_ID),
		dummyToken(0, 1, 3, `<-`, token.TT_ASSIGN),
		dummyToken(0, 3, 7, `3.14`, token.TT_NUMBER),
		dummyToken(0, 7, 8, `*`, token.TT_MULTIPLY),
		dummyToken(0, 8, 9, `(`, token.TT_CURVY_OPEN),
		dummyToken(0, 9, 10, `1`, token.TT_NUMBER),
		dummyToken(0, 10, 11, `-`, token.TT_SUBTRACT),
		dummyToken(0, 11, 12, `2`, token.TT_NUMBER),
		dummyToken(0, 12, 13, `+`, token.TT_ADD),
		dummyToken(0, 13, 14, `3`, token.TT_NUMBER),
		dummyToken(0, 14, 15, `)`, token.TT_CURVY_CLOSE),
	}
	doTestScanner(t, in, exp)
}

func TestScanner_12(t *testing.T) {
	in := `!x => y <- 0`
	exp := []token.Token{
		dummyToken(0, 0, 1, `!`, token.TT_NOT),
		dummyToken(0, 1, 2, `x`, token.TT_ID),
		dummyToken(0, 2, 3, ` `, token.TT_SPACE),
		dummyToken(0, 3, 5, `=>`, token.TT_MATCH),
		dummyToken(0, 5, 6, ` `, token.TT_SPACE),
		dummyToken(0, 6, 7, `y`, token.TT_ID),
		dummyToken(0, 7, 8, ` `, token.TT_SPACE),
		dummyToken(0, 8, 10, `<-`, token.TT_ASSIGN),
		dummyToken(0, 10, 11, ` `, token.TT_SPACE),
		dummyToken(0, 11, 12, `0`, token.TT_NUMBER),
	}
	doTestScanner(t, in, exp)
}
