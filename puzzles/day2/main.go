package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nzdjb/adventofcode2020/util"
)

type entry struct {
	min      int
	max      int
	letter   rune
	password string
}

func newEntry(line string) *entry {
	reg := getRegexp()
	matches := reg.FindStringSubmatch(line)
	min, err := strconv.Atoi(matches[1])
	util.Check(err)
	max, err := strconv.Atoi(matches[2])
	util.Check(err)
	letter := []rune(matches[3])[0]
	password := matches[4]

	return &entry{min, max, letter, password}
}

var reg *regexp.Regexp

func getRegexp() *regexp.Regexp {
	if reg == nil {
		reg = regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
	}
	return reg
}

func checkPasswordPart1(entry entry) bool {
	count := 0
	for _, p := range entry.password {
		if p == entry.letter {
			count++
		}
	}
	return entry.min <= count && count <= entry.max
}

func checkPasswordPart2(entry entry) bool {
	rPassword := []rune(entry.password)
	return checkPasswordLocation(rPassword, entry.letter, entry.min) + checkPasswordLocation(rPassword, entry.letter, entry.max) == 1
}

func checkPasswordLocation(password []rune, letter rune, location int) int {
	if password[location - 1] == letter {
		return 1
	}
	return 0
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	okPart1, okPart2 := 0, 0
	for _, line := range lines {
		entry := newEntry(line)
		if checkPasswordPart1(*entry) {
			okPart1++
		}
		if checkPasswordPart2(*entry) {
			okPart2++
		}
	}
	fmt.Println(okPart1)
	fmt.Println(okPart2)
}
