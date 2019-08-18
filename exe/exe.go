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

	done := make(chan bool)
	scanChan := make(chan token.Token)
	strimChan := make(chan token.Token)

	go token.PrintlnTokenChan(done, strimChan, tokenToType)
	go strimmer.Strim(scanChan, strimChan)

	r := newRuner(sc.Data)
	scanner.Scan(r, true, scanChan)

	<-done

	return 0
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
