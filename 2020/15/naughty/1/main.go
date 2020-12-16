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
	input, err := os.Open(path.Join("2020", "15", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	chomp := func() string {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
		}
		return scanner.Text()
	}

	raw := strings.Split(chomp(), ",")

	numbers := make([]int, len(raw))

	for i, rawNum := range raw {
		intNum, err := strconv.Atoi(rawNum)
		if err != nil {
			panic(err)
		}

		numbers[i] = intNum
	}

	fmt.Println(game(numbers, 2020))
}

func game(starting []int, limit int) int {
	memory := make(map[int]int)

	next := 0

	for i := 0; i < limit-1; i++ {
		if i < len(starting) {
			next = starting[i]
		}

		oldNext := next

		age, ok := memory[next]
		if !ok {
			next = 0
		} else {
			next = i - age
		}

		memory[oldNext] = i
	}

	return next
}

func main() {
	solve(strings.NewReader("0,3,6"))
	solve(strings.NewReader("1,3,2"))
	solve(strings.NewReader("2,1,3"))
	solve(strings.NewReader("1,2,3"))
	solve(strings.NewReader("2,3,1"))
	solve(strings.NewReader("3,2,1"))
	solve(strings.NewReader("3,1,2"))
	solve(input())
}
