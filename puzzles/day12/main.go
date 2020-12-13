package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nzdjb/adventofcode2020/util"
)

func parseInstruction(instruction string) (string, int) {
	reg := regexp.MustCompile("^([NSEWLRF])([0-9]+)$")
	matches := reg.FindStringSubmatch(instruction)
	quantity, _ := strconv.Atoi(matches[2])
	return matches[1], quantity
}

func enactInstruction(instruction string, quantity, x, y, bearing int) (int, int, int) {
	newX, newY, newBearing := x, y, bearing
	switch instruction {
	case "N":
		newX = x + quantity
	case "S":
		newX = x - quantity
	case "E":
		newY = y + quantity
	case "W":
		newY = y - quantity
	case "L":
		newBearing = (bearing - quantity) % 360
		if newBearing < 0 {
			newBearing += 360
		}
	case "R":
		newBearing = (bearing + quantity) % 360
		if newBearing < 0 {
			newBearing += 360
		}
	case "F":
		switch bearing {
		case 0:
			newX = x + quantity
		case 90:
			newY = y + quantity
		case 180:
			newX = x - quantity
		case 270:
			newY = y - quantity
		default:
			panic(fmt.Errorf("i wasn't prepared for that, bearing of %v", bearing))
		}
	default:
		panic(fmt.Errorf("bad instruction"))
	}
	return newX, newY, newBearing
}

func manhattan(x, y, dx, dy int) int {
	return util.Abs(dx) - util.Abs(x) + util.Abs(dy) - util.Abs(y)
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	x, y, bearing := 0, 0, 90
	for _, line := range lines {
		instruction, quantity := parseInstruction((line))
		x, y, bearing = enactInstruction(instruction, quantity, x, y, bearing)
	}
	fmt.Println(manhattan(0, 0, x, y))
}
