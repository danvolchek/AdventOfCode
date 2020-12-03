package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func parse(r io.Reader) [][]bool {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := bytes.Split(bytes.TrimSpace(raw), []byte{'\r', '\n'})

	geology := make([][]bool, len(rows))

	for i, row := range rows {
		geology[i] = make([]bool, len(row))

		for j := 0; j < len(row); j++ {
			geology[i][j] = row[j] == '#'
		}
	}

	return geology
}

type slope struct {
	deltaX, deltaY int
}

func findTrees(geology [][]bool) int {
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

	for row := 0; row < len(geology); row+=slope.deltaY {
		if geology[row][col%len(geology[row])] {
			trees += 1
		}

		col += slope.deltaX
	}

	return trees
}

func main() {
	input, err := os.Open(path.Join("2020", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(findTrees(parse(input)))
}
