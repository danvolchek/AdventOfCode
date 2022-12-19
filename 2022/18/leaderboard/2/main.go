package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type Pos struct {
	x, y, z int
}

func (p Pos) Add(o Pos) Pos {
	return Pos{
		x: p.x + o.x,
		y: p.y + o.y,
		z: p.z + o.z,
	}
}

func (p Pos) Magnitude() int {
	return lib.Abs(p.x) + lib.Abs(p.y) + lib.Abs(p.z)
}

func parse(line string) Pos {
	nums := lib.Ints(line)

	return Pos{
		x: nums[0],
		y: nums[1],
		z: nums[2],
	}
}

var offsets = genOffsets()

func genOffsets() [6]Pos {
	var result [6]Pos

	j := 0

	for i := -1; i <= 1; i += 2 {
		result[j] = Pos{i, 0, 0}
		j++
		result[j] = Pos{0, i, 0}
		j++
		result[j] = Pos{0, 0, i}
		j++
	}

	return result
}

var cubeSet lib.Set[Pos]

type Air struct {
	pos Pos
}

func (a Air) Adjacent() []Air {
	var result []Air
	for _, offset := range offsets {
		airPos := a.pos.Add(offset)

		if !cubeSet.Contains(airPos) {
			result = append(result, Air{pos: airPos})
		}
	}

	return result
}

// The approach is to do a BFS search block by block from a known external position (0,0,0) and look
// for all the faces that are reached
func solve(positions []Pos) int {
	cubeSet = lib.Set[Pos]{}
	for _, cube := range positions {
		cubeSet.Add(cube)
	}

	searchRange := lib.MaxSlice(lib.Map(positions, Pos.Magnitude)) + 10

	externalFaces := 0
	lib.BFS(Air{}, func(a Air) bool {
		if a.pos.Magnitude() >= searchRange {
			return true
		}

		for _, offset := range offsets {
			cube := a.pos.Add(offset)
			if cubeSet.Contains(cube) {
				externalFaces += 1
			}
		}

		return false
	})

	return externalFaces
}

func main() {
	solver := lib.Solver[[]Pos, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("2,2,2\n1,2,2\n3,2,2\n2,1,2\n2,3,2\n2,2,1\n2,2,3\n2,2,4\n2,2,6\n1,2,5\n3,2,5\n2,1,5\n2,3,5", 58)
	solver.Verify(2118)
}
