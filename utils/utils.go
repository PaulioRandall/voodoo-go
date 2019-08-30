package utils

// LogicalConjunction performs a logical conjunction of the operand set return
// true if, and only if, operands are true.
func LogicalConjunction(operands ...bool) bool {
	for _, b := range operands {
		if b == false {
			return false
		}
	}
	return true
}
