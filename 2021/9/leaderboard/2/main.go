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

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var cave [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var row []int
		for _, rawHeight := range line {
			height := rawHeight - 48
			row = append(row, int(height))
		}
		cave = append(cave, row)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	result := assignBasins(cave)

	basins := getBasins(cave)

	sizes := basinSizes(result, basins)

	fmt.Println(sizes, sizes[0]*sizes[1]*sizes[2])
}

type pos struct {
	i, j int
}

func basinSizes(assignment [][]pos, basins []pos) []int {
	var ret []int

	for _, basin := range basins {
		ret = append(ret, count(assignment, basin))
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i] > ret[j]
	})

	return ret
}

func count(assignment [][]pos, basin pos) int {
	sum := 0

	for i := 0; i < len(assignment); i++ {
		for j := 0; j < len(assignment[i]); j++ {
			if assignment[i][j] == basin {
				sum += 1
			}
		}
	}

	return sum
}

func assignBasins(cave [][]int) [][]pos {
	assignment := make([][]pos, len(cave))
	for i, row := range cave {
		assignment[i] = make([]pos, len(row))

		for j := range row {
			assignment[i][j] = pos{
				i: -1,
				j: -1,
			}
		}
	}

	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			flow(cave, assignment, i, j)
		}
	}

	return assignment
}

var sentinel = pos{i: -1, j: -1}

func flow(cave [][]int, assignment [][]pos, i, j int) pos {
	if assignment[i][j] != sentinel || cave[i][j] == 9 {
		return assignment[i][j]
	}

	values, positions := adjacent(cave, i, j)
	smaller, ok := firstSmallerPos(cave[i][j], values, positions)
	if !ok {
		assignment[i][j] = pos{i, j}
		return pos{i, j}
	}

	basin := flow(cave, assignment, smaller.i, smaller.j)
	assignment[i][j] = basin
	return basin
}

func firstSmallerPos(v int, values []int, positions []pos) (pos, bool) {
	for i, vv := range values {
		if vv < v {
			return positions[i], true
		}
	}

	return pos{}, false
}

func getBasins(cave [][]int) []pos {
	var ret []pos

	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			vals, _ := adjacent(cave, i, j)
			if allSmaller(cave[i][j], vals) {
				ret = append(ret, pos{i: i, j: j})
			}
		}
	}

	return ret
}

func allSmaller(v int, values []int) bool {
	for _, vv := range values {
		if v >= vv {
			return false
		}
	}

	return true
}

func adjacent(cave [][]int, i, j int) ([]int, []pos) {
	var ret []int
	var ret2 []pos

	for _, offset := range []pos{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	} {
		v, ok := get(cave, i+offset.i, j+offset.j)
		if ok {
			ret = append(ret, v)
			ret2 = append(ret2, pos{
				i: i + offset.i,
				j: j + offset.j,
			})
		}
	}

	return ret, ret2
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
