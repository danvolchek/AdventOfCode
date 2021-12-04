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

	big := 1

	for i, item := range []int{1, 3, 5, 7, 1} {
		found := 0
		q := 0
		for k := 0; k < len(rows); k++ {

			if grid[k][q%len(grid[k])] == '#' {
				found++
			}

			q += item

			if i == 4 {
				k++
			}
		}
		big *= found

	}

	fmt.Println(big)
}

func main() {
	input, err := os.Open(path.Join("2020", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	solve(input)
}
