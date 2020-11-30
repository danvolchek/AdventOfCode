package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func mass(val int) int {
	curr := val
	var total int

	for {
		curr = (curr / 3) - 2
		if curr <= 0 {
			break
		}

		total += curr
	}

	return total
}

func parse() []int {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(input)

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var res []int

	for _, row := range rows {
		stringVal := row[0]
		if len(stringVal) == 0 {
			continue
		}

		val, err := strconv.Atoi(stringVal)
		if err != nil {
			panic(err)
		}

		res = append(res, val)
	}

	return res
}

func main() {
	var sum int
	for _, val := range parse() {
		sum += mass(val)
	}

	fmt.Println(mass(14))
	fmt.Println(mass(1969))
	fmt.Println(mass(100756))

	fmt.Println(sum)
}
