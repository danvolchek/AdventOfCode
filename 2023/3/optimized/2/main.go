package main

import (
	"bytes"
	"github.com/danvolchek/AdventOfCode/lib"
)

type Grid struct {
	lib.MapGrid[Cell]

	symbols []Cell
}

type Cell struct {
	number int
	symbol string

	pos lib.Pos
}

func (c Cell) isNum() bool {
	return c.number != 0
}

func (c Cell) num() int {
	return c.number
}

func (c Cell) isGear() bool {
	return c.symbol == "*"
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

			if !lib.IsDigit(char) {
				symbol := Cell{
					symbol: string(char),
					pos: lib.Pos{
						Row: row,
						Col: col,
					},
				}
				set(row, col, symbol)
				symbols = append(symbols, symbol)
				continue
			}

			digitIndex := col
			var digits string
			for {
				if digitIndex == len(line) || !lib.IsDigit(line[digitIndex]) {
					break
				}

				digits += string(line[digitIndex])
				digitIndex++
			}

			cell := Cell{
				number: lib.Atoi(digits),
				pos: lib.Pos{
					Row: row,
					Col: col,
				},
			}

			for i := col; i < digitIndex; i++ {
				set(row, i, cell)
			}

			col += len(digits) - 1
		}
	}

	return Grid{
		MapGrid: lib.MapGrid[Cell]{
			Rows: len(rows),
			Cols: len(rows[0]),
			Grid: cells,
		},
		symbols: symbols,
	}
}

func getPartNumbers(grid Grid) lib.Set[Cell] {
	var partNumbers lib.Set[Cell]

	for _, symbol := range grid.symbols {
		adjacentCells := lib.Adjacent[Cell](symbol.pos, grid, true)

		adjacentPartNumbers := lib.Filter(adjacentCells, Cell.isNum)

		partNumbers.Add(adjacentPartNumbers...)
	}

	return partNumbers
}

func solve(grid Grid) int {
	var gearRatios []int

	partNumbers := getPartNumbers(grid)

	for _, gear := range lib.Filter(grid.symbols, Cell.isGear) {
		adjacentCells := lib.Adjacent[Cell](gear.pos, grid, true)

		adjacentPartNumbers := lib.Unique(lib.Filter(adjacentCells, partNumbers.Contains))

		if len(adjacentPartNumbers) == 2 {
			gearRatios = append(gearRatios, lib.MulSlice(lib.Map(adjacentPartNumbers, Cell.num)))
		}
	}

	return lib.SumSlice(gearRatios)
}

func main() {
	solver := lib.Solver[Grid, int]{
		ParseF: lib.ParseBytesFunc(parse),
		SolveF: solve,
	}

	solver.Expect("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..", 467835)
	solver.Verify(81296995)
}
