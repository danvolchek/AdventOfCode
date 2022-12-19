package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type Pos struct {
	x, y, z int
}

func parse(line string) Pos {
	nums := lib.Ints(line)

	return Pos{
		x: nums[0],
		y: nums[1],
		z: nums[2],
	}
}

func notConnected(cube Pos, cubes lib.Set[Pos]) int {
	sum := 0

	for i := -1; i <= 1; i += 2 {
		if !cubes.Contains(Pos{x: cube.x + i, y: cube.y, z: cube.z}) {
			sum += 1
		}

		if !cubes.Contains(Pos{x: cube.x, y: cube.y + i, z: cube.z}) {
			sum += 1
		}

		if !cubes.Contains(Pos{x: cube.x, y: cube.y, z: cube.z + i}) {
			sum += 1
		}

	}

	return sum
}

func solve(cubes []Pos) int {
	var cubeMap lib.Set[Pos]
	for _, cube := range cubes {
		cubeMap.Add(cube)
	}

	return lib.SumSlice(lib.Map(cubes, func(p Pos) int {
		return notConnected(p, cubeMap)
	}))
}

func main() {
	solver := lib.Solver[[]Pos, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("2,2,2\n1,2,2\n3,2,2\n2,1,2\n2,3,2\n2,2,1\n2,2,3\n2,2,4\n2,2,6\n1,2,5\n3,2,5\n2,1,5\n2,3,5", 64)
	solver.Verify(3650)
}
