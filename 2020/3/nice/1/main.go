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

func solve(geology [][]bool) int {
	trees := 0
	col := 0
	numCols := len(geology[0])

	for row := 0; row < len(geology); row++ {
		if geology[row][col] {
			trees += 1
		}

		col = (col + 3) % numCols
	}

	return trees
}

func main() {
	fmt.Println(solve(parse(input())))
}
