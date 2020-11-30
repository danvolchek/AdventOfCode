package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path"
	"strconv"
)

func parse() (int, int) {
	input, err := os.Open(path.Join("2019", "4", "input.txt"))
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(input)

	row, err := csvReader.Read()
	if err != nil {
		panic(err)
	}

	min, err := strconv.Atoi(row[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(row[1])
	if err != nil {
		panic(err)
	}

	return min, max
}

func intAtIndex(val, index int) int {
	return (val / int(math.Pow(10, float64(5-index)))) % 10
}

func numPasswords(min, max int) int {
	ret := 0
	for i := min; i <= max; i++ {

		double := false
		for j := 1; j < 6; j++ {
			span := 1
			for j < 6 && intAtIndex(i, j) == intAtIndex(i, j-1) {
				j += 1
				span += 1
			}

			if span == 2 {
				double = true
				break
			}
		}

		if !double {
			continue
		}

		nonDecreasing := true
		for j := 0; j < 6; j++ {
			if intAtIndex(i, j) < intAtIndex(i, j-1) {
				nonDecreasing = false
				break
			}
		}

		if !nonDecreasing {
			continue
		}

		ret += 1
	}

	return ret
}

func main() {
	fmt.Println(numPasswords(112233, 112233))
	fmt.Println(numPasswords(123444, 123444))
	fmt.Println(numPasswords(111122, 111122))

	min, max := parse()

	fmt.Println(numPasswords(min, max))
}
