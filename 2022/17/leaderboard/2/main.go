package main

import (
	"bytes"
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"time"
)

func parse(line []byte) []bool {
	line = bytes.TrimSpace(line)

	return lib.Map(line, func(b byte) bool {
		return b == '<'
	})
}

type rock struct {
	shape  []Pos
	height int

	wind *globalWind
	cave *lib.Set[Pos]
}

// 0 -> left, 1 -> right, 2 -> down
func (r *rock) valid(pos Pos) bool {
	for _, off := range r.shape {
		ppx := pos.x + off.x
		ppy := pos.y - off.y
		if ppx < 0 || ppx > 6 {
			return false
		}

		if ppy <= -1 {
			return false
		}

		if r.cave.Contains(Pos{x: ppx, y: ppy}) {
			return false
		}
	}
	return true
}

type globalWind struct {
	pattern []bool
	i       int
}

func (g *globalWind) get() bool {
	ret := g.pattern[g.i]

	g.i = (g.i + 1) % len(g.pattern)
	return ret
}

func (r *rock) Fall(highestY int) Pos {
	curr := Pos{x: 2, y: highestY + 3 + r.height}

	for {
		dirWind := -1
		if !r.wind.get() {
			dirWind = 1
		}

		curr.x += dirWind
		if ok := r.valid(curr); !ok {
			curr.x -= dirWind
		}

		curr.y -= 1
		if ok := r.valid(curr); !ok {
			curr.y += 1
			break
		}
	}

	for _, off := range r.shape {
		finalP := Pos{x: curr.x + off.x, y: curr.y - off.y}
		r.cave.Add(finalP)
	}

	return curr
}

type Pos struct {
	x, y int
}

func solve(lines []bool) int {
	cave := &lib.Set[Pos]{}

	wind := &globalWind{pattern: lines}

	rocks := []*rock{
		{
			shape:  []Pos{{x: 0}, {x: 1}, {x: 2}, {x: 3}},
			height: 1,
			wind:   wind,
			cave:   cave,
		},
		{
			shape:  []Pos{{x: 1}, {x: 0, y: 1}, {x: 1, y: 1}, {x: 2, y: 1}, {x: 1, y: 2}},
			height: 3,
			wind:   wind,
			cave:   cave,
		},
		{
			shape:  []Pos{{x: 2}, {x: 2, y: 1}, {x: 0, y: 2}, {x: 1, y: 2}, {x: 2, y: 2}},
			height: 3,
			wind:   wind,
			cave:   cave,
		},
		{
			shape:  []Pos{{y: 0}, {y: 1}, {y: 2}, {y: 3}},
			height: 4,
			wind:   wind,
			cave:   cave,
		},
		{
			shape:  []Pos{{}, {y: 1}, {x: 1}, {x: 1, y: 1}},
			height: 2,
			wind:   wind,
			cave:   cave,
		},
	}

	last := time.Now()
	highest := Pos{x: 0, y: -1}
	for i := 0; i < 1000000000000; i++ {
		if i%1000000 == 0 {
			nowNow := time.Now()
			rem := nowNow.Sub(last) * (time.Duration(1000000000000-i) / 1000000)
			fmt.Println(i, (100 * float64(i) / 1000000000000), nowNow.Sub(last), rem, nowNow.Add(rem))
			last = nowNow
		}
		// assume jet cycle starts over per rock
		landed := rocks[i%len(rocks)].Fall(highest.y)

		if landed.y > highest.y {
			highest = landed
		}
	}

	return highest.y + 1
}

func main() {
	solver := lib.Solver[[]bool, int]{
		ParseF: lib.ParseBytesFunc(parse),
		SolveF: solve,
	}

	solver.Expect(">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>", 1514285714288)
	solver.Solve()
}
