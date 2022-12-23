package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Tile int

const (
	open Tile = iota
	wall
	outOfBounds
)

type Dir int

const (
	right Dir = iota
	down
	left
	up
)

type Pos struct {
	x, y int
}

type Input struct {
	grove        [][]Tile
	instructions []Instruction
}

func parseTile(char string) Tile {
	switch char {
	case " ":
		return outOfBounds
	case ".":
		return open
	case "#":
		return wall
	default:
		panic(char)
	}
}

func parseGrove(chunk string, input *Input) {
	// Fill in lines that don't go all the way with spaces
	lines := strings.Split(chunk, "\n")
	var max int
	for _, line := range lines {
		max = lib.Max(len(line), max)
	}
	for i, line := range lines {
		lines[i] = line + strings.Repeat(" ", max-len(line))
	}
	chunk = strings.Join(lines, "\n")

	input.grove = lib.ParseGrid(parseTile)(chunk)
}

func parseInstructions(line string, input *Input) {
	line = strings.TrimSpace(line)

	var instructions []Instruction
	for i := 0; i < len(line); {

		if line[i] == 'L' || line[i] == 'R' {
			instructions = append(instructions, turnInstr{left: line[i] == 'L'})
			i += 1
			continue
		}

		j := i + 1
		for j < len(line) && line[j] >= '0' && line[j] <= '9' {
			j++
		}

		instructions = append(instructions, moveInstr{amount: lib.Atoi(line[i:j])})
		i = j
	}

	input.instructions = instructions
}

type State struct {
	wrap  map[Pos]map[Dir]Pos
	grove [][]Tile

	dir Dir
	pos Pos
}

func (s *State) String() string {
	var sb strings.Builder

	for y, row := range s.grove {
		for x, tile := range row {
			if s.pos.x == x && s.pos.y == y {
				switch s.dir {
				case left:
					sb.WriteByte('<')
				case right:
					sb.WriteByte('>')
				case down:
					sb.WriteByte('v')
				case up:
					sb.WriteByte('^')
				}
			} else {
				switch tile {
				case outOfBounds:
					sb.WriteByte(' ')
				case wall:
					sb.WriteByte('#')
				case open:
					sb.WriteByte('.')
				}
			}
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

type Instruction interface {
	Act(state *State)
}

type moveInstr struct {
	amount int
}

func (m moveInstr) Next(state *State) Pos {
	nextPos := state.pos
	switch state.dir {
	case up:
		nextPos.y -= 1
	case down:
		nextPos.y += 1
	case left:
		nextPos.x -= 1
	case right:
		nextPos.x += 1
	}

	if wraps, ok := state.wrap[nextPos]; ok {
		if nextPos.x >= 0 && nextPos.y >= 0 && nextPos.y < len(state.grove) && nextPos.x < len(state.grove[nextPos.y]) && state.grove[nextPos.y][nextPos.x] != outOfBounds {
			panic("sir")
		}

		nextPos = wraps[state.dir]
	}

	switch state.grove[nextPos.y][nextPos.x] {
	case wall:
		return state.pos
	case open:
		return nextPos
	case outOfBounds:
		fallthrough
	default:
		panic("can't walk into oob my dude")
	}
}

func (m moveInstr) Act(state *State) {
	for i := 0; i < m.amount; i++ {
		next := m.Next(state)

		//fmt.Println(state)
		if next == state.pos {
			break
		}

		state.pos = next
	}
}

type turnInstr struct {
	left bool
}

func (t turnInstr) Act(state *State) {
	switch state.dir {
	case left:
		switch t.left {
		case true:
			state.dir = down
		case false:
			state.dir = up
		}
	case right:
		switch t.left {
		case true:
			state.dir = up
		case false:
			state.dir = down
		}
	case up:
		switch t.left {
		case true:
			state.dir = left
		case false:
			state.dir = right
		}
	case down:
		switch t.left {
		case true:
			state.dir = right
		case false:
			state.dir = left
		}
	}
}

func firstNonEmptyTile(start Pos, dir Dir, grove [][]Tile) Pos {
	for {
		if start.y == -1 {
			start.y = len(grove) - 1
		} else if start.y == len(grove) {
			start.y = 0
		}

		if start.x == -1 {
			start.x = len(grove[start.y]) - 1
		} else if start.x == len(grove[start.y]) {
			start.x = 0
		}

		if grove[start.y][start.x] != outOfBounds {
			return start
		}

		switch dir {
		case up:
			start.y -= 1
		case down:
			start.y += 1
		case left:
			start.x -= 1
		case right:
			start.x += 1
		}
	}
}

func buildWraps(grove [][]Tile) map[Pos]map[Dir]Pos {
	wraps := make(map[Pos]map[Dir]Pos)

	for row, contents := range grove {
		for col, tile := range contents {
			if tile != open {
				continue
			}

			if row == 0 || grove[row-1][col] == outOfBounds {
				p := Pos{col, row - 1}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]Pos)
				}
				wraps[p][up] = firstNonEmptyTile(p, up, grove)
			}

			if row == len(grove)-1 || grove[row+1][col] == outOfBounds {
				p := Pos{col, row + 1}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]Pos)
				}
				wraps[p][down] = firstNonEmptyTile(p, down, grove)
			}

			if col == 0 || grove[row][col-1] == outOfBounds {
				p := Pos{col - 1, row}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]Pos)
				}
				wraps[p][left] = firstNonEmptyTile(p, left, grove)
			}

			if col == len(contents)-1 || grove[row][col+1] == outOfBounds {
				p := Pos{col + 1, row}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]Pos)
				}
				wraps[p][right] = firstNonEmptyTile(p, right, grove)
			}
		}
	}

	return wraps
}

func solve(input Input) int {
	startY := 0
	var startX int

	for x, tile := range input.grove[startY] {
		if tile == open {
			startX = x
			break
		}
	}

	if input.grove[startY][startX] != open {
		panic("bad start position")
	}

	state := &State{
		grove: input.grove,
		wrap:  buildWraps(input.grove),
		dir:   right,
		pos: Pos{
			x: startX,
			y: startY,
		},
	}

	for _, instr := range input.instructions {
		instr.Act(state)
	}

	return 1000*(state.pos.y+1) + 4*(state.pos.x+1) + int(state.dir-right)
}

func main() {
	solver := lib.Solver[Input, int]{
		ParseF: lib.ParseChunksUnique(parseGrove, parseInstructions),
		SolveF: solve,
	}

	solver.Expect("        ...#\n        .#..\n        #...\n        ....\n...#.......#\n........#...\n..#....#....\n..........#.\n        ...#....\n        .....#..\n        .#......\n        ......#.\n\n10R5L5R10L4R5L5", 6032)
	solver.Verify(80392)
}
