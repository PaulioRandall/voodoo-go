package exe

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PaulioRandall/voodoo-go/parser/ctx"
	"github.com/PaulioRandall/voodoo-go/parser/expr"
	"github.com/PaulioRandall/voodoo-go/parser/parser"
	"github.com/PaulioRandall/voodoo-go/parser/perror"
	"github.com/PaulioRandall/voodoo-go/parser/scan"
	"github.com/PaulioRandall/voodoo-go/parser/scan/runer"
	"github.com/PaulioRandall/voodoo-go/parser/strim"
	"github.com/PaulioRandall/voodoo-go/parser/token"
)

// Execute runs a Voodoo scroll.
func Execute(file string, scArgs []string) int {

	r, e := newRuner(file)
	if e != nil {
		print(e)
		return 1
	}

	if e := scanAndPrintShebang(r, file); e != nil {
		perror.PrintError(file, e)
		return 1
	}

	c := ctx.New(nil)
	if e := scanExpr(r, c); e != nil {
		perror.PrintError(file, e)
		return 1
	}

	fmt.Print("\n" + c.String())

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
func scanAndPrintShebang(r *runer.Runer, file string) perror.Perror {
	s, e := scan.ShebangScanner()(r)
	if e != nil {
		return e
	}

	fmt.Println(`[` + token.KindName(s.Kind()) + `]`)
	return nil
}

// scanExpr scans, strims, parses, and prints the tokens into expression trees.
func scanExpr(r *runer.Runer, c ctx.Context) perror.Perror {
	p := parser.New()

	for f, e := scan.Next(r); f != nil; f, e = scan.Next(r) {
		if e != nil {
			return e
		}

		tk, e := f(r)
		if e != nil {
			return e
		}

		tk = strim.Strim(tk)
		if tk == nil {
			continue
		}

		ex, e := p.Parse(tk)
		if e != nil {
			return e
		}

		if ex != nil {
			printExprTree(ex)

			if _, e = ex.Exe(c); e != nil {
				return e
			}
		}
	}

	return nil
}

// printExprTree prints an expression tree.
func printExprTree(ex expr.Expr) {
	fmt.Print(`[`)
	fmt.Print(ex.String())
	fmt.Println("]")
}
