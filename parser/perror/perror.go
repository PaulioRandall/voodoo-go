package perror

// Perror represents an error found while parsing.
type Perror interface {

	// Line returns the offending line index within the scroill.
	Line() int

	// Cols returns the offending column indices within the scroll.
	Cols() []int

	// Errors returns an array of error messages.
	Errors() []string
}

// perror is an implementation of Perror.
type perror struct {
	l int
	c []int
	e []string
}

// New returns a new initialised Perror.
func New(l int, c []int, e []string) Perror {
	return perror{
		l: l,
		c: c,
		e: e,
	}
}

// Line satisfies the Perror interface.
func (e perror) Line() int {
	return e.l
}

// Cols satisfies the Perror interface.
func (e perror) Cols() []int {
	return e.c
}

// Errors satisfies the Perror interface.
func (e perror) Errors() []string {
	return e.e
}

// NewByError creates a new Perror from an error.
func NewByError(l, c int, e error) Perror {
	return perror{
		l: l,
		c: []int{c},
		e: []string{e.Error()},
	}
}
