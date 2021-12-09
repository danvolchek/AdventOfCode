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

	totalRisk := 0
	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			if smaller(cave[i][j], adjacent(cave, i, j)) {
				totalRisk += cave[i][j] + 1
			}
		}
	}

	fmt.Println(totalRisk)
}

func smaller(v int, values []int) bool {
	for _, vv := range values {
		if v >= vv {
			return false
		}
	}

	return true
}

func adjacent(cave [][]int, i, j int) []int {
	var ret []int
	for _, offset := range []struct{ x, y int }{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	} {
		v, ok := get(cave, i+offset.x, j+offset.y)
		if ok {
			ret = append(ret, v)
		}
	}

	return ret
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
