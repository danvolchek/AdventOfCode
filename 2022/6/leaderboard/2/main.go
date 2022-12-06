package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) any {
	return line
}

func solve(lines string) int {
	for j := range lines {
		if j < 13 {
			continue
		}

		var s lib.Set[byte]
		for i := 0; i < 14; i++ {
			s.Add(lines[j-i])
		}

		if s.Size() == 14 {
			return j + 1
		}
	}

	panic("asd")
}

func main() {
	solver := lib.Solver[string, int]{
		ParseF: lib.ParseString,
		SolveF: solve,
	}

	solver.Expect("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19)
	solver.Expect("bvwbjplbgvbhsrlpgdmjqwftvncz", 23)
	solver.Expect("nppdvjthqldpwncqszvftbrmjlhg", 23)
	solver.Expect("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29)
	solver.Expect("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26)

	solver.Verify(3613)
}
