package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
)

type node struct {
	x, y int

	adjacent []*node
}

func (n *node) Id() string {
	return fmt.Sprintf("%d,%d", n.y, n.x)
}

func (n *node) Adjacent() []*node {
	return n.adjacent
}

func (n *node) String() string {
	return n.Id()
}

func parse(line string) int {
	if line == "S" {
		return 9999999
	}

	if line == "E" {
		return -2
	}

	return int(line[0] - 'a')
}

func solve(grid [][]int) int {

	var start *node
	var end string

	gridMap := make(map[int]map[int]*node)

	for y, line := range grid {
		for x, height := range line {
			n := &node{
				x:        x,
				y:        y,
				adjacent: nil,
			}

			if height == 9999999 {
				start = n
			}

			if height == -2 {
				end = n.Id()
			}

			if gridMap[y] == nil {
				gridMap[y] = make(map[int]*node)
			}

			gridMap[y][x] = n

			if height == -2 {
				grid[y][x] = parse("z")
			}
		}
	}

	for y, line := range grid {
		for x, height := range line {
			for iy := -1; iy <= 1; iy += 1 {
				for ix := -1; ix <= 1; ix += 1 {
					if iy == 0 && ix == 0 {
						continue
					}
					if lib.Abs(iy)+lib.Abs(ix) == 2 {
						continue
					}

					ey := y + iy
					ex := x + ix

					if ey < 0 || ex < 0 || ey >= len(grid) || ex >= len(grid[ey]) {
						continue
					}

					if height >= grid[ey][ex]-1 {
						gridMap[y][x].adjacent = append(gridMap[y][x].adjacent, gridMap[ey][ex])
					}
				}
			}
		}
	}

	result := lib.BFS(start, end)
	fmt.Println(result)
	return len(result) - 1
}

func main() {
	solver := lib.Solver[[][]int, int]{
		ParseF: lib.ParseGrid(parse),
		SolveF: solve,
	}

	solver.Expect("Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi", 31)
	solver.Verify(330)
}
