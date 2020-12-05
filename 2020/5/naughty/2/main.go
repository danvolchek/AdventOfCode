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

	present := make(map[int]bool)
	for _, row := range rows {
		present[getId(row[0])] = true
	}

	for i := 0; i < 128*8; i++ {
		_, pres := present[i]

		_, presNext := present[i+1]
		_, presPrev := present[i-1]

		if !pres && presPrev && presNext {
			fmt.Println(i)
		}
	}

}

func getId(pass string) int {
	min, max := 0, 128

	for i := 0; i < 7; i++ {
		if pass[i] == 'F' {
			max -= (max - min) / 2
		} else {
			min += (max - min) / 2
		}
	}

	row := min

	colMin, colMax := 0, 8
	for i := 0; i < 3; i++ {
		if pass[7+i] == 'L' {
			colMax -= (colMax - colMin) / 2
		} else {
			colMin += (colMax - colMin) / 2
		}
	}

	col := colMin

	return row*8 + col
}

func main() {
	input, err := os.Open(path.Join("2020", "5", "input.txt"))
	if err != nil {
		panic(err)
	}

	solve(input)

}
