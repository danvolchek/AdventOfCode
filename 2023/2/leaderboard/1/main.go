package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strconv"
	"strings"
)

type game struct {
	id   int
	sets []map[string]int
}

func parse(line string) game {
	before, after, ok := strings.Cut(line, ": ")
	if !ok {
		panic(line)
	}

	id := lib.Int(before)

	var groups []map[string]int

	parts := strings.Split(after, "; ")

	for _, part := range parts {
		group := make(map[string]int)

		for _, cube := range strings.Split(part, ", ") {
			num := lib.Int(cube)
			color := strings.TrimSpace(strings.ReplaceAll(cube, strconv.Itoa(num), " "))

			group[color] = num
		}

		groups = append(groups, group)
	}

	return game{
		id:   id,
		sets: groups,
	}
}

func (g game) valid(constraints map[string]int) bool {
	for _, pull := range g.sets {
		for constraintColor, constraintAmount := range constraints {
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

func solve(lines []game) int {
	return lib.SumSlice(lib.Map(lib.Filter(lines, func(g game) bool {
		return g.valid(map[string]int{"red": 12, "green": 13, "blue": 14})
	}), func(g game) int {
		return g.id
	}))
}

func main() {
	solver := lib.Solver[[]game, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 8)
	solver.Verify(2505)
}
