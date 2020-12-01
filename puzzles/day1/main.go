package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Index(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func Include(haystack []int, needle int) bool {
	return Index(haystack, needle) >= 0
}

func ScanFileToIntSlice(filename string) []int {
        file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		i, err := strconv.Atoi(entry)
		check(err)
		lines = append(lines, i)
	}
        
        return lines
}

func main() {
        lines := ScanFileToIntSlice("./input.txt")

	for _, v := range lines {
		diff := 2020 - v
		match := Include(lines, diff)
		if match {
			product := v * diff
			fmt.Printf("%v: 2020 - %v = %v, match = %t, product = %v \n", v, v, diff, match, product)
			return
		}
	}
}
