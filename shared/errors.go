package shared

import (
	"fmt"
	"os"
	"strings"
)

// DEPRECATED
//
// CompilerBug writes a compiler bug to output then exits the program
// with code 1.
func CompilerBug(lineNum int, msg string) {
	fmt.Print("[COMPILER BUG]")
	info := fmt.Sprintf("...when parsing line '%d'", lineNum)
	fmt.Println(info)

	msgLines := strings.Split(msg, "\n")
	for _, v := range msgLines {
		fmt.Println("\t..." + v)
	}

	os.Exit(1)
}

// DEPRECATED
//
// SyntaxError writes a syntax error to output then exits the program
// with code 1.
func SyntaxError(lineNum int, start, end int, err error) {
	SyntaxErr(lineNum, start, end, err.Error())
}

// DEPRECATED
//
// SyntaxErr writes a syntax error to output then exits the program
// with code 1.
func SyntaxErr(lineNum int, start, end int, msg string) {
	fmt.Print("[SYNTAX BUG]")
	info := fmt.Sprintf("...at line %d, columns %d -> %d ", lineNum, start, end)
	fmt.Println(info)

	msgLines := strings.Split(msg, "\n")
	for _, v := range msgLines {
		fmt.Println("\t..." + v)
	}

	os.Exit(1)
}
