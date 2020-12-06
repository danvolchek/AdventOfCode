package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) [][]bool {
	var geology [][]bool

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := scanner.Text()

		trees := make([]bool, len(row))

		for i := 0; i < len(row); i++ {
			trees[i] = row[i] == '#'
		}

		geology = append(geology, trees)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return geology
}

type slope struct {
	deltaX, deltaY int
}

func solve(geology [][]bool) int {
	answer := 1

	for _, slope := range []slope{
		{deltaX: 1, deltaY: 1},
		{deltaX: 3, deltaY: 1},
		{deltaX: 5, deltaY: 1},
		{deltaX: 7, deltaY: 1},
		{deltaX: 1, deltaY: 2}} {
		answer *= findTreesInSlope(geology, slope)
	}

	return answer
}

func findTreesInSlope(geology [][]bool, slope slope) int {
	trees := 0
	col := 0
	numCols := len(geology[0])

	for row := 0; row < len(geology); row += slope.deltaY {
		if geology[row][col] {
			trees += 1
		}

		col = (col + slope.deltaX) % numCols
	}

	return trees
}

func main() {
	fmt.Println(solve(parse(input())))
}
