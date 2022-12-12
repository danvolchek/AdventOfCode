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

func neighbors(replacements []Replacement, current string) []string {
	var s lib.OrderedSet[string]

	for _, replacement := range replacements {
		for i := 0; i < len(current)-len(replacement.in)+1; i++ {
			if current[i:i+len(replacement.in)] == replacement.in {
				molecule := current[:i] + replacement.out + current[i+len(replacement.in):]
				s.Add(molecule)
			}
		}
	}

	return s.Items()
}

type node struct {
	molecule string

	replacements []Replacement
}

func (n *node) Id() string {
	return n.molecule
}

func (n *node) Adjacent() []*node {
	neighborMolecules := neighbors(n.replacements, n.molecule)
	return lib.Map(neighborMolecules, func(molecule string) *node { return &node{molecule: molecule, replacements: n.replacements} })
}

func explore(input Input) []string {
	/*start := &node{
		molecule:     "e",
		replacements: input.replacements,
	}*/

	// Removing the id field from lib.BFS makes this at-runtime creation of adjacent nodes not work anymore
	// This approach didn't work anyway - the search space got way too large, BFS probably isn't right
	//return lib.Map(lib.BFS(start, input.target), func(n *node) string { return n.molecule })
	return nil
}

func solve(input Input) int {
	return len(explore(input)) - 1
}

func main() {
	graphSolver := lib.Solver[Input, []string]{
		ParseF: parse,
		SolveF: explore,
	}

	graphSolver.Expect("e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOH\n", []string{"e", "O", "HH", "HOH"})
	graphSolver.Expect("e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOHOHO\n", []string{"e", "H", "HO", "HHH", "HHHO", "HHOHO", "HOHOHO"})

	solver := lib.Solver[Input, int]{
		ParseF: parse,
		SolveF: solve,
	}

	solver.ParseExpect("H => HO\nH => OH\n\nHOHOHO\n", Input{
		replacements: []Replacement{{in: "H", out: "HO"}, {in: "H", out: "OH"}},
		target:       "HOHOHO",
	})

	solver.Expect("e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOH\n", 3)
	solver.Expect("e => H\ne => O\nH => HO\nH => OH\nO => HH\n\nHOHOHO\n", 6)

	// Note for me later: the search space is so big that solving the input runs out of memory and the process crashes
	// the search space needs to be reduced somehow, or the problem be formulated not as graph search.
	// solver.Solve()
}
