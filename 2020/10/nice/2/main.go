package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "10", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) []int {
	// start with outlet
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
	// convert slice of adapters to map for faster lookup, while also finding maximum
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
	// a map of joltage level -> number of ways the joltage level can be reached
	ways := make(map[int]int)

	// there's only 1 way to reach 0 joltage: to use the outlet
	ways[0] = 1

	for target := 1; target <= maximumAdapter; target++ {
		// if there's no adapter for this target joltage then it cannot be reached
		if !adapters[target] {
			continue
		}

		// if there are adapters at a joltage lower than but still supported by this one,
		// we can reach this joltage using any of those adapters
		for step := 1; step <= 3; step++ {
			if adapters[target - step] {
				ways[target] += ways[target-step]
			}
		}
	}

	// once we're at the highest joltage, there's only one more way to reach the device: to use the device
	// this doesn't affect the number of combinations

	return ways[maximumAdapter]
}

func main() {
	fmt.Println(solve(parse(input())))
}
