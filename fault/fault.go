package fault

// Fault represents an error produced by this program
// rather than a library. The error could be due to a bug
// or could detail a problem with code being parsed.
type Fault interface {
	error

	// Line returns the line index where the error occurred.
	Line() int

	// From returns the column index where the error starts.
	From() int

	// To returns the column index where the error ends.
	To() int

	// Type returns the type of the fault.
	Type() FaultType
}

// SetLine sets the line index of a fault within the scroll.
func SetLine(f Fault, i int) Fault {
	sf, ok := f.(stdFault)
	if !ok {
		return Bug("Can't deal with this unknown concrete fault type")
	}

	sf.line = i
	return Fault(sf)
}

// SetFrom sets the column index where the error starts.
func SetFrom(f Fault, i int) Fault {
	sf, ok := f.(stdFault)
	if !ok {
		return Bug("Can't deal with this unknown concrete fault type")
	}

	sf.from = i
	return Fault(sf)
}

// SetTo sets the column index where the error ends.
func SetTo(f Fault, i int) Fault {
	sf, ok := f.(stdFault)
	if !ok {
		return Bug("Can't deal with this unknown concrete fault type")
	}

	sf.to = i
	return Fault(sf)
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
