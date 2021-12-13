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
	paper := make(map[pos]bool)
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

			paper[pos{x: x, y: y}] = true
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

	return paper, folds
}

func solve(r io.Reader) {
	paper, folds := parse(r)

	for _, f := range folds {
		performFold(paper, f)
	}

	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if paper[pos{x: x, y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func performFold(paper map[pos]bool, f fold) {
	if f.isX {
		for p := range paper {
			if p.x < f.value {
				continue
			}

			delete(paper, p)
			paper[pos{
				x: f.value - (p.x - f.value),
				y: p.y,
			}] = true
		}
	} else {
		for p := range paper {
			if p.y < f.value {
				continue
			}

			delete(paper, p)
			paper[pos{
				x: p.x,
				y: f.value - (p.y - f.value),
			}] = true
		}
	}
}

func main() {
	solve(strings.NewReader("6,10\n0,14\n9,10\n0,3\n10,4\n4,11\n6,0\n6,12\n4,1\n0,13\n10,12\n3,4\n3,0\n8,4\n1,10\n2,14\n8,10\n9,0\n\nfold along y=7\nfold along x=5"))
	solve(input())
}
