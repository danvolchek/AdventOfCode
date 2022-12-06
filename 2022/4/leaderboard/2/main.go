package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
)

type SectionRange struct {
	start, stop int
}

func (s SectionRange) Overlap(o SectionRange) bool {
	return !(s.start > o.stop || s.stop < o.start)
}

type Pair struct {
	first, second SectionRange
}

var reg = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func parse(parts []string) Pair {
	return Pair{
		first: SectionRange{
			start: lib.Atoi(parts[0]),
			stop:  lib.Atoi(parts[1]),
		},
		second: SectionRange{
			start: lib.Atoi(parts[2]),
			stop:  lib.Atoi(parts[3]),
		},
	}
}

func solve(lines []Pair) int {
	amount := 0
	for _, line := range lines {
		if line.second.Overlap(line.first) {
			amount += 1
		}
	}

	return amount
}

func main() {
	solver := lib.Solver[[]Pair, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(reg, parse)),
		SolveF: solve,
	}

	solver.Expect("2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8\n", 4)
	solver.Verify(849)
}
