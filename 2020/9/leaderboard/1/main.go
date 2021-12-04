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
	input, err := os.Open(path.Join("2020", "9", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func find(target int, nums []int) bool {
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
func solve(r io.Reader) {
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

	preamble := 25

	for i := preamble; i < len(numbers); i++ {
		if !find(numbers[i], numbers[i-preamble:i]) {
			fmt.Println(numbers[i])
			return
		}
	}

	panic("nothing")
}

func main() {
	//solve(strings.NewReader("35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"))
	solve(input())
}
