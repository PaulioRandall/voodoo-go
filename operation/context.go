package operation

// Context represents the state of an executing code block
// such as a scroll, spell, loop, or when block.
type Context interface {

	// Counter returns the index of the next line within the
	// scroll to be executed.
	Counter() int

	// Jump sets the counter to the specified value. It represents
	// a jump to a specific line within the executing scroll.
	Jump(i int)
}
