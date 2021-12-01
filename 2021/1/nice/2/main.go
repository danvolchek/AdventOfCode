package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	// can't use math.MaxInt on its own, because the sum will cause overflow
	measurementWindow := []int{math.MaxInt >> 2, math.MaxInt >> 2, math.MaxInt >> 2}
	windowIndex := 0
	numIncreases := 0

	for scanner.Scan() {
		row := scanner.Text()

		measurement, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}

		lastWindowSum := sum(measurementWindow)

		measurementWindow[windowIndex] = measurement
		windowIndex = (windowIndex + 1) % len(measurementWindow)

		currentWindowSum := sum(measurementWindow)
		if currentWindowSum > lastWindowSum {
			numIncreases++
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(numIncreases)
}

func sum(ints []int) int {
	var result int
	for _, value := range ints {
		result += value
	}

	return result
}

func main() {
	solve(input())
}
