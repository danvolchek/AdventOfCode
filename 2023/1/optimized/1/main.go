package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

func parse(line string) int {
	nums := lib.Map(strings.Split(line, ""), func(charStr string) int {
		char := charStr[0]

		if char < '0' || char > '9' {
			return -1
		}

		return int(char) - '0'
	})

	nums = lib.Filter(nums, func(v int) bool {
		return v != -1
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
