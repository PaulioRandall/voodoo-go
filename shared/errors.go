package shared

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// ExitCode represents a program exit code
type ExitCode int

const (
	OK ExitCode = iota
	CatchAllErr
)

// ExeError represents an error when executing a scroll.
type ExeError interface {
	error

	// Code returns the exit code.
	Code() ExitCode

	// Err returns the Go error.
	Err() error
}

// simpleExeError is an implementation of ExeError.
type simpleExeError struct {
	code ExitCode
	err  error
}

// Code satisfies the ExeError interface.
func (see simpleExeError) Code() ExitCode {
	return see.code
}

// Err satisfies the ExeError interface.
func (see simpleExeError) Err() error {
	return see.err
}

// Error satisfies the error interface.
func (see simpleExeError) Error() string {
	return see.err.Error()
}

// WrapError returns a error wrapped as a new ExeError.
func WrapError(code ExitCode, err error) ExeError {
	// TODO: do I need to check exit code is valid?
	return simpleExeError{
		code: code,
		err:  err,
	}
}

// NewError returns a new exeError.
func NewError(code ExitCode, msg string) ExeError {
	// TODO: do I need to check exit code is valid?
	return simpleExeError{
		code: code,
		err:  errors.New(msg),
	}
}

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

// SyntaxError writes a syntax error to output then exits the program
// with code 1.
func SyntaxError(lineNum int, start, end int, err error) {
	SyntaxErr(lineNum, start, end, err.Error())
}

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
