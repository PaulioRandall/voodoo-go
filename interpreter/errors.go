
package interpreter

import (
	"fmt"
	"os"
)

// compilerBug writes a compiler bug to output then exits the program
// with code 1.
func compilerBug(scroll *Scroll, msg string) {
	fmt.Println("[COMPILER BUG]")
	info := fmt.Sprintf("\t...when parsing line '%d' of '%s'", scroll.Number, scroll.File)
	fmt.Println(info)
	fmt.Print("\t..." + msg)
	os.Exit(1)
}
