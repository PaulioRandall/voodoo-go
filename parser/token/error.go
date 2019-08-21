package token

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	leadingLines  = 5
	trailingLines = 5
	// Modify carefully
	leadLines  = 0 - leadingLines
	trailLines = 1 + trailingLines
)

// PrintErrorToken prints an error token including the lines of code where the
// occurrs.
func PrintErrorToken(file string, tk Token) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = prettyPrintErrorToken(f, tk)
	if err != nil {
		panic(err)
	}
}

// prettyPrintErrorToken prints an informative error message for a specific line
// within the scroll.
func prettyPrintErrorToken(f *os.File, tk Token) error {
	sb := &strings.Builder{}

	addHeader(f, sb)
	err := addLines(f, sb, tk.Line+leadLines, tk.Line+1)
	if err != nil {
		return err
	}

	addMessages(sb, tk.End, tk.Errors...)
	err = addLines(f, sb, tk.Line+1, tk.Line+trailLines)
	if err != nil {
		return err
	}

	fmt.Println(sb.String())
	return nil
}

// addHeader adds a header to the print message.
func addHeader(f *os.File, sb *strings.Builder) {
	sb.WriteString("\n[SYNTAX ERROR]\n")

	sb.WriteString(`Line: `)
	sb.WriteRune('`')
	sb.WriteString(f.Name())
	sb.WriteRune('`')
	sb.WriteRune('\n')

	pad := strings.Repeat(`-`, 4)
	sb.WriteString(pad)
	sb.WriteString(`:`)
	pad = strings.Repeat(`-`, len(f.Name())+3)
	sb.WriteString(pad)
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
func readNextLine(f *os.File) (string, bool, error) {

	sb := strings.Builder{}
	utfLineFeed := byte(10)

	b := []byte{0}
	var err error

	for {
		_, err = f.Read(b)
		if err == io.EOF {
			return ``, true, nil
		}

		if err != nil {
			return ``, false, err
		}

		if b[0] == utfLineFeed {
			break
		}

		sb.WriteByte(b[0])
	}

	return sb.String(), false, nil
}

// addLines adds the specified number of lines to the builder. It will stop if
// EOF is encountered.
func addLines(f *os.File, sb *strings.Builder, start, end int) error {
	if start < 0 {
		start = 0
	}
	seekLine(f, start)

	for i := start; i < end; i++ {
		s, eof, err := readNextLine(f)
		if eof || err != nil {
			return err
		}

		s = fmt.Sprintf("%4d: %s\n", i+1, s)
		sb.WriteString(s)
	}

	return nil
}

// addMessages adds each input message to the print message.
func addMessages(sb *strings.Builder, col int, msgs ...string) {
	first := true
	for _, s := range msgs {
		addIndent(sb, 6, col)

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
