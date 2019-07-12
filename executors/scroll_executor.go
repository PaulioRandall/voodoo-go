package executors

import (
	cl "github.com/PaulioRandall/voodoo-go/lexer"
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

type TokenType string

const (
	Keyword  TokenType = "keyword"
	VarNames TokenType = "variable-names"
	Void     TokenType = "void"
)

// ScrollExecutor represents the scroll itself as an activity.
type ScrollExecutor struct {
	vars map[string]sc.VooValue
}

// NewScrollExecutor returns a new scroll executor.
func NewScrollExecutor() *ScrollExecutor {
	return &ScrollExecutor{
		vars: make(map[string]sc.VooValue),
	}
}

// ExeLine satisfies the Executor interface.
func (sa *ScrollExecutor) Exe(scroll *sc.Scroll, line sc.Line) (sh.ExitCode, Executor, sh.ExeError) {
	exitCode := sh.OK
	next := Executor(sa)
	var err sh.ExeError = nil

	// TODO
	frags := cl.Cleave(line.Val, line.Num)
	cl.PrintlnSymbols(frags)

	return exitCode, next, err
}
