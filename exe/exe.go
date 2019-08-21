package exe

import (
	"bufio"
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/scanner"
	"github.com/PaulioRandall/voodoo-go/parser/strimmer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Execute runs a Voodoo scroll.
func Execute(sc *Scroll, scArgs []string) int {

	done := make(chan *token.Token)
	scanChan := make(chan token.Token)

	go token.PrintlnTokenChan(done, scanChan, tokenToType)
	scan(sc.Data, scanChan)
	errTk := <-done

	if errTk != nil {
		token.PrintErrorToken(sc.File, *errTk)
		return 1
	}

	return 0
}

// scan scans the input string for tokens and places them onto the channel.
func scan(data string, out chan token.Token) {
	defer close(out)

	r := newRuner(data)
	f, ok := parseShebang(r, out)
	if !ok {
		return
	}

	var tk *token.Token
	t := token.TT_UNDEFINED

	for f != nil {
		tk, f, ok = scanToken(r, f, out)
		if !ok {
			return
		}

		t = strimToken(tk, t, out)
	}
}

// strimToken strims the token and places the result on the output channel.
func strimToken(tk *token.Token, prevType token.TokenType, out chan token.Token) token.TokenType {
	tk = strimmer.Strim(*tk, prevType)

	if tk != nil {
		out <- *tk
		return tk.Type
	}

	return token.TT_UNDEFINED
}

// scanToken gets the next token from the runer.
func scanToken(r *scanner.Runer, f scanner.ParseToken, out chan token.Token) (*token.Token, scanner.ParseToken, bool) {
	tk, f, errTk := f(r)
	if errTk != nil {
		out <- *errTk
		return nil, nil, false
	}
	return tk, f, true
}

// parseShebang scans the first line of the scroll returning a SHEBANG token.
func parseShebang(r *scanner.Runer, out chan token.Token) (scanner.ParseToken, bool) {
	_, f, errTk := scanner.ScanShebang(r)
	if errTk != nil {
		out <- *errTk
		return nil, false
	}
	return f, true
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

// tokenToType is used by token.PrintlnTokenChan() to determine what should
// be printed for each supplied token.
func tokenToType(tk token.Token) string {
	n := token.TokenName(tk.Type)

	if tk.Type == token.TT_EOS {
		return n + "\n"
	}

	return n
}
