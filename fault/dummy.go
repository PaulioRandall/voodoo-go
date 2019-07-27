package fault

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Dummy returns a new Fault intended for test use only.
func Dummy(t FaultType) Fault {
	return stdFault{
		errType: t,
	}
}

// Assert asserts that the `act` fault is the same as the
// `exp` fault except for the error message.
func Assert(t *testing.T, exp Fault, act Fault) {
	e := exp.(stdFault)
	a := act.(stdFault)
	assert.Equal(t, e.errType, a.errType)
	assert.Equal(t, e.line, a.line)
	assert.Equal(t, e.from, a.from)
	assert.Equal(t, e.to, a.to)
}
