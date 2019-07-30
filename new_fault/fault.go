package fault_new

// Fault represents an error produced by this program
// rather than a library. The error could be due to a bug
// or could detail a problem with code being parsed.
type Fault interface {
	error

	// Print prints the fault to console with the line number
	// of the scroll where the error originated.
	Print(line int)
}
