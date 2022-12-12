package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"math"
)

func parse(line string) int {
	return lib.Atoi(line)
}

func presents(house int) int {
	if house == 1 {
		return 10
	}

	total := 0
	for i := 1; i <= int(math.Sqrt(float64(house))); i++ {
		if house%i == 0 {
			if i*50 >= house {
				total += i * 11
			}

			if i != house/i {
				if (house/i)*50 >= house {
					total += (house / i) * 11
				}
			}
		}
	}

	return total
}

func solve(target int) int {
	i := target / 50
	for {
		result := presents(i)
		if result >= target {
			return i
		}

		i += 1
	}
}

func main() {
	solver := lib.Solver[int, int]{
		ParseF: parse,
		SolveF: solve,
	}

	solver.Verify(705600)
}
