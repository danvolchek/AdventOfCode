package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Assignment struct {
	start, stop int
}

type Pair struct {
	first, second Assignment
}

func parse(line string) Pair {
	pairs := strings.Split(line, ",")
	numsF, numsS := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")

	return Pair{
		first: Assignment{
			start: lib.Atoi(numsF[0]),
			stop:  lib.Atoi(numsF[1]),
		},
		second: Assignment{
			start: lib.Atoi(numsS[0]),
			stop:  lib.Atoi(numsS[1]),
		},
	}
}

func solve(elfPairs []Pair) int {
	sum := 0

	// contains returns whether s contains f
	contains := func(f, s Assignment) bool {
		return s.start <= f.start && s.stop >= f.stop
	}

	// count the pairs where one elf contains the other
	for _, elfPair := range elfPairs {
		if contains(elfPair.first, elfPair.second) || contains(elfPair.second, elfPair.first) {
			sum += 1
		}
	}

	return sum
}

func main() {
	solver := lib.Solver[[]Pair, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8\n", 2)
	solver.Verify(487)
}
