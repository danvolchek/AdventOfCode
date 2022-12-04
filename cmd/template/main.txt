package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) any {
	return line
}

func solve(lines []any) int {
	for _, line := range lines {

		fmt.Println(line)
	}

	return 0
}

func main() {
	solver := lib.Solver[[]any, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("foo", 1)
	solver.Solve()
}
