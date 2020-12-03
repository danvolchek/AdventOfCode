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

func findTrees(geology [][]bool) int {
	trees := 0
	col := 0

	for row := 0; row < len(geology); row++ {
		if geology[row][col%len(geology[row])] {
			trees += 1
		}

		col += 3
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
