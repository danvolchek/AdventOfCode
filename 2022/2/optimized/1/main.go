package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

func parse(line string) Round {
	parts := strings.Split(line, " ")
	return Round{
		opponent: parseChoice(parts[0]),
		you:      parseChoice(parts[1]),
	}
}

func parseChoice(str string) Choice {
	switch str {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
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

// beats[a][b] returns whether a beats b.
var beats = map[Choice]map[Choice]bool{
	Rock:     {Scissors: true},
	Paper:    {Rock: true},
	Scissors: {Paper: true},
}

// Beats returns the outcome of c against other.
func (c Choice) Beats(other Choice) Outcome {
	if beats[c][other] {
		return Win
	} else if c == other {
		return Tie
	} else {
		return Lose
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
	opponent, you Choice
}

// Score returns the score of a round for you.
func (r Round) Score() int {
	// The outcome of a round is whether your choice beats your opponents choice.
	outcome := r.you.Beats(r.opponent)

	// The score of a round is the score of the outcome plus the score of your choice.
	return outcome.Score() + r.you.Score()
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
