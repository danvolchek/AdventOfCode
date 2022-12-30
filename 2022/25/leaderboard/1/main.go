package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
)

func fromSnafu(num string) uint64 {
	var val uint64

	var multiplier uint64 = 1
	for i := len(num) - 1; i >= 0; i-- {

		switch num[i] {
		case '-':
			val -= multiplier
		case '=':
			val -= multiplier * 2
		default:
			val += multiplier * uint64(num[i]-'0')
		}

		multiplier *= 5
	}

	return val
}

func toBase5(num uint64) []int {
	var result []int

	for num >= 0 {
		result = append([]int{int(num % 5)}, result...)
		num /= 5

		if num == 0 {
			break
		}
	}

	return result
}

func toSnafu(num uint64) string {
	base5 := toBase5(num)

	fmt.Println(base5)
	for i := 0; i < len(base5); i++ {
		if base5[i] > 2 {
			base5[i] -= 5
			base5[i-1] += 1
			fmt.Println(base5)

			i = 0
		}
	}

	var result string

	for _, val := range base5 {
		var char string
		switch val {
		case -2:
			char = "="
		case -1:
			char = "-"
		case 0:
			char = "0"
		case 1:
			char = "1"
		case 2:
			char = "2"
		}

		result += char
	}

	return result
}

func solve(snafus []string) string {
	nums := lib.Map(snafus, fromSnafu)
	sum := lib.SumSlice(nums)

	return toSnafu(sum)
}

func main() {
	solver := lib.Solver[[]string, string]{
		ParseF: lib.ParseLine(lib.AsIs),
		SolveF: solve,
	}

	solver.Expect("1=-0-2\n12111\n2=0=\n21\n2=01\n111\n20012\n112\n1=-1=\n1-12\n12\n1=\n122", "2=-1=0")
	solver.Solve()
}
