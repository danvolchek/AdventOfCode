package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

// common returns the items that are common between first and second
func common(first, second []byte) []byte {
	var result lib.Set[byte]

	for _, f := range first {
		for _, s := range second {
			if f == s {
				result.Add(f)
			}
		}
	}

	return result.Items()
}

// priority returns the priority of an item.
func priority(item byte) int {
	if item <= 'Z' {
		return int(item-'A') + 27
	}

	return int(item-'a') + 1
}

func solve(rucksacks [][]byte) int {
	sum := 0

	// sum up the priority of the common item between every three elves
	for i := 0; i < len(rucksacks); i += 3 {
		first := rucksacks[i]
		second := rucksacks[i+1]
		third := rucksacks[i+2]

		sum += priority(common(first, common(second, third))[0])
	}

	return sum
}

func main() {
	solver := lib.Solver[[][]byte, int]{
		ParseF: lib.ParseLine(lib.ToByteSlice),
		SolveF: solve,
	}

	solver.Expect("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 70)
	solver.Verify(2567)
}
