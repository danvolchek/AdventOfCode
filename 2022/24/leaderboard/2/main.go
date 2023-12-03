package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Tile int

const (
	Wall Tile = 1 << iota
	Empty
	BlizzardLeft
	BlizzardRight
	BlizzardDown
	BlizzardUp
)

func parse(char byte) Tile {
	switch char {
	case '#':
		return Wall
	case '.':
		return Empty
	case '<':
		return BlizzardLeft
	case '>':
		return BlizzardRight
	case '^':
		return BlizzardUp
	case 'v':
		return BlizzardDown
	default:
		panic(string(char))
	}
}

var totalWorlds int
var worldMap map[int]*Valley

type Valley struct {
	minute int
	lib.SliceGrid[Tile]
}

func (w *Valley) String() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Minute: %v\n", w.minute))

	for _, row := range w.Grid {
		for _, tile := range row {
			if tile&Wall != 0 {
				s.WriteByte('#')
				continue
			}

			numBlizz := 0

			if tile&BlizzardLeft != 0 {
				numBlizz += 1
			}
			if tile&BlizzardRight != 0 {
				numBlizz += 1
			}
			if tile&BlizzardUp != 0 {
				numBlizz += 1
			}
			if tile&BlizzardDown != 0 {
				numBlizz += 1
			}

			if numBlizz == 0 {
				s.WriteByte('.')
				continue
			}

			if numBlizz != 1 {
				s.WriteByte(byte('0' + numBlizz))
				continue
			}

			if tile&BlizzardLeft != 0 {
				s.WriteByte('<')
			}
			if tile&BlizzardRight != 0 {
				s.WriteByte('>')
			}
			if tile&BlizzardUp != 0 {
				s.WriteByte('^')
			}
			if tile&BlizzardDown != 0 {
				s.WriteByte('v')
			}

		}
		s.WriteString("\n")
	}

	return s.String()
}

type Node struct {
	p lib.Pos

	w *Valley
}

func (n Node) Adjacent() []Node {
	var nodes []Node

	nextWorldMin := (n.w.minute + 1) % totalWorlds
	nextWorld := worldMap[nextWorldMin]
	for _, p := range lib.AdjacentPos[Tile](false, n.p.Row, n.p.Col, nextWorld) {
		if nextWorld.Grid[p.Row][p.Col]&Empty != 0 {
			node := Node{
				p: p,
				w: nextWorld,
			}

			nodes = append(nodes, node)
		}
	}

	if nextWorld.Grid[n.p.Row][n.p.Col]&Empty != 0 {
		node := Node{
			p: n.p,
			w: nextWorld,
		}

		nodes = append(nodes, node)
	}

	return nodes
}

func blow(valley [][]Tile) [][]Tile {
	valleyMap := make(map[lib.Pos]Tile)

	getPosInOffset := func(start, offset lib.Pos) lib.Pos {
		result := start.Add(offset)

		if result.Row == 0 {
			result.Row = len(valley) - 2
		} else if result.Row == len(valley)-1 {
			result.Row = 1
		}

		if result.Col == 0 {
			result.Col = len(valley[result.Row]) - 2
		} else if result.Col == len(valley[result.Row])-1 {
			result.Col = 1
		}

		return result
	}

	for rowNum, row := range valley {
		for colNum, tile := range row {
			pos := lib.Pos{Row: rowNum, Col: colNum}
			if tile&Wall != 0 {
				valleyMap[pos] = Wall
				continue
			}

			if tile&BlizzardLeft != 0 {
				valleyMap[getPosInOffset(pos, lib.Pos{Row: 0, Col: -1})] |= BlizzardLeft
			}
			if tile&BlizzardRight != 0 {
				valleyMap[getPosInOffset(pos, lib.Pos{Row: 0, Col: 1})] |= BlizzardRight
			}
			if tile&BlizzardUp != 0 {
				valleyMap[getPosInOffset(pos, lib.Pos{Row: -1, Col: 0})] |= BlizzardUp
			}
			if tile&BlizzardDown != 0 {
				valleyMap[getPosInOffset(pos, lib.Pos{Row: 1, Col: 0})] |= BlizzardDown
			}
		}
	}

	newValley := make([][]Tile, len(valley))
	for rowNum, row := range valley {
		newValley[rowNum] = make([]Tile, len(valley[rowNum]))
		for colNum := range row {
			pos := lib.Pos{Row: rowNum, Col: colNum}
			if val, ok := valleyMap[pos]; ok {
				newValley[rowNum][colNum] = val
			} else {
				newValley[rowNum][colNum] = Empty
			}
		}
	}

	return newValley
}

func solve(valley [][]Tile) int {
	var start, end lib.Pos
	for col := 0; col < len(valley[0]); col++ {
		if valley[0][col]&Empty != 0 {
			start = lib.Pos{Row: 0, Col: col}
		}

		if valley[len(valley)-1][col]&Empty != 0 {
			end = lib.Pos{Row: len(valley) - 1, Col: col}
		}
	}

	totalWorlds = len(valley) * len(valley[0])
	worldMap = make(map[int]*Valley)

	worldMap[0] = &Valley{
		minute: 0,
		SliceGrid: lib.SliceGrid[Tile]{
			Grid: valley,
		},
	}

	for i := 1; i < totalWorlds; i++ {
		newValley := blow(valley)

		worldMap[i] = &Valley{
			minute: i,
			SliceGrid: lib.SliceGrid[Tile]{
				Grid: newValley,
			},
		}

		valley = newValley
	}

	startToGoal, ok := lib.BFS(Node{
		p: start,
		w: worldMap[0],
	}, func(n Node) bool {
		return n.p == end
	})
	if !ok {
		panic("path not found")
	}

	goalToStart, ok := lib.BFS(startToGoal[len(startToGoal)-1],
		func(n Node) bool {
			return n.p == start
		})
	if !ok {
		panic("path not found")
	}

	startToGoalAgain, ok := lib.BFS(goalToStart[len(goalToStart)-1],
		func(n Node) bool {
			return n.p == end
		})
	if !ok {
		panic("path not found")
	}

	return len(startToGoal) - 1 + len(goalToStart) - 1 + len(startToGoalAgain) - 1
}

func main() {
	solver := lib.Solver[[][]Tile, int]{
		ParseF: lib.TrimSpace(lib.ParseGrid(parse)),
		SolveF: solve,
	}

	solver.Expect("#.######\n#>>.<^<#\n#.<..<<#\n#>v.><>#\n#<^v^^>#\n######.#", 54)
	solver.Verify(717)
}
