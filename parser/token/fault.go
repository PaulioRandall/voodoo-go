package token

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SyntaxFault represents a generic fault with syntax.
type SyntaxFault struct {
	Line int      // Line where the error occurred
	Col  int      // Index where the error actually occurred on the line
	Msgs []string // Description of the error
}

// Print satisfies the Fault interface.
func (flt SyntaxFault) Print(file string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = flt.prettyPrintError(f)
	if err != nil {
		panic(err)
	}
}

// prettyPrintError prints an informative error message for a specific line
// within the scroll.
func (flt SyntaxFault) prettyPrintError(f *os.File) error {
	sb := &strings.Builder{}

	addHeader(f, sb)
	err := addLines(f, sb, flt.Line-4, flt.Line+1)
	if err != nil {
		return err
	}
	addMessages(sb, flt.Col, flt.Msgs...)

	fmt.Println(sb.String())
	return nil
}

// addHeader adds a header to the print message.
func addHeader(f *os.File, sb *strings.Builder) {
	sb.WriteString("\n[SYNTAX ERROR] `")
	sb.WriteString(f.Name())
	sb.WriteRune('`')
	sb.WriteRune('\n')
}

// seekLine sets the file cursor to point at the specified line. If the index is
// less than zero then the cursor will be set to the first line. An error
// returns if the index exceeds the last line.
func seekLine(f *os.File, index int) (err error) {
	f.Seek(0, io.SeekStart)

	if index == 0 {
		return
	}

	n := 0
	for {
		err = jumpToNextLine(f)
		n++

		if err != nil || n >= index {
			break
		}
	}

	return
}

// jumpToNextLine jumps to the next new line in the file.
func jumpToNextLine(f *os.File) (err error) {
	utfLineFeed := byte(10)
	b := []byte{0}

	for {
		_, err = f.Read(b)
		if err != nil || b[0] == utfLineFeed {
			break
		}
	}

	return
}

// readNextLine jumps to the next new line in the file.
func readNextLine(f *os.File) (string, error) {

	sb := strings.Builder{}
	utfLineFeed := byte(10)

	b := []byte{0}
	var err error

	for {
		_, err = f.Read(b)
		if err != io.EOF && err != nil {
			return ``, err
		}

		if b[0] == utfLineFeed {
			break
		}

		sb.WriteByte(b[0])
	}

	return sb.String(), nil
}

// addLines adds the specified number of lines to the builder. It will stop if
// EOF is encountered.
func addLines(f *os.File, sb *strings.Builder, start, end int) (err error) {
	seekLine(f, start)

	for i := start; i < end; i++ {
		s, err := readNextLine(f)
		if err != nil {
			break
		}

		s = fmt.Sprintf("%3d: %s\n", i, s)
		sb.WriteString(s)
	}

	return
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

// addIndent adds a defined number of spaces to the print message.
func addIndent(sb *strings.Builder, base, specific int) {
	indent := base + specific
	s := strings.Repeat(" ", indent)
	sb.WriteString(s)
}
