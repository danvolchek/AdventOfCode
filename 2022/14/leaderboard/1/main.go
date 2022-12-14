package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Coord struct {
	x, y int
}

type Path struct {
	coords []Coord
}

func parse(line string) Path {
	parts := strings.Split(line, " -> ")

	coords := lib.Map(parts, func(part string) Coord {
		ints := lib.Ints(part)
		return Coord{
			x: ints[0],
			y: ints[1],
		}
	})

	return Path{coords: coords}
}

type Tile int

const (
	Air Tile = iota
	Rock
	Sand
)

func doSand(grid map[Coord]Tile, spot Coord, sX, lX, lY int) bool {
	if spot.x < sX || spot.x > lX || spot.y > lY {
		return true
	}
	oneBelow := Coord{x: spot.x, y: spot.y + 1}
	below, ok := grid[oneBelow]
	if !ok || below == Air {
		return doSand(grid, oneBelow, sX, lX, lY)
	}

	diagLeft := Coord{x: spot.x - 1, y: spot.y + 1}
	below, ok = grid[diagLeft]
	if !ok || below == Air {
		return doSand(grid, diagLeft, sX, lX, lY)
	}

	diagRight := Coord{x: spot.x + 1, y: spot.y + 1}
	below, ok = grid[diagRight]
	if !ok || below == Air {
		return doSand(grid, diagRight, sX, lX, lY)
	}

	grid[spot] = Sand
	return false
}

func solve(paths []Path) int {
	grid := make(map[Coord]Tile)

	smallestX, largestX := 9999999, -9999999
	largestY := -9999999

	for _, line := range paths {
		curr := line.coords[0]

		for _, item := range line.coords {
			smallestX = lib.Min(smallestX, item.x)
			largestX = lib.Max(largestX, item.x)
			largestY = lib.Max(largestY, item.y)
		}

		for _, next := range line.coords[1:] {

			if curr.x != next.x {
				for i := lib.Min(curr.x, next.x); i <= lib.Max(curr.x, next.x); i++ {
					grid[Coord{x: i, y: curr.y}] = Rock
				}
			} else {
				for i := lib.Min(curr.y, next.y); i <= lib.Max(curr.y, next.y); i++ {
					grid[Coord{x: curr.x, y: i}] = Rock
				}
			}

			curr = next
		}
	}

	i := 0
	for !doSand(grid, Coord{x: 500, y: 0}, smallestX, largestX, largestY) {
		i += 1
	}

	return i
}

func main() {
	solver := lib.Solver[[]Path, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9\n", 24)
	solver.Solve()
}
