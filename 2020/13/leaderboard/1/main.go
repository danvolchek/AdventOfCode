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

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	earliest := 0
	var busses []int

	scanner.Scan()
	first := scanner.Text()
	earliest, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	second := scanner.Text()
	for _, item := range strings.Split(second, ",") {
		if item == "x" {
			continue
		}

		val, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		busses = append(busses, val)
	}

	id := 0
	wait := 999999999

	for _, bus := range busses {
		if bus-earliest%bus < wait {
			id = bus
			wait = bus - earliest%bus
		}
	}

	fmt.Println(id, wait, id*wait)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(strings.NewReader("939\n7,13,x,x,59,x,31,19"))
	solve(input())
}
