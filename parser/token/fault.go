package token

// SyntaxFault represents a generic fault with syntax.
type SyntaxFault struct {
	Index int      // Index where the error actually occurred
	Msgs  []string // Description of the error
}

// Print satisfies the Fault interface.
func (err SyntaxFault) Print(file string) {
	//sc.PrettyPrintError(-1, err.Index, err.Msgs...)
}
