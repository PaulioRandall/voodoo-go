
package instruction

// Instruction is a unit of activity. Instructions are
// executed within a given context; usually, it will make
// a single change to the specified context.
type Instruction interface {

	// Exe executes the instruction within the given context
	// returning true if the context has changed.
	Exe(Context) (bool, InsError)
}
