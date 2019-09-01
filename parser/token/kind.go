package token

// Kind represents the type of a token.
type Kind int

const (
	TT_UNDEFINED Kind = iota
	TT_NEWLINE        // '\n' or '\r\n'
	TT_SPACE          // All whitespace characters except newlines
	TT_SHEBANG        // Always the first line in a file
	TT_ID             // Identifier
	TT_ASSIGN         // :
	TT_NUMBER         // 123.456
	TT_ADD            // +
	TT_SUBTRACT       // -
	TT_MULTIPLY       // *
	TT_DIVIDE         // /
	TT_MODULO         // %
)

// KindName returns the name of the token type.
func KindName(t Kind) string {
	switch t {
	case TT_SHEBANG:
		return `SHEBANG`
	case TT_NEWLINE:
		return `NEWLINE`
	case TT_SPACE:
		return `SPACE`
	case TT_ID:
		return `ID`
	case TT_ASSIGN:
		return `ASSIGN`
	case TT_NUMBER:
		return `NUMBER`
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
	default:
		return `UNDEFINED`
	}
}