package new_parser

// StatName satisfies the Statement interface.
func (e *Assignment) StatName() string {
	return `assignment`
}

// ExprName satisfies the Expression interface.
func (e *Value) ExprName() string {
	return `value`
}

// ExprName satisfies the Expression interface.
func (e *Join) ExprName() string {
	return `join`
}

// ExprName satisfies the Expression interface.
func (e *Operation) ExprName() string {
	return `operation`
}
