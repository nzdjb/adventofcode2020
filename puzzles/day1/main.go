package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/nzdjb/adventofcode2020/util"
)

func scanFileToIntSlice(filename string) []int {
  file, err := os.Open(filename)
	util.Check(err)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := scanner.Text()
		i, err := strconv.Atoi(entry)
		util.Check(err)
		lines = append(lines, i)
	}
        
  return lines
}

func outputResult(nums ...int) {
	sum := 0
	product := 1
	for _, num := range nums {
		sum += num
		product *= num
	}
	fmt.Printf("%v: sum: %v product: %v \n", nums, sum, product) 
}

func findTwo(lines []int) {
  for _, v := range lines {
    diff := 2020 - v
    match := util.Include(lines, diff)
    if match {
      outputResult(v, diff)
      return
    }
  }
}

func findThree(lines []int) {
  for _, i := range lines {
    for _, j := range lines {
      diff := 2020 - i - j
      match := util.Include(lines, diff)
      if match {
        outputResult(i, j, diff)
        return
      }
    }
  }
}

func main() {
  lines := scanFileToIntSlice("./input.txt")

	findTwo(lines)
	findThree(lines)
}
