package scanner

import (
	"bufio"
	"io"
)

const NUL = rune(0)
const EOF = rune(3)

// Runer wraps a bufio.Reader to provide easy reading and peeking of runes. It
// allows a look ahead of two runes by using a temp array. It also keeps a track
// of the current line and column index.
type Runer struct {
	line    int
	col     int
	newLine bool
	reader  *bufio.Reader
	buf     [2]rune
}

// NewRuner returns a new initialised Runer instance.
func NewRuner(reader *bufio.Reader) *Runer {
	return &Runer{
		reader:  reader,
		line:    -1,
		col:     -1,
		newLine: true,
	}
}

// Line returns the line index, number of newline runes incountered.
func (r *Runer) Line() int {
	l := r.line
	if l == -1 {
		return 0
	}
	return l
}

// Col returns the column index of the last rune returned or -1 if no calls to
// read runes has been made yet.
func (r *Runer) Col() int {
	return r.col
}

// ReadRune reads the next rune from the reader. EOF is returned if the end of
// the file has been reached.
func (r *Runer) ReadRune() (rune, error) {
	if r.newLine {
		r.newLine = false
		r.line++
		r.col = -1
	}

	ru, err := r.nextRune()
	r.col++

	if ru == '\n' {
		r.newLine = true
	}

	return ru, err
}

// SkipRune skips the next rune in the reader. It still may produce an error as
// the reader may still be read in order to do this.
func (r *Runer) SkipRune() error {
	_, err := r.ReadRune()
	return err
}

// LookAhead returns the next two runes in the sequence without incrementing the
// 'cursor'. After a call to LookAhead() it is safe to ignore the error returned
// on the next two calls to ReadRune() or SkipRune().
func (r *Runer) LookAhead() (rune, rune, error) {
	var err error

	if r.buf[0] == NUL {
		r.buf[0], err = r.readRune()
		if err != nil {
			return NUL, NUL, err
		}
	}

	if r.buf[1] == NUL {
		r.buf[1], err = r.readRune()
		if err != nil {
			return NUL, NUL, err
		}
	}

	return r.buf[0], r.buf[1], nil
}

// nextRune returns the next rune in the sequence. It will check the temp buffer
// before trying the reader.
func (r *Runer) nextRune() (rune, error) {
	ru := r.buf[0]
	if ru == NUL {
		return r.readRune()
	}

	r.buf[0] = r.buf[1]
	r.buf[1] = NUL
	return ru, nil
}

// readRune reads the next rune in the sequence returning EOF if the end of the
// reader has been reached.
func (r *Runer) readRune() (rune, error) {
	ru, _, err := r.reader.ReadRune()

	if err == io.EOF {
		return EOF, nil
	}

	if err != nil {
		return NUL, err
	}

	return ru, nil
}
