package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
)

type Assignment struct {
	start, stop int
}

type Pair struct {
	first, second Assignment
}

var reg = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func parse(parts []string) Pair {
	return Pair{
		first: Assignment{
			start: lib.Atoi(parts[0]),
			stop:  lib.Atoi(parts[1]),
		},
		second: Assignment{
			start: lib.Atoi(parts[2]),
			stop:  lib.Atoi(parts[3]),
		},
	}
}

func solve(elfPairs []Pair) int {
	sum := 0

	// overlap returns whether s overlaps with f
	overlap := func(f, s Assignment) bool {
		return !(s.start > f.stop || s.stop < f.start)
	}

	// count the pairs where the elves overlap
	for _, elfPair := range elfPairs {
		if overlap(elfPair.first, elfPair.second) {
			sum += 1
		}
	}

	return sum
}

func main() {
	solver := lib.Solver[[]Pair, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(reg, parse)),
		SolveF: solve,
	}

	solver.Expect("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8\n", 4)
	solver.Verify(849)
}
