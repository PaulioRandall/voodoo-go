package interpreter

import (
	"bufio"
	"strings"

	"github.com/PaulioRandall/voodoo-go/lexer/scanner"
	"github.com/PaulioRandall/voodoo-go/lexer/strimmer"
	"github.com/PaulioRandall/voodoo-go/scroll"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Execute runs a Voodoo scroll.
func Execute(sc *scroll.Scroll, scArgs []string) int {

	text := ""
	for i, line := range sc.Lines {
		if i == 0 {
			continue // Ignoring first line: shebang
		}

		text += "\n" + line
	}

	done := make(chan bool)
	scanChan := make(chan token.Token)
	strimChan := make(chan token.Token)

	go token.PrintlnTokenChan(done, strimChan, tokenToString)
	go strimmer.Strim(scanChan, strimChan)

	br := bufio.NewReader(strings.NewReader(text))
	err := scanner.Scan(br, scanChan)
	if err != nil {
		err.Print(sc, -1)
		return 1
	}

	<-done

	return 0
}

// tokenToString is used by token.PrintlnTokenChan() to determine what should
// be printed for each supplied token.
func tokenToString(tk token.Token) string {
	if tk.Type == token.NEWLINE {
		return `\n`
	}

	return tk.Val //token.TokenName(tk.Type)
}
