package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

func parse(line string) Round {
	parts := strings.Split(line, " ")
	return Round{
		opponent: parts[0],
		outcome:  parts[1],
	}
}

const (
	rock     = 1
	paper    = 2
	scissors = 3

	win  = 6
	tie  = 3
	lose = 0
)

var score = map[string]map[string]int{
	// Rock
	"A": {
		"X": lose + scissors,
		"Y": tie + rock,
		"Z": win + paper,
	},
	// Paper
	"B": {
		"X": lose + rock,
		"Y": tie + paper,
		"Z": win + scissors,
	},
	// Scissors
	"C": {
		"X": lose + paper,
		"Y": tie + scissors,
		"Z": win + rock,
	},
}

// Round represents a round of rock paper scissors.
type Round struct {
	opponent, outcome string
}

// Score returns the score of a round for you.
func (r Round) Score() int {
	return score[r.opponent][r.outcome]
}

func solve(rounds []Round) int {
	// Add up all the scores of each round.
	return lib.SumSlice(lib.Map(rounds, Round.Score))
}

func main() {
	solver := lib.Solver[[]Round, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("A Y\nB X\nC Z", 12)
	solver.Verify(11998)
}
