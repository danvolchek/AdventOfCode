package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
)

func parse(r io.Reader) {
	csvReader := csv.NewReader(r)

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
}

func main() {
	input, err := os.Open(path.Join("2020", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	parse(input)
}
