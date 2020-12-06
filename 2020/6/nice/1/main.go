package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "6", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) []map[byte]bool {
	var groups []map[byte]bool

	currentGroup := make(map[byte]bool)
	finishGroup := func() {
		groups = append(groups, currentGroup)

		currentGroup = make(map[byte]bool)
	}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := scanner.Bytes()

		if len(row) == 0 {
			finishGroup()
			continue
		}

		for i := 0; i < len(row); i++ {
			currentGroup[row[i]] = true
		}

	}

	finishGroup()

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return groups
}

func solve(groups []map[byte]bool) int {
	total := 0

	for _, group := range groups {
		total += len(group)
	}

	return total
}

func main() {
	fmt.Println(solve(parse(input())))
}
