package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "8", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

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
		ParseF: lib.ParseLine(lib.AsIs[string]()),
		SolveF: solve,
	}

	solver.Expect(`""`, 4)
	solver.Expect(`"abc"`, 4)
	solver.Expect(`"aaa\"aaa"`, 6)
	solver.Expect(`"\x27"`, 5)
	solver.Verify(input(), 2085)
}
