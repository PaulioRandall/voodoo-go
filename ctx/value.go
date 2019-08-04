package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// ValueType represents the name of the type of a value.
type ValueType string

const (
	UNDEFINED_TYPE  = `undefined`
	IDENTIFIER_TYPE = `identifier`
	BOOL_TYPE       = `bool`
	NUMBER_TYPE     = `number`
	STRING_TYPE     = `string`
	LIST_TYPE       = `list`
)

// Value represents an expression which simple evaluates
// to itself; an actual value or identifier.
type Value struct {
	Token token.Token
	Val   interface{}
	Type  string
}

// ExprName satisfies the Expression interface.
func (v Value) ExprName() string {
	return `value`
}

// Evaluate satisfies the Expression interface.
func (v Value) Evaluate(c *Context) (Value, fault.Fault) {
	return v, nil
}

// IsSameType returns true if the value types are the same.
func (v Value) IsSameType(other Value) bool {
	return v.Type == other.Type
}
