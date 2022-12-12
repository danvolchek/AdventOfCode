package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Replacement struct {
	in, out string
}

type Input struct {
	replacements []Replacement
	target       string
}

func parse(data string) Input {
	before, after, ok := strings.Cut(data, "\n\n")
	if !ok {
		panic(ok)
	}

	input := Input{
		replacements: nil,
		target:       strings.TrimSpace(after),
	}

	for _, line := range strings.Split(before, "\n") {
		parts := strings.Split(line, " => ")
		input.replacements = append(input.replacements, Replacement{in: parts[0], out: parts[1]})
	}

	return input
}

func solve(input Input) int {
	molecules := make(map[string]bool)

	for _, replacement := range input.replacements {
		for i := 0; i < len(input.target)-len(replacement.in)+1; i++ {
			if input.target[i:i+len(replacement.in)] == replacement.in {
				molecule := input.target[:i] + replacement.out + input.target[i+len(replacement.in):]
				molecules[molecule] = true
			}
		}
	}

	return len(molecules)
}

func main() {
	solver := lib.Solver[Input, int]{
		ParseF: parse,
		SolveF: solve,
	}

	solver.ParseExpect("H => HO\nH => OH\n\nHOHOHO\n", Input{
		replacements: []Replacement{{in: "H", out: "HO"}, {in: "H", out: "OH"}},
		target:       "HOHOHO",
	})

	solver.Expect("H => HO\nH => OH\nO => HH\n\nHOH\n", 4)
	solver.Expect("H => HO\nH => OH\nO => HH\n\nHOHOHO\n", 7)
	solver.Verify(509)
}
