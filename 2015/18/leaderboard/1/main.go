package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(data []byte) [][]bool {
	var grid [][]bool

	var line []bool

	for _, d := range data {
		switch d {
		case '#':
			line = append(line, true)
		case '.':
			line = append(line, false)
		case '\n':
			grid = append(grid, line)
			line = nil
		default:
			panic(d)
		}
	}

	if len(line) != 0 {
		grid = append(grid, line)
	}

	return grid
}

func neighbors(grid [][]bool, x, y int) []bool {
	var result []bool

	for xOff := -1; xOff <= 1; xOff++ {
		for yOff := -1; yOff <= 1; yOff++ {
			if xOff == 0 && yOff == 0 {
				continue
			}

			xNeighbor := x + xOff
			yNeighbor := y + yOff

			if xNeighbor < 0 || xNeighbor >= len(grid) {
				continue
			}

			if yNeighbor < 0 || yNeighbor >= len(grid[xNeighbor]) {
				continue
			}

			result = append(result, grid[xNeighbor][yNeighbor])
		}
	}

	return result
}

type pos struct {
	x, y int
}

func step(grid [][]bool) [][]bool {
	var flips []pos

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			neighborOnCount := len(lib.Filter(neighbors(grid, x, y), func(neighbor bool) bool { return neighbor }))

			switch grid[x][y] {
			case true:
				if !(neighborOnCount == 2 || neighborOnCount == 3) {
					flips = append(flips, pos{x: x, y: y})
				}
			case false:
				if neighborOnCount == 3 {
					flips = append(flips, pos{x: x, y: y})
				}
			}
		}
	}

	for _, flip := range flips {
		grid[flip.x][flip.y] = !grid[flip.x][flip.y]
	}

	// only for step solver
	return grid
}

func solve(grid [][]bool) int {
	for i := 0; i < 100; i++ {
		step(grid)
	}

	sum := 0
	for _, line := range grid {
		for _, elem := range line {
			if elem {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	stepSolver := lib.Solver[[][]bool, [][]bool]{
		ParseF: lib.ParseBytesFunc(parse),
		SolveF: step,
	}
	stepSolver.ParseExpect(".#.#\n..##", [][]bool{{false, true, false, true}, {false, false, true, true}})
	stepSolver.Expect(".#.#.#\n...##.\n#....#\n..#...\n#.#..#\n####..", parse([]byte("..##..\n..##.#\n...##.\n......\n#.....\n#.##..")))
	stepSolver.Expect("..##..\n..##.#\n...##.\n......\n#.....\n#.##..", parse([]byte("..###.\n......\n..###.\n......\n.#....\n.#....")))
	stepSolver.Expect("..###.\n......\n..###.\n......\n.#....\n.#....", parse([]byte("...#..\n......\n...#..\n..##..\n......\n......")))
	stepSolver.Expect("...#..\n......\n...#..\n..##..\n......\n......", parse([]byte("......\n......\n..##..\n..##..\n......\n......")))

	solver := lib.Solver[[][]bool, int]{
		ParseF: lib.ParseBytesFunc(parse),
		SolveF: solve,
	}

	solver.Verify(768)
}
