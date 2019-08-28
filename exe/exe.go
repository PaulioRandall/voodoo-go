package exe

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/evaluator"
	"github.com/PaulioRandall/voodoo-go/parser/scanner"
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

	strimmer := evaluator.NewStrimmer()
	stat := evaluator.NewStatement()
	stats := [][]token.Token{}
	var tk *token.Token
	var ok bool

	for f != nil {

		tk, f, ok = scanToken(r, f)
		if !ok {
			return tk
		}

		strimmer.Strim(stat, tk)

		if stat.IsComplete() {
			stats = append(stats, stat.Tokens())
			stat.Reset()
		}
	}

	if !stat.IsEmpty() {
		stat.SetComplete()
		stats = append(stats, stat.Tokens())
	}

	printStatements(stats)
	return nil
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
		last := len(stat) - 1
		if last < 0 {
			continue
		}

		fmt.Print("\n  ")
		for i, tk := range stat {
			s := token.TokenName(tk.Type)
			fmt.Print(s)

			if i < last {
				fmt.Print(`, `)
			}
		}
	}

	fmt.Println("\n]")
}
