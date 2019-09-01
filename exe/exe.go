package exe

import (
	"fmt"

	"github.com/PaulioRandall/voodoo-go/parser/farm"
	"github.com/PaulioRandall/voodoo-go/parser/scan"
	"github.com/PaulioRandall/voodoo-go/parser/scan/err"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Execute runs a Voodoo scroll.
func Execute(sc *Scroll, scArgs []string) int {

	r := runer.NewByStr(sc.Data)
	s := scan.New(true)

	for tks, exit := scanNext(s, r, sc); !exit; {
		if tks == nil {
			return 1
		}

		printStatement(tks)
		tks, exit = scanNext(s, r, sc)
	}

	return 0
}

// scanNextStat scans the next statement from the scanner passing each token
// through the farm in the process.
func scanNext(s *scan.Scanner, r *runer.Runer, sc *Scroll) (_ []token.Token, last bool) {
	frm := farm.New()

	for f, e := s.Next(r); f != nil; f, e = s.Next(r) {
		if e != nil {
			err.PrintScanError(sc.File, e)
			return nil, false
		}

		tk, e := f(r)
		if e != nil {
			err.PrintScanError(sc.File, e)
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

	return frm.SaltHarvest(), true
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
