package scanner

import (
	"github.com/PaulioRandall/voodoo-go/parser_2/scanner/runer"
	"github.com/PaulioRandall/voodoo-go/parser_2/token"
)

// scanWord scans word tokens returning a keyword or identifier.
func scanWord(r *runer.Runer) (token.Token, ScanError) {
	start := r.NextCol()

	text, err := scanWordStr(r)
	if err != nil {
		return nil, err
	}

	tk := scanTok{
		text:  text,
		line:  r.Line(),
		start: start,
		end:   r.NextCol(),
		kind:  findWordKind(text),
	}

	return tk, nil
}

// findWordKind finds the kind of the word.
func findWordKind(text string) token.Kind {
	return token.TT_ID
	/*
		switch strings.ToLower(text) {
		case `func`:
			return token.TT_FUNC
		case `loop`:
			return token.TT_LOOP
		case `match`:
			return token.TT_MATCH
		case `true`:
			return token.TT_TRUE
		case `false`:
			return token.TT_FALSE
		default:
			return token.TT_ID
		}
	*/
}
