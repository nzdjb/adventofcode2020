package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

type field struct {
	amin int
	amax int
	bmin int
	bmax int
}

func part1(lines []string) {
	fieldReg := regexp.MustCompile("^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$")
	fields := map[string]field{}
	var i int
	for i = 0; lines[i] != ""; i++ {
		match := fieldReg.FindStringSubmatch(lines[i])
		amin, _ := strconv.Atoi(match[2])
		amax, _ := strconv.Atoi(match[3])
		bmin, _ := strconv.Atoi(match[4])
		bmax, _ := strconv.Atoi(match[5])
		fields[match[1]] = field{amin, amax, bmin, bmax}
	}
	
	i += 2
	// myTicket := strings.Split(lines[i], ",")

	invalids := []int{}
	tickets := [][]int{}
	for i += 2; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		ticket := strings.Split(lines[i], ",")
		iTicket := []int{}
		for _, s := range ticket {
			invalid := true
			f, _ := strconv.Atoi(s)
			for _, field := range fields {
				if ((f >= field.amin && f <= field.amax) || (f >= field.bmin && f <= field.bmax)){
					invalid = false
					break
				}
			}
			if invalid {
				invalids = append(invalids, f)
			}
			iTicket = append(iTicket, f)
		}
		tickets = append(tickets, iTicket)
	}

	sum := 0
	for _, v := range invalids {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	part1(lines)
}