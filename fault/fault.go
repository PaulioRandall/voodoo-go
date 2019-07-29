package fault

// Fault represents an error produced by this program
// rather than a library. The error could be due to a bug
// or could detail a problem with code being parsed.
type Fault interface {
	error

	// SetLine sets the line index where the error occurred.
	SetLine(i int) Fault

	// SetFrom sets the inclusive column index where the error starts.
	SetFrom(i int) Fault

	// SetTo sets the exclusive column index where the error ends.
	SetTo(i int) Fault

	// Line returns the line index where the error occurred.
	Line() int

	// From returns the column index where the error starts.
	From() int

	// To returns the column index where the error ends.
	To() int

	// Type returns the type of the fault.
	Type() FaultType
}

// stdFault is the standard implementation of the
// Fault interface.
type stdFault struct {
	msg     string
	line    int
	from    int
	to      int // Exclusive
	errType FaultType
}

// Error satisfies the error interface.
func (err stdFault) Error() string {
	return err.msg
}

// SetLine satisfies the Fault interface.
func (err stdFault) SetLine(i int) Fault {
	err.line = i
	return err
}

// SetFrom satisfies the Fault interface.
func (err stdFault) SetFrom(i int) Fault {
	err.from = i
	return err
}

// SetTo satisfies the Fault interface.
func (err stdFault) SetTo(i int) Fault {
	err.to = i
	return err
}

// Line satisfies the Fault interface.
func (err stdFault) Line() int {
	return err.line
}

// From satisfies the Fault interface.
func (err stdFault) From() int {
	return err.from
}

// To satisfies the Fault interface.
func (err stdFault) To() int {
	return err.to
}

// Type satisfies the Fault interface.
func (err stdFault) Type() FaultType {
	return err.errType
}
