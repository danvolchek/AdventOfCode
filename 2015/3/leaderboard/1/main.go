package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type pos struct {
	x, y int
}

func solve(instructions []byte) int {
	position := pos{x: 0, y: 0}

	gifts := make(map[pos]int)

	gifts[position] = 1

	for _, instruction := range instructions {
		switch instruction {
		case '^':
			position.y -= 1
		case 'v':
			position.y += 1
		case '>':
			position.x += 1
		case '<':
			position.x -= 1
		default:
			panic(string(instruction))
		}

		gifts[position] += 1
	}

	return len(gifts)
}

func main() {
	solver := lib.Solver[[]byte, int]{
		ParseF: lib.ParseBytes,
		SolveF: solve,
	}

	solver.Expect(">", 2)
	solver.Expect("^>v<", 4)
	solver.Expect("^v^v^v^v^v", 2)
	solver.Verify(2572)
}
