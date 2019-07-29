package token

// TokenType represents the type of a token.
type TokenType int

const (
	UNDEFINED TokenType = iota
	// Fully or partly alphabetic
	KEYWORD_FUNC // func
	KEYWORD_LOOP // loop
	KEYWORD_WHEN // when
	KEYWORD_END  // end
	// Identifiers
	IDENTIFIER_EXPLICIT // Declared by the coder
	IDENTIFIER_IMPLICIT // Inserted during syntax analysis
	// Booleans
	BOOLEAN_TRUE  // true
	BOOLEAN_FALSE // false
	SOURCERY      // @Blahblah
	// Literals
	LITERAL_NUMBER // ##.###
	LITERAL_STRING // "blahblah"
	COMMENT        // // blahblah
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
	ASSIGNMENT          // <-
	SEPARATOR_VALUE     // ,
	SEPARATOR_KEY_VALUE // :
	RANGE               // ..
	// Whitespace
	WHITESPACE
	// Ignoramuses
	VOID // _
)

// TokenName returns the name of the token type.
func TokenName(t TokenType) string {
	switch t {
	case KEYWORD_FUNC:
		return `KEYWORD_FUNC`
	case KEYWORD_LOOP:
		return `KEYWORD_LOOP`
	case KEYWORD_WHEN:
		return `KEYWORD_WHEN`
	case KEYWORD_END:
		return `KEYWORD_END`
	case IDENTIFIER_EXPLICIT:
		return `IDENTIFIER_EXPLICIT`
	case IDENTIFIER_IMPLICIT:
		return `IDENTIFIER_IMPLICIT`
	case BOOLEAN_TRUE:
		return `BOOLEAN_TRUE`
	case BOOLEAN_FALSE:
		return `BOOLEAN_FALSE`
	case SOURCERY:
		return `SOURCERY`
	case LITERAL_NUMBER:
		return `LITERAL_NUMBER`
	case LITERAL_STRING:
		return `LITERAL_STRING`
	case COMMENT:
		return `COMMENT`
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
	case ASSIGNMENT:
		return `ASSIGNMENT`
	case SEPARATOR_VALUE:
		return `SEPARATOR_VALUE`
	case SEPARATOR_KEY_VALUE:
		return `SEPARATOR_KEY_VALUE`
	case RANGE:
		return `RANGE`
	case WHITESPACE:
		return `WHITESPACE`
	case VOID:
		return `VOID`
	}

	return `UNDEFINED`
}
