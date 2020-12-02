package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"
)

// parse input
func Parse() (int, int) {
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

// returns the digits that make up value
func digits(val int) []int {
	ret := []int{0, 0, 0, 0, 0, 0}

	for i := 5; i >= 0; i-- {
		ret[i] = val % 10
		val /= 10
	}

	return ret
}

// modifies digs to represent the next value after the one digs represents which still has all digits non-decreasing
// faster, but requires the input to already be non-decreasing
func incrementFast(digs []int) {
	if digs[5] != 9 {
		digs[5] += 1
	} else if digs[4] != 9 {
		digs[4] += 1
		digs[5] = digs[4]
	} else if digs[3] != 9 {
		digs[3] += 1
		digs[4] = digs[3]
		digs[5] = digs[3]
	} else if digs[2] != 9 {
		digs[2] += 1
		digs[3] = digs[2]
		digs[4] = digs[2]
		digs[5] = digs[2]
	} else if digs[1] != 9 {
		digs[1] += 1
		digs[2] = digs[1]
		digs[3] = digs[1]
		digs[4] = digs[1]
		digs[5] = digs[1]
	} else if digs[0] != 9 {
		digs[0] += 1
		digs[1] = digs[0]
		digs[2] = digs[0]
		digs[3] = digs[0]
		digs[4] = digs[0]
		digs[5] = digs[0]
	} else {
		panic("can't increment 999999")
	}
}

// modifies digs to represent the next value after the one digs represents which still has all digits non-decreasing
// slower, but has no preconditions on the input
func incrementSlow(digs []int) {
	// increase by 1
	for i := 5; i >= 0; i-- {
		digs[i] += 1
		if digs[i] == 10 {
			digs[i] = 0
		} else {
			break
		}
	}

	// set to next non-decreasing number
	for i := 1; i < 6; i++ {
		if digs[i] < digs[i-1] {
			for k := i; k < 6; k++ {
				digs[k] = digs[i-1]
			}
		}
	}
}

// returns whether the number a represents is less than the number b represents
func isLessThanOrEqual(a, b []int) bool {
	for i := 0; i < 6; i++ {
		if a[i] < b[i] {
			return true
		}

		if a[i] > b[i] {
			return false
		}
	}

	return true
}

// solves the puzzle
func NumPasswords(min, max int) int {
	result := 0
	maxDigs := digits(max)

	// start with digits that meet the non-decreasing rule
	digs := digits(min - 1)
	incrementSlow(digs)

	// check range rule
	for isLessThanOrEqual(digs, maxDigs) {
		// check double rule
		double := false
		for j := 1; j < 6; j++ {
			if digs[j] == digs[j-1] {
				double = true
				break
			}
		}

		// non-decreasing and range rules are already met
		if double {
			result += 1
		}

		// get next non-decreasing rule following number
		incrementFast(digs)
	}

	return result
}

func main() {
	fmt.Printf("%v\n", digits(123456))

	fmt.Println(NumPasswords(111111, 111111))
	fmt.Println(NumPasswords(223450, 223450))
	fmt.Println(NumPasswords(123789, 123789))

	min, max := Parse()

	fmt.Println(NumPasswords(min, max))
}
