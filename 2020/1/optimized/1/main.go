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
	input, err := os.Open(path.Join("2020", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) map[int]bool {
	expenses := make(map[int]bool)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		expenses[expense] = true
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return expenses
}

func solve(expenses map[int]bool) int {
	const target = 2020

	for expense := range expenses {
		// look for whether the other needed value exists
		needed := target - expense

		if _, ok := expenses[needed]; ok {
			return expense * needed
		}
	}

	panic("no matches")
}

func main() {
	fmt.Println(solve(parse(input())))
}
