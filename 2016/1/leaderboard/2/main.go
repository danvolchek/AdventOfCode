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
			distance: lib.Ints(p)[0],
		}
	})
}

func solve(lines []instruction) int {
	type pos struct{ x, y int }
	var seen lib.Set[pos]

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

		newX, newY := x, y
		switch dir {
		case 0:
			newY += line.distance
		case 1:
			newX += line.distance
		case 2:
			newY -= line.distance
		case 3:
			newX -= line.distance
		default:
			panic(dir)
		}

		xDir, yDir := 1, 1
		if dir == 2 {
			yDir = -1
		}

		if dir == 3 {
			xDir = -1
		}
		for xx := x + xDir; xx != newX+xDir; xx += xDir {
			if seen.Contains(pos{xx, y}) {
				return lib.Abs(xx) + lib.Abs(y)
			}
			seen.Add(pos{xx, y})
		}

		for yy := y + yDir; yy != newY+yDir; yy += yDir {
			if seen.Contains(pos{x, yy}) {
				return lib.Abs(x) + lib.Abs(yy)
			}
			seen.Add(pos{x, yy})
		}

		x, y = newX, newY
	}

	panic("no dupe")
}

func main() {
	solver := lib.Solver[[]instruction, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("R8, R4, R4, R8", 4)
	solver.Verify(153)
}
