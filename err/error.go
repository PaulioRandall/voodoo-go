package err

// VooError represents an error produced by this program
// rather than a library.
type VooError interface {
  error
  
  // Line sets the line index where the error occurred.
  Line(i int) VooError
  
  // From sets the inclusive column index where the error starts.
  From(i int) VooError
  
  // To sets the exclusive column index where the error ends.
  To(i int) VooError
}

// errorType represents the type of VooError
type errorType int

const (
  std errorType = iota + 1
  bug
)

// stdVooError is the standard implementation of the
// VooError interface.
type stdVooError struct {
  msg string
  line int
  from int
  to int    // Exclusive
  errType errorType
}

// Error satisfies the error interface.
func (err stdVooError) Error() string {
  return err.msg
}

// Line satisfies the VooError interface.
func (err stdVooError) Line(i int) VooError {
  err.line = i
  return err
}

// From satisfies the VooError interface.
func (err stdVooError) From(i int) VooError {
  err.from = i
  return err
}

// To satisfies the VooError interface.
func (err stdVooError) To(i int) VooError {
  err.to = i
  return err
}

// Bug returns a new VooError with a message formatted
// to present a bug with this program.
func Bug(m string) VooError {
  return stdVooError{
    msg: m,
    errType: bug,
  }
}
