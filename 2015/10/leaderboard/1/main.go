package main

import (
	"strconv"

	"github.com/danvolchek/AdventOfCode/lib"
)

func expand(input []byte) []byte {
	var result []byte

	for i := 0; i < len(input); {
		j := i + 1

		for j < len(input) && input[i] == input[j] {
			j++
		}

		result = append(result, []byte(strconv.Itoa(j-i))...)
		result = append(result, input[i])
		i = j
	}

	return result
}

func solve(input []byte) int {
	for i := 0; i < 40; i++ {
		input = expand(input)
	}

	return len(input)
}

func main() {
	expandTest := lib.Solver[[]byte, string]{
		ParseF: lib.ParseBytes,
		SolveF: lib.ToString(expand),
	}

	expandTest.Expect("1", "11")
	expandTest.Expect("11", "21")
	expandTest.Expect("21", "1211")
	expandTest.Expect("1211", "111221")
	expandTest.Expect("111221", "312211")

	solver := lib.Solver[[]byte, int]{
		ParseF: lib.ParseBytes,
		SolveF: solve,
	}

	solver.Verify(360154)
}
