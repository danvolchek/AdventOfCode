package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) int {
	rawNums := lib.Filter([]byte(line), func(char byte) bool {
		return char >= '0' && char <= '9'
	})

	nums := lib.Map(rawNums, func(char byte) int {
		return int(char) - '0'
	})

	return nums[0]*10 + nums[len(nums)-1]
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: lib.SumSlice[int],
	}

	solver.Expect("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n", 142)
	solver.Verify(55108)
}
