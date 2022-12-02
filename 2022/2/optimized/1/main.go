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

const (
	rock     = 1
	paper    = 2
	scissors = 3

	win  = 6
	tie  = 3
	lose = 0
)

// score[x][y] returns the score for you when your opponent chooses x, and you choose y
var score = map[string]map[string]int{
	// Rock
	"A": {
		"X": tie + rock,
		"Y": win + paper,
		"Z": lose + scissors,
	},
	// Paper
	"B": {
		"X": lose + rock,
		"Y": tie + paper,
		"Z": win + scissors,
	},
	// Scissors
	"C": {
		"X": win + rock,
		"Y": lose + paper,
		"Z": tie + scissors,
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
