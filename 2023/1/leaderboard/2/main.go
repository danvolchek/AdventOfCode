package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

var nums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func parse(line string) int {
	first := 0
	second := 0
outer:
	for i := 0; i < len(line); i++ {
		for j, num := range nums {
			if strings.Index(line[i:], num) == 0 {
				first = j + 1
				break outer
			}
		}

		if line[i] >= '0' && line[i] <= '9' {
			first = int(line[i]) - '0'
			break
		}
	}

outer2:
	for j := len(line) - 1; j >= 0; j-- {
		for jj, num := range nums {
			if strings.Index(line[j:], num) == 0 {
				second = jj + 1
				break outer2
			}
		}

		if line[j] >= '0' && line[j] <= '9' {
			second = int(line[j]) - '0'
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

	solver.Expect("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n", 281)
	solver.Verify(56324)
}
