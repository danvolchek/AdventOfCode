package main

import (
	"strings"

	"github.com/danvolchek/AdventOfCode/lib"
)

func encode(line string) string {
	line = strings.ReplaceAll(line, `\`, `\\`)
	line = strings.ReplaceAll(line, `"`, `\"`)

	return `"` + line + `"`
}

func solve(lines []string) int {
	total := 0

	for _, line := range lines {
		total += len(encode(line)) - len(line)
	}

	return total
}

func main() {
	solver := lib.Solver[[]string, int]{
		ParseF: lib.ParseLine(lib.AsIs),
		SolveF: solve,
	}

	solver.Expect(`""`, 4)
	solver.Expect(`"abc"`, 4)
	solver.Expect(`"aaa\"aaa"`, 6)
	solver.Expect(`"\x27"`, 5)
	solver.Verify(2085)
}
