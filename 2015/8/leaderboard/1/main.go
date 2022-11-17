package main

import (
	"regexp"
	"strings"

	"github.com/danvolchek/AdventOfCode/lib"
)

func decode(line string) string {
	line = line[1 : len(line)-1]
	line = strings.ReplaceAll(line, `\\`, `\`)
	line = strings.ReplaceAll(line, `\"`, `"`)
	line = regexp.MustCompile(`\\x[0-9a-f][0-9a-f]`).ReplaceAllString(line, "f")

	return line
}

func solve(lines []string) int {
	total := 0

	for _, line := range lines {
		total += len(line) - len(decode(line))
	}

	return total
}

func main() {
	solver := lib.Solver[[]string, int]{
		ParseF: lib.ParseLine(lib.AsIs),
		SolveF: solve,
	}

	solver.Expect(`""`, 2)
	solver.Expect(`"abc"`, 2)
	solver.Expect(`"aaa\"aaa"`, 3)
	solver.Expect(`"\x27"`, 5)
	solver.Expect(`"\x27"`, 5)
	solver.Verify(1350)
}
