package main

import (
	"fmt"

	"github.com/nzdjb/adventofcode2020/util"
)

func tick(room []string) []string {
	result := []string{}
	for i := range room {
		line := ""
		for j := range room[0] {
			switch string(room[i][j]) {
			case ".":
				line += "."
			case "#":
				if checkSurroundings(room, i, j) < 4 {
					line += "#"
				} else {
					line += "L"
				}
			case "L":
				if checkSurroundings(room, i, j) == 0 {
					line += "#"
				} else {
					line += "L"
				}
			default:
				panic(fmt.Errorf("Fail"))
			}
		}
		result = append(result, line)
	}
	return result
}

func checkSurroundings(room []string, i, j int) int {
	count := 0
	// if (i == 0 || i == len(room)-1) && (j == 0 || j == len(room[0])-1) {
	// 	return false
	// }
	di, dj := genOffsets(len(room), i), genOffsets(len(room[0]), j)
	for _, vali := range di {
		for _, valj := range dj {
			if vali == i && valj == j {
				continue
			} else if string(room[vali][valj]) == "#" {
				count++
			}
		}
	}
	return count
}

func genOffsets(length, i int) []int {
	result := []int{}
	switch i {
	case 0:
		result = []int{i, i + 1}
	case length - 1:
		result = []int{i - 1, i}
	default:
		result = []int{i - 1, i, i + 1}
	}
	return result
}

func checkEqual(a, b []string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func countOccupied(room []string) int {
	result := 0
	for i := range room {
		for j := range room[0] {
			if string(room[i][j]) == "#" {
				result++
			}
		}
	}
	return result
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	prev := lines
	ticks := 0
	for true {
		new := tick(prev)
		ticks++
		if checkEqual(prev, new) {
			break
		}
		prev = new
		for _, x := range prev {
			fmt.Println(x)
		}
		fmt.Println("Tick", ticks, countOccupied(new))
		// if ticks > 3 {
		// 	break
		// }
	}
	// fmt.Println(prev)
	fmt.Println("Stable after", ticks, "ticks.")
	fmt.Println(countOccupied(prev))
}
