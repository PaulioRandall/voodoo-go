package perror

// Perror represents an error found while parsing.
type Perror interface {

	// Line returns the offending line within the scroill.
	Line() int

	// Index returns the offending index within the scroll.
	Index() int

	// Errors returns an array of errors messages.
	Errors() []string
}

// perror is an implementation of Perror.
type perror struct {
	l int
	i int
	e []string
}

// New returns a new initialised Perror.
func New(l, i int, e []string) Perror {
	return perror{
		l: l,
		i: i,
		e: e,
	}
}

// Line satisfies the Perror interface.
func (e perror) Line() int {
	return e.l
}

// Index satisfies the Perror interface.
func (e perror) Index() int {
	return e.i
}

// Errors satisfies the Perror interface.
func (e perror) Errors() []string {
	return e.e
}

// NewByError creates a new Perror from an error.
func NewByError(l, i int, e error) Perror {
	return perror{
		l: l,
		i: i,
		e: []string{e.Error()},
	}
}
