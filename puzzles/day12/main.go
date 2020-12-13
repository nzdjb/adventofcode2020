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

func enactInstruction2(instruction string, quantity, shipX, shipY, wayX, wayY int) (int, int, int, int) {
	newShipX, newShipY, newWayX, newWayY := shipX, shipY, wayX, wayY
	switch instruction {
	case "N":
		newWayY = wayY + quantity
	case "S":
		newWayY = wayY- quantity
	case "E":
		newWayX = wayX + quantity
	case "W":
		newWayX = wayX - quantity
	case "L":
		switch quantity {
		case 90:
			newWayX = wayY * -1
			newWayY = wayX
		case 180:
			newWayX = wayX * -1
			newWayY = wayY * -1
		case 270:
			newWayX = wayY
			newWayY = wayX * -1
		default:
			panic(fmt.Errorf("bad angle %v", quantity))
		}
	case "R":
		switch quantity {
		case 90:
			newWayX = wayY
			newWayY = wayX * -1
		case 180:
			newWayX = wayX * -1
			newWayY = wayY * -1
		case 270:
			newWayX = wayY * -1
			newWayY = wayX
		default:
			panic(fmt.Errorf("bad angle %v", quantity))
		}
	case "F":
		newShipX = shipX + wayX*quantity
		newShipY = shipY + wayY*quantity
	default:
		panic(fmt.Errorf("bad instruction"))
	}
	fmt.Println(newShipX, newShipY, newWayX, newWayY)
	return newShipX, newShipY, newWayX, newWayY
}

func manhattan(x, y, dx, dy int) int {
	return util.Abs(dx) - util.Abs(x) + util.Abs(dy) - util.Abs(y)
}

func part1(lines []string) {
	x, y, bearing := 0, 0, 90
	for _, line := range lines {
		instruction, quantity := parseInstruction((line))
		x, y, bearing = enactInstruction(instruction, quantity, x, y, bearing)
	}
	fmt.Println("part1:", manhattan(0, 0, x, y))
}

func part2(lines []string) {
	shipX, shipY, wayX, wayY := 0, 0, 10, 1
	for _, line := range lines {
		instruction, quantity := parseInstruction((line))
		shipX, shipY, wayX, wayY = enactInstruction2(instruction, quantity, shipX, shipY, wayX, wayY)
	}
	fmt.Println("part2:", manhattan(0, 0, shipX, shipY))
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	part1(lines)
	part2(lines)
}
