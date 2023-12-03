package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func parse(char byte) bool {
	return char == '#'
}

type Dir int

const (
	North Dir = iota
	South
	West
	East
)

func round(elves *lib.Set[lib.Pos], dirs []Dir) bool {
	proposed := make(map[lib.Pos][]lib.Pos)

	for _, elf := range elves.Items() {
		adjacent := lib.Filter(
			lib.AdjacentPosNoBoundsChecks(true, elf.Row, elf.Col),
			func(p lib.Pos) bool {
				return elves.Contains(p)
			},
		)

		if len(adjacent) == 0 {
			continue
		}

		check := func(pos lib.Pos) bool {
			target := elf.Add(pos)

			allEmpty := true
			for i := -1; i <= 1; i++ {
				shouldBeEmpty := target
				if pos.Row == 0 {
					shouldBeEmpty.Row += i
				} else {
					shouldBeEmpty.Col += i
				}

				if elves.Contains(shouldBeEmpty) {
					allEmpty = false
					break
				}
			}

			if allEmpty {
				proposed[target] = append(proposed[target], elf)
			}

			return allEmpty
		}

	outer:
		for _, dir := range dirs {
			switch dir {
			case North:
				if check(lib.Pos{Row: -1}) {
					break outer
				}
			case South:
				if check(lib.Pos{Row: 1}) {
					break outer
				}
			case West:
				if check(lib.Pos{Col: -1}) {
					break outer
				}
			case East:
				if check(lib.Pos{Col: 1}) {
					break outer
				}
			}
		}
	}

	for pos, proposers := range proposed {
		if len(proposers) == 1 {
			elves.Remove(proposers[0])
			elves.Add(pos)
		}
	}

	return len(proposed) == 0
}

func solve(lines [][]bool) int {
	elfMap := &lib.Set[lib.Pos]{}

	for y, row := range lines {
		for x, hasElf := range row {
			if hasElf {
				elfMap.Add(lib.Pos{Row: y, Col: x})
			}
		}
	}

	dirs := []Dir{North, South, West, East}

	i := 1
	for !round(elfMap, dirs) {
		i++
		dirs = append(dirs[1:], dirs[0])
	}

	return i
}

func main() {
	solver := lib.Solver[[][]bool, int]{
		ParseF: lib.ParseGrid(parse),
		SolveF: solve,
	}

	solver.Expect("....#..\n..###.#\n#...#.#\n.#...##\n#.###..\n##.#.##\n.#..#..", 20)
	solver.Verify(980)
}
