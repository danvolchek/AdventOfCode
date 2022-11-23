package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) int {
	return lib.Atoi(line)
}

var target = 150

func solve(containers []int) int {
	subsets := lib.Filter(lib.Subsets(containers), func(subset []int) bool {
		return lib.SumSlice(subset) == target
	})

	subsetLengths := lib.Map(subsets, func(subset []int) int {
		return len(subset)
	})

	min := lib.MinSlice(subsetLengths)

	return len(lib.Filter(subsetLengths, func(length int) bool { return length == min }))
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	target = 25
	solver.Expect("20\n15\n10\n5\n5", 3)

	target = 150
	solver.Verify(4)
}
