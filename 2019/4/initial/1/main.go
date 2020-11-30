package main

import (
	"fmt"
	"math"
)

func intAtIndex(val, index int) int {
	return (val / int(math.Pow(10, float64(5-index)))) % 10
}

func numPasswords(min, max int) int {
	ret := 0
	for i := min; i <= max; i++ {

		double := false
		for j := 1; j < 6; j++ {
			if intAtIndex(i, j) == intAtIndex(i, j-1) {
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
	fmt.Printf("%d%d%d%d%d%d\n", intAtIndex(123456, 0), intAtIndex(123456, 1),
		intAtIndex(123456, 2), intAtIndex(123456, 3), intAtIndex(123456, 4),
		intAtIndex(123456, 5))

	fmt.Println(numPasswords(111111, 111111))
	fmt.Println(numPasswords(223450, 223450))
	fmt.Println(numPasswords(123789, 123789))

	fmt.Println(numPasswords(382345, 843167))
}
