package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) string {
	return line
}

func solve(lines []string) int {
	sum := 0

	for _, line := range lines {

		first := line[:len(line)/2]
		second := line[len(line)/2:]

	outer:
		for _, f := range first {
			for _, s := range second {
				if f == s {
					if f >= 'a' && f <= 'z' {
						sum += int(f-'a') + 1
					} else {
						sum += int(f-'A') + 27
					}

					break outer
				}
			}
		}
	}

	return sum
}

func main() {
	solver := lib.Solver[[]string, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 157)
	solver.Solve()
}
