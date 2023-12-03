package main

import (
	"bytes"
	"github.com/danvolchek/AdventOfCode/lib"
	"strconv"
	"strings"
)

type Grid struct {
	rows, cols int
	cells      map[int]map[int]Cell

	symbols []Cell
}

type Cell struct {
	number int
	symbol string

	row, col int
}

func (c Cell) num() int {
	return c.number
}

func parse(input []byte) Grid {
	cells := make(map[int]map[int]Cell)
	var symbols []Cell

	set := func(row, col int, cell Cell) {
		rowCells, ok := cells[row]
		if !ok {
			rowCells = make(map[int]Cell)
			cells[row] = rowCells
		}

		rowCells[col] = cell
	}

	rows := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	for row, line := range rows {
		for col := 0; col < len(line); col++ {
			char := line[col]

			if char == '.' {
				continue
			}

			if _, ok := lib.AsDigit(char); !ok {
				symbol := Cell{
					symbol: string(char),
					row:    row,
					col:    col,
				}
				set(row, col, symbol)
				symbols = append(symbols, symbol)
				continue
			}

			digitIndex := col
			var digits []int
			for {
				if digitIndex == len(line) {
					break
				}
				digit, ok := lib.AsDigit(line[digitIndex])
				if !ok {
					break
				}

				digits = append(digits, digit)
				digitIndex++
			}

			cell := Cell{
				number: lib.Atoi(strings.Join(lib.Map(digits, strconv.Itoa), "")),
				row:    row,
				col:    col,
			}

			for i := col; i < digitIndex; i++ {
				set(row, i, cell)
			}

			col += len(digits) - 1
		}
	}

	return Grid{
		rows:    len(rows),
		cols:    len(rows[0]),
		cells:   cells,
		symbols: symbols,
	}
}

func solve(grid Grid) int {
	var partNumbers lib.Set[Cell]

	for _, symbol := range grid.symbols {
		for _, adj := range lib.AdjacentPosBounds(true, symbol.row, symbol.col, grid.rows, grid.cols) {
			adjCell := grid.cells[adj.Row][adj.Col]

			if adjCell.number != 0 {
				partNumbers.Add(adjCell)
			}
		}
	}

	partNums := lib.Map(partNumbers.Items(), Cell.num)

	return lib.SumSlice(partNums)
}

func main() {
	solver := lib.Solver[Grid, int]{
		ParseF: lib.ParseBytesFunc(parse),
		SolveF: solve,
	}

	solver.Expect("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..", 4361)
	solver.Incorrect(326862) // duplicate numbers were being combined in the partNumbers map, fixed by adding row and col to Cell struct
	solver.Verify(517021)
}
