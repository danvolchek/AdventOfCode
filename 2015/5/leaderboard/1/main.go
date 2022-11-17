package main

import (
	"strings"
	"unicode/utf8"

	"github.com/danvolchek/AdventOfCode/lib"
)

func countVowels(line string) int {
	vowels := 0

	for _, r := range line {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowels += 1
		}
	}

	return vowels
}

func hasDouble(line string) bool {
	curr, _ := utf8.DecodeRune([]byte(line[0:1]))

	for _, next := range line[1:] {
		if curr == next {
			return true
		}

		curr = next
	}

	return false
}

func doesNotHave(line string, bad []string) bool {
	for _, badString := range bad {
		if strings.Contains(line, badString) {
			return false
		}
	}

	return true
}

func solve(lines []string) int {
	nice := 0

	isNice := func(line string) bool {
		return countVowels(line) >= 3 && hasDouble(line) && doesNotHave(line, []string{"ab", "cd", "pq", "xy"})
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
		ParseF: lib.ParseLine(lib.AsIs),
		SolveF: solve,
	}

	solver.Expect("tugknbfddgicrmopn", 1)
	solver.Expect("taaa", 1)
	solver.Expect("jchzalrnumimnmhp", 0)
	solver.Expect("haegwjzuvuyypxyu", 0)
	solver.Expect("dvszwmarrgswjxmb", 0)
	solver.Verify(258)
}
