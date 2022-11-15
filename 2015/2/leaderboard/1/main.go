package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"strings"
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

func parse(line string) present {
	parts := strings.Split(line, "x")
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
	lib.TestSolveParseLines("2x3x4", parse, solve)
	lib.SolveParseLines(input(), parse, solve)
}
