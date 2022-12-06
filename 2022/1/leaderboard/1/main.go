package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) int {
	if line == "" {
		return -1
	}

	return lib.Atoi(line)
}

func solve(lines []int) int {
	max := 0
	sum := 0
	for _, line := range lines {

		if line != -1 {
			sum += line
		} else {
			if sum > max {
				max = sum
			}
			sum = 0
		}
	}

	return max
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", 24000)
	solver.Verify(71471)
}
