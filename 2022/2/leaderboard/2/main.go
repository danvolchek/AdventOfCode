package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type round struct {
	opponent, you string
}

func parse(line string) round {
	parts := strings.Split(line, " ")
	return round{
		opponent: toTerm(parts[0]),
		you:      toAct(parts[1]),
	}
}

func toTerm(line string) string {
	switch line {
	case "A", "X":
		return "R"
	case "B", "Y":
		return "P"
	case "C", "Z":
		return "S"
	default:
		panic(line)
	}
}

func toAct(line string) string {
	switch line {
	case "A", "X":
		return "L"
	case "B", "Y":
		return "D"
	case "C", "Z":
		return "W"
	default:
		panic(line)
	}
}

func solve(lines []round) int {
	score := 0

	for _, line := range lines {
		win := line.you == "W"
		tie := line.you == "D"

		you := ""
		if win {
			switch line.opponent {
			case "S":
				you = "R"
			case "R":
				you = "P"
			case "P":
				you = "S"
			}
		} else if tie {
			you = line.opponent
		} else {
			switch line.opponent {
			case "R":
				you = "S"
			case "P":
				you = "R"
			case "S":
				you = "P"
			}
		}
		switch you {
		case "R":
			score += 1
		case "P":
			score += 2
		case "S":
			score += 3
		}

		if win {
			score += 6
		} else if tie {
			score += 3
		}
	}

	return score
}

func main() {
	solver := lib.Solver[[]round, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("A Y\nB X\nC Z", 12)
	solver.Verify(11998)
}
