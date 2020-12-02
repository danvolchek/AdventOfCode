package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

func parse(r io.Reader) map[int]bool {
	csvReader := csv.NewReader(r)

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	expenses := make(map[int]bool, len(rows))
	for _, row := range rows {
		expense, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		expenses[expense] = true
	}

	return expenses

}

// solves the puzzle
func findEntries(expenses map[int]bool) int {
	const target = 2020

	for expense := range expenses {
		newTarget := target - expense

		for expense2 := range expenses {
			// look for whether the other needed value exists
			needed := newTarget-expense2

			if _, ok := expenses[needed]; ok {
				return expense * expense2 * needed
			}
		}
	}

	panic("no matches")
}

func main() {
	input, err := os.Open(path.Join("2020", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	expenses := parse(input)

	fmt.Println(findEntries(expenses))
}
