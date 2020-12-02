package main

import (
	"fmt"
	"github.com/nzdjb/adventofcode2020/util"
	"regexp"
	"strconv"
)

type entry struct {
	min int
	max int
	letter rune
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
	password :=  matches[4]

	return &entry { min, max, letter, password }
}

var reg *regexp.Regexp
func getRegexp() *regexp.Regexp {
	if(reg == nil) {
		reg = regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
	}
	return reg
}

func checkPassword(entry entry) bool {
	count := 0
	for _, p := range entry.password {
		if(p == entry.letter) {
			count++
		}
	}
	return entry.min <= count && count <= entry.max
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	ok := 0
	for _, line := range lines {
		entry := newEntry(line)
		if(checkPassword(*entry)) {
			ok++
		}
	}
	fmt.Println(ok)
}