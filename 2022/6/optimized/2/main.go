package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

const numUnique = 14

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

	solver.Expect("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19)
	solver.Expect("bvwbjplbgvbhsrlpgdmjqwftvncz", 23)
	solver.Expect("nppdvjthqldpwncqszvftbrmjlhg", 23)
	solver.Expect("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29)
	solver.Expect("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26)

	solver.Verify(3613)
}
