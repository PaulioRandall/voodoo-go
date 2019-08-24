package token

// TokenType represents the type of a token.
type TokenType int

const (
	TT_UNDEFINED      TokenType = iota
	TT_ERROR_UPSTREAM           // An error occurred upstream so close gracefully
	//
	TT_NEWLINE      // '\n', often converts to END_OF_STATEMENT token
	TT_SHEBANG      // Always the first line in a file
	TT_EOS          // END OF STATEMENT
	TT_FUNC         // "func"
	TT_LOOP         // "loop"
	TT_WHEN         // "when"
	TT_TRUE         // "true"
	TT_FALSE        // "false"
	TT_ID           // Identifier
	TT_SPELL        // EBNF: "@", IDENTIFIER
	TT_NUMBER       // EBNF: whole part, [ fractional part ]
	TT_STRING       // EBNF: '"', { string character }, '"'
	TT_SPACE        // All whitespace characters except newlines
	TT_COMMENT      // Same as Go line comment
	TT_ASSIGN       // <-
	TT_CMP_EQ       // ==
	TT_CMP_NOT_EQ   // !=
	TT_CMP_LT       // <
	TT_CMP_LT_OR_EQ // <=
	TT_CMP_MT       // >
	TT_CMP_MT_OR_EQ // >=
	TT_OR           // ||
	TT_AND          // &&
	TT_NOT          // !
	TT_MATCH        // =>
	TT_ADD          // +
	TT_SUBTRACT     // -
	TT_MULTIPLY     // *
	TT_DIVIDE       // /
	TT_MODULO       // %
	TT_CURLY_OPEN   // {
	TT_CURLY_CLOSE  // }
	TT_CURVED_OPEN  // (
	TT_CURVED_CLOSE // )
	TT_SQUARE_OPEN  // [
	TT_SQUARE_CLOSE // ]
	TT_VALUE_DELIM  // ,
	TT_VOID         // _
)

// IsOperator returns true if the input type is an operator.
func IsOperator(t TokenType) bool {
	return t >= TT_CMP_EQ && t <= TT_MODULO
}

// TokenName returns the name of the token type.
func TokenName(t TokenType) string {
	switch t {
	case TT_ERROR_UPSTREAM:
		return `ERROR UPSTREAM`
	case TT_SHEBANG:
		return `SHEBANG`
	case TT_NEWLINE:
		return `NEWLINE`
	case TT_EOS:
		return `EOS`
	case TT_FUNC:
		return `FUNC`
	case TT_LOOP:
		return `LOOP`
	case TT_WHEN:
		return `WHEN`
	case TT_TRUE:
		return `TRUE`
	case TT_FALSE:
		return `FALSE`
	case TT_ID:
		return `ID`
	case TT_SPELL:
		return `SPELL`
	case TT_NUMBER:
		return `NUMBER`
	case TT_STRING:
		return `STRING`
	case TT_SPACE:
		return `SPACE`
	case TT_COMMENT:
		return `COMMENT`
	case TT_ASSIGN:
		return `ASSIGN`
	case TT_CMP_EQ:
		return `EQUAL`
	case TT_CMP_NOT_EQ:
		return `NOT EQUAL`
	case TT_CMP_LT:
		return `LESS THAN`
	case TT_CMP_LT_OR_EQ:
		return `LESS THAN OR EQUAL`
	case TT_CMP_MT:
		return `MORE THAN`
	case TT_CMP_MT_OR_EQ:
		return `MORE THAN OR EQUAL`
	case TT_OR:
		return `OR`
	case TT_AND:
		return `AND`
	case TT_NOT:
		return `NOT`
	case TT_MATCH:
		return `MATCH THEN`
	case TT_ADD:
		return `ADD`
	case TT_SUBTRACT:
		return `SUBTRACT`
	case TT_MULTIPLY:
		return `MULTIPLY`
	case TT_DIVIDE:
		return `DIVIDE`
	case TT_MODULO:
		return `MODULO`
	case TT_CURLY_OPEN:
		return `CURLY OPEN`
	case TT_CURLY_CLOSE:
		return `CURLY CLOSE`
	case TT_CURVED_OPEN:
		return `CURVED OPEN`
	case TT_CURVED_CLOSE:
		return `CURVED CLOSE`
	case TT_SQUARE_OPEN:
		return `SQUARE OPEN`
	case TT_SQUARE_CLOSE:
		return `SQUARE CLOSE`
	case TT_VALUE_DELIM:
		return `DELIM`
	case TT_VOID:
		return `VOID`
	}

	return `UNDEFINED`
}
