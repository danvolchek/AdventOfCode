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

func parseFile() [][]bool {
	input, err := os.Open(path.Join("2020", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	return parse(input)
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
	fmt.Println(solve(parseFile()))
}
