package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"strings"
)

func parse(chunk string) int {
	return lib.SumSlice(lib.Map(strings.Split(strings.TrimSpace(chunk), "\n"), lib.Atoi))
}

func solve(chunks []int) int {
	slices.Sort(chunks)

	length := len(chunks)

	return chunks[length-1] + chunks[length-2] + chunks[length-3]
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseChunks(parse),
		SolveF: solve,
	}

	solver.Expect("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", 45000)
	solver.Verify(211189)
}
