package scroll

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/voodoo-go/shared"
)

// Scroll represents a scroll.
type Scroll struct {
	File   string   // File path to the scroll
	Lines  []string // Raw lines from the scroll
	Length int      // Length of the scroll
}

// LoadScroll reads the lines of the scroll and creates a
// new Scroll instance for it.
func LoadScroll(path string) (*Scroll, error) {
	lines, err := shared.ReadLines(path)
	if err != nil {
		return nil, err
	}
	sc := &Scroll{
		File:   path,
		Lines:  lines,
		Length: len(lines),
	}
	return sc, nil
}

// PrettyPrintError prints an informative error message for a specific
// line within the scroll.
func (sc *Scroll) PrettyPrintError(line, index int, msgs ...string) {
	sb := &strings.Builder{}
	addHeader(sc, sb)
	addLines(sc, sb, line-4, line+1)
	addMessages(sb, index, msgs...)
	addLines(sc, sb, line+1, line+5)
	fmt.Println(sb.String())
}

// addHeader adds a header to the print message.
func addHeader(sc *Scroll, sb *strings.Builder) {
	sb.WriteString("\n[SYNTAX ERROR]")
	sb.WriteString(" `")
	sb.WriteString(sc.File)
	sb.WriteString("`")
	sb.WriteRune('\n')
}

// addLines adds the specified range of lines within the
// scroll to the print message. It will ignore lines that
// are out of bounds.
func addLines(sc *Scroll, sb *strings.Builder, start, end int) {
	size := len(sc.Lines)

	for i := start; i < end; i++ {
		if i >= 0 && i < size {
			s := fmt.Sprintf("%3d: %s\n", i, sc.Lines[i])
			sb.WriteString(s)
		}
	}
}

// addMessages adds each input message to the print message.
func addMessages(sb *strings.Builder, col int, msgs ...string) {
	first := true
	for _, s := range msgs {
		addIndent(sb, 4, col)

		if first {
			first = false
			sb.WriteString("^...")
		} else {
			sb.WriteString("    ")
		}

		sb.WriteString(s)
		sb.WriteRune('\n')
	}
}

// addIndent adds a defined number of spaces to the print
// message.
func addIndent(sb *strings.Builder, base, specific int) {
	indent := base + specific
	s := strings.Repeat(" ", indent)
	sb.WriteString(s)
}
