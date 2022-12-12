package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) any {
	return line
}

func solve(lines string) int {
	for j := range lines {
		if j < 3 {
			continue
		}

		var s lib.Set[byte]
		s.Add(lines[j])
		s.Add(lines[j-1])
		s.Add(lines[j-2])
		s.Add(lines[j-3])

		if s.Size() == 4 {
			return j + 1
		}
	}

	panic("asd")
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
