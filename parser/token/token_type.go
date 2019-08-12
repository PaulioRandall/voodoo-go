package token

// TokenType represents the type of a token.
type TokenType int

const (
	UNDEFINED TokenType = iota
	// Shebang
	SHEBANG
	// Flag tokens
	END_OF_STATEMENT
	// Keywords
	KEYWORD_FUNC // func
	KEYWORD_LOOP // loop
	KEYWORD_WHEN // when
	KEYWORD_DONE // done, end of block
	// Identifiers
	IDENTIFIER
	// Booleans
	BOOLEAN_TRUE  // true
	BOOLEAN_FALSE // false
	// Spells
	SPELL // @Blahblah
	// Literals
	LITERAL_NUMBER // ##.###
	LITERAL_STRING // "blahblah"
	COMMENT        // // blahblah
	// Assignment
	ASSIGNMENT // <-
	// Comparison operators
	CMP_EQUAL                 // ==
	CMP_NOT_EQUAL             // !=
	CMP_LESS_THAN             // <
	CMP_LESS_THAN_OR_EQUAL    // <=
	CMP_GREATER_THAN          // >
	CMP_GREATER_THAN_OR_EQUAL // >=
	// Logical operators
	LOGICAL_OR    // ||
	LOGICAL_AND   // &&
	LOGICAL_NOT   // !
	LOGICAL_MATCH // =>
	// Arithmetic operators
	CALC_ADD      // +
	CALC_SUBTRACT // -
	CALC_MULTIPLY // *
	CALC_DIVIDE   // /
	CALC_MODULO   // %
	// Context boundries
	PAREN_CURVY_OPEN   // (
	PAREN_CURVY_CLOSE  // )
	PAREN_SQUARE_OPEN  // [
	PAREN_SQUARE_CLOSE // ]
	// Punctuation
	VALUE_DELIM // ,
	// Whitespace
	NEWLINE // \n
	WHITESPACE
	// Ignoramuses
	VOID // _
)

// IsOperator returns true if the input type is an operator.
func IsOperator(t TokenType) bool {
	return t >= CMP_EQUAL && t <= CALC_MODULO
}

// TokenName returns the name of the token type.
func TokenName(t TokenType) string {
	switch t {
	case SHEBANG:
		return `SHEBANG`
	case END_OF_STATEMENT:
		return `END_OF_STATEMENT`
	case KEYWORD_FUNC:
		return `KEYWORD_FUNC`
	case KEYWORD_LOOP:
		return `KEYWORD_LOOP`
	case KEYWORD_WHEN:
		return `KEYWORD_WHEN`
	case KEYWORD_DONE:
		return `KEYWORD_DONE`
	case IDENTIFIER:
		return `IDENTIFIER`
	case BOOLEAN_TRUE:
		return `BOOLEAN_TRUE`
	case BOOLEAN_FALSE:
		return `BOOLEAN_FALSE`
	case SPELL:
		return `SPELL`
	case LITERAL_NUMBER:
		return `LITERAL_NUMBER`
	case LITERAL_STRING:
		return `LITERAL_STRING`
	case COMMENT:
		return `COMMENT`
	case ASSIGNMENT:
		return `ASSIGNMENT`
	case CMP_EQUAL:
		return `CMP_EQUAL`
	case CMP_NOT_EQUAL:
		return `CMP_NOT_EQUAL`
	case CMP_LESS_THAN:
		return `CMP_LESS_THAN`
	case CMP_LESS_THAN_OR_EQUAL:
		return `CMP_LESS_THAN_OR_EQUAL`
	case CMP_GREATER_THAN:
		return `CMP_GREATER_THAN`
	case CMP_GREATER_THAN_OR_EQUAL:
		return `CMP_GREATER_THAN_OR_EQUAL`
	case LOGICAL_OR:
		return `OR`
	case LOGICAL_AND:
		return `AND`
	case LOGICAL_NOT:
		return `NOT`
	case LOGICAL_MATCH:
		return `LOGICAL_MATCH`
	case CALC_ADD:
		return `CALC_ADD`
	case CALC_SUBTRACT:
		return `CALC_SUBTRACT`
	case CALC_MULTIPLY:
		return `CALC_MULTIPLY`
	case CALC_DIVIDE:
		return `CALC_DIVIDE`
	case CALC_MODULO:
		return `CALC_MODULO`
	case PAREN_CURVY_OPEN:
		return `PAREN_CURVY_OPEN`
	case PAREN_CURVY_CLOSE:
		return `PAREN_CURVY_CLOSE`
	case PAREN_SQUARE_OPEN:
		return `PAREN_SQUARE_OPEN`
	case PAREN_SQUARE_CLOSE:
		return `PAREN_SQUARE_CLOSE`
	case VALUE_DELIM:
		return `VALUE_DELIM`
	case NEWLINE:
		return `NEWLINE`
	case WHITESPACE:
		return `WHITESPACE`
	case VOID:
		return `VOID`
	}

	return `UNDEFINED`
}