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
	input, err := os.Open(path.Join("2021", "5", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type point struct {
	x, y int
}

func parsePoint(rawPoint string) point {
	parts := strings.Split(rawPoint, ",")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return point{
		x: x,
		y: y,
	}
}

func parseLine(rawLine string) (point, point) {
	parts := strings.Split(rawLine, " -> ")

	return parsePoint(parts[0]), parsePoint(parts[1])
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	field := make(map[point]int)

	for scanner.Scan() {
		line := scanner.Text()

		start, end := parseLine(line)

		walk(start, end, func(p point) {
			field[p] += 1
		})
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	overlap := 0
	for _, numberOfLines := range field {
		if numberOfLines > 1 {
			overlap += 1
		}
	}

	fmt.Println(overlap)
}

func walk(start, end point, action func(p point)) {
	xIncrement := sign(end.x - start.x)
	yIncrement := sign(end.y - start.y)

	current := start
	for current != end {
		action(current)

		current.x += xIncrement
		current.y += yIncrement
	}

	action(end)
}

func sign(value int) int {
	if value < 0 {
		return -1
	} else if value == 0 {
		return 0
	} else {
		return 1
	}
}

func main() {
	solve(strings.NewReader("0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2\n"))
	solve(input())
}
