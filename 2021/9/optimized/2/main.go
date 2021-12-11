package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
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

	lowPoints := findLowPoints(cave)

	basins := findBasins(cave, lowPoints)

	sizes := basinSizes(basins)

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}

func findLowPoints(cave [][]int) []pos {
	var lowPoints []pos

	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			p := pos{
				i: i,
				j: j,
			}
			if isSmallerThanAdjacent(cave, p) {
				lowPoints = append(lowPoints, p)
			}
		}
	}

	return lowPoints
}

func findBasins(cave [][]int, lowPoints []pos) [][]int {
	basins := make([][]int, len(cave))
	for i, row := range cave {
		basins[i] = make([]int, len(row))
	}

	for basinNum, lowPoint := range lowPoints {
		flood(cave, basins, basinNum+1, lowPoint)
	}

	return basins
}

func flood(cave, basins [][]int, basinNum int, p pos) {
	if cave[p.i][p.j] != 9 {
		basins[p.i][p.j] = basinNum
	}

	for _, adj := range largerAdjacent(cave, p) {
		if cave[p.i][p.j] == 9 {
			continue
		}

		flood(cave, basins, basinNum, adj)
	}
}

type pos struct {
	i, j int
}

func basinSizes(basins [][]int) []int {
	sizes := make(map[int]int)

	for i := 0; i < len(basins); i++ {
		for j := 0; j < len(basins[i]); j++ {
			sizes[basins[i][j]] += 1
		}
	}

	var sizeSlice []int
	for basinNum, basinSize := range sizes {
		if basinNum != 0 {
			sizeSlice = append(sizeSlice, basinSize)
		}
	}

	return sizeSlice
}

func largerAdjacent(cave [][]int, p pos) []pos {
	height := cave[p.i][p.j]

	var filtered []pos
	for _, other := range adjacent(cave, p) {
		if cave[other.i][other.j] > height {
			filtered = append(filtered, other)
		}
	}

	return filtered
}

func isSmallerThanAdjacent(cave [][]int, p pos) bool {
	height := cave[p.i][p.j]

	for _, other := range adjacent(cave, p) {
		if height >= cave[other.i][other.j] {
			return false
		}
	}

	return true
}

func adjacent(cave [][]int, p pos) []pos {
	var adjacentHeights []pos

	for _, offset := range []pos{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	} {
		adj := pos{
			i: p.i + offset.i,
			j: p.j + offset.j,
		}

		if inBounds(cave, adj) {
			adjacentHeights = append(adjacentHeights, adj)
		}
	}

	return adjacentHeights
}

func inBounds(cave [][]int, p pos) bool {
	return p.i >= 0 && p.j >= 0 && p.i < len(cave) && p.j < len(cave[p.i])
}

func main() {
	solve(strings.NewReader("2199943210\n3987894921\n9856789892\n8767896789\n9899965678"))
	solve(input())
}
