package new_scanner

import (
	"bufio"
	"io"

	"github.com/PaulioRandall/voodoo-go/fault"
)

const NUL = rune(0)
const EOF = rune(3)

// Runer wraps a bufio.Reader to provide easy reading and peeking of runes. It
// allows a look ahead of two runes by using a temp array. It also keeps a track
// of the number of lines read so far.
type Runer struct {
	line   int
	reader *bufio.Reader
	buf    [2]rune
}

// NewRuner returns a new initialised Runer instance.
func NewRuner(reader *bufio.Reader) *Runer {
	return &Runer{
		reader: reader,
	}
}

// Line returns the number of newline runes incountered.
func (r *Runer) Line() int {
	return r.line
}

// ReadRune reads the next rune from the reader. EOF is returned if the end of
// the file has been reached.
func (r *Runer) ReadRune() (rune, fault.Fault) {
	ru, err := r.nextRune()

	if ru == '\n' {
		r.line++
	}

	return ru, err
}

// PeekRunes returns the next two runes in the sequence without incrementing the
// 'cursor'. It will check the temp buffer first and if there are NUL, populate
// them with new values from the reader before returning their contents.
func (r *Runer) PeekRunes() (rune, rune, fault.Fault) {
	var err fault.Fault

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
func (r *Runer) nextRune() (rune, fault.Fault) {
	ru := r.buf[0]
	if ru == NUL {
		return r.readRune()
	}

	r.buf[0] = r.buf[1]
	r.buf[1] = NUL
	return ru, nil
}

// readRune reads the next rune from the reader returning EOF if the end of the
// reader has been reached.
func (r *Runer) readRune() (rune, fault.Fault) {
	ru, _, err := r.reader.ReadRune()

	if err == io.EOF {
		return EOF, nil
	}

	if err != nil {
		return NUL, fault.ReaderFault(err.Error())
	}

	return ru, nil
}
