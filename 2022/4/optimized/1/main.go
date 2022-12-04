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
		ParseF: lib.ParseLine(lib.ParseRegexp(reg, parse)),
		SolveF: solve,
	}

	solver.Expect("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8\n", 2)
	solver.Verify(487)
}
