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

	diff1 := 0
	diff3 := 0
	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		switch diff {
		case 1:
			diff1 += 1
		case 3:
			diff3 += 1
		default:
			panic(fmt.Sprintf("%d: %d - %d is %d", i, adapters[i], adapters[i-1], diff))
		}
	}

	// add the device, which is always a 3 joltage difference
	diff3 += 1

	return diff1 * diff3
}

func main() {
	fmt.Println(solve(parse(input())))
}
