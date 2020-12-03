package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
)

func solve(r io.Reader) {
	csvReader := csv.NewReader(r)

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var grid [][]byte

	grid = make([][]byte, len(rows))

	for i, row := range rows {
		actual := row[0]

		grid[i] = make([]byte, len(actual))

		for j := 0; j < len(actual); j++ {
			grid[i][j] = actual[j]
		}
	}

	found := 0
	q := 0
	for k := 0; k < len(grid); k++ {

		if grid[k][q%len(grid[k])] == '#' {
			found++
		}

		q += 3
	}

	fmt.Println(found)
}

func main() {
	input, err := os.Open(path.Join("2020", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	solve(input)
}
