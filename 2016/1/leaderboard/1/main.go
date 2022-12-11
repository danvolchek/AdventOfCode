package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type instruction struct {
	left     bool
	distance int
}

func parse(line string) []instruction {
	parts := strings.Split(strings.TrimSpace(line), ", ")

	return lib.Map(parts, func(p string) instruction {
		return instruction{
			left:     p[0] == 'L',
			distance: lib.Int(p),
		}
	})
}

func solve(lines []instruction) int {
	var x, y int

	var dir int
	// 0 north
	// 1 east
	// 2 south
	// 3 west

	for _, line := range lines {
		//fmt.Printf("%+v\n", line)
		if line.left {
			dir = (dir - 1 + 4) % 4
		} else {
			dir = (dir + 1) % 4
		}

		//fmt.Println(dir)

		switch dir {
		case 0:
			y += line.distance
		case 1:
			x += line.distance
		case 2:
			y -= line.distance
		case 3:
			x -= line.distance
		default:
			panic(dir)
		}

		//fmt.Println(x, y)
	}

	return lib.Abs(x) + lib.Abs(y)
}

func main() {
	solver := lib.Solver[[]instruction, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("R2, L3", 5)
	solver.Expect("R2, R2, R2", 2)
	solver.Expect("R5, L5, R5, R3", 12)
	solver.Verify(271)
}
