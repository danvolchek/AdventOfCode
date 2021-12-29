package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "25", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type pos struct {
	x, y int
}

const (
	indicatorEast  = '>'
	indicatorSouth = 'v'
	indicatorEmpty = '.'
)

type cucumberRegion [][]byte

func (c cucumberRegion) move() bool {
	eastMoved := c.moveHerd(indicatorEast, pos{
		x: 1,
		y: 0,
	})

	southMoved := c.moveHerd(indicatorSouth, pos{
		x: 0,
		y: 1,
	})

	return eastMoved || southMoved
}

func (c cucumberRegion) moveHerd(herdIndicator byte, direction pos) bool {
	type move struct {
		from, to pos
	}
	var moves []move

	for y := 0; y < len(c); y++ {
		for x := 0; x < len(c[y]); x++ {
			if c[y][x] != herdIndicator {
				continue
			}

			nextX, nextY := x+direction.x, y+direction.y
			if nextX == len(c[y]) {
				nextX = 0
			}

			if nextY == len(c) {
				nextY = 0
			}

			if c[nextY][nextX] == indicatorEmpty {
				moves = append(moves, move{
					from: pos{x, y},
					to:   pos{nextX, nextY},
				})
			}
		}
	}

	for _, mv := range moves {
		c[mv.to.y][mv.to.x] = herdIndicator
		c[mv.from.y][mv.from.x] = indicatorEmpty
	}

	return len(moves) != 0
}

func (c cucumberRegion) String() string {
	var sb strings.Builder

	for _, row := range c {
		for _, cucumber := range row {
			sb.WriteByte(cucumber)
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func parse(r io.Reader) cucumberRegion {
	var cucumbers cucumberRegion

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()

		row := make([]byte, len(line))
		copy(row, line)

		cucumbers = append(cucumbers, row)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return cucumbers
}

func solve(r io.Reader) {
	cucumbers := parse(r)

	steps := 0

	for cucumbers.move() {
		steps += 1
	}

	fmt.Println(steps + 1)
}

func main() {
	solve(strings.NewReader("v...>>.vv>\n.vv>>.vv..\n>>.>v>...v\n>>v>>.>.v.\nv>v.vv.v..\n>.>>..v...\n.vv..>.>v.\nv.v..>>v.v\n....v..v.>"))
	solve(input())
}
