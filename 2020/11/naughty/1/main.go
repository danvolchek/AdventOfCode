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
func adjacentSeats(seats [][]byte, i, j int) int {
	numOccupied := 0

	for _, xOffset := range []int{1, 0, -1} {
		for _, yOffset := range []int{1, 0, -1} {
			x := i + xOffset
			y := j + yOffset

			if xOffset == 0 && yOffset == 0 {
				continue
			}

			if x < 0 || x >= len(seats) || y < 0 || y >= len(seats[x]) {
				continue
			}

			if seats[x][y] == '#' {
				numOccupied += 1
			}
		}
	}

	return numOccupied
}
func copyy(seats [][]byte) [][]byte {
	newSeats := make([][]byte, len(seats))
	for i := 0; i < len(seats); i++ {
		newSeats[i] = make([]byte, len(seats[i]))
		copy(newSeats[i], seats[i])
	}

	return newSeats
}

func copyyy(seats [][]byte) {
	for i := 0; i < len(seats); i++ {
		copy(toUse[i], seats[i])
	}
}

var toUse [][]byte

type pos struct {
	x, y int
}

func update(seats [][]byte) bool {
	changed := false

	changes := make(map[pos]byte)

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '.' {
				continue
			}

			numOccupied := adjacentSeats(seats, i, j)
			if seats[i][j] == 'L' && numOccupied == 0 {
				changes[pos{i, j}] = '#'
				changed = true
			} else if seats[i][j] == '#' && numOccupied >= 4 {
				changes[pos{i, j}] = 'L'
				changed = true
			}
		}
	}

	for pos, v := range changes {
		seats[pos.x][pos.y] = v
	}

	return changed
}

func numOccupied(seats [][]byte) int {
	occupied := 0

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '#' {
				occupied += 1
			}
		}
	}

	return occupied
}

func printy(seats [][]byte) {
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			fmt.Print(string(seats[i][j]))
		}
		fmt.Println()
	}

	fmt.Println()
}
func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var seats [][]byte
	for scanner.Scan() {
		q := scanner.Bytes()
		j := make([]byte, len(q))
		copy(j, q)
		seats = append(seats, j)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	//toUse = copyy(seats)

	for update(seats) {
		//printy(seats)

	}

	fmt.Println(numOccupied(seats))

}

func main() {
	solve(strings.NewReader("L.LL.LL.LL\nLLLLLLL.LL\nL.L.L..L..\nLLLL.LL.LL\nL.LL.LL.LL\nL.LLLLL.LL\n..L.L.....\nLLLLLLLLLL\nL.LLLLLL.L\nL.LLLLL.LL"))
	toUse = nil
	solve(input())
}
