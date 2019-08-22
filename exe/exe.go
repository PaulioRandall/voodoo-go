package exe

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/scanner"
	"github.com/PaulioRandall/voodoo-go/parser/strimmer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Execute runs a Voodoo scroll.
func Execute(sc *Scroll, scArgs []string) int {

	errTk := scan(sc.Data)
	if errTk != nil {
		token.PrintErrorToken(sc.File, *errTk)
		return 1
	}

	return 0
}

// scan scans the input string for tokens and places them onto the channel.
func scan(data string) *token.Token {

	r := newRuner(data)
	f, errTk := parseShebang(r)
	if errTk != nil {
		return errTk
	}

	var tk *token.Token
	stats := [][]token.Token{}
	stat := []token.Token{}
	var eos bool
	t := token.TT_UNDEFINED

	for f != nil {
		var ok bool
		tk, f, ok = scanToken(r, f)
		if !ok {
			return tk
		}

		tk, t = strimToken(tk, t)
		if tk == nil {
			continue
		}

		stat, eos = appendToken(stat, tk)

		if eos {
			stats = append(stats, stat)
			stat = []token.Token{}
		}
	}

	printStatements(stats)
	return nil
}

// strimToken strims the token and places the result on the output channel.
func strimToken(tk *token.Token, prevType token.TokenType) (*token.Token, token.TokenType) {
	t := tk.Type
	tk = strimmer.Strim(*tk, prevType)
	if tk != nil {
		t = tk.Type
	}
	return tk, t
}

// scanToken gets the next token from the runer.
func scanToken(r *scanner.Runer, f scanner.ParseToken) (*token.Token, scanner.ParseToken, bool) {
	tk, f, errTk := f(r)
	if errTk != nil {
		return errTk, nil, false
	}
	return tk, f, true
}

// parseShebang scans the first line of the scroll returning a SHEBANG token.
func parseShebang(r *scanner.Runer) (scanner.ParseToken, *token.Token) {
	_, f, errTk := scanner.ScanShebang(r)
	if errTk != nil {
		return nil, errTk
	}
	return f, nil
}

// newRuner makes a new Runer instance.
func newRuner(text string) *scanner.Runer {
	sr := strings.NewReader(text)
	br := bufio.NewReader(sr)
	return scanner.NewRuner(br)
}

// tokenToVal is used by token.PrintlnTokenChan() to determine what should
// be printed for each supplied token.
func tokenToVal(tk token.Token) string {
	if tk.Type == token.TT_EOS {
		return `\n`
	}

	return tk.Val //token.TokenName(tk.Type)
}

// appendToken appends the token to the token array if it forms part of the next
// statement and returns true only if the token array now represents a full
// statement.
func appendToken(a []token.Token, tk *token.Token) ([]token.Token, bool) {
	if tk.Type == token.TT_EOS {
		return a, true
	}
	a = append(a, *tk)
	return a, false
}

// printStatements prints each statmant.
func printStatements(stats [][]token.Token) {
	fmt.Print(`[`)

	for _, stat := range stats {
		if len(stat) < 1 {
			continue
		}

		fmt.Print("\n  ")
		size := len(stat) - 1

		for i, tk := range stat {
			s := token.TokenName(tk.Type)
			fmt.Print(s)

			if i < size {
				fmt.Print(`, `)
			}
		}
	}

	fmt.Println("\n]")
}
