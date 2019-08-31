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
	reader   *bufio.Reader
	line     int
	col      int
	newline  bool
	eof      bool
	buf1     rune
	buf1_eof bool
	buf2     rune
	buf2_eof bool
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
		buf1:    -1,
		buf2:    -1,
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

// Col returns the column index of the last rune returned.
func (r *Runer) Col() int {
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
	e := r.ensureBufferInit()
	return r.buf1, r.buf1_eof, e
}

// PeekMore returns the rune after the next rune in the sequence without
// incrementing the 'cursor'.
func (r *Runer) PeekMore() (rune, bool, error) {
	e := r.ensureBufferInit()
	return r.buf2, r.buf2_eof, e
}

// PeekBoth returns the next rune and the rune after the next rune in the
// sequence without incrementing the 'cursor'.
func (r *Runer) PeekBoth() (rune, rune, bool, error) {
	e := r.ensureBufferInit()
	return r.buf1, r.buf2, r.buf1_eof, e
}

// ensureBufferInit checks if the buffer has been initialised, if it hasn't it
// initialises it.
func (r *Runer) ensureBufferInit() error {
	if r.buf1 == -1 || r.buf2 == -1 {
		return r.buffer()
	}
	return nil
}

// buffer reads a rune from the reader and places the it in the buffer along
// with the buffer EOF flag.
func (r *Runer) buffer() error {
	if r.eof {
		return nil
	}

	var e error
	r.buf1, r.buf1_eof = r.buf2, r.buf2_eof
	r.buf2, r.buf2_eof, e = r.read()
	if e != nil {
		return e
	}

	if r.buf1 == -1 {
		return r.buffer()
	}

	return nil
}

// readRune reads the next rune in the sequence returning the rune followed by
// an EOF flag.
func (r *Runer) read() (rune, bool, error) {
	ru, _, e := r.reader.ReadRune()

	if e == io.EOF {
		return 0, true, nil
	}

	return ru, false, e
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

	if e := r.ensureBufferInit(); e != nil {
		return 0, false, e
	}

	return r.next()
}

// next returns the rune in the buffer and then reads the next rune from the
// reader into the buffer.
func (r *Runer) next() (rune, bool, error) {
	var ru rune
	ru, r.eof = r.buf1, r.buf1_eof
	r.col++

	if ru == '\n' {
		r.newline = true
	}

	return ru, r.eof, r.buffer()
}

// Skip skips the next rune in the reader. True is returned if the end of the
// file has been reached.
func (r *Runer) Skip() (bool, error) {
	_, eof, e := r.Read()
	return eof, e
}

// predicate returns true if the first rune passed is part of the token being
// scanned. The second rune iss the one after the next.
type predicate func(rune, rune) (bool, error)

// ReadIf reads the next rune only if the predicate function returns true. If
// the second return value is true then the a value was succesfully read from
// the reader and the predicate function evaluated to true.
func (r *Runer) ReadIf(f predicate) (rune, bool, error) {
	ru1, eof1, e := r.Peek()
	if eof1 || e != nil {
		return 0, !eof1, e
	}

	ru2, _, e := r.PeekMore()
	if e != nil {
		return 0, !eof1, e
	}

	want, e := f(ru1, ru2)
	if e != nil {
		return 0, false, e
	}

	if want {
		if _, e = r.Skip(); e != nil {
			return 0, false, e
		}
		return ru1, true, nil
	}

	return 0, false, nil
}

// ReadWhile reads runes until the predicate returns false. The runes are
// returned as a string.
func (r *Runer) ReadWhile(f predicate) (string, error) {
	sb := strings.Builder{}

	for ru, read, e := r.ReadIf(f); read; ru, read, e = r.ReadIf(f) {
		if e != nil {
			return ``, e
		}

		sb.WriteRune(ru)
	}

	return sb.String(), nil
}
