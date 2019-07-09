
package interpreter

import (
	"os"
	"bufio"
	
	sc "github.com/PaulioRandall/voodoo-go/scroll"
	ex "github.com/PaulioRandall/voodoo-go/executors"
	sh "github.com/PaulioRandall/voodoo-go/shared"
)

// LoadScroll reads the lines of the scroll and creates a
// new Scroll instance for it.
func LoadScroll(path string) (scroll *sc.Scroll, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	lines, err := scanLines(file)
	if err == nil {
		scroll = sc.NewScroll(path, lines)
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

// Execute runs a voodoo scroll.
func Execute(scroll *sc.Scroll, scrollArgs []string) (code sh.ExitCode, exErr sh.ExeError) {
	
	var exe ex.Executor
	exe = ex.NewScrollExecutor()
	
	ignoreFirstLine(scroll)
	
	for scroll.Next() {
		stat := sc.Statement{
			Val: scroll.Code,
			Row: scroll.Line,
			Col: 1,
		}
		
		code, exe, exErr = exe.Exe(scroll, stat)
		if code != sh.OK {
			return
		}
	}
	
	return
}

// ignoreFirstLine skips the first line of a scroll so shebang's can
// be used easily.
func ignoreFirstLine(scroll *sc.Scroll) {
	scroll.JumpToLine(1)
}

