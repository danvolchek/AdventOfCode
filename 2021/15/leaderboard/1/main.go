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
	input, err := os.Open(path.Join("2021", "15", "input.txt"))
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

		//fmt.Println(line)

		var ints []int

		for _, c := range line {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			ints = append(ints, v)
		}

		grid = append(grid, ints)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(explore(grid, pos{}))
}

type pos struct{ x, y int }

func explore(grid [][]int, start pos) int {
	dist := make(map[pos]int)
	prev := make(map[pos]pos)
	vertices := make(map[pos]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			p := pos{x: i, y: j}

			dist[p] = 999999999
			vertices[p] = true
		}
	}

	dist[start] = 0

	for len(vertices) != 0 {
		minVertex := minDist(vertices, dist)

		delete(vertices, minVertex)

		for _, offset := range []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			v := pos{minVertex.x + offset.x, minVertex.y + offset.y}

			if !vertices[v] {
				continue
			}

			alt := dist[minVertex] + grid[v.x][v.y]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = minVertex
			}
		}
	}

	return dist[pos{len(grid) - 1, len(grid[0]) - 1}]
}

func minDist(vertices map[pos]bool, dist map[pos]int) pos {
	minVertex, minValue := pos{}, -1

	for vertex := range vertices {
		val := dist[vertex]
		if minValue == -1 || val < minValue {
			minVertex = vertex
			minValue = val
		}
	}

	return minVertex
}

func valid(grid [][]int, i, j int) bool {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[i])-1 {
		return false
	}

	return true
}

func main() {
	solve(strings.NewReader("1163751742\n1381373672\n2136511328\n3694931569\n7463417111\n1319128137\n1359912421\n3125421639\n1293138521\n2311944581"))
	solve(input())
}
