
package interpreter

import (
	"os"
	"bufio"
)

// LoadScroll reads the lines of the scroll and creates a
// new Scroll instance for it.
func LoadScroll(path string) (scroll *Scroll, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	lines, err := scanLines(file)
	if err == nil {
		scroll = newScroll(path, lines)
	}
	
	return
}

// scanLines reads in the lines of an opened file.
func scanLines(file *os.File) ([]string, error) {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Execute runs the voodoo scroll.
func Execute(scroll *Scroll, scrollArgs []string) (exitCode int, err error) {
	
	scroll.JumpToLine(1) // Ignore first line
	executeLines(scroll)
	return 1, nil
}
