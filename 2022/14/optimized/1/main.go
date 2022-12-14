package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Coord struct {
	x, y int
}

func parse(line string) []Coord {
	parts := strings.Split(line, " -> ")

	coords := lib.Map(parts, func(part string) Coord {
		ints := lib.Ints(part)
		return Coord{
			x: ints[0],
			y: ints[1],
		}
	})

	return coords
}

type Cave struct {
	blocked lib.Set[Coord]

	minX, maxX, maxY int

	sandPoured int
}

func (c *Cave) Construct(paths [][]Coord) {
	for i, path := range paths {
		curr := path[0]

		for j, next := range path[1:] {
			// record blocked positions
			minX, maxX := lib.Min(curr.x, next.x), lib.Max(curr.x, next.x)
			minY, maxY := lib.Min(curr.y, next.y), lib.Max(curr.y, next.y)

			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					c.blocked.Add(Coord{x: x, y: y})
				}
			}

			curr = next

			// store cave bounds to determine if fallen into abyss
			if i == 0 && j == 0 {
				c.minX = minX
			}
			c.minX = lib.Min(c.minX, minX)
			c.maxX = lib.Max(c.maxX, maxX)
			c.maxY = lib.Max(c.maxY, maxY)
		}
	}
}

// PourUnitOfSand drops one sand into the cave and returns whether
// the sand falls into the abyss.
func (c *Cave) PourUnitOfSand() bool {
	pos := Coord{x: 500, y: 0}

	canMove := func(newPos Coord) bool {
		if !c.blocked.Contains(newPos) {
			pos = newPos
			return true
		}

		return false
	}

	for {
		// fallen into abyss
		if pos.x < c.minX || pos.x > c.maxX || pos.y > c.maxY {
			return true
		}

		// down
		if canMove(Coord{x: pos.x, y: pos.y + 1}) {
			continue
		}

		// down left
		if canMove(Coord{x: pos.x - 1, y: pos.y + 1}) {
			continue
		}

		// down right
		if canMove(Coord{x: pos.x + 1, y: pos.y + 1}) {
			continue
		}

		break
	}

	// landed on a spot
	c.blocked.Add(pos)
	c.sandPoured += 1

	return false
}

func solve(paths [][]Coord) int {
	cave := &Cave{}
	cave.Construct(paths)

	for !cave.PourUnitOfSand() {
	}

	return cave.sandPoured
}

func main() {
	solver := lib.Solver[[][]Coord, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9\n", 24)
	solver.Verify(1001)
}
