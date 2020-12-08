package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

type instruction struct {
	command  string
	argument int
}

func interpret(line string) instruction {
	splitLine := strings.Split(line, " ")
	value, _ := strconv.Atoi(splitLine[1])
	return instruction{splitLine[0], value}
}

func contains(needle int, haystack []int) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

func run(program []instruction) int {
	var visited []int
	acc := 0
	pc := 0

	for true {
		if contains(pc, visited) {
			return acc
		}
		visited = append(visited, pc)
		switch command := program[pc].command; command {
		case "jmp":
			pc += program[pc].argument
		case "acc":
			acc += program[pc].argument
			pc++
		case "nop":
			pc++
		default:
			panic(fmt.Errorf("Unrecognised instruction! %v at %v", command, pc))
		}
	}
	return -1
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	var program []instruction
	for _, line := range lines {
		program = append(program, interpret(line))
	}
	fmt.Println(run(program))
}
