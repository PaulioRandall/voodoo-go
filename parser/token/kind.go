package token

// Kind represents the type of a token.
type Kind int

const (
	TK_UNDEFINED Kind = iota
	TK_NEWLINE        // '\n' or '\r\n'
	TK_SPACE          // All whitespace characters except newlines
	TK_SHEBANG        // Always the first line in a file
	TK_ID             // Identifier
	TK_ASSIGN         // <-, :=
	TK_BOOL           // true, false
	TK_NUMBER         // 123.456
	TK_STRING         // `...`
	TK_ADD            // +
	TK_SUBTRACT       // -
	TK_MULTIPLY       // *
	TK_DIVIDE         // /
	TK_MODULO         // %
	TK_VOID           // _
	TK_DELIM          // ,
)

// KindName returns the name of the token type.
func KindName(t Kind) string {
	switch t {
	case TK_SHEBANG:
		return `SHEBANG`
	case TK_NEWLINE:
		return `NEWLINE`
	case TK_SPACE:
		return `SPACE`
	case TK_ID:
		return `ID`
	case TK_ASSIGN:
		return `ASSIGN`
	case TK_BOOL:
		return `BOOL`
	case TK_NUMBER:
		return `NUMBER`
	case TK_STRING:
		return `STRING`
	case TK_ADD:
		return `ADD`
	case TK_SUBTRACT:
		return `SUBTRACT`
	case TK_MULTIPLY:
		return `MULTIPLY`
	case TK_DIVIDE:
		return `DIVIDE`
	case TK_MODULO:
		return `MODULO`
	case TK_VOID:
		return `VOID`
	case TK_DELIM:
		return `DELIM`
	default:
		return `UNDEFINED`
	}
}
