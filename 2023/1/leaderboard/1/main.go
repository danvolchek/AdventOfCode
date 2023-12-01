package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(line string) int {
	first := 0
	second := 0
	for _, i := range line {
		if i >= '0' && i <= '9' {
			first = int(i) - '0'
			break
		}
	}

	for j := len(line) - 1; j >= 0; j-- {
		i := line[j]
		if i >= '0' && i <= '9' {
			second = int(i) - '0'
			break
		}
	}

	fmt.Println(line, first, second)

	return first*10 + second

}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: lib.SumSlice[int],
	}

	solver.Expect("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n", 142)
	solver.Verify(55108)
}
