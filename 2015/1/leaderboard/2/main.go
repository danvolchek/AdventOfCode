package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	floor := 0

	instructions, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	for i, instruction := range instructions {
		switch instruction {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			panic(string(instruction))
		}

		if floor == -1 {
			fmt.Println(i + 1)
			return
		}
	}
	panic("not found")
}

func main() {
	solve(strings.NewReader("()())"))
	solve(input())
}
