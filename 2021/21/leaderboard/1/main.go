package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "21", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()

	first := scanner.Text()

	scanner.Scan()

	second := scanner.Text()

	first = strings.Split(first, ": ")[1]
	second = strings.Split(second, ": ")[1]

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	firstPos, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}

	secondPos, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}

	firstScore, secondScore := 0, 0

	nextDieValue := 1
	dieRolls := 0
	for firstScore <= 1000 && secondScore <= 1000 {
		toAdd := 0
		for i := 0; i < 3; i++ {
			toAdd += nextDieValue
			nextDieValue += 1

			if nextDieValue == 101 {
				nextDieValue = 1
			}
			dieRolls += 1
		}
		for i := 0; i<toAdd; i++ {
			firstPos ++
			if firstPos == 11 {
				firstPos = 1
			}
		}

		firstScore += firstPos

		if firstScore >= 1000 {
			break
		}

		toAdd = 0
		for i := 0; i < 3; i++ {
			toAdd += nextDieValue
			nextDieValue += 1
			if nextDieValue == 101 {
				nextDieValue = 1
			}
			dieRolls += 1
		}
		for i := 0; i<toAdd; i++ {
			secondPos ++
			if secondPos == 11 {
				secondPos = 1
			}
		}

		secondScore += secondPos
	}

	fmt.Println(firstScore, secondScore, dieRolls)
	if firstScore < 1000 {
		fmt.Println(firstScore * dieRolls)
	} else {
		fmt.Println(secondScore * dieRolls)
	}
}

func main() {
	solve(strings.NewReader("Player 1 starting position: 4\nPlayer 2 starting position: 8\n"))
	solve(input())
}
