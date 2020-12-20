package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "17", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type coord struct {
	x, y, z int
}

func neighbors(val coord) []coord {
	var coords []coord

	for _, xOffset := range []int{-1, 0, 1} {
		for _, yOffset := range []int{-1, 0, 1} {
			for _, zOffset := range []int{-1, 0, 1} {
				if xOffset == 0 && yOffset == 0 && zOffset == 0 {
					continue
				}

				pos := coord{val.x + xOffset, val.y + yOffset, val.z + zOffset}

				coords = append(coords, pos)

			}
		}
	}

	return coords
}

func numActive(val coord, coords map[coord]bool) int {
	active := 0
	for _, neighbor := range neighbors(val) {
		if coords[neighbor] {
			active += 1
		}
	}

	return active
}

func process(coords map[coord]bool) {
	processed := make(map[coord]bool)
	changes := make(map[coord]bool)

	for curr, _ := range coords {
		for _, neighbor2 := range neighbors(curr) {
			for _, neighbor := range neighbors(neighbor2) {
				if processed[neighbor] {
					continue
				}

				processed[neighbor] = true

				num := numActive(neighbor, coords)
				switch coords[neighbor] {
				case true:
					changes[neighbor] = num == 2 || num == 3
				case false:
					changes[neighbor] = num == 3
				}
			}
		}

	}

	for pos, newVal := range changes {
		coords[pos] = newVal
	}
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	coords := make(map[coord]bool)
	y := 0
	for scanner.Scan() {
		row := scanner.Text()

		for x := 0; x < len(row); x++ {
			if row[x] == '#' {
				coords[coord{x, y, 0}] = true
			}
		}

		y += 1

	}

	for i := 0; i < 6; i++ {
		process(coords)
	}

	active := 0
	for _, isActive := range coords {
		if isActive {
			active += 1
		}
	}

	fmt.Println(active)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(strings.NewReader(".#.\n..#\n###"))
	solve(input())
}
