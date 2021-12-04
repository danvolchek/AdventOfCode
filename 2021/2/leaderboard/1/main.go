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
	input, err := os.Open(path.Join("2021", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	depth := 0
	horiz := 0

	for scanner.Scan() {
		row := scanner.Text()

		if strings.Index(row, "down") == 0 {
			v, _ := strconv.Atoi(row[len("down")+1:])
			depth += v
		}

		if strings.Index(row, "up") == 0 {
			v, _ := strconv.Atoi(row[len("up")+1:])
			depth -= v
		}

		if strings.Index(row, "forward") == 0 {
			v, _ := strconv.Atoi(row[len("forward")+1:])
			horiz += v
		}
	}

	fmt.Printf("%v %v %v\n", depth, horiz, depth*horiz)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(input())
}
