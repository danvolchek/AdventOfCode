package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "7", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) ([]int, int, int) {
	line, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	var crabs []int
	minCrab, maxCrab := -1, -1

	rawCrabs := strings.Split(strings.TrimSpace(string(line)), ",")
	for _, rawCrab := range rawCrabs {
		crab, err := strconv.Atoi(rawCrab)
		if err != nil {
			panic(err)
		}
		crabs = append(crabs, crab)

		if minCrab < 0 || crab < minCrab {
			minCrab = crab
		}

		if maxCrab < 0 || crab > maxCrab {
			maxCrab = crab
		}
	}

	return crabs, minCrab, maxCrab
}

func solve(r io.Reader) {
	crabs, minCrab, maxCrab := parse(r)

	minFuel := -1
	for position := minCrab; position <= maxCrab; position += 1 {
		fuel := totalFuel(position, crabs)
		if minFuel < 0 || fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)
}

func totalFuel(position int, crabs []int) int {
	fuel := 0
	for _, crab := range crabs {
		n := abs(crab - position)
		fuel += n * (n + 1) / 2
	}

	return fuel
}

func abs(value int) int {
	if value < 0 {
		return -1 * value
	}

	return value
}

func main() {
	solve(strings.NewReader("16,1,2,0,4,2,7,1,2,14"))
	solve(input())
}
