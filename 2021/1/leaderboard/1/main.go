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
	input, err := os.Open(path.Join("2021", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var last int
	first := true

	inc := 0
	for scanner.Scan() {
		row := scanner.Text()

		v, err := strconv.ParseInt(row, 10, 32)
		vv := int(v)
		if err != nil {
			panic(err)
		}

		if !first {
			if vv > last {
				inc++
			}
		}

		first = false
		last = vv

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(inc)

}

func main() {
	solve(input())
}
