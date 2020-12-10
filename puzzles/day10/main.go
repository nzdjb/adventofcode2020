package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/nzdjb/adventofcode2020/util"
)

func part1(adaptors []int) {
	ones := 0
	threes := 0
	adaptors = append([]int{0}, adaptors...)
	for i, val := range adaptors {
		if i == 0 {
			continue
		}
		switch diff := val - adaptors[i-1]; diff {
		case 1:
			ones++
		case 3:
			threes++
		default:
			fmt.Println(adaptors[i], val)
			panic(fmt.Errorf("chain broken"))
		}
	}
	threes++
	fmt.Println("Ones:", ones, "Threes:", threes, "Product:", ones*threes)
}

func part2(adaptors []int) {
	adaptors = append([]int{0}, adaptors...)
	paths := []int{1}
	for i := range adaptors {
		n, j, iplus, x := 0, 1, i + 1, 0
		if iplus >= len(adaptors) {
			x = adaptors[len(adaptors) - 1] + 3
		} else {
			x = adaptors[iplus]
		}
		for j <= 3 {
			if iplus - j >= 0 && x - adaptors[iplus - j] <= 3 {
				n += paths[iplus - j]
			}
			j++
		}
		paths = append(paths, n)
	}
	fmt.Println(paths[len(paths) - 1])
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	adaptors := []int{}
	for _, line := range lines {
		joltage, _ := strconv.Atoi(line)
		adaptors = append(adaptors, joltage)
	}
	sort.Ints(adaptors)
	part1(adaptors)
	part2(adaptors)
}
