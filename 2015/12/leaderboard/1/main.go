package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"regexp"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "12", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

var numberRegexp = regexp.MustCompile(`-?[0-9]+`)

func solve(input string) int {
	matches := numberRegexp.FindAllString(input, -1)

	sum := 0

	for _, match := range matches {
		sum += lib.Atoi(match)
	}

	return sum
}

func main() {
	solver := lib.Solver[string, int]{
		ParseF: lib.ParseString,
		SolveF: solve,
	}

	solver.Expect(`[1,2,3]`, 6)
	solver.Expect(`{"a":2,"b":4}`, 6)
	solver.Expect(`[[[3]]]`, 3)
	solver.Expect(`{"a":{"b":4},"c":-1}`, 3)
	solver.Expect(`{"a":[-1,1]}`, 0)
	solver.Expect(`[-1,{"a":1}]`, 0)
	solver.Expect(`[]`, 0)
	solver.Expect(`{}`, 0)

	solver.Verify(input(), 119433)
}
