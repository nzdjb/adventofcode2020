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
			panic(fmt.Errorf("Loop detected! acc: %v, pc: %v", acc, pc))
		}
		if pc >= len(program) {
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

func part1(program []instruction) {
	runWrapper(program, "Part1:")
}

func part2(program []instruction) int {
	for ind, ins := range program {
		mutated := make([]instruction, len(program))
		copy(mutated, program)
		switch command := ins.command; command {
		case "jmp":
			mutated[ind] = instruction{"nop", ins.argument}
		case "nop":
			mutated[ind] = instruction{"jmp", ins.argument}
		default:
			continue
		}
		result := runWrapper(mutated, "Part2:")
		if result == 0 {
			continue
		}
		return result
	}
	return -1
}

func runWrapper(program []instruction, prefix string) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(prefix, r)
		}
	}()
	return run(program)
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	var program []instruction
	for _, line := range lines {
		program = append(program, interpret(line))
	}
	part1(program)
	fmt.Println(part2(program))
}
