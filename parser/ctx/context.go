package ctx

import (
	"strings"

	"github.com/PaulioRandall/voodoo-go/parser/value"
)

// Context represents the current scope including all variables within scope and
// can perform system activities allowable by the scope.
type Context struct {
	Args []string
	Vars map[string]value.Value
}

// New creates a new context with the given arguments.
func New(args []string) Context {
	return Context{
		Args: args,
		Vars: map[string]value.Value{},
	}
}

// String returns the string representation of the context.
func (c Context) String() string {
	sb := strings.Builder{}

	for k, v := range c.Vars {
		sb.WriteString(k)
		sb.WriteString(`.`)
		sb.WriteString(v.String())
		sb.WriteRune('\n')
	}

	return sb.String()
}
