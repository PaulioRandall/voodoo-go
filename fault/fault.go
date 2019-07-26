package fault

// Fault represents an error produced by this program
// rather than a library. The error could be due to a bug
// or could detail a problem with code being parsed.
type Fault interface {
  error
  
  // Line sets the line index where the error occurred.
  Line(i int) Fault
  
  // From sets the inclusive column index where the error starts.
  From(i int) Fault
  
  // To sets the exclusive column index where the error ends.
  To(i int) Fault
  
  // Type returns the type of the fault.
  Type() FaultType
}

// stdFault is the standard implementation of the
// Fault interface.
type stdFault struct {
  msg string
  line int
  from int
  to int    // Exclusive
  errType FaultType
}

// Error satisfies the error interface.
func (err stdFault) Error() string {
  return err.msg
}

// Line satisfies the Fault interface.
func (err stdFault) Line(i int) Fault {
  err.line = i
  return err
}

// From satisfies the Fault interface.
func (err stdFault) From(i int) Fault {
  err.from = i
  return err
}

// To satisfies the Fault interface.
func (err stdFault) To(i int) Fault {
  err.to = i
  return err
}

// Type satisfies the Fault interface.
func (err stdFault) Type() FaultType {
  return err.errType
}
