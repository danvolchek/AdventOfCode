package main

import (
	"fmt"
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

func round(elfMap *lib.Set[lib.Pos], dirs []Dir) {
	proposed := make(map[lib.Pos][]lib.Pos)

	for _, elf := range elfMap.Items() {
		adj := lib.Filter(lib.AdjacentPos(true, elf.Row, elf.Col),
			func(p lib.Pos) bool {
				return elfMap.Contains(p)
			},
		)

		if len(adj) == 0 {
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

				if elfMap.Contains(shouldBeEmpty) {
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
			elfMap.Remove(proposers[0])
			elfMap.Add(pos)
		}
	}
}

func bounds(elfMap *lib.Set[lib.Pos]) (lib.Pos, lib.Pos) {
	var min, max lib.Pos

	for _, elf := range elfMap.Items() {
		min = min.Min(elf)
		max = max.Max(elf)
	}

	return min, max
}

func printy(elfMap *lib.Set[lib.Pos]) {
	min, max := bounds(elfMap)

	for y := min.Row; y <= max.Row; y++ {
		for x := min.Col; x <= max.Col; x++ {
			if elfMap.Contains(lib.Pos{Row: y, Col: x}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
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

	//printy(elfMap)

	dirs := []Dir{North, South, West, East}

	for i := 0; i < 10; i++ {
		round(elfMap, dirs)
		dirs = append(dirs[1:], dirs[0])

		//printy(elfMap)
	}

	min, max := bounds(elfMap)
	sum := 0

	for y := min.Row; y <= max.Row; y++ {
		for x := min.Col; x <= max.Col; x++ {
			if !elfMap.Contains(lib.Pos{Row: y, Col: x}) {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	solver := lib.Solver[[][]bool, int]{
		ParseF: lib.ParseGrid(parse),
		SolveF: solve,
	}

	solver.Test(".....\n..##.\n..#..\n.....\n..##.\n.....")
	solver.Expect("....#..\n..###.#\n#...#.#\n.#...##\n#.###..\n##.#.##\n.#..#..", 110)
	solver.Verify(4249)
}
