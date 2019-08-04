package ctx

import (
	"github.com/PaulioRandall/voodoo-go/fault"
	"github.com/PaulioRandall/voodoo-go/token"
)

// ValueType represents the name of the type of a value.
type ValueType string

const (
	UNDEFINED_TYPE  ValueType = `undefined`
	IDENTIFIER_TYPE ValueType = `identifier`
	BOOL_TYPE       ValueType = `bool`
	NUMBER_TYPE     ValueType = `number`
	STRING_TYPE     ValueType = `string`
	LIST_TYPE       ValueType = `list`
)

// Value represents an expression which simple evaluates
// to itself; an actual value or identifier.
type Value struct {
	Token token.Token
	Val   interface{}
	Type  ValueType
}

// Expr satisfies the Expression interface.
func (v Value) Expr() ExprType {
	return VALUE
}

// Evaluate satisfies the Expression interface.
func (v Value) Evaluate(c *Context) (Value, fault.Fault) {
	return v, nil
}

// IsSameType returns true if the value types are the same.
func (v Value) IsSameType(other Value) bool {
	return v.Type == other.Type
}
