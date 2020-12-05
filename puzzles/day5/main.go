package main

import (
	"fmt"
	"sort"
	"errors"

	"github.com/nzdjb/adventofcode2020/util"
)

func seatID(row, col int) int {
	return row*8 + col
}

func binarySearch(min, max int, remainder string) int {
	if min == max {
		return min
	}
	if remainder[:1] == "F" || remainder[:1] == "L" {
		return binarySearch(min, min+(max-min)/2, remainder[1:])
	}
	return binarySearch(min+(max-min)/2+1, max, remainder[1:])
}

func findGap(seats []int) int {
	sort.Ints(seats)
	for i, seat := range seats {
		if seats[i+1] != seat+1 {
			return seat + 1
		}
	}
	panic(errors.New("Couldn't find gap."))
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	var seats []int
	max := 0
	for _, line := range lines {
		row := binarySearch(0, 127, line[:7])
		col := binarySearch(0, 7, line[7:])
		seat := seatID(row, col)
		if seat > max {
			max = seat
		}
		seats = append(seats, seat)
	}
	fmt.Println(max)
	fmt.Println(findGap(seats))
}
