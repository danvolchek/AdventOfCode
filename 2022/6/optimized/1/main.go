package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

const numUnique = 4

func solve(input string) int {
	for index := range input {
		if index < numUnique-1 {
			continue
		}

		var s lib.Set[byte]
		for i := 0; i < numUnique; i++ {
			s.Add(input[index-i])
		}

		if s.Size() == numUnique {
			return index + 1
		}
	}

	panic("not found")
}

func main() {
	solver := lib.Solver[string, int]{
		ParseF: lib.AsIs,
		SolveF: solve,
	}

	solver.Expect("bvwbjplbgvbhsrlpgdmjqwftvncz", 5)
	solver.Expect("nppdvjthqldpwncqszvftbrmjlhg", 6)
	solver.Expect("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10)

	solver.Verify(1640)
}
