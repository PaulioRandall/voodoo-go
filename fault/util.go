package fault

import (
	"runtime"
)

func CurrLine() int {
	_, _, line, ok := runtime.Caller(1)

	if !ok {
		return -1
	}

	return line
}
