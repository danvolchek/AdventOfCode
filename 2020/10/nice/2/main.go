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
		// we can reach the target joltage using any smaller supported joltage adapter
		for step := minAdapterDifference; step <= maxAdapterDifference; step++ {
			arrangements[target] += arrangements[target-step]
		}
	}

	// once we're at the highest joltage, there's only one way to reach the device's joltage: to use the device
	// this doesn't affect the number of arrangements, so we can just return the number of ways to reach the largest adapter
	return arrangements[adapters[len(adapters)-1]]
}

// optimized for space complexity: besides the input, this uses constant space
// it's the same algorithm as solve, but only stores the last 3 joltages as none are needed beyond that
func solve2(adapters []int) int {
	sort.Ints(adapters)

	// a sliding window of the ways to reach four joltage levels, in decreasing joltage value as index increases
	// initialized to the number of ways to reach joltage 0, -1, -2, -3
	arrangements := [maxAdapterDifference - minAdapterDifference + 2]int{}

	// same comment as in solve
	arrangements[outletJoltage] = 1

	// the joltage that the first item in the window represents
	// used to shift the sliding window the proper number of steps to get to the current adapter
	windowIndex := 0

	// shifts the sliding window one step
	shift := func() {
		// move elements one to the right
		for i := 0; i < maxAdapterDifference; i++ {
			arrangements[maxAdapterDifference-i] = arrangements[maxAdapterDifference-i-1]
		}

		// reset first element
		arrangements[0] = 0

		windowIndex += 1
	}

	for _, target := range adapters {
		// shift window so that the first index is the target joltage
		for windowIndex != target {
			shift()
		}

		// same comment as in solve
		for step := minAdapterDifference; step <= maxAdapterDifference; step++ {
			arrangements[0] += arrangements[step]
		}
	}

	// same comment as in solve
	return arrangements[0]
}

func main() {
	fmt.Print(solve(parse(strings.NewReader("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"))), " ")
	fmt.Println(solve2(parse(strings.NewReader("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"))))

	fmt.Print(solve(parse(input())), " ")
	fmt.Println(solve2(parse(input())))
}
