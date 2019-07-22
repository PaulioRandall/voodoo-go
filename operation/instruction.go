package operation

// Instruction is a unit of activity within this
// implementation. Instructions are executed within a given
// context. Usually, it will make a single change to the
// specified context. Instructions are considered atomic
// like CPU instructions but not as fine grained since
// activities such as fetching and placing values from and
// to memory, or rather their entry in the context, are
// considered part of the instruction.
type Instruction interface {

	// Exe executes the instruction within the given context
	// returning true if the context has changed.
	Exe(Context) (bool, InsError)
}
