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

type pos struct {
	num    string
	marked bool
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var nums []string
	var boards [][][]*pos

	var currentBoard [][]*pos
	for scanner.Scan() {
		row := scanner.Text()

		if nums == nil {
			nums = strings.Split(row, ",")
			continue
		}

		if len(row) == 0 {
			if currentBoard != nil {
				boards = append(boards, currentBoard)
				currentBoard = nil
			}

			continue
		}

		currentBoard = append(currentBoard, makePos(strings.Split(row, " ")))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	if currentBoard != nil {
		boards = append(boards, currentBoard)
		currentBoard = nil
	}

	for _, num := range nums {
		mark(num, boards)
		winner := getWinner(boards)
		if winner != nil {
			numv, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			fmt.Println(sumUnmarked(winner), numv, sumUnmarked(winner)*numv)
			return
		}
	}
}

func mark(num string, boards [][][]*pos) {
	for _, board := range boards {
		for _, row := range board {
			for _, pos := range row {
				if pos.num == num {
					pos.marked = true
				}
			}
		}
	}
}

func sumUnmarked(board [][]*pos) int {
	v := 0
	for _, row := range board {
		for _, pos := range row {
			if !pos.marked {
				vv, err := strconv.Atoi(pos.num)
				if err != nil {
					panic(err)
				}
				v += vv

			}
		}
	}

	return v
}

func getWinner(boards [][][]*pos) [][]*pos {
	for _, board := range boards {
		if isWinner(board) {
			return board
		}
	}

	return nil
}

func isWinner(board [][]*pos) bool {
	for i := 0; i < len(board); i++ {
		if areNumsMarked(board[i]) {
			return true
		}

		var col []*pos
		for _, row := range board {
			col = append(col, row[i])
		}
		if areNumsMarked(col) {
			return true
		}
	}

	return false
}

func areNumsMarked(row []*pos) bool {
	for _, p := range row {
		if !p.marked {
			return false
		}
	}

	return true
}

func makePos(nums []string) []*pos {
	var result []*pos
	for _, v := range nums {
		if v == "" {
			continue
		}
		result = append(result, &pos{
			num:    v,
			marked: false,
		})
	}

	return result
}

func main() {
	solve(strings.NewReader("7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1\n\n22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n\n 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n\n14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7"))
	solve(input())
}
