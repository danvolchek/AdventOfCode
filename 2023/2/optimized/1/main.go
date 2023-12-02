package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
	"strings"
)

type game struct {
	id       int
	cubeSets []map[string]int
}

var constraintBag = map[string]int{"red": 12, "green": 13, "blue": 14}

func (g game) valid() bool {
	for _, pull := range g.cubeSets {
		for constraintColor, constraintAmount := range constraintBag {
			pullAmount, ok := pull[constraintColor]
			if !ok {
				continue
			}

			if pullAmount > constraintAmount {
				return false
			}
		}
	}

	return true
}

func (g game) Id() int {
	return g.id
}

var gameRegexp = regexp.MustCompile(`Game (\d+): (.*)`)

func parseGame(parts []string) game {
	id := lib.Atoi(parts[0])

	var cubeSets []map[string]int

	for _, rawPull := range strings.Split(parts[1], "; ") {
		cubeSet := make(map[string]int)

		for _, cubes := range strings.Split(rawPull, ", ") {
			description := strings.Split(cubes, " ")

			amount, color := lib.Atoi(description[0]), description[1]
			cubeSet[color] = amount
		}

		cubeSets = append(cubeSets, cubeSet)
	}

	return game{
		id:       id,
		cubeSets: cubeSets,
	}
}

func solve(lines []game) int {
	return lib.SumSlice(lib.Map(lib.Filter(lines, game.valid), game.Id))
}

func main() {
	solver := lib.Solver[[]game, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(gameRegexp, parseGame)),
		SolveF: solve,
	}

	solver.Expect("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 8)
	solver.Verify(2505)
}
