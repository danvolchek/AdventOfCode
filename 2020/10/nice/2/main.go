package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "10", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := scanner.Bytes()

		fmt.Println(row)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(input())
}
