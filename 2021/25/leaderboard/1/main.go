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
	input, err := os.Open(path.Join("2021", "25", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var cucumbers [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]byte, len(line))
		for i := range row {
			row[i] = line[i]
		}

		cucumbers = append(cucumbers, row)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	i := 0

	for move(cucumbers) {
		i += 1
	}

	fmt.Println(i)
}

type pos struct{ i, j int }
type swap struct{ p1, p2 pos }

func move(cucumbers [][]byte) bool {
	moved := false

	var swaps []swap

	for i := 0; i < len(cucumbers); i++ {
		for j := 0; j < len(cucumbers[i]); j++ {
			if cucumbers[i][j] != '>' {
				continue
			}

			nextJ := j + 1
			if nextJ == len(cucumbers[i]) {
				nextJ = 0
			}

			if cucumbers[i][nextJ] == '.' {
				moved = true
				swaps = append(swaps, swap{p1: pos{i, j}, p2: pos{i, nextJ}})
			}
		}
	}

	for _, sw := range swaps {
		cucumbers[sw.p1.i][sw.p1.j], cucumbers[sw.p2.i][sw.p2.j] = cucumbers[sw.p2.i][sw.p2.j], cucumbers[sw.p1.i][sw.p1.j]
	}

	swaps = nil

	for i := 0; i < len(cucumbers); i++ {
		for j := 0; j < len(cucumbers[i]); j++ {
			if cucumbers[i][j] != 'v' {
				continue
			}

			nextI := i + 1
			if nextI == len(cucumbers) {
				nextI = 0
			}

			if cucumbers[nextI][j] == '.' {
				moved = true
				swaps = append(swaps, swap{p1: pos{i, j}, p2: pos{nextI, j}})
			}
		}
	}

	for _, sw := range swaps {
		cucumbers[sw.p1.i][sw.p1.j], cucumbers[sw.p2.i][sw.p2.j] = cucumbers[sw.p2.i][sw.p2.j], cucumbers[sw.p1.i][sw.p1.j]
	}

	return moved
}

func main() {
	solve(strings.NewReader("v...>>.vv>\n.vv>>.vv..\n>>.>v>...v\n>>v>>.>.v.\nv>v.vv.v..\n>.>>..v...\n.vv..>.>v.\nv.v..>>v.v\n....v..v.>"))
	solve(input())
}
