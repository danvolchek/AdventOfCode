package main

import (
	"regexp"

	"github.com/danvolchek/AdventOfCode/lib"
)

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
		ParseF: lib.AsIs,
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

	solver.Verify(119433)
}
