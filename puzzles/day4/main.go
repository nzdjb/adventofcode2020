package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

func isValid(entry []string) bool {
	m := map[string]string{}
	for _, chunk := range entry {
		s := strings.Split(chunk, ":")
		m[s[0]] = s[1]
	}
	hasFields := (checkField("ecl", m) && checkField("pid", m) && checkField("eyr", m) && checkField("hcl", m) && checkField("byr", m) && checkField("iyr", m) && checkField("hgt", m))
	if !hasFields {
		return false
	}
	return checkByr(m["byr"]) && checkIyr(m["iyr"]) && checkEyr(m["eyr"]) && checkHgt(m["hgt"]) && checkHcl(m["hcl"]) && checkEcl(m["ecl"]) && checkPid(m["pid"])
}

func checkByr(val string) bool {
	matched, _ := regexp.Match("^[0-9]{4}$", []byte(val))
	return matched && val >= "1920" && val <= "2002"
}

func checkIyr(val string) bool {
	matched, _ := regexp.Match("^[0-9]{4}$", []byte(val))
	return matched && val >= "2010" && val <= "2020"
}

func checkEyr(val string) bool {
	matched, _ := regexp.Match("^[0-9]{4}$", []byte(val))
	return matched && val >= "2020" && val <= "2030"
}

func checkHgt(val string) bool {
	reg := regexp.MustCompile("^([0-9]{2,3})(cm|in)$")
	matches := reg.FindStringSubmatch(val)
	if matches == nil {
		return false
	}
	if matches[2] == "cm" {
		return matches[1] >= "150" && matches[1] <= "193"
	} else {
		return matches[1] >= "59" && matches[1] <= "76"
	}
}

func checkHcl(val string) bool {
	matched, _ := regexp.Match("^#[0-9a-f]{6}$", []byte(val))
	return matched
}

func checkEcl(val string) bool {
	matched, _ := regexp.Match("^(?:amb|blu|brn|gry|grn|hzl|oth)$", []byte(val))
	return matched
}

func checkPid(val string) bool {
	matched, _ := regexp.Match("^[0-9]{9}$", []byte(val))
	return matched
}

func checkField(needle string, m map[string]string) bool {
	_, prs := m[needle]
	return prs
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	count := 0
	var possibles [][]string
	var blob []string
	for count < len(lines) {
		if len(lines[count]) == 0 {
			possibles = append(possibles, blob)
			blob = nil
		} else {
			for _, entry := range strings.Fields(lines[count]) {
				blob = append(blob, entry)
			}
		}
		count++
	}
	possibles = append(possibles, blob)
	result := 0
	for _, entry := range possibles {
		if isValid(entry) {
			result++
		}
	}
	fmt.Println(result)
}
