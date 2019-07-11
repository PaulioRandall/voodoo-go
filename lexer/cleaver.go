
package cleaver

import (
	"fmt"
	"strings"
	"strconv"
)

// The purpose of the cleaver is to split up a string, usually a line,
// into fragments of elements that represent either a potential
// lexeme/token or part of a potential one. Assuming valid syntax,
// each fragment can be one of:
// - terminal UTF-8 rune symbol such as '=', '+', a whitespace rune, etc
// - part of a terminal UTF-8 string symbol such as '=', '>', etc
//	 (resolving to '=>')
// - terminal UTF-8 string symbol such as 'spell', 'true', etc
// - non-terminal UTF-8 string symbol such as 'i', 'isAlive', etc
// - part of a non-terminal UTf-8 string symbol such as 'is', '_', 'alive'
//	 that may share runes with terminal symbols such that the examples come
//	 together to form the identifier 'is_alive'.
//
// Fragments are joined back together later to form elements which are
// scanned to form lexemes which are evaluated to form tokens. Tokens
// represent a meaningful symbol, and if applicable, with a value.
//
// A nice property of cleaving is that a returned Fragment array containing
// a single empty fragment, one that contains no runes, always represents
// an empty line in the scroll. This will be useful for formatting or
// reccreating the scroll later if we wish; allthough all lines would be
// void of leading and trailing whiitespace if so but correct indentation
// can be inferred.
//
// Example:
// Original: true => @Print("It's true!")
// Cleaved_: [
//             `true`, ` `, `=`, `>`, ` `, `@`,      // `true => @`
//             `Print`, `(`, `"`, `It`, `'`, `s`,    // `Print("It's`
//             ` `, `true`, `!`, `"`, `)`            // ` true!")`
//           ]

// Fragment are described above. 
type Fragment struct {
	Val string
	Start int
	End int				// Exclusive
}

// String creates a string representation of the Fragment.
// TODO: Is this a duplicate of an Element?
func (frag Fragment) String() string {
	start := strconv.Itoa(frag.Start)
	start = strings.Repeat(` `, 3 - len(start)) + start
	return fmt.Sprintf("[%s:%-3d] `%s`", start, frag.End, frag.Val)
}

// PrintFragments prints an array of Fragments.
func PrintFragments(frags []Fragment) {
	for _, v := range frags {
		fmt.Println(v)
	}
}

// TODO: Create SQL flow diagram of what the cleaver does!
//
// Cleave splits a string into Fragments to make the rest of the
// scanning process an additive one rather than a splitting one.
// Paulio: I found it easier this way.
func Cleave(s string) []Fragment {
	result := []Fragment{}
	fragType := none
	var frag Fragment
	var val strings.Builder
	
	push := func(end int) {
		frag.Val = val.String()
		frag.End = end
		result = append(result, frag)
	}
	
	reset := func(start int, rt runeType) {
		val.Reset()
		frag = Fragment{
			Start: start,
		}
		fragType = rt
	}
	
	for i, r := range s {
		rt := runeTypeOf(r)
		
		switch {
		case (fragType == none):
			reset(i, rt)
		case (fragType == letter) && (rt == letter):
		case (fragType == number) && (rt == number):
		default:
			push(i)
			reset(i, rt)
		}
		
		val.WriteRune(r)
	}
	
	push(len(s))
	return result
}
