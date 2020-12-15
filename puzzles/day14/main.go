package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

func part1(lines []string) {
	memory := map[int]string{}
	var mask [36]string
	for _, line := range lines {
		command, args := interpret(line)
		switch command {
		case "mask":
			m := strings.Split(args[0], "")
			for i, s := range m {
				mask[i] = s
			}
		case "mem":
			address, _ := strconv.Atoi(args[0])
			ba := intStringToBinaryArray(args[1])
			// fmt.Println(ba)
			masked := applyMask(mask, ba)
			memory[address] = binaryArrayToIntString(masked)
		default:
			panic(fmt.Errorf("bad instruction"))
		}
	}
	result := 0
	for _, v := range memory {
		vi, _ := strconv.Atoi(v)
		result += vi
	}
	fmt.Println(result)
}

func part2(lines []string) {
	memory := map[int]string{}
	var mask [36]string
	for _, line := range lines {
		command, args := interpret(line)
		switch command {
		case "mask":
			m := strings.Split(args[0], "")
			for i, s := range m {
				mask[i] = s
			}
		case "mem":
			ba := intStringToBinaryArray(args[1])
			addresses := addressMask(mask, args[0])
			for _, a := range addresses {
				memory[a] = binaryArrayToIntString(ba)
			}
		default:
			panic(fmt.Errorf("bad instruction"))
		}
	}
	result := 0
	for _, v := range memory {
		vi, _ := strconv.Atoi(v)
		result += vi
	}
	fmt.Println(result)
}

func intStringToBinaryArray(num string) [36]string {
	dec, _ := strconv.Atoi(num)
	var result [36]string
	i := 35
	for i >= 0 {
		result[i] = strconv.Itoa(dec & 1)
		dec = dec >> 1
		i--
	}
	return result
}

func binaryArrayToIntString(bin [36]string) string {
	result := 0
	for _, v := range bin {
		vi, _ := strconv.Atoi(v)
		result = result << 1
		result = result | vi
	}
	return strconv.Itoa(result)
}

func applyMask(mask, arg [36]string) [36]string {
	var result [36]string
	for i, v := range mask {
		if v == "X" {
			result[i] = arg[i]
		} else {
			result[i] = v
		}
	}
	return result
}

func addressMask(mask [36]string, arg string) []int {
	result := []int{}
	var ones [36]string
	argba := intStringToBinaryArray(arg)
	for i, m := range mask {
		if m == "0" {
			ones[i] = argba[i]
		} else {
			ones[i] = m
		}
	}
	todo := [][36]string{ones}
	for len(todo) > 0 {
		t := todo[0]
		todo = todo[1:]
		for i, v := range t {
			if v == "X" {
				one := maskBit(t, i, "1")
				zero := maskBit(t, i, "0")
				todo = append(todo, one, zero)
				break
			}
			r, _ := strconv.Atoi(binaryArrayToIntString(t))
			result = append(result, r)
		}
	}
	return result
}

func maskBit(t [36]string, bit int, val string) [36]string {
	num := t
	num[bit] = val
	return num
}

func interpret(line string) (string, []string) {
	reg := regexp.MustCompile(`^(?:(?:(mask) = ([01X]+))|(?:(mem)\[([0-9]+)\] = ([0-9]+)))$`)
	matches := reg.FindStringSubmatch(line)
	if matches[1] == "mask" {
		return "mask", []string{matches[2]}
	} else {
		return "mem", []string{matches[4], matches[5]}
	}
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	part1(lines)
	part2(lines)
}
