package token

// Kind represents the type of a token.
type Kind int

const (
	TT_UNDEFINED Kind = iota
	TT_NEWLINE        // '\n' or '\r\n'
	TT_SHEBANG        // Always the first line in a file
	//TT_FUNC         // "func"
	//TT_LOOP         // "loop"
	//TT_MATCH        // "match"
	//TT_TRUE         // "true"
	//TT_FALSE        // "false"
	TT_ID // Identifier
	//TT_SPELL        // Inbuilt function
	TT_NUMBER // 123.456
	//TT_STRING       // "blahblah"
	TT_SPACE // All whitespace characters except newlines
	//TT_COMMENT      // Same as Go line comment
	TT_ASSIGN // <-, :=, =
	//TT_CMP_EQ       // ==
	//TT_CMP_NOT_EQ   // !=
	//TT_CMP_LT       // <
	//TT_CMP_LT_OR_EQ // <=
	//TT_CMP_MT       // >
	//TT_CMP_MT_OR_EQ // >=
	//TT_OR           // ||
	//TT_AND          // &&
	//TT_NOT          // !
	//TT_IF_THEN      // =>
	TT_ADD         // +
	TT_SUBTRACT    // -
	TT_MULTIPLY    // *
	TT_DIVIDE      // /
	TT_MODULO      // %
	TT_BODY_OPEN   // {
	TT_BODY_CLOSE  // }
	TT_PARAM_OPEN  // (
	TT_PARAM_CLOSE // )
	TT_LIST_OPEN   // [
	TT_LIST_CLOSE  // ]
	TT_VALUE_DELIM // ,
	//TT_VOID         // _
)

// KindName returns the name of the token type.
func KindName(t Kind) string {
	switch t {
	case TT_SHEBANG:
		return `SHEBANG`
	case TT_NEWLINE:
		return `NEWLINE`
	//case TT_FUNC:
	//return `FUNC`
	//case TT_LOOP:
	//return `LOOP`
	//case TT_MATCH:
	//return `MATCH`
	//case TT_TRUE:
	//return `TRUE`
	//case TT_FALSE:
	//return `FALSE`
	case TT_ID:
		return `ID`
	//case TT_SPELL:
	//return `SPELL`
	case TT_NUMBER:
		return `NUMBER`
	//case TT_STRING:
	//return `STRING`
	case TT_SPACE:
		return `SPACE`
	//case TT_COMMENT:
	//return `COMMENT`
	case TT_ASSIGN:
		return `ASSIGN`
	//case TT_CMP_EQ:
	//return `EQUAL`
	//case TT_CMP_NOT_EQ:
	//return `NOT EQUAL`
	//case TT_CMP_LT:
	//return `LESS THAN`
	//case TT_CMP_LT_OR_EQ:
	//return `LESS THAN OR EQUAL`
	//case TT_CMP_MT:
	//return `MORE THAN`
	//case TT_CMP_MT_OR_EQ:
	//return `MORE THAN OR EQUAL`
	//case TT_OR:
	//return `OR`
	//case TT_AND:
	//return `AND`
	//case TT_NOT:
	//return `NOT`
	//case TT_IF_THEN:
	//return `IF THEN`
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
	//case TT_BODY_OPEN:
	//return `BODY OPEN`
	//case TT_BODY_CLOSE:
	//return `BODY CLOSE`
	//case TT_PARAM_OPEN:
	//return `PARAM OPEN`
	//case TT_PARAM_CLOSE:
	//return `PARAM CLOSE`
	//case TT_LIST_OPEN:
	//return `LIST OPEN`
	//case TT_LIST_CLOSE:
	//return `LIST CLOSE`
	case TT_VALUE_DELIM:
		return `DELIM`
		//case TT_VOID:
		//return `VOID`
	}

	return `UNDEFINED`
}
