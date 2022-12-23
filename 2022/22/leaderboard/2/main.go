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

func (p Pos) Step(dir Dir) Pos {
	switch dir {
	case up:
		p.y -= 1
	case down:
		p.y += 1
	case left:
		p.x -= 1
	case right:
		p.x += 1
	}

	return p
}

type PosDir struct {
	pos Pos
	dir Dir
}

type Input struct {
	grove        [][]Tile
	instructions []Instruction
}

func parseTile(char byte) Tile {
	switch char {
	case ' ':
		return outOfBounds
	case '.':
		return open
	case '#':
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
	wrap  map[Pos]map[Dir]PosDir
	grove [][]Tile

	pd PosDir
}

func (s *State) String() string {
	var sb strings.Builder

	for y, row := range s.grove {
		for x, tile := range row {
			if s.pd.pos.x == x && s.pd.pos.y == y {
				switch s.pd.dir {
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

func (m moveInstr) Next(state *State) PosDir {
	nextPd := PosDir{
		pos: state.pd.pos.Step(state.pd.dir),
		dir: state.pd.dir,
	}

	if wraps, ok := state.wrap[nextPd.pos]; ok {
		nextPd = wraps[state.pd.dir]
	}

	switch state.grove[nextPd.pos.y][nextPd.pos.x] {
	case wall:
		return state.pd
	case open:
		return nextPd
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
		if next == state.pd {
			break
		}

		state.pd = next
	}
}

type turnInstr struct {
	left bool
}

func (t turnInstr) Act(state *State) {
	switch state.pd.dir {
	case left:
		switch t.left {
		case true:
			state.pd.dir = down
		case false:
			state.pd.dir = up
		}
	case right:
		switch t.left {
		case true:
			state.pd.dir = up
		case false:
			state.pd.dir = down
		}
	case up:
		switch t.left {
		case true:
			state.pd.dir = left
		case false:
			state.pd.dir = right
		}
	case down:
		switch t.left {
		case true:
			state.pd.dir = right
		case false:
			state.pd.dir = left
		}
	}
}

func face(p Pos) int {
	return 0
}

func firstNonEmptyTile(start Pos, dir Dir, grove [][]Tile) PosDir {

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
			return PosDir{}
		}

		start.Step(dir)
	}
}

func buildWraps(grove [][]Tile) map[Pos]map[Dir]PosDir {
	wraps := make(map[Pos]map[Dir]PosDir)

	for row, contents := range grove {
		for col, tile := range contents {
			if tile != open {
				continue
			}

			if row == 0 || grove[row-1][col] == outOfBounds {
				p := Pos{col, row - 1}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]PosDir)
				}
				wraps[p][up] = firstNonEmptyTile(p, up, grove)
			}

			if row == len(grove)-1 || grove[row+1][col] == outOfBounds {
				p := Pos{col, row + 1}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]PosDir)
				}
				wraps[p][down] = firstNonEmptyTile(p, down, grove)
			}

			if col == 0 || grove[row][col-1] == outOfBounds {
				p := Pos{col - 1, row}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]PosDir)
				}
				wraps[p][left] = firstNonEmptyTile(p, left, grove)
			}

			if col == len(contents)-1 || grove[row][col+1] == outOfBounds {
				p := Pos{col + 1, row}
				if wraps[p] == nil {
					wraps[p] = make(map[Dir]PosDir)
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
		pd: PosDir{
			pos: Pos{
				x: startX,
				y: startY,
			},
			dir: right,
		},
	}

	for _, instr := range input.instructions {
		instr.Act(state)
	}

	return 1000*(state.pd.pos.y+1) + 4*(state.pd.pos.x+1) + int(state.pd.dir-right)
}

func main() {
	solver := lib.Solver[Input, int]{
		ParseF: lib.ParseChunksUnique(parseGrove, parseInstructions),
		SolveF: solve,
	}

	solver.Expect("        ...#\n        .#..\n        #...\n        ....\n...#.......#\n........#...\n..#....#....\n..........#.\n        ...#....\n        .....#..\n        .#......\n        ......#.\n\n10R5L5R10L4R5L5", 5031)
	solver.Solve()
}
