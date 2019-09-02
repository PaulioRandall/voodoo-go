package exe

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PaulioRandall/voodoo-go/parser/farm"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Execute runs a Voodoo scroll.
func Execute(file string, scArgs []string) int {

	r, e := newRuner(file)
	if e != nil {
		print(e)
		return 1
	}

	if scanAndPrintShebang(r, file) {
		return 1
	}

	for tks, exit := scanNext(r, file); !exit; {
		if tks == nil {
			return 1
		}

		printStatement(tks)
		tks, exit = scanNext(r, file)
	}

	return 0
}

// newRuner creates a new Runer from a file.
func newRuner(file string) (*runer.Runer, error) {
	r, e := os.Open(file)
	if e != nil {
		return nil, e
	}

	bf := bufio.NewReader(r)
	return runer.New(bf), nil
}

// scanAndPrintShebang scans the shebang line and prints it.
func scanAndPrintShebang(r *runer.Runer, file string) bool {
	s, e := scan.ShebangScanner()(r)
	if e != nil {
		perror.PrintError(file, e)
		return true
	}

	fmt.Println(`[` + token.KindName(s.Kind()) + `]`)
	return false
}

// scanNextStat scans the next statement from the scanner passing each token
// through the farm in the process.
func scanNext(r *runer.Runer, file string) (_ []token.Token, last bool) {
	frm := farm.New()

	for f, e := scan.Next(r); f != nil; f, e = scan.Next(r) {
		if e != nil {
			perror.PrintError(file, e)
			return nil, false
		}

		tk, e := f(r)
		if e != nil {
			perror.PrintError(file, e)
			return nil, false
		}

		ready, er := frm.Feed(tk)
		if er != nil {
			print(er)
			return nil, false
		}

		if ready {
			return frm.Harvest(), false
		}
	}

	return frm.FinalHarvest(), true
}

// printStatement prints a statmant.
func printStatement(tks []token.Token) {
	fmt.Print(`[`)

	last := len(tks) - 1
	for i, tk := range tks {
		s := token.KindName(tk.Kind())
		fmt.Print(s)

		if i < last {
			fmt.Print(`, `)
		}
	}

	fmt.Println("]")
}
