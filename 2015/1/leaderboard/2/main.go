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
	lib.TestSolveBytes("()())", solve)
	lib.SolveBytes(input(), solve)
}
