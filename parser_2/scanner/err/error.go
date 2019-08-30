package err

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
)

// ScanError represents an error found while scanning text.
type ScanError interface {

	// Line returns the line of the error.
	Line() int

	// Index returns the index of the error.
	Index() int

	// Errors returns an array of errors messages.
	Errors() []string
}

// scanErr is an implementation of ScanError.
type scanErr struct {
	l int
	i int
	e []string
}

// New returns a new initialised ScanError.
func New(l, i int, e []string) ScanError {
	return scanErr{
		l: l,
		i: i,
		e: e,
	}
}

// Line satisfies the ScanError interface.
func (e scanErr) Line() int {
	return e.l
}

// Index satisfies the ScanError interface.
func (e scanErr) Index() int {
	return e.i
}

// Errors satisfies the ScanError interface.
func (e scanErr) Errors() []string {
	return e.e
}

// NewByRuner creates a new ScanError from an error returned by a Runer.
func NewByRuner(r *runer.Runer, e error) ScanError {
	return &scanErr{
		l: r.Line(),
		i: r.NextCol(),
		e: []string{e.Error()},
	}
}
