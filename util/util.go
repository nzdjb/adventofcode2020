package util

import (
	"os"
	"bufio"
)

// Check Panic if errror.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// ScanFileToStringSlice Read a file and return its rows as a slice.
func ScanFileToStringSlice(filename string) []string {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
        
  return lines
}