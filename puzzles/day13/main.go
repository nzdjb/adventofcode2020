package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

func part1(start int, busString string) {
	busses := []int{}
	for _, bus := range strings.Split(busString, ","){
		if bus == "x" {
			continue
		}
		iBus, _ := strconv.Atoi(bus)
		busses = append(busses, iBus)
	}
	bestBus := 0
	smallestWait := math.MaxInt64
	for _, bus := range busses {
		wait := bus - (start % bus)
		if wait < smallestWait {
			smallestWait = wait
			bestBus = bus
		}
	}
	fmt.Println("best bus", bestBus, "wait", smallestWait, "product", bestBus * smallestWait)
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	start, _ := strconv.Atoi(lines[0])
	part1(start, lines[1])
}
