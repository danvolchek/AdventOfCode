package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "11", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()

		var lineIO []int

		for i := 0; i < len(line); i++ {
			v, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}

			lineIO = append(lineIO, v)
		}

		grid = append(grid, lineIO)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	totalFlashes := 0

	for i := 0; i < 100; i++ {
		//print(grid)
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				grid[x][y] += 1
			}
		}

		flashed := make([][]bool, len(grid))
		for q, row := range grid {
			flashed[q] = make([]bool, len(row))
		}

		for {
			flashedIter := false

			for x := 0; x < len(grid); x++ {
				for y := 0; y < len(grid[x]); y++ {
					if grid[x][y] > 9 && !flashed[x][y] {
						flashedIter = true
						flashed[x][y] = true
						totalFlashes += 1
					} else {
						continue
					}

					for _, offset := range []struct{ x, y int }{
						{-1, 1},
						{0, 1},
						{1, 1},
						{-1, 0},
						{1, 0},
						{-1, -1},
						{0, -1},
						{1, -1},
					} {
						inc(grid, x+offset.x, y+offset.y, 1)
					}
				}
			}

			if !flashedIter {
				break
			}
		}

		for x := 0; x < len(flashed); x++ {
			for y := 0; y < len(flashed[x]); y++ {
				if flashed[x][y] {
					grid[x][y] = 0
				}

			}
		}
	}
	fmt.Println(totalFlashes)
}

func print(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

func inc(grid [][]int, x, y, v int) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) {
		grid[x][y] += v
	}
}

func main() {
	solve(strings.NewReader("11111\n19991\n19191\n19991\n11111\n"))
	solve(strings.NewReader("5483143223\n2745854711\n5264556173\n6141336146\n6357385478\n4167524645\n2176841721\n6882881134\n4846848554\n5283751526"))
	solve(input())
}
