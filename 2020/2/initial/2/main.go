package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

type thingy struct {
	min, max int
	char uint8
	pass string
}

func parse(r io.Reader) {
	csvReader := csv.NewReader(r)

	csvReader.Comma = ':'
	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	matching := 0

	for _, row := range rows {
		t := thingy{}
		first := row[0]

		parts := strings.Split(first, " ")

		amount := parts[0]

		nums := strings.Split(amount, "-")

		min, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		t.min = min

		max, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		t.max = max

		t.char = parts[1][0]

		t.pass = strings.TrimSpace(row[1])


		a := t.pass[t.min - 1] == t.char
		b := t.pass[t.max - 1] == t.char

		if a && !b || b && !a {
			matching++
			fmt.Println(t.min, t.max, t.char, t.pass)
		}
	}

	fmt.Println(matching)

}

func main() {
	input, err := os.Open(path.Join("2020", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	parse(input)
}
