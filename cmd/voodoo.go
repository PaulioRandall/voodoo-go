//usr/bin/env go run "$0" "$@"; exit "$?"

package main

import (
	"fmt"
	"os"
	"time"
	
	inter "github.com/PaulioRandall/voodoo-go/interpreter"
)

// main is the entry point for this script. It wraps the standard Go format,
// build, test, run, and install operations specifically for this project.
func main() {
	stopWatch := inter.StopWatch{}
	stopWatch.Start()
	fmt.Printf("Started\t%v\n\n", stopWatch.Started.UTC())

	// Don't abstract the build workflows! They are more readable and extendable
	// this way.
	option := getOption()
	switch option {
	case "exe":
		exeScroll()

	default:
		badSyntax()
	}

	stopWatch.Stop()
	fmt.Printf("\nDone\t")
	stopWatch.PrintElapsed(time.Microsecond)

	os.Exit(0)
}

// exeScroll loads then executes the scroll supplied as a parameter.
func exeScroll() {
	scrollPath := getScroll()
	scrollArgs := getScrollArgs()
	
	scroll, err := inter.LoadScroll(scrollPath)
	if err != nil {
		panic(err)
	}
	
	// TODO: Handle when error returned.
	exitCode, err := inter.Execute(scroll, scrollArgs)
	if err != nil {
		panic(err)
	}
	
	// TODO: What to do when non-zero exit code?
	exitMsg := fmt.Sprintf("\nExit %d", exitCode)
	fmt.Println(exitMsg)
}

// getOption returns the argument passed that represents the operation to
// perform.
func getOption() string {
	if len(os.Args) < 2 {
		badSyntax()
	}
	return os.Args[1]
}

// getScroll returns the scroll path argument.
func getScroll() string {
	if len(os.Args) < 3 {
		badSyntax()
	}
	return os.Args[2]
}

// getScrollArgs returns the scrolls arguments.
func getScrollArgs() []string {
	return os.Args[3:]
}

// badSyntax prints the scripts syntax to console then exits the application
// with code 1.
func badSyntax() {
	syntax := `syntax options:
1) ./voodoo.exe run [scroll-name]`

	fmt.Println(syntax + "\n")
	os.Exit(1)
}
