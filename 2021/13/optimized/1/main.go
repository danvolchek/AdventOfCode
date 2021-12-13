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
	input, err := os.Open(path.Join("2021", "13", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type fold struct {
	value int
	isX   bool
}

type pos struct {
	x, y int
}

func parse(r io.Reader) (map[pos]bool, []fold) {
	dots := make(map[pos]bool)
	var folds []fold

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, ",") != -1 {
			parts := strings.Split(line, ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			dots[pos{x: x, y: y}] = true
		} else if strings.Index(line, "=") != -1 {
			parts := strings.Split(line, "=")

			value, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			folds = append(folds, fold{
				value: value,
				isX:   parts[0][len(parts[0])-1] == 'x',
			})
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return dots, folds
}

func solve(r io.Reader) {
	dots, folds := parse(r)

	performFold(dots, folds[0])

	fmt.Println(len(dots))
}

func getFunc(x bool) func(pos) int {
	return func(p pos) int {
		if x {
			return p.x
		}

		return p.y
	}
}

func setFunc(x bool) func(*pos, int) {
	return func(p *pos, value int) {
		if x {
			p.x = value
		} else {
			p.y = value
		}
	}
}

func performFold(dots map[pos]bool, f fold) {
	get, set := getFunc(f.isX), setFunc(f.isX)

	for dot := range dots {
		changingValue := get(dot)

		// the fold is either left or up, and the dot is already to the left or above the fold line, so it would not move
		if changingValue < f.value {
			continue
		}

		// remove old dot
		delete(dots, dot)

		// update dot position
		set(&dot, 2*f.value-changingValue)

		// add new dot
		dots[dot] = true
	}
}

func main() {
	solve(strings.NewReader("6,10\n0,14\n9,10\n0,3\n10,4\n4,11\n6,0\n6,12\n4,1\n0,13\n10,12\n3,4\n3,0\n8,4\n1,10\n2,14\n8,10\n9,0\n\nfold along y=7\nfold along x=5"))
	solve(input())
}
