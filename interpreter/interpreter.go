package interpreter

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/lexer/scanner"
	"github.com/PaulioRandall/voodoo-go/lexer/strimmer"
	"github.com/PaulioRandall/voodoo-go/scroll"
	"github.com/PaulioRandall/voodoo-go/token"
)

// Execute runs a Voodoo scroll.
func Execute(sc *scroll.Scroll, scArgs []string) int {

	for i, line := range sc.Lines {
		if i == 0 {
			continue // Ignoring first line: shebang
		}

		scanChan := make(chan token.Token)
		colChan := make(chan []token.Token)
		go collateLine(scanChan, colChan)

		err := scanner.Scan(line, scanChan)
		if err != nil {
			err.Print(sc, i+1)
			return 1
		}

		tks := <-colChan

		strimChan := make(chan token.Token)
		colChan = make(chan []token.Token)
		go collateLine(strimChan, colChan)

		strimmer.Strim(tks, strimChan)
		tks = <-colChan

		token.PrintlnTokenValues(tks)
		token.PrintlnTokenTypes(tks)
		fmt.Println()
	}

	return 0
}

// collateLine collates a single line from a channel of tokens.
func collateLine(in chan token.Token, out chan []token.Token) {
	defer close(out)
	tks := []token.Token{}
	for tk := range in {
		tks = append(tks, tk)
	}
	out <- tks
}
