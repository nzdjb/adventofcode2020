package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Index(haystack []int, needle int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

func Include(haystack []int, needle int) bool {
	return Index(haystack, needle) >= 0
}

func ScanFileToIntSlice(filename string) []int {
        file, err := os.Open(filename)
	check(err)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		i, err := strconv.Atoi(entry)
		check(err)
		lines = append(lines, i)
	}
        
        return lines
}

func OutputResult(nums ...int) {
	sum := 0
	product := 1
	for _, num := range nums {
		sum += num
		product *= num
	}
	fmt.Printf("%v: sum: %v product: %v \n", nums, sum, product) 
}

func FindTwo(lines []int) {
        for _, v := range lines {
                diff := 2020 - v
                match := Include(lines, diff)
                if match {
                        OutputResult(v, diff)
                        return
                }
        }
}

func FindThree(lines []int) {
        for _, i := range lines {
                for _, j := range lines {
                        diff := 2020 - i - j
                        match := Include(lines, diff)
                        if match {
                                OutputResult(i, j, diff)
                                return
                        }
                }
        }
}

func main() {
        lines := ScanFileToIntSlice("./input.txt")

	FindTwo(lines)
	FindThree(lines)
}
