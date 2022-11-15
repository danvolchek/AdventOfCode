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

func solve(r io.Reader) {
	instructions := lib.Must(io.ReadAll(r))

	position := pos{x: 0, y: 0}

	gifts := make(map[pos]int)

	gifts[position] = 1

	for _, instruction := range instructions {
		switch instruction {
		case '^':
			position.y -= 1
		case 'v':
			position.y += 1
		case '>':
			position.x += 1
		case '<':
			position.x -= 1
		default:
			panic(string(instruction))
		}

		gifts[position] += 1
	}

	fmt.Println(len(gifts))
}

func main() {
	solve(strings.NewReader("^>v<"))
	solve(strings.NewReader("^v^v^v^v^v"))
	solve(input())
}
