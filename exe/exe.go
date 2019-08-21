package exe

import (
	"bufio"
	"strings"

	scanner "github.com/PaulioRandall/voodoo-go/parser/scanner_new"
	"github.com/PaulioRandall/voodoo-go/parser/strimmer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Execute runs a Voodoo scroll.
func Execute(sc *Scroll, scArgs []string) int {

	done := make(chan bool)
	scanChan := make(chan token.Token)
	strimChan := make(chan token.Token)

	go token.PrintlnTokenChan(done, strimChan, tokenToType)
	go strimmer.Strim(scanChan, strimChan)

	scan(sc.Data, scanChan)

	<-done

	return 0
}

// scan scans the input string for tokens and places them onto the channel.
func scan(data string, out chan token.Token) {
	defer close(out)

	r := newRuner(data)
	f := scanToken(r, scanner.ScanShebang, out)

	for f != nil {
		f = scanToken(r, f, out)
	}
}

// scanToken scans the next token handling any errors and returning the next
// token parsing function. If nil is returned then an error occurred or no more
// tokens are left to parse.
func scanToken(r *scanner.Runer, f scanner.ParseToken, out chan token.Token) scanner.ParseToken {
	tk, f, errTk := f(r)
	if errTk != nil {
		out <- *errTk
		return nil
	}

	if tk != nil {
		out <- *tk
	}

	return f
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
