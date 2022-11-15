package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type present struct {
	l, w, h int
}

const split = "x"

func parse(parts []string) present {
	return present{
		l: lib.Atoi(parts[0]),
		w: lib.Atoi(parts[1]),
		h: lib.Atoi(parts[2]),
	}
}

func solve(presents []present) int {
	totalPaper := 0

	paperForPresent := func(l, w, h int) int {
		return 2*l*w + 2*w*h + 2*h*l + lib.Min(l*w, l*h, h*w)
	}

	for _, present := range presents {
		totalPaper += paperForPresent(present.l, present.w, present.h)
	}

	return totalPaper
}

func main() {
	solver := lib.Solver[[]present, int]{
		ParseF: lib.ParseLine(lib.ParseSplit(split, parse)),
		SolveF: solve,
	}

	solver.Expect("2x3x4", 58)
	solver.Expect("1x1x10", 43)
	solver.Solve(input())
}
