package fault

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Dummy returns a new Fault intended for test use only.
func Dummy(t FaultType, line, from, to int) Fault {
	return stdFault{
		line:    line,
		from:    from,
		to:      to,
		errType: t,
	}
}

// Assert asserts that the `act` fault is the same as the
// `exp` fault except for the error message.
func Assert(t *testing.T, exp Fault, act Fault) {
	e := exp.(stdFault)
	a := act.(stdFault)
	var m string

	m = "Expected fault type `" + FaultName(e.errType) + "` but got `" + FaultName(a.errType) + "`"
	assert.Equal(t, e.errType, a.errType, m)

	m = "Expected fault line `" + strconv.Itoa(e.line) + "` but got `" + strconv.Itoa(a.line) + "`"
	assert.Equal(t, e.line, a.line, m)

	m = "Expected fault at columns [" + strconv.Itoa(e.from) + ":" + strconv.Itoa(e.to) + "]"
	m += " but got [" + strconv.Itoa(a.from) + ":" + strconv.Itoa(a.to) + "]"
	assert.Equal(t, e.from, a.from, m)
	assert.Equal(t, e.to, a.to, m)
}
