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
	input, err := os.Open(path.Join("2020", "13", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) (int, []int) {
	scanner := bufio.NewScanner(r)

	chomp := func() string {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
		}
		return scanner.Text()
	}

	earliestTime, err := strconv.Atoi(chomp())
	if err != nil {
		panic(err)
	}

	var busses []int
	rawBusses := chomp()
	for _, item := range strings.Split(rawBusses, ",") {
		if item == "x" {
			continue
		}

		val, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		busses = append(busses, val)
	}

	return earliestTime, busses
}

func solve(earliestTime int, busses []int) int {
	earliestBus := 0
	minimumWaitTime := -1

	for _, bus := range busses {
		waitTime := bus - earliestTime%bus
		if minimumWaitTime == -1 || waitTime < minimumWaitTime {
			earliestBus = bus
			minimumWaitTime = waitTime
		}
	}

	return earliestBus * minimumWaitTime
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("939\n7,13,x,x,59,x,31,19"))))
	fmt.Println(solve(parse(input())))
}
