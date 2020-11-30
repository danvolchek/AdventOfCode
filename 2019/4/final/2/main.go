package main

import (
	"fmt"
)

// returns the digits that make up value
func digits(val int) []int {
	ret := []int{0, 0, 0, 0, 0, 0}

	for i := 0; i < 6; i++ {
		ret[5-i] = val % 10
		val /= 10
	}

	return ret
}

// returns the value after digs which still has all digits non-decreasing
func increment(digs []int) {
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

// solves the problem
func NumPasswords(min, max int) int {
	result := 0
	maxDigs := digits(max)

	// start with digits that meet the non-decreasing rule
	digs := digits(min - 1)
	increment(digs)

	// check range rule
	for isLessThanOrEqual(digs, maxDigs) {
		// check double rule
		double := false
		for j := 1; j < 6; j++ {
			span := 1
			for j < 6 && digs[j] == digs[j-1] {
				j += 1
				span += 1
			}

			if span == 2 {
				double = true
				break
			}
		}

		// non-decreasing and range rules are already met
		if double {
			result += 1
		}

		// get next non-decreasing rule following number
		increment(digs)
	}

	return result
}

// splits the work to solve the passwords problem into chunks and spawns goroutines to solve each one
func NumPasswordsGoroutines(min, max, numGoroutines int) int {
	results := make(chan int, numGoroutines)
	step := (max - min) / numGoroutines

	for {
		// create max
		newMax := min + step
		if newMax > max {
			newMax = max
		}

		// start goroutine
		go func(a, b int) {
			results <- NumPasswords(a, b)
		}(min, newMax)

		// set new min and check to exit
		min = newMax + 1
		if min > max {
			break
		}
	}

	// aggregate results
	sum := 0
	for i := 0; i < numGoroutines; i++ {
		sum += <-results
	}

	return sum
}

func main() {
	fmt.Printf("%v\n", digits(123456))

	fmt.Println(NumPasswords(112233, 112233))
	fmt.Println(NumPasswords(123444, 123444))
	fmt.Println(NumPasswords(111122, 111122))

	fmt.Println(NumPasswords(382345, 843167))
	fmt.Println(NumPasswordsGoroutines(382345, 843167, 2))
}
