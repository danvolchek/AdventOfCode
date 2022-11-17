package main

import (
	"regexp"

	"github.com/danvolchek/AdventOfCode/lib"
)

type happinessChange struct {
	target, neighbor string
	change           int
}

var parseRegExp = regexp.MustCompile(`(.*) would (.*) (.*) happiness units by sitting next to (.*)\.`)

func parse(parts []string) happinessChange {
	var multiplier int
	switch parts[1] {
	case "gain":
		multiplier = 1
	case "lose":
		multiplier = -1
	default:
		panic(parts)
	}

	return happinessChange{
		target:   parts[0],
		neighbor: parts[3],
		change:   multiplier * lib.Atoi(parts[2]),
	}
}

func solve(changes []happinessChange) int {
	peopleMap := map[string]map[string]int{}
	for _, change := range changes {
		if _, ok := peopleMap[change.target]; !ok {
			peopleMap[change.target] = map[string]int{}
		}
		peopleMap[change.target][change.neighbor] = change.change
	}

	people := lib.Keys(peopleMap)

	peopleMap["you"] = map[string]int{}
	for _, person := range people {
		peopleMap["you"][person] = 0
		peopleMap[person]["you"] = 0
	}

	people = append(people, "you")

	calcHappy := func(seatingArrangement []string) int {
		total := 0
		for i := range seatingArrangement {
			total += peopleMap[seatingArrangement[i]][seatingArrangement[(i+1)%len(peopleMap)]]
			total += peopleMap[seatingArrangement[i]][seatingArrangement[(i-1+len(peopleMap))%len(peopleMap)]]
		}
		return total
	}

	seatingArrangements := lib.Permutations(people)

	happinessForArrangements := lib.Map(seatingArrangements, calcHappy)

	return lib.Max(happinessForArrangements...)
}

func main() {
	solver := lib.Solver[[]happinessChange, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseRegExp, parse)),
		SolveF: solve,
	}

	solver.Verify(601)
}
