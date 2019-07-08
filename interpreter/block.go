

package interpreter

// Block represents a block of multiple code lines with their own set
// of variables and context rules. I.e. Scrolls, spells, fors, whens,
// when cases, and list definitions are all current block types. 
type Block interface {

	// ExecuteLines continues execution of the scroll lines until an error,
	// an exit scroll command, or the end of the block is encountered. 
	ExecuteLines(scroll *Scroll)
	
	// Variables returns the variables used by the block.
	Variables() map[string]VoodooValue
	
	// DoesShareVariables returns true if the block uses and updates it's
	// parents/callers variables.
	DoesShareVariables() bool
}