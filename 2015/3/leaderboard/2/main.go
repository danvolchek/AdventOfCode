package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type pos struct {
	x, y int
}

type gifter struct {
	position pos
	gifts    map[pos]int
}

func (g *gifter) Move(instruction byte) {
	switch instruction {
	case '^':
		g.position.y -= 1
	case 'v':
		g.position.y += 1
	case '>':
		g.position.x += 1
	case '<':
		g.position.x -= 1
	default:
		panic(string(instruction))
	}

	g.gifts[g.position] += 1
}

func solve(r io.Reader) {
	instructions := lib.Must(io.ReadAll(r))

	santa := gifter{
		gifts: map[pos]int{
			{x: 0, y: 0}: 1,
		},
	}
	roboSanta := gifter{
		gifts: map[pos]int{
			{x: 0, y: 0}: 1,
		},
	}

	santaTurn := true

	for _, instruction := range instructions {
		if santaTurn {
			santa.Move(instruction)
		} else {
			roboSanta.Move(instruction)
		}

		santaTurn = !santaTurn
	}

	gotGifts := make(map[pos]bool)

	for position := range santa.gifts {
		gotGifts[position] = true
	}

	for position := range roboSanta.gifts {
		gotGifts[position] = true
	}

	fmt.Println(len(gotGifts))
}

func main() {
	solve(strings.NewReader("^>v<"))
	solve(strings.NewReader("^v^v^v^v^v"))
	solve(input())
}
