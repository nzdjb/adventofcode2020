package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bytes"

	"github.com/nzdjb/adventofcode2020/util"
)

func copy(src, dest string) {
	copyAndSub(src, dest, "", "")
}

func copyAndSub(src, dest, needle, replacement string) {
	data, err := ioutil.ReadFile(src)
	util.Check(err)
	data = bytes.ReplaceAll(data, []byte(needle), []byte(replacement))
	err = ioutil.WriteFile(dest, data, 0644)
	util.Check(err)
}

func main() {
	args := os.Args[1:]
	newDir := fmt.Sprintf("%v/puzzles/day%v/", args[0], args[1])
	templateDir := fmt.Sprintf("%v/gen/templates/", args[0])
	err := os.Mkdir(newDir, 0755)
	util.Check(err)
	copyAndSub(templateDir+"go.mod", newDir+"go.mod", "###", args[1])
	copy(templateDir+"main.go", newDir+"main.go")
}
