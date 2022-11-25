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
			total += i * 10

			if i != house/i {
				total += (house / i) * 10
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
	presentsSolver := lib.Solver[int, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: presents,
	}
	presentsSolver.Expect("1", 10)
	presentsSolver.Expect("2", 30)
	presentsSolver.Expect("3", 40)
	presentsSolver.Expect("4", 70)
	presentsSolver.Expect("5", 60)
	presentsSolver.Expect("6", 120)
	presentsSolver.Expect("7", 80)
	presentsSolver.Expect("8", 150)
	presentsSolver.Expect("9", 130)

	solver := lib.Solver[int, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("100", 6)
	solver.Verify(665280)
}
