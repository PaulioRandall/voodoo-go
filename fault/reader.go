package fault

import (
	"fmt"
)

// ReaderFault represents an error with reading wrapped as a fault.
type ReaderFault string

// Print satisfies the Fault interface.
func (err ReaderFault) Print(file string) {
	fmt.Println(err)
}
