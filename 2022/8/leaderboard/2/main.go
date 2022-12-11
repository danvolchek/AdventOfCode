package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

func parse(input string) [][]int {
	var result [][]int

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, "")

		var ret []int

		for _, p := range parts {
			ret = append(ret, lib.Int(p))
		}

		result = append(result, ret)
	}

	return result
}

func viewingDistance(x, y int, grid [][]int) int {
	v := grid[x][y]

	visibleX1 := 0
	for i := x - 1; i > -1; i-- {
		visibleX1 += 1
		if grid[i][y] >= v {
			break
		}
	}

	visibleX2 := 0
	for i := x + 1; i < len(grid); i++ {
		visibleX2 += 1
		if grid[i][y] >= v {
			break
		}
	}

	visibleY1 := 0
	for i := y - 1; i > -1; i-- {
		visibleY1 += 1
		if grid[x][i] >= v {
			break
		}
	}

	visibleY2 := 0
	for i := y + 1; i < len(grid[x]); i++ {
		visibleY2 += 1
		if grid[x][i] >= v {
			break
		}
	}

	return visibleX1 * visibleX2 * visibleY1 * visibleY2
}

func solve(lines [][]int) int {
	maxView := 0
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines); y++ {
			if dist := viewingDistance(x, y, lines); dist > maxView {
				maxView = dist
			}
		}
	}

	return maxView
}

func main() {
	solver := lib.Solver[[][]int, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("30373\n25512\n65332\n33549\n35390\n", 8)
	solver.Verify(288120)
}
