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

type group struct {
	answers map[byte]int
	people  int
}

func parse(r io.Reader) []group {
	var groups []group

	currentGroup := group{
		answers: make(map[byte]int),
	}
	finishGroup := func() {
		groups = append(groups, currentGroup)

		currentGroup = group{
			answers: make(map[byte]int),
		}
	}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := scanner.Bytes()

		if len(row) == 0 {
			finishGroup()
			continue
		}

		currentGroup.people += 1
		for i := 0; i < len(row); i++ {
			currentGroup.answers[row[i]] += 1
		}

	}

	finishGroup()

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return groups
}

func solve(groups []group) int {
	total := 0

	for _, group := range groups {
		for _, numAnswers := range group.answers {
			if numAnswers == group.people {
				total += 1
			}
		}
	}

	return total
}

func main() {
	fmt.Println(solve(parse(input())))
}
