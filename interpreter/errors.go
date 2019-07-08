
package interpreter

import (
	"fmt"
	"os"
	"strings"
)

// compilerBug writes a compiler bug to output then exits the program
// with code 1.
func compilerBug(lineNum int, msg string) {
	fmt.Print("[COMPILER BUG]")
	info := fmt.Sprintf("...when parsing line '%d'", lineNum)
	fmt.Println(info)
	
	msgLines := strings.Split(msg, "\n")
	for _, v := range msgLines {
		fmt.Print("\t..." + v)
	}
	
	os.Exit(1)
}
