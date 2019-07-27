package interpreter

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/lexer/scanner"
	"github.com/PaulioRandall/voodoo-go/lexer/strimmer"
	"github.com/PaulioRandall/voodoo-go/scroll"
	"github.com/PaulioRandall/voodoo-go/symbol"
)

// ExitCode represents a program exit code
type ExitCode int

// Execute runs a Voodoo scroll.
func Execute(sc *scroll.Scroll, scArgs []string) (ExitCode, error) {

	line := sc.Next(nil)
	line = sc.Next(line) // Ignoring first line: shebang

	for line != nil {

		lexemes, err := scanner.Scan(line.Val, line.Num)
		if err != nil {
			return 1, err
		}

		tokens := strimmer.Strim(lexemes)
		symbol.PrintlnTokenValues(tokens)
		symbol.PrintlnTokenTypes(tokens)
		fmt.Println()

		// NEXT: Rename lexeme pkg to symbol
		// NEXT: Syntax Analyser

		line = sc.Next(line)
	}

	return 0, nil
}
