package main

import (
	"fmt"

	"github.com/nzdjb/adventofcode2020/util"
)

func tick(room []string, checkFunction func([]string, int, int) int, threshold int) []string {
	result := []string{}
	for i := range room {
		line := ""
		for j := range room[0] {
			switch string(room[i][j]) {
			case ".":
				line += "."
			case "#":
				if checkFunction(room, i, j) < threshold {
					line += "#"
				} else {
					line += "L"
				}
			case "L":
				if checkFunction(room, i, j) == 0 {
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

func checkSurroundings2(room []string, i, j int) int {
	count := 0
	seats := findVisibleSeats(room, i, j)
	for _, seat := range seats {
		if seat == "#" {
			count++
		}
	}
	return count
}

func findVisibleSeats(room []string, i, j int) []string {
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	seats := []string{}
	for _, direction := range directions {
		di, dj := i, j
		for true {
			di += direction[0]
			dj += direction[1]
			if !checkInsideBounds(di, dj, len(room), len(room[i])) {
				break
			} else if string(room[di][dj]) == "." {
				continue
			} else {
				seats = append(seats, string(room[di][dj]))
				break
			}
		}
	}
	return seats
}

func checkInsideBounds(x, y, lenx, leny int) bool {
	return x >= 0 && y >= 0 && x < lenx && y < leny
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

func part1(room []string) {
	prev := room
	ticks := 0
	for true {
		new := tick(prev, checkSurroundings, 4)
		ticks++
		if checkEqual(prev, new) {
			break
		}
		prev = new
	}
	fmt.Println("Stable after", ticks, "ticks.")
	fmt.Println(countOccupied(prev))
}

func part2(room []string) {
	prev := room
	ticks := 0
	for true {
		new := tick(prev, checkSurroundings2, 5)
		ticks++
		if checkEqual(prev, new) {
			break
		}
		prev = new
	}
	fmt.Println("Stable after", ticks, "ticks.")
	fmt.Println(countOccupied(prev))
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	part1(lines)
	part2(lines)
}
