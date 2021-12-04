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
	input, err := os.Open(path.Join("2021", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type board struct {
	grid [][]*number

	winningNumber string
}

func (b *board) Call(number string) bool {
	if b.isWinner() {
		panic("can't call number on an already won board")
	}
	b.mark(number)

	won := b.isWinner()
	if won {
		b.winningNumber = number
	}

	return won
}

func (b *board) Score() int {
	if b.winningNumber == "" {
		panic("can't get score of a board that hasn't won")
	}

	winningNumberInt, err := strconv.Atoi(b.winningNumber)
	if err != nil {
		panic(err)
	}

	return winningNumberInt * b.sumUnmarked()

}

func (b *board) sumUnmarked() int {
	score := 0
	for _, row := range b.grid {
		for _, number := range row {
			if !number.marked {
				intValue, err := strconv.Atoi(number.value)
				if err != nil {
					panic(err)
				}
				score += intValue

			}
		}
	}

	return score
}

func (b *board) mark(num string) {
	for _, row := range b.grid {
		for _, pos := range row {
			if pos.value == num {
				pos.marked = true
			}
		}
	}
}

func (b board) isWinner() bool {
	for i := 0; i < len(b.grid); i++ {
		if areAllNumbersMarked(b.grid[i]) {
			return true
		}

		var col []*number
		for _, row := range b.grid {
			col = append(col, row[i])
		}

		if areAllNumbersMarked(col) {
			return true
		}
	}

	return false
}

func areAllNumbersMarked(numbers []*number) bool {
	for _, p := range numbers {
		if !p.marked {
			return false
		}
	}

	return true
}

type number struct {
	value  string
	marked bool
}

func parse(r io.Reader) ([]string, []*board) {
	scanner := bufio.NewScanner(r)

	var numbers []string
	var boards []*board

	var currentBoard *board
	for scanner.Scan() {
		row := scanner.Text()

		if numbers == nil {
			numbers = strings.Split(row, ",")
			continue
		}

		if len(row) == 0 {
			if currentBoard != nil {
				boards = append(boards, currentBoard)
				currentBoard = nil
			}

			continue
		}

		if currentBoard == nil {
			currentBoard = &board{}
		}

		currentBoard.grid = append(currentBoard.grid, makeUnmarkedNumbers(strings.Split(row, " ")))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	if currentBoard != nil {
		boards = append(boards, currentBoard)
	}

	return numbers, boards
}

func makeUnmarkedNumbers(numbers []string) []*number {
	var result []*number
	for _, value := range numbers {
		if value == "" {
			continue
		}

		result = append(result, &number{
			value:  value,
			marked: false,
		})
	}

	return result
}

func solve(r io.Reader) {
	numbers, boards := parse(r)

	var lastWinner *board

	for _, number := range numbers {
		for i := 0; i < len(boards); i++ {

			b := boards[i]
			won := b.Call(number)

			if won {
				lastWinner = b

				boards[i] = boards[len(boards)-1]
				boards = boards[:len(boards)-1]
				i--
			}
		}
	}

	fmt.Println(lastWinner.Score())
}

func main() {
	solve(strings.NewReader("7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1\n\n22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n\n 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n\n14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7"))
	solve(input())
}
