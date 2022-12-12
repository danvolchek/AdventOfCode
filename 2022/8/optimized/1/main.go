package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func isVisible(x, y int, grid [][]int) bool {
	v := grid[x][y]

	visibleX1 := true
	for i := x - 1; i > -1; i-- {
		if grid[i][y] >= v {
			visibleX1 = false
			break
		}
	}

	visibleX2 := true
	for i := x + 1; i < len(grid); i++ {
		if grid[i][y] >= v {
			visibleX2 = false
			break
		}
	}

	visibleY1 := true
	for i := y - 1; i > -1; i-- {
		if grid[x][i] >= v {
			visibleY1 = false
			break
		}
	}

	visibleY2 := true
	for i := y + 1; i < len(grid[x]); i++ {
		if grid[x][i] >= v {
			visibleY2 = false
			break
		}
	}

	return visibleX1 || visibleX2 || visibleY1 || visibleY2
}

func solve(lines [][]int) int {

	sum := 0
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines); y++ {
			if isVisible(x, y, lines) {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	solver := lib.Solver[[][]int, int]{
		ParseF: lib.ParseGrid(lib.Atoi),
		SolveF: solve,
	}

	solver.Expect("30373\n25512\n65332\n33549\n35390\n", 21)
	solver.Verify(1796)
}