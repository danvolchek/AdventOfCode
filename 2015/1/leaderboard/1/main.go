package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

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
		ParseF: lib.ParseBytes,
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

	solver.Verify(232)
}
