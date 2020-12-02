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

	fmt.Println(rows)
}

func main() {
	input, err := os.Open(path.Join("2020", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	solve(input)
}
