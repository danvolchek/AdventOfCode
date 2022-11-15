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

type instruction struct {
	act        action
	start, end pos
}

func (p pos) Range(o pos, action func(p pos)) {
	for x := p.x; x <= o.x; x += 1 {
		for y := p.y; y <= o.y; y += 1 {
			action(pos{x: x, y: y})
		}
	}
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
	grid := make(map[pos]int)

	for _, instr := range instructions {
		switch instr.act {
		case on:
			instr.start.Range(instr.end, func(p pos) {
				grid[p] += 1
			})
		case off:
			instr.start.Range(instr.end, func(p pos) {
				if grid[p] > 0 {
					grid[p] -= 1
				}
			})
		case toggle:
			instr.start.Range(instr.end, func(p pos) {
				grid[p] += 2
			})
		default:
			panic(instr.act)
		}
	}

	totalBrightness := 0
	for _, isLit := range grid {
		totalBrightness += isLit
	}

	return totalBrightness
}

func main() {
	solver := lib.Solver[[]instruction, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseReg, parse)),
		SolveF: solve,
	}

	solver.Expect("turn on 0,0 through 999,999", 1000000)
	solver.Expect("toggle 0,0 through 999,0", 2000)
	solver.Verify(input(), 14110788)
}
