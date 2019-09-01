package err

import (
	"fmt"
	"io"
	"os"
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

// PrintScanError prints a ScanError.
func PrintScanError(file string, sce ScanError) {
	f, e := os.Open(file)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	e = printScanError(f, sce)
	if e != nil {
		panic(e)
	}
}

// printScanError prints a ScanError.
func printScanError(f *os.File, sce ScanError) error {
	sb := &strings.Builder{}
	line := sce.Line()

	addHeader(f, sb)
	e := addLines(f, sb, line+leadLines, line+1)
	if e != nil {
		return e
	}

	addMsgs(sb, sce.Index(), sce.Errors()...)
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

	sb.WriteString(`Line: `)
	sb.WriteRune('`')
	sb.WriteString(f.Name())
	sb.WriteRune('`')
	sb.WriteRune('\n')

	pad := strings.Repeat(`-`, lenOfLine)
	sb.WriteString(pad)
	sb.WriteString(`:`)

	padLen := len(f.Name()) + lenOfQuotesAndSpace
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
func addMsgs(sb *strings.Builder, col int, msgs ...string) {
	first := true
	for _, s := range msgs {
		addIndent(sb, msgIndent+col)

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
func addIndent(sb *strings.Builder, indent int) {
	s := strings.Repeat(" ", indent)
	sb.WriteString(s)
}
