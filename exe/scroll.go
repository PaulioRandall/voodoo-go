package exe

import (
	"io/ioutil"
)

// TODO: This pkg is redundant, move the preety print code to the parser pkg

// Scroll represents a scroll.
type Scroll struct {
	File string // File path to the scroll
	Data string // File data from the scroll
}

// LoadScroll loads the data from a scroll.
func LoadScroll(f string) (*Scroll, error) {
	d, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	sc := &Scroll{
		File: f,
		Data: string(d),
	}

	return sc, nil
}
