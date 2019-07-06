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
	stopWatch := StopWatch{}
	stopWatch.Start()
	fmt.Printf("Started\t%v\n\n", stopWatch.Started.UTC())

	// Don't abstract the build workflows! They are more readable and extendable
	// this way.
	option := getOption()
	switch option {
	case "run":
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

/******************************************************************************
	github.com/PaulioRandall/go-cookies/cookies
******************************************************************************/

// StopWatch represents a process timer with a few common operations.
type StopWatch struct {
	Started time.Time
	Stopped time.Time
}

// Start starts the stop watch overwritting any previously start time.
func (sw *StopWatch) Start() {
	sw.Started = time.Now().UTC()
}

// Stop stops the stop watch overwriting any previous stop time.
func (sw *StopWatch) Stop() {
	sw.Stopped = time.Now().UTC()
}

// Lap restarts the stop watch but returns a copy of the stop watch with the
// previous laps times.
func (sw *StopWatch) Lap() StopWatch {
	sw.Stop()
	r := StopWatch{
		Started: sw.Started,
		Stopped: sw.Stopped,
	}
	sw.Started = sw.Stopped
	return r
}

// Elapsed returns the elapsed time between the started and stopped times.
func (sw *StopWatch) Elapsed() time.Duration {
	return sw.Stopped.Sub(sw.Started)
}

// PrintElapsed prints the elapsed time as returned by Elapsed(). 'radix' may be
// passed to print the result in a more appropriate manner.
//
// e.g. PrintElapsed(time.Millisecond) will print as milliseconds.
//
// e.g. PrintElapsed(10 * 1000 * 1000) will print using a custom radix,
// microseconds * 10 in this case.
func (sw *StopWatch) PrintElapsed(radix time.Duration) {
	t := sw.Elapsed()
	f := float64(t) / float64(radix)

	switch radix {
	case time.Nanosecond:
		fmt.Printf("%.3f ns\n", f)
	case time.Microsecond:
		fmt.Printf("%.3f us\n", f)
	case time.Millisecond:
		fmt.Printf("%.3f ms\n", f)
	case time.Second:
		fmt.Printf("%.3f s\n", f)
	case time.Minute:
		fmt.Printf("%.3f m\n", f)
	case time.Hour:
		fmt.Printf("%.3f hr\n", f)
	default:
		fmt.Printf("%.3f\n", f)
	}
}