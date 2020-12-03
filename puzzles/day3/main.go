package main

import (
	"fmt"

	"github.com/nzdjb/adventofcode2020/util"
)

func trySlope(m []string, dx, dy int) int {
	x, y, count := 0, 0, 0
	for y <= len(m)-1 {
		if m[y][x] == '#' {
			count++
		}
		y += dy
		x = (x + dx) % len(m[0])
	}
	return count
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	fmt.Println("Soln 1:", trySlope(lines, 3, 1))
	product := 1
	product *= trySlope(lines, 1, 1)
	product *= trySlope(lines, 3, 1)
	product *= trySlope(lines, 5, 1)
	product *= trySlope(lines, 7, 1)
	product *= trySlope(lines, 1, 2)
	fmt.Println("Soln 2:", product)
}
