package main

import (
	"regexp"

	"github.com/danvolchek/AdventOfCode/lib"
)

type ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

var parseRegExp = regexp.MustCompile(`(.*): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

func parse(matches []string) ingredient {
	return ingredient{
		name:       matches[0],
		capacity:   lib.Atoi(matches[1]),
		durability: lib.Atoi(matches[2]),
		flavor:     lib.Atoi(matches[3]),
		texture:    lib.Atoi(matches[4]),
		calories:   lib.Atoi(matches[5]),
	}
}

func score(amounts []int, ingredients []ingredient) int {
	capacity, durability, flavor, texture, calories := 0, 0, 0, 0, 0

	for i := 0; i < len(ingredients); i++ {
		capacity += amounts[i] * ingredients[i].capacity
		durability += amounts[i] * ingredients[i].durability
		flavor += amounts[i] * ingredients[i].flavor
		texture += amounts[i] * ingredients[i].texture
		calories += amounts[i] * ingredients[i].calories
	}

	if capacity < 0 || durability < 0 || flavor < 0 || texture < 0 {
		return 0
	}

	if calories != 500 {
		return 0
	}

	total := capacity * durability * flavor * texture
	return lib.Max(total, 0)
}

func increment(amounts []int) bool {
	for index := len(amounts) - 1; index > -1; index-- {
		switch amounts[index] {
		case 100:
			amounts[index] = 0
		default:
			amounts[index] += 1
		}

		if amounts[index] != 0 {
			return true
		}
	}

	return false
}

func solve(ingredients []ingredient) int {
	max := 0
	amounts := make([]int, len(ingredients))

	for increment(amounts) {
		sum := 0
		for _, v := range amounts {
			sum += v
		}
		if sum != 100 {
			continue
		}

		max = lib.Max(max, score(amounts, ingredients))
	}

	return max
}

func main() {
	solver := lib.Solver[[]ingredient, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseRegExp, parse)),
		SolveF: solve,
	}

	solver.Expect("Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8\nCinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3", 57600000)
	solver.Verify(15862900)
}
