package main

import (
	"bufio"
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	totalPaper := 0

	paperForPresent := func(l, w, h int) int {
		return 2*l*w + 2*w*h + 2*h*l + lib.Min(l*w, l*h, h*w)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, "x")
		l, w, h := lib.Must(strconv.Atoi(parts[0])), lib.Must(strconv.Atoi(parts[1])), lib.Must(strconv.Atoi(parts[2]))

		totalPaper += paperForPresent(l, w, h)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(totalPaper)
}

func main() {
	solve(strings.NewReader("2x3x4"))
	solve(input())
}
