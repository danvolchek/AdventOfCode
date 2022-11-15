package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(instructions []byte) int {
	floor := 0

	for _, instruction := range instructions {
		switch instruction {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			panic(string(instruction))
		}
	}

	return floor
}

func main() {
	solver := lib.Solver[[]byte, int]{
		ParseF: lib.ParseBytes(),
		SolveF: solve,
	}

	for _, tc := range []struct {
		in  string
		out int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	} {
		solver.Expect(tc.in, tc.out)
	}

	solver.Solve(input())
}
