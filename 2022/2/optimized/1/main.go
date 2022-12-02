package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

func parse(line string) Round {
	parts := strings.Split(line, " ")
	return Round{
		opponent: parts[0],
		you:      parts[1],
	}
}

var score = map[string]map[string]int{
	// Rock
	"A": {
		// Tie, Rock
		"X": 3 + 1,
		// Win, Paper
		"Y": 6 + 2,
		// Lose, Scissors
		"Z": 0 + 3,
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
		// Win, Rock
		"X": 6 + 1,
		// Lose, Paper
		"Y": 0 + 2,
		// Tie, Scissors
		"Z": 3 + 3,
	},
}

// Round represents a round of rock paper scissors.
type Round struct {
	opponent, you string
}

// Score returns the score of a round for you.
func (r Round) Score() int {
	return score[r.opponent][r.you]
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

	solver.Expect("A Y\nB X\nC Z", 15)
	solver.Verify(8933)
}
