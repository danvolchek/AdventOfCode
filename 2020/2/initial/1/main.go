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

		t.pass = row[1]

		found := count(t.char, t.pass)

		if found >= t.min && found <= t.max {
			matching++
		}
	}

	fmt.Println(matching)

}

func count(char uint8, pass string) int {
	num := 0
	for i := 0; i < len(pass); i++{
		if pass[i] == char {
			num++
		}
	}

	return num
}

func main() {
	input, err := os.Open(path.Join("2020", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	parse(input)
}
