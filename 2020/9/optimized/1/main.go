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
	input, err := os.Open(path.Join("2020", "9", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) []int {
	scanner := bufio.NewScanner(r)

	var numbers []int
	for scanner.Scan() {
		row := scanner.Text()

		v, _ := strconv.Atoi(row)
		numbers = append(numbers, v)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return numbers
}

func solve(numbers []int, preambleSize int) int {
	for i := preambleSize; i < len(numbers); i++ {
		if !hasUniqueComposite(numbers[i], numbers[i-preambleSize:i]) {
			return numbers[i]
		}
	}

	panic("no solution")
}

func hasUniqueComposite(target int, nums []int) bool {
	for _, i := range nums {
		for _, j := range nums {
			if i == j {
				continue
			}

			if i+j == target {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576")), 5))
	fmt.Println(solve(parse(input()), 25))
}
