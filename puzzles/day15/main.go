package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nzdjb/adventofcode2020/util"
)

func rfind(haystack []int) int {
	needle := haystack[len(haystack)-1]
	for i := len(haystack) - 2; i >= 0; i-- {
		if haystack[i] == needle {
			return len(haystack) - i -1
		}
	}
	return 0
}

func calcNaive(nums []int, target int) int {
	i := 0
	for len(nums) < target {
		if i > 0 && i % 100000 == 0 {
			fmt.Println("Progress:", i)
		}
		nums = append(nums, rfind(nums))
		i++
	}
	return nums[len(nums)-1]
}

func calcMap(nums []int, target int) int {
	m := map[int]int{}
	for i, v := range nums {
		m[v] = i + 1
	}
	last := nums[len(nums) - 1]
	for i := len(nums); i < target; i++ {
		next, found := m[last]
		if found {
			next = i - next
		}
		m[last] = i
		last = next
	}
	return last
}

func main() {
	input := util.ScanFileToStringSlice("./input.txt")[0]
	nums := []int{}
	for _, v := range strings.Split(input, ",") {
		vi, _ := strconv.Atoi(v)
		nums = append(nums, vi)
	}
	fmt.Println(calcMap(nums, 2020))
	fmt.Println(calcMap(nums, 30000000))
}
