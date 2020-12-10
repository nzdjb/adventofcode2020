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

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	adaptors := []int{}
	for _, line := range lines {
		joltage, _ := strconv.Atoi(line)
		adaptors = append(adaptors, joltage)
	}
	sort.Ints(adaptors)
	part1(adaptors)
}
