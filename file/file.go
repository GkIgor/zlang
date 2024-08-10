package file

import (
	"bufio"
	"os"
)

func ReadFiles(filename string) []string {
	var lines []string

	input, err := os.Open(filename)

	if err != nil {
		panic("Error opening file")
	}

	scan := bufio.NewScanner(input)

	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	defer input.Close()

	return lines
}
