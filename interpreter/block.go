
package interpreter

// BlockExe represents a block of multiple code lines usually with their
// own set of variables and context rules. I.e. Scrolls, spells, fors, whens,
// when cases, and list definitions are all current block types. The
// block is used to execute code for that block.
type BlockExe interface {

	// ExeLines continues execution of the scroll lines until an error,
	// an exit scroll command, or the end of the block is encountered. 
	ExeLines(scroll *Scroll)
	
	// Vars returns the variables used by the block.
	Vars() map[string]VooValue
	
	// DoesShareVars returns true if the block uses and updates it's
	// parents/callers variables rather than having it's own set.
	DoesShareVars() bool
}