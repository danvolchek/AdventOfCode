package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
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
	var adapters []int

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

// optimized for readability: it's pretty easy to tell what's going on
func solve(adapters []int) int {
	sort.Ints(adapters)

	// a map of joltage adapter -> the number of distinct ways that joltage can be reached
	arrangements := make(map[int]int, len(adapters))

	// there's only one way to reach the outlet's joltage: to use the outlet
	arrangements[outletJoltage] = 1

	for _, target := range adapters {
		// we can reach this joltage level using any smaller supported joltage adapter
		for step := minAdapterDifference; step <= maxAdapterDifference; step++ {
			arrangements[target] += arrangements[target-step]
		}
	}

	// once we're at the highest joltage, there's only one way to reach the device's joltage: to use the device
	// this doesn't affect the number of arrangements, so we can just return the number of ways to reach the largest adapter
	return arrangements[adapters[len(adapters)-1]]
}

// optimized for space complexity: besides the input, this uses constant space
// this works because we only ever need to check the last 3 joltage levels
func solve2(adapters []int) int {
	sort.Ints(adapters)

	// a sliding window of the ways to reach four joltage levels, in decreasing joltage value as index increases
	// initialized to the number of ways to reach joltage 0, -1, -2, -3
	arrangements := [maxAdapterDifference - minAdapterDifference + 2]int{}

	// there's only one way to reach the outlet's joltage: to use the outlet
	arrangements[outletJoltage] = 1

	// the joltage level that the first item in the window represents
	// used to shift the sliding window the proper number of steps on non-contiguous adapter sequences
	windowIndex := 0

	// shifts the sliding window one step
	shift := func() {
		for i := 0; i < maxAdapterDifference; i++ {
			arrangements[maxAdapterDifference-i] = arrangements[maxAdapterDifference-i-1]
		}
		arrangements[0] = 0

		windowIndex += 1
	}

	for _, target := range adapters {
		// shift window so that the first index is the target joltage
		for windowIndex != target {
			shift()
		}

		// calculate ways to reach the target (using any smaller supported joltage adapter)
		for step := minAdapterDifference; step <= maxAdapterDifference; step++ {
			arrangements[0] += arrangements[step]
		}
	}

	// same comment on highest adapter joltage as the other solver
	return arrangements[0]
}

func main() {
	fmt.Print(solve(parse(strings.NewReader("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"))), " ")
	fmt.Println(solve2(parse(strings.NewReader("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"))))

	fmt.Print(solve(parse(input())), " ")
	fmt.Println(solve2(parse(input())))
}
