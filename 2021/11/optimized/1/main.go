package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "11", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type pos struct {
	x, y int
}

func parse(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)

	var octopi [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, rawHeight := range line {
			row = append(row, int(rawHeight-48))
		}
		octopi = append(octopi, row)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return octopi
}

func simulateStep(octopi [][]int) int {
	flashed := make(map[pos]bool)
	var q queue

	increment := func(p pos) {
		octopi[p.x][p.y] += 1

		if !flashed[p] && octopi[p.x][p.y] > 9 {
			flashed[p] = true

			q.push(p)
		}
	}

	for x := 0; x < len(octopi); x++ {
		for y := 0; y < len(octopi[x]); y++ {
			increment(pos{
				x: x,
				y: y,
			})
		}
	}

	for !q.empty() {
		p := q.pop()

		for _, adj := range getAdjacentPositions(octopi, p) {
			increment(adj)
		}
	}

	for p := range flashed {
		octopi[p.x][p.y] = 0
	}

	return len(flashed)
}

func getAdjacentPositions(octopi [][]int, p pos) []pos {
	var adjacentPositions []pos

	for _, offset := range []pos{
		{x: -1, y: 1},
		{x: 0, y: 1},
		{x: 1, y: 1},
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},
	} {
		adjacent := pos{
			x: p.x + offset.x,
			y: p.y + offset.y,
		}

		if adjacent.x >= 0 && adjacent.x < len(octopi) && adjacent.y >= 0 && adjacent.y < len(octopi[adjacent.x]) {
			adjacentPositions = append(adjacentPositions, adjacent)
		}
	}

	return adjacentPositions
}

func solve(r io.Reader) {
	octopi := parse(r)

	totalFlashes := 0
	for i := 0; i < 100; i++ {
		totalFlashes += simulateStep(octopi)
	}

	fmt.Println(totalFlashes)
}

func main() {
	solve(strings.NewReader("5483143223\n2745854711\n5264556173\n6141336146\n6357385478\n4167524645\n2176841721\n6882881134\n4846848554\n5283751526"))
	solve(input())
}
