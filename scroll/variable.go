package scroll

import (
	"strconv"
	"strings"
)

// VooType represents the type of a voodoo value.
type VooType string

// Declaration of voo types.
const (
	BoolType   VooType = "bool"
	NumType    VooType = "number"
	StrType    VooType = "string"
	KeyValType VooType = "key-value"
	ListType   VooType = "list"
	SpellType  VooType = "spell"
)

// VooValue represents a value within the current block.
type VooValue interface {

	// Type returns the type of the value
	Type() VooType

	// String returns the string representation of the value.
	String() string
}

// VooPrimitive represents a primitive type. I.e. bool, number,
// or string.
type VooPrimitive interface {
	VooValue

	// PrimitiveOnly does nothing when invoked. It ensures
	// compile failure if non-primitive types used where they
	// should not.
	PrimitiveOnly()
}

// *************************************************************************************

// VooBool represents a boolean value.
type VooBool bool

// Type satisfies the VooValue interface.
func (v VooBool) Type() VooType {
	return BoolType
}

// String satisfies the VooValue interface.
func (v VooBool) String() string {
	b := bool(v)
	return strconv.FormatBool(b)
}

// PrimitiveOnly satisfies the VooPrimitive interface.
func (v VooBool) PrimitiveOnly() {
}

// *************************************************************************************

// VooNum represents a number value.
type VooNum float64

// Type satisfies the VooValue interface.
func (v VooNum) Type() VooType {
	return NumType
}

// String satisfies the VooValue interface.
func (v VooNum) String() string {
	f := float64(v)
	return strconv.FormatFloat(f, 'f', 6, 64)
}

// PrimitiveOnly satisfies the VooPrimitive interface.
func (v VooNum) PrimitiveOnly() {
}

// *************************************************************************************

// VooStr represents a string value.
type VooStr string

// Type satisfies the VooValue interface.
func (v VooStr) Type() VooType {
	return StrType
}

// String satisfies the VooValue interface.
func (v VooStr) String() string {
	return string(v)
}

// PrimitiveOnly satisfies the VooPrimitive interface.
func (v VooStr) PrimitiveOnly() {
}

// *************************************************************************************

// VooKeyVal represents a key value type.
type VooKeyVal struct {
	Key VooPrimitive
	Val VooValue
}

// Type satisfies the VooValue interface.
func (v VooKeyVal) Type() VooType {
	return KeyValType
}

// String satisfies the VooValue interface.
func (kv VooKeyVal) String() string {
	k := kv.Key.String()
	v := kv.Val.String()
	return k + ": " + v
}

// *************************************************************************************

// VooList represents a list type.
type VooList struct {
	ValueType VooType
	Val       []VooValue
}

// Type satisfies the VooValue interface.
func (v VooList) Type() VooType {
	return ListType
}

// String satisfies the VooValue interface.
func (vl VooList) String() string {
	var sb strings.Builder
	sb.WriteString("[\n")

	list := vl.Val
	last := len(list) - 1
	for i, v := range list {
		t := v.Type()

		if t == BoolType || t == NumType {
			sb.WriteString(v.String())
		} else {
			sb.WriteRune('\'')
			sb.WriteString(v.String())
			sb.WriteRune('\'')
		}

		if i < last {
			sb.WriteString(", ")
		}
	}

	sb.WriteRune(']')
	return sb.String()
}

// *************************************************************************************

// VooSpell represents a spell type.
type VooSpell struct {
	// TODO
}

// Type satisfies the VooValue interface.
func (v VooSpell) Type() VooType {
	return SpellType
}

// String satisfies the VooValue interface.
func (v VooSpell) String() string {
	return "TODO"
}
