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

	rows := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})

	set := func(row, col int, val Cell) {
		c, ok := cells[row]
		if !ok {
			c = make(map[int]Cell)
			cells[row] = c
		}

		c[col] = val
	}

	for y, row := range rows {
		for i := 0; i < len(row); i++ {
			val := row[i]

			if val == '.' {
				continue
			}

			if _, ok := lib.AsDigit(val); !ok {
				set(y, i, Cell{
					symbol: string(val),
					row:    y,
					col:    i,
				})
				continue
			}

			j := i
			var digits []int
			for {
				if j == len(row) {
					break
				}
				digit, ok := lib.AsDigit(row[j])
				if !ok {
					break
				}

				digits = append(digits, digit)
				j++
			}

			cellV := Cell{
				number: lib.Atoi(strings.Join(lib.Map(digits, strconv.Itoa), "")),
				row:    y,
				col:    i,
			}

			for k := i; k < j; k++ {
				set(y, k, cellV)
			}

			i += len(digits) - 1
		}
	}

	return Grid{
		rows:  len(rows),
		cols:  len(rows[0]),
		cells: cells,
	}
}

func getPartNumbers(grid Grid) map[Cell]bool {
	partNumbers := make(map[Cell]bool)

	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			cell := grid.cells[row][col]
			if cell.symbol == "" {
				continue
			}

			for _, adj := range adjacentPosBounds(true, row, col, grid.rows, grid.cols) {
				adjCell := grid.cells[adj.Row][adj.Col]

				if adjCell.number != 0 {
					partNumbers[adjCell] = true
				}
			}
		}
	}

	return partNumbers
}

func solve(grid Grid) int {
	var gearRatios []int

	partNumbers := getPartNumbers(grid)

	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			cell := grid.cells[row][col]
			if cell.symbol != "*" {
				continue
			}

			adjacentPartNumbers := make(map[Cell]bool)

			for _, adj := range adjacentPosBounds(true, row, col, grid.rows, grid.cols) {
				adjCell := grid.cells[adj.Row][adj.Col]

				if _, ok := partNumbers[adjCell]; ok {
					adjacentPartNumbers[adjCell] = true
				}
			}

			if len(adjacentPartNumbers) == 2 {
				ratio := 1
				for partNumberCell := range adjacentPartNumbers {
					ratio *= partNumberCell.number
				}
				gearRatios = append(gearRatios, ratio)
			}
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

// copied over from lib, as this was made private later
func adjacentPosBounds(diag bool, row, col, rows, cols int) []lib.Pos {
	return lib.Filter(lib.AdjacentPosNoBoundsChecks(lib.Pos{Row: row, Col: col}, diag), func(pos lib.Pos) bool {
		return !(pos.Row < 0 || pos.Col < 0 || pos.Row >= rows || pos.Col >= cols)
	})
}
