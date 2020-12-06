package main

import (
	"fmt"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

func contains(haystack []string, needle string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

func groupTotal(group []string) int {
	var groupAnswers []string
	for _, person := range group {
		for _, answer := range person {
			if !contains(groupAnswers, string(answer)) {
				groupAnswers = append(groupAnswers, string(answer))
			}
		}
	}
	return len(groupAnswers)
}

func groupUnionTotal(group []string) int {
	var groupAnswers []string
	for question := 'a'; question <= 'z'; question++ {
		answer := true
		for _, person := range group {
			if !contains(strings.Split(person, ""), string(question)) {
				answer = false
			}
		}
		if answer {
			groupAnswers = append(groupAnswers, string(question))
		}
	}
	return len(groupAnswers)
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	var groups [][]string
	var blob []string
	for _, line := range lines {
		if len(line) == 0 {
			groups = append(groups, blob)
			blob = nil
		} else {
			blob = append(blob, line)
		}
	}
	groups = append(groups, blob)
	total, unionTotal := 0, 0
	for _, group := range groups {
		total += groupTotal(group)
		unionTotal += groupUnionTotal(group)
	}
	fmt.Println(total)
	fmt.Println(unionTotal)
}
