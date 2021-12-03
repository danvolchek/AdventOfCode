package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

const width = 12

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var sum []int = make([]int, width)
	for scanner.Scan() {
		row := scanner.Text()

		//fmt.Println(row)

		for i, v := range row {
			if v == '0' {
				sum[i] -= 1
			} else {
				sum[i] += 1
			}
		}
	}

	gamma := 0
	eps := 0

	fmt.Printf("%+v\n", sum)

	for i, v := range sum {
		if v > 0 {
			gamma += (1 << (width - i - 1))
		} else {
			eps += (1 << (width - i - 1))
		}
	}

	fmt.Println(gamma, eps, gamma*eps)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	//solve(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"))
	solve(input())
}
