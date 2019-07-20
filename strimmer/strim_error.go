package strimmer

import (
	"github.com/PaulioRandall/voodoo-go/lexer"
)

// TODO: Remove dependency on lexer

// StrimError represents an error returned by the strimmer
// and contains additional error information such as the
// line and column number.
type StrimError lexer.LexError
