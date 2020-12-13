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
	input, err := os.Open(path.Join("2020", "11", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) [][]byte {
	var seats [][]byte

	clone := func(source []byte) []byte {
		cloned := make([]byte, len(source))
		copy(cloned, source)
		return cloned
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		seats = append(seats, clone(scanner.Bytes()))
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return seats
}

const (
	occupied = '#'
	empty    = 'L'
	floor    = '.'
)

func occupiedAdjacentSeats(seats [][]byte, i, j int) int {
	numOccupied := 0

	isInvalid := func(x, y int) bool {
		return x < 0 || x >= len(seats) || y < 0 || y >= len(seats[x])
	}

	for _, xOffset := range []int{1, 0, -1} {
		for _, yOffset := range []int{1, 0, -1} {
			if xOffset == 0 && yOffset == 0 {
				continue
			}

			x := i + xOffset
			y := j + yOffset

			if isInvalid(x, y) {
				continue
			}

			for !isInvalid(x, y) && seats[x][y] == floor {
				x += xOffset
				y += yOffset
			}

			if isInvalid(x, y) {
				continue
			}

			if seats[x][y] == occupied {
				numOccupied += 1
			}
		}
	}

	return numOccupied
}

type pos struct {
	x, y int
}

func update(seats [][]byte) bool {
	changes := make(map[pos]byte)

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			seat := seats[i][j]

			if seat == floor {
				continue
			}

			numOccupied := occupiedAdjacentSeats(seats, i, j)
			if seat == empty && numOccupied == 0 {
				changes[pos{i, j}] = occupied
			} else if seat == occupied && numOccupied >= 5 {
				changes[pos{i, j}] = empty
			}
		}
	}

	for pos, v := range changes {
		seats[pos.x][pos.y] = v
	}

	return len(changes) != 0
}

func numOccupied(seats [][]byte) int {
	occupiedSeats := 0

	for _, row := range seats {
		for _, seat := range row {
			if seat == occupied {
				occupiedSeats += 1
			}
		}
	}

	return occupiedSeats
}

func solve(seats [][]byte) int {
	for update(seats) {
	}

	return numOccupied(seats)
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"))))
	fmt.Println(solve(parse(input())))
}
