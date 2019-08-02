package new_parser

// Name satisfies the Expression interface.
func (e *Operation) Name() string {
	return `operation`
}

// Name satisfies the Expression interface.
func (e *Value) Name() string {
	return `value`
}

// Name satisfies the Expression interface.
func (e *Join) Name() string {
	return `join`
}
