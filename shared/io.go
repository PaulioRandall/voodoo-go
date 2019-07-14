package shared

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines reads in the lines of a file.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// PrintlnWithLineNum pretty prints the line number given a line
// index then prints the string.
func PrintlnWithLineNum(i int, s string) {
	PrintLineNumber(i)
	fmt.Println(s)
}

// PrintLineNumber pretty prints the line number given an index.
func PrintLineNumber(index int) {
	num := index + 1
	out := fmt.Sprintf("%-3d: ", num)
	fmt.Print(out)
}
