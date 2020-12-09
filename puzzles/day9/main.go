package main

import (
	"fmt"
	"strconv"

	"github.com/nzdjb/adventofcode2020/util"
)

func checkValid(x int, candidates []int) bool {
	for _, i := range candidates {
		for _, j := range candidates {
			if i != j && i+j == x {
				return true
			}
		}
	}
	return false
}

func min(nums []int) int {
	result := nums[0]
	for _, x := range nums[1:] {
		if x < result {
			result = x
		}
	}
	return result
}

func max(nums []int) int {
	result := nums[0]
	for _, x := range nums[1:] {
		if x > result {
			result = x
		}
	}
	return result
}

func sum(nums []int) int {
	result := 0
	for _, x := range nums {
		result += x
	}
	return result
}

func part1(nums []int) int {
	preamble := 25
	for i, x := range nums {
		if i >= preamble {
			if !checkValid(x, nums[i-preamble:i]) {
				return x
			}
		}
	}
	return -1
}

func part2(target int, nums []int) int {
	result := []int{}
	start, end := 0, 0
	for end < len(nums) {
		x := nums[start:end]
		if sum(x) == target {
			result = x
			break
		} else if sum(x) > target {
			start++
		} else {
			end++
		}
	}
	return min(result) + max(result)
}

func main() {
	lines := util.ScanFileToStringSlice("./input.txt")
	nums := []int{}
	for _, x := range lines {
		xInt, _ := strconv.Atoi(x)
		nums = append(nums, xInt)
	}
	result := part1(nums)
	fmt.Println(result)
	fmt.Println(part2(result, nums))
}
