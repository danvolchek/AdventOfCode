package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
)

const (
	outletJoltage        = 0
	minAdapterDifference = 1
	maxAdapterDifference = 3
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "10", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) []int {
	// start with the outlet
	adapters := []int{0}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Text()

		adapter, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, adapter)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return adapters
}

func solve(adapters []int) int {
	sort.Ints(adapters)

	// a map of joltage adapter -> the number of distinct ways that joltage can be reached
	arrangements := make(map[int]int, len(adapters))

	// there's only one way to reach the outlet's joltage: to use the outlet
	arrangements[outletJoltage] = 1

	for _, target := range adapters {
		// we can reach this joltage using any smaller supported joltage adapter
		for step := minAdapterDifference; step <= maxAdapterDifference; step++ {
			arrangements[target] += arrangements[target-step]
		}
	}

	// once we're at the highest joltage, there's only one way to reach the device's joltage: to use the device
	// this doesn't affect the number of arrangements, so we can just return the number of ways to reach the largest adapter

	return arrangements[adapters[len(adapters) - 1]]
}

func main() {
	fmt.Println(solve(parse(input())))
}
