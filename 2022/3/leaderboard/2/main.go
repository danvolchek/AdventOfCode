package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) string {
	return line
}

func solve(lines []string) int {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]

	outer:
		for _, f := range first {
			for _, s := range second {
				for _, t := range third {
					if f == s && s == t {
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
	}

	return sum
}

func main() {
	solver := lib.Solver[[]string, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 70)
	solver.Solve()
}
