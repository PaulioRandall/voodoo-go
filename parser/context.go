package parser

// Context represents the working environment in which
// expressions can be evaluated. It contains the
// identifiers and their values available to an
// expression and provide means to add or modify them.
type Context struct {
	IDs map[string]Value // Map of identifiers available to expressions
}

// Value represents a value within a scroll, whether
// explicit referenced or implicit.
type Value struct {
}
