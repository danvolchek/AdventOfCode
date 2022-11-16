package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "10", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func expand(input []byte) []byte {
	var result []byte

	for i := 0; i < len(input); {
		j := i + 1

		for ; j < len(input) && input[i] == input[j]; j++ {
		}

		result = append(result, []byte(strconv.Itoa(j-i))...)
		result = append(result, input[i])
		i = j
	}

	return result
}

func solve(input []byte) int {
	for i := 0; i < 50; i++ {
		input = expand(input)
	}

	return len(input)
}

func main() {
	stepSolver := lib.Solver[[]byte, []byte]{
		ParseF: lib.ParseBytes,
		SolveF: expand,
	}

	stepSolver.Expect("1", []byte("11"))
	stepSolver.Expect("11", []byte("21"))
	stepSolver.Expect("21", []byte("1211"))
	stepSolver.Expect("1211", []byte("111221"))
	stepSolver.Expect("111221", []byte("312211"))

	solver := lib.Solver[[]byte, int]{
		ParseF: lib.ParseBytes,
		SolveF: solve,
	}

	solver.Solve(input())
}
