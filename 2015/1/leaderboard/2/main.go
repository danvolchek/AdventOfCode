package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func solve(instructions []byte) int {
	floor := 0

	for i, instruction := range instructions {
		switch instruction {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			panic(string(instruction))
		}

		if floor == -1 {
			return i + 1
		}
	}

	panic("not found")
}

func main() {
	solver := lib.Solver[[]byte, int]{
		ParseF: lib.ParseBytes,
		SolveF: solve,
	}

	solver.Expect(")", 1)
	solver.Expect("()())", 5)

	solver.Verify(1783)
}
