package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
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
	// convert the slice of adapter joltages to map for faster lookup, while also finding maximum
	max := 0
	adaptersMap := make(map[int]bool, len(adapters))

	for _, adapter := range adapters {
		adaptersMap[adapter] = true

		if adapter > max {
			max = adapter
		}
	}

	return numArrangements(adaptersMap, max)
}

func numArrangements(adapters map[int]bool, maximumAdapter int) int {
	// a map of joltage level -> the number of distinct ways that joltage level can be reached
	arrangements := make(map[int]int)

	// there's only one way to reach the outlet's joltage: to use the outlet
	arrangements[outletJoltage] = 1

	for target := 1; target <= maximumAdapter; target++ {
		// if there's no adapter for this target joltage then it cannot be reached
		if !adapters[target] {
			continue
		}

		// if there are adapters at a joltage lower than but still supported by this one,
		// we can reach this joltage using any of those adapters
		for step := minAdapterDifference; step <= maxAdapterDifference; step++ {
			lowerJoltage := target - step

			if adapters[lowerJoltage] {
				arrangements[target] += arrangements[lowerJoltage]
			}
		}
	}

	// once we're at the highest joltage, there's only one way to reach the device's joltage: to use the device
	// this doesn't affect the number of arrangements, so we can just return the number of ways to reach the largest adapter

	return arrangements[maximumAdapter]
}

func main() {
	fmt.Println(solve(parse(input())))
}
