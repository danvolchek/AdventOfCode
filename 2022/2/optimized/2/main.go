package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

func parse(line string) Round {
	parts := strings.Split(line, " ")
	return Round{
		opponent: parseChoice(parts[0]),
		outcome:  parseOutcome(parts[1]),
	}
}

func parseChoice(str string) Choice {
	switch str {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	default:
		panic(str)
	}
}

func parseOutcome(str string) Outcome {
	switch str {
	case "X":
		return Lose
	case "Y":
		return Tie
	case "Z":
		return Win
	default:
		panic(str)
	}
}

//go:generate go run golang.org/x/tools/cmd/stringer -type=Choice,Outcome -output enums_string.go

// Choice represents a choice in the game of rock paper scissors.
type Choice int

const (
	Rock Choice = iota
	Paper
	Scissors
)

// Outcome represents the outcome of a round of rock paper scissors.
type Outcome int

const (
	Win Outcome = iota
	Tie
	Lose
)

// Score returns the score of an outcome for you.
func (o Outcome) Score() int {
	switch o {
	case Win:
		return 6
	case Tie:
		return 3
	case Lose:
		return 0
	default:
		panic(o)
	}
}

// beats[a] returns the choice that a beats.
var beats = map[Choice]Choice{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

// Pick returns a choice such that the returned choice against c is outcome.
func (c Choice) Pick(outcome Outcome) Choice {
	switch outcome {
	case Win:
		return beats[beats[c]]
	case Tie:
		return c
	case Lose:
		return beats[c]
	default:
		panic(c)
	}
}

// Score returns the score of a choice for you.
func (c Choice) Score() int {
	switch c {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		panic(c)
	}
}

// Round represents a round of rock paper scissors.
type Round struct {
	opponent Choice
	outcome  Outcome
}

// Score returns the score of a round for you.
func (r Round) Score() int {
	// Your choice is the one that results in the outcome based on your opponents choice.
	you := r.opponent.Pick(r.outcome)

	// The score of a round is the score of the outcome plus the score of your choice.
	return r.outcome.Score() + you.Score()
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
