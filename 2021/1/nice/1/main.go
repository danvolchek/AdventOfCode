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

	lastMeasurement := math.MaxInt
	numIncreases := 0

	for scanner.Scan() {
		row := scanner.Text()

		measurement, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}

		if measurement > lastMeasurement {
			numIncreases += 1
		}

		lastMeasurement = measurement
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(numIncreases)
}

func main() {
	solve(input())
}
