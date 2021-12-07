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
	input, err := os.Open(path.Join("2021", "7", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var crabs []int

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		for _, parse := range parts {
			v, _ := strconv.Atoi(parse)
			crabs = append(crabs, v)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	min := 9999999
	for _, c := range crabs {
		if align(c, crabs) < min {
			min = align(c, crabs)
		}
	}

	fmt.Println(min)
}

func align(v int, crabs []int) int {
	sum := 0
	for _, c := range crabs {
		sum += abs(c - v)
	}

	return sum
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}

	return v
}

func main() {
	solve(strings.NewReader("16,1,2,0,4,2,7,1,2,14"))
	solve(input())
}
