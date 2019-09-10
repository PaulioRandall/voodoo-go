package perror

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	leadingLines        = 5
	trailingLines       = 5
	leadLines           = 0 - leadingLines
	trailLines          = 1 + trailingLines
	utfLineFeed         = byte(10)
	lenOfLine           = len(`line`)
	lenOfQuotesAndSpace = len(` ""`)
	msgIndent           = len(`line: `)
)

// PrintError prints a Perror.
func PrintError(file string, perr Perror) {
	f, e := os.Open(file)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	e = printError(f, perr)
	if e != nil {
		panic(e)
	}
}

// printError prints a Perror.
func printError(f *os.File, perr Perror) error {
	sb := &strings.Builder{}
	line := perr.Line()

	addHeader(f, sb)
	e := addLines(f, sb, line+leadLines, line+1)
	if e != nil {
		return e
	}

	addMsgs(sb, perr.Cols(), perr.Errors()...)
	e = addLines(f, sb, line+1, line+trailLines)
	if e != nil {
		return e
	}

	fmt.Println(sb.String())
	return nil
}

// addHeader adds a header to the print message.
func addHeader(f *os.File, sb *strings.Builder) {
	sb.WriteString("\n[SYNTAX ERROR]\n")
	path, e := filepath.Abs(f.Name())
	if e != nil {
		panic(e)
	}

	sb.WriteString(`Line: `)
	sb.WriteRune('`')
	sb.WriteString(path)
	sb.WriteRune('`')
	sb.WriteRune('\n')

	pad := strings.Repeat(`-`, lenOfLine)
	sb.WriteString(pad)
	sb.WriteString(`:`)

	padLen := len(path) + lenOfQuotesAndSpace
	pad = strings.Repeat(`-`, padLen)
	sb.WriteString(pad)
	sb.WriteRune('\n')
}

// seekLine sets the file cursor to point at the specified line. If the index is
// less than zero then the cursor will be set to the first line. An error
// returns if the index exceeds the last line.
func seekLine(f *os.File, index int) (e error) {
	f.Seek(0, io.SeekStart)

	if index == 0 {
		return
	}

	n := 0
	for {
		e = jumpToNextLine(f)
		n++

		if e != nil || n >= index {
			break
		}
	}

	return
}

// jumpToNextLine jumps to the next new line in the file.
func jumpToNextLine(f *os.File) (e error) {
	b := []byte{0}

	for {
		_, e = f.Read(b)
		if e != nil || b[0] == utfLineFeed {
			break
		}
	}

	return
}

// readNextLine jumps to the next new line in the file.
func readNextLine(f *os.File) (string, bool, error) {

	sb := strings.Builder{}

	b := []byte{0}
	var e error

	for {
		_, e = f.Read(b)
		if e == io.EOF {
			return ``, true, nil
		}

		if e != nil {
			return ``, false, e
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
		s, eof, e := readNextLine(f)
		if eof || e != nil {
			return e
		}

		s = fmt.Sprintf("%4d: %s\n", i+1, s)
		sb.WriteString(s)
	}

	return nil
}

// addMsgs adds each input message to the print message.
func addMsgs(sb *strings.Builder, cols []int, msgs ...string) {
	sort.Ints(cols)

	writeLineOfRunes(sb, msgIndent, cols, '^')

	for _, s := range msgs {
		addIndent(sb, msgIndent)

		sb.WriteString(`  `)

		s = fmt.Sprintf("\033[1;33m%s\033[0m", s)
		sb.WriteString(s)
		sb.WriteRune('\n')
	}
}

// writeLineOfRunes writes a rune to the string builder at every column index
// in the column array before finishing with a line feed.
func writeLineOfRunes(sb *strings.Builder, baseIndent int, cols []int, ru rune) {
	addIndent(sb, baseIndent)
	for i, c := range cols {
		addIndent(sb, c-i)
		s := fmt.Sprintf("\033[1;31m%s\033[0m", string(ru))
		sb.WriteString(s)
	}
	sb.WriteRune('\n')
}

// addIndent adds a defined number of spaces to the print message.
func addIndent(sb *strings.Builder, indent int) {
	s := strings.Repeat(" ", indent)
	sb.WriteString(s)
}
