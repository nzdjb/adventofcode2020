package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

func parseLine(line string) []string {
	reg := regexp.MustCompile("^([a-z]+ [a-z]+) bags contain (no other bags|(?:[0-9a-z, ])+).$")
	return reg.FindStringSubmatch(line)[1:]
}

func findContained(rules map[string]string, needle string) []string {
	var result []string
	for k, v := range rules {
		matched, _ := regexp.MatchString(needle, v)
		if matched {
			result = append(result, k)
		}
	}
	return result
}

func stringContains(needle string, haystack []string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

func reverseSearch(rules map[string]string, target string) []string {
	queue := []string{"shiny gold"}
	searched := []string{}
	for len(queue) > 0 {
		var searchBag string
		searchBag, queue = queue[0], queue[1:]
		for _, bag := range findContained(rules, searchBag) {
			if !stringContains(bag, searched) {
				searched = append(searched, bag)
				queue = append(queue, bag)
			}
		}
	}
	return searched
}

func findContentCount(rules map[string]string, bag string) int {
	if rules[bag] == "no other bags" {
		return 1
	}
	reg := regexp.MustCompile("^([0-9]) ([a-z]+ [a-z]+) bags?$")
	total := 0
	for _, child := range strings.Split(rules[bag], ", ") {
		details := reg.FindStringSubmatch(child)
		mult, err := strconv.Atoi(details[1])
		util.Check(err)
		total += mult * findContentCount(rules, details[2])
	}
	return 1 + total
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	rules := make(map[string]string)
	for _, line := range lines {
		result := parseLine(line)
		rules[result[0]] = result[1]
	}
	fmt.Println(len(reverseSearch(rules, "shiny gold")))
	fmt.Println(findContentCount(rules, "shiny gold") - 1)
}
