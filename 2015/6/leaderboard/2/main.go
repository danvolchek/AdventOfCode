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

var lineRegexp = regexp.MustCompile(`(.+) (\d+),(\d+) through (\d+),(\d+)`)

func parse(line string) instruction {
	result := lineRegexp.FindAllStringSubmatch(line, -1)

	parts := result[0][1:]

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
			for x := instr.start.x; x <= instr.end.x; x += 1 {
				for y := instr.start.y; y <= instr.end.y; y += 1 {
					grid[pos{x: x, y: y}] += 1
				}
			}
		case off:
			for x := instr.start.x; x <= instr.end.x; x += 1 {
				for y := instr.start.y; y <= instr.end.y; y += 1 {
					if grid[pos{x: x, y: y}] > 0 {
						grid[pos{x: x, y: y}] -= 1
					}
				}
			}
		case toggle:
			for x := instr.start.x; x <= instr.end.x; x += 1 {
				for y := instr.start.y; y <= instr.end.y; y += 1 {
					grid[pos{x: x, y: y}] += 2
				}
			}
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
	lib.TestSolveParseLines("turn on 0,0 through 999,999", parse, solve)
	lib.TestSolveParseLines("toggle 0,0 through 999,0", parse, solve)
	lib.TestSolveParseLines("turn on 499,499 through 500,500", parse, solve)
	lib.SolveParseLines(input(), parse, solve)
}
