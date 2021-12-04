package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

var rowRegex = regexp.MustCompile(`(\w+) (\d+)`)

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	depth := 0
	horizontalPosition := 0

	for scanner.Scan() {
		row := scanner.Text()

		parts := rowRegex.FindStringSubmatch(row)
		if len(parts) != 3 {
			panic(fmt.Sprintf("bad row %s", row))
		}

		operation := parts[1]
		rawValue := parts[2]
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			panic(err)
		}

		switch operation {
		case "down":
			depth += value
		case "up":
			depth -= value
		case "forward":
			horizontalPosition += value
		}
	}

	fmt.Printf("Depth: %v Horizontal Position: %v Answer: %v\n", depth, horizontalPosition, depth*horizontalPosition)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(input())
}
