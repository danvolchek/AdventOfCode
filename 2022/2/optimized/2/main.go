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

var score = map[string]map[string]int{
	// Rock
	"A": {
		// Lose, Scissors
		"X": 0 + 3,
		// Tie, Rock
		"Y": 3 + 1,
		// Win, Paper
		"Z": 6 + 2,
	},
	// Paper
	"B": {
		// Lose, Rock
		"X": 0 + 1,
		// Tie, Paper
		"Y": 3 + 2,
		// Win, Scissors
		"Z": 6 + 3,
	},
	// Scissors
	"C": {
		// Lose, Paper
		"X": 0 + 2,
		// Tie, Scissors
		"Y": 3 + 3,
		// Win, Rock
		"Z": 6 + 1,
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
