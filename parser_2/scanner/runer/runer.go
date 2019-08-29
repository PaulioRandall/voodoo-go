package runer

import (
	"bufio"
	"io"
	"strings"
)

// Runer wraps a bufio.Reader to provide easy reading and peeking of runes. It
// allows a look ahead of two runes. It also keeps a track of the current line
// and column indexes.
type Runer struct {
	reader  *bufio.Reader
	line    int
	col     int
	newline bool
	eof     bool
	buf     rune
	bufEOF  bool
}

// New returns a new initialised Runer.
func New(r *bufio.Reader) *Runer {
	if r == nil {
		panic(`Can't construct new Runer: *bufio.Reader == nil`)
	}

	return &Runer{
		reader:  r,
		line:    -1,
		col:     -1,
		newline: true,
		buf:     -1,
	}
}

// NewByStr creates a new Runer from the input string.
func NewByStr(s string) *Runer {
	sr := strings.NewReader(s)
	br := bufio.NewReader(sr)
	return New(br)
}

// Line returns the current line index, number of line feeds encountered.
func (r *Runer) Line() int {
	if r.newline {
		return r.line + 1
	}
	return r.line
}

// Col returns the column index of the last rune returned or 0 if no calls to
// read runes has been made yet.
func (r *Runer) Col() int {
	if r.newline {
		return -1
	}
	return r.col
}

// NextCol returns the column index after the last rune returned or 0 if no
// runes have been read yet.
func (r *Runer) NextCol() int {
	if r.newline {
		return 0
	}
	return r.col + 1
}

// Peek returns the next rune in the sequence without incrementing the 'cursor'.
func (r *Runer) Peek() (rune, bool, error) {
	err := r.ensureBufferInit()
	return r.buf, r.bufEOF, err
}

// ensureBufferInit checks if the buffer has been initialised, if it hasn't it
// initialises it.
func (r *Runer) ensureBufferInit() error {
	if r.buf == -1 {
		return r.buffer()
	}
	return nil
}

// buffer reads a rune from the reader and places the it in the buffer along
// with the buffer EOF flag.
func (r *Runer) buffer() error {
	var err error
	r.buf, r.bufEOF, err = r.read()
	return err
}

// readRune reads the next rune in the sequence returning the rune followed by
// an EOF flag.
func (r *Runer) read() (rune, bool, error) {
	ru, _, err := r.reader.ReadRune()

	if err == io.EOF {
		return 0, true, nil
	}

	return ru, false, err
}

// Read reads the next rune from the reader followed by a flag indicating
// the end of the file.
func (r *Runer) Read() (rune, bool, error) {
	if r.eof {
		return 0, true, nil
	}

	if r.newline {
		r.newline = false
		r.line++
		r.col = -1
	}

	if err := r.ensureBufferInit(); err != nil {
		return 0, false, err
	}

	return r.next()
}

// next returns the rune in the buffer and then reads the next rune from the
// reader into the buffer.
func (r *Runer) next() (rune, bool, error) {
	var ru rune
	ru, r.eof = r.buf, r.bufEOF
	r.col++

	if ru == '\n' {
		r.newline = true
	}

	return ru, r.eof, r.buffer()
}

// Skip skips the next rune in the reader. True is returned if the end of the
// file has been reached.
func (r *Runer) Skip() (bool, error) {
	_, eof, err := r.Read()
	return eof, err
}
