package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"regexp"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "6", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type action int

const (
	on action = iota
	off
	toggle
)

type pos struct {
	x, y int
}

func (p pos) Range(o pos, action func(p pos)) {
	for x := p.x; x <= o.x; x += 1 {
		for y := p.y; y <= o.y; y += 1 {
			action(pos{x: x, y: y})
		}
	}
}

type instruction struct {
	act        action
	start, end pos
}

var parseReg = regexp.MustCompile(`(.+) (\d+),(\d+) through (\d+),(\d+)`)

func parse(parts []string) instruction {
	// line format: "(turn on|turn off|toggle) 123,456 through 789,100"

	var act action
	switch parts[0] {
	case "turn on":
		act = on
	case "turn off":
		act = off
	case "toggle":
		act = toggle
	default:
		panic(parts[0])
	}

	return instruction{
		act: act,
		start: pos{
			x: lib.Atoi(parts[1]),
			y: lib.Atoi(parts[2]),
		},
		end: pos{
			x: lib.Atoi(parts[3]),
			y: lib.Atoi(parts[4]),
		},
	}
}

func solve(instructions []instruction) int {
	grid := make(map[pos]bool)

	for _, instr := range instructions {
		switch instr.act {
		case on:
			instr.start.Range(instr.end, func(p pos) {
				grid[p] = true
			})
		case off:
			instr.start.Range(instr.end, func(p pos) {
				grid[p] = false
			})
		case toggle:
			instr.start.Range(instr.end, func(p pos) {
				grid[p] = !grid[p]
			})
		default:
			panic(instr.act)
		}
	}

	totalLit := 0
	for _, isLit := range grid {
		if isLit {
			totalLit += 1
		}
	}

	return totalLit
}

func main() {
	solver := lib.Solver[[]instruction, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseReg, parse)),
		SolveF: solve,
	}

	solver.Expect("turn on 0,0 through 999,999", 1000000)
	solver.Expect("toggle 0,0 through 999,0", 1000)
	solver.Expect("turn on 499,499 through 500,500", 4)
	solver.Verify(input(), 377891)
}
