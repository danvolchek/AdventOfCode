package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "5", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func find(line string, substr string, not int) bool {
	line = line[0:not] + strings.Repeat("*", len(substr)) + line[not+len(substr):]

	return strings.Contains(line, substr)
}

func hasRepeat(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if find(line, line[i:i+2], i) {
			return true
		}
	}

	return false
}

func hasBetween(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}

	return false
}

func solve(lines []string) int {
	nice := 0

	isNice := func(line string) bool {
		return hasRepeat(line) && hasBetween(line)
	}

	for _, line := range lines {
		if isNice(line) {
			nice += 1
		}
	}

	return nice
}

func main() {
	solver := lib.Solver[[]string, int]{
		ParseF: lib.ParseLine(lib.AsIs[string]()),
		SolveF: solve,
	}

	solver.Expect("qjhvhtzxzqqjkmpb", 1)
	solver.Expect("xxyxx", 1)
	solver.Expect("uurcxstgmygtbstg", 0)
	solver.Expect("ieodomkazucvgmuy", 0)
	solver.Solve(input())
}
