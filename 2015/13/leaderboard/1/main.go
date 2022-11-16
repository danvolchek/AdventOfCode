package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"regexp"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "13", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

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

	var testCase = `
Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.
`
	solver.Expect(strings.TrimSpace(testCase), 330)
	solver.Verify(input(), 618)
}
