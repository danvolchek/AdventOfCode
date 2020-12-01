package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

func parse(r io.Reader) {
	csvReader := csv.NewReader(r)

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var items []int
	for _, item := range rows {
		item, err := strconv.Atoi(item[0])
		if err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	for i, item1 := range items {
		for j, item2 := range items {
			for _, item3 := range items {
				if i != j && item1+item2+item3 == 2020 {
					fmt.Println(item1 * item2 * item3)
					return
				}
			}
		}
	}
}

func main() {
	input, err := os.Open(path.Join("2020", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	parse(input)
}
