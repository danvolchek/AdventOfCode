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
		adj := lib.Filter(lib.AdjacentPosNoBoundsChecks(elf, true),
			func(p lib.Pos) bool {
				return elfMap.Contains(p)
			},
		)

		if len(adj) == 0 {
			continue
		}

	outer:
		for _, dir := range dirs {
			switch dir {
			case North:
				allEmpty := true
				for i := -1; i <= 1; i++ {
					if elfMap.Contains(lib.Pos{Row: elf.Row - 1, Col: elf.Col + i}) {
						allEmpty = false
						break
					}
				}
				if allEmpty {
					target := lib.Pos{Row: elf.Row - 1, Col: elf.Col}
					proposed[target] = append(proposed[target], elf)
					break outer
				}
			case South:
				allEmpty := true
				for i := -1; i <= 1; i++ {
					if elfMap.Contains(lib.Pos{Row: elf.Row + 1, Col: elf.Col + i}) {
						allEmpty = false
						break
					}
				}
				if allEmpty {
					target := lib.Pos{Row: elf.Row + 1, Col: elf.Col}
					proposed[target] = append(proposed[target], elf)
					break outer
				}
			case West:
				allEmpty := true
				for i := -1; i <= 1; i++ {
					if elfMap.Contains(lib.Pos{Row: elf.Row + i, Col: elf.Col - 1}) {
						allEmpty = false
						break
					}
				}
				if allEmpty {
					target := lib.Pos{Row: elf.Row, Col: elf.Col - 1}
					proposed[target] = append(proposed[target], elf)
					break outer
				}
			case East:
				allEmpty := true
				for i := -1; i <= 1; i++ {
					if elfMap.Contains(lib.Pos{Row: elf.Row + i, Col: elf.Col + 1}) {
						allEmpty = false
						break
					}
				}
				if allEmpty {
					target := lib.Pos{Row: elf.Row, Col: elf.Col + 1}
					proposed[target] = append(proposed[target], elf)
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
		if elf.Row < min.Row {
			min.Row = elf.Row
		}

		if elf.Row > max.Row {
			max.Row = elf.Row
		}

		if elf.Col < min.Col {
			min.Col = elf.Col
		}

		if elf.Col > max.Col {
			max.Col = elf.Col
		}
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
