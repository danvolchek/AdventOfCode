package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type Rucksack struct {
	first, second []byte
}

func parse(line string) Rucksack {
	half := len(line) / 2

	return Rucksack{
		first:  []byte(line[:half]),
		second: []byte(line[half:]),
	}
}

// common returns the item that is common between first and second.
func common(first, second []byte) byte {
	for _, f := range first {
		for _, s := range second {
			if f == s {
				return f
			}
		}
	}

	panic("no common item")
}

// priority returns the priority of an item.
func priority(item byte) int {
	if item <= 'Z' {
		return int(item-'A') + 27
	}

	return int(item-'a') + 1
}

func solve(rucksacks []Rucksack) int {
	sum := 0

	// sum up the priority of all the common items
	for _, rucksack := range rucksacks {
		sum += priority(common(rucksack.first, rucksack.second))
	}

	return sum
}

func main() {
	solver := lib.Solver[[]Rucksack, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw", 157)
	solver.Verify(8072)
}
