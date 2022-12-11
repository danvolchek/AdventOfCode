package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

func parse(chunk string) int {
	return lib.SumSlice(lib.Map(strings.Split(strings.TrimSpace(chunk), "\n"), lib.Atoi))
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseChunks(parse),
		SolveF: lib.MaxSlice[int],
	}

	solver.Expect("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", 24000)
	solver.Verify(71471)
}
