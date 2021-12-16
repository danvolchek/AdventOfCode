package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
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

// this takes ~35 minutes - the problem is minDist iterates over all vertices, of which there are 250000
// a priority queue would be much more efficient, but there isn't a suitable one in the Go stdlib or other libraries,
// and I don't want to write one
func explore(grid [][]int, start pos) int {
	dist := make(map[pos]int)
	prev := make(map[pos]pos)
	vertices := make(map[pos]bool)

	for i := 0; i < len(grid)*5; i++ {
		for j := 0; j < len(grid)*5; j++ {
			p := pos{x: i, y: j}

			dist[p] = 999999999
			vertices[p] = true
		}
	}

	dist[start] = 0

	st := time.Now()
	j := 0
	for len(vertices) != 0 {
		j++
		if j == 500 {
			cur := time.Now()
			fmt.Println(len(vertices), "500 took", cur.Sub(st))
			fmt.Println("Estimated", time.Duration(cur.Sub(st).Nanoseconds()*(int64(len(vertices))/500)))
			st = cur
			j = 0
		}
		minVertex := minDist(vertices, dist)

		delete(vertices, minVertex)

		for _, offset := range []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			v := pos{minVertex.x + offset.x, minVertex.y + offset.y}

			if !vertices[v] {
				continue
			}

			alt := dist[minVertex] + risk(grid, v.x, v.y)
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = minVertex
			}
		}
	}

	return dist[pos{len(grid)*5 - 1, len(grid[0])*5 - 1}]
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

func risk(grid [][]int, i, j int) int {
	actI, actJ := i%len(grid), j%len(grid[i%len(grid)])
	v := grid[actI][actJ]

	for k := 0; k < (i/len(grid))+(j/len(grid[i%len(grid)])); k++ {
		v++
		if v == 10 {
			v = 1
		}
	}

	return v
}

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			q, c := 0, 3

			steve := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {8, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
			steve[q][c] = 8
			fmt.Print(risk(steve,
				q+10*i, c+10*j), " ")
		}
		fmt.Println()
	}
	solve(strings.NewReader("1163751742\n1381373672\n2136511328\n3694931569\n7463417111\n1319128137\n1359912421\n3125421639\n1293138521\n2311944581"))
	solve(input())
}
