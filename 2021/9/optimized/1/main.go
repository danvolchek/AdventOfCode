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
	input, err := os.Open(path.Join("2021", "9", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)

	var cave [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, rawHeight := range line {
			row = append(row, int(rawHeight-48))
		}
		cave = append(cave, row)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return cave
}

func solve(r io.Reader) {
	cave := parse(r)

	totalRisk := 0
	for _, lowPoint := range findLowPoints(cave) {
		totalRisk += lowPoint + 1
	}

	fmt.Println(totalRisk)
}

func findLowPoints(cave [][]int) []int {
	var lowPoints []int

	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			height := cave[i][j]

			if smaller(height, adjacent(cave, i, j)) {
				lowPoints = append(lowPoints, height)
			}
		}
	}

	return lowPoints
}

func smaller(value int, values []int) bool {
	for _, otherValue := range values {
		if value >= otherValue {
			return false
		}
	}

	return true
}

func adjacent(cave [][]int, i, j int) []int {
	var adjacentHeights []int

	for _, offset := range []struct{ i, j int }{
		{i: 1, j: 0},
		{i: 0, j: 1},
		{i: -1, j: 0},
		{i: 0, j: -1},
	} {
		height, ok := get(cave, i+offset.i, j+offset.j)
		if ok {
			adjacentHeights = append(adjacentHeights, height)
		}
	}

	return adjacentHeights
}

func get(cave [][]int, i, j int) (int, bool) {
	if i < 0 || j < 0 || i > len(cave)-1 || j > len(cave[0])-1 {
		return 0, false
	}

	return cave[i][j], true
}

func main() {
	solve(strings.NewReader("2199943210\n3987894921\n9856789892\n8767896789\n9899965678"))
	solve(input())
}
