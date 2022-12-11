package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Motion struct {
	dir    string
	amount int
}

func parse(line string) Motion {
	return Motion{
		dir:    string(line[0]),
		amount: lib.Int(line),
	}
}

type Pos struct {
	x, y int
}

func (p *Pos) Move(xx, yy int) {
	p.x += xx
	p.y += yy
}

func (p *Pos) At(x, y int) bool {
	return p.x == x && p.y == y
}

func (p *Pos) Dist(o Pos) int {
	return lib.Abs(p.x-o.x) + lib.Abs(p.y-o.y)
}

type Grid struct {
	Head, Tail *Pos

	visited *lib.Set[Pos]
}

func (g *Grid) Step(dir string) {
	xStep, yStep := 0, 0
	switch dir {
	case "R":
		xStep = 1
	case "D":
		yStep = 1
	case "L":
		xStep = -1
	case "U":
		yStep = -1
	default:
		panic(dir)
	}

	g.Head.Move(xStep, yStep)

	dist := g.Head.Dist(*g.Tail)

	if dist == 2 && (g.Head.x == g.Tail.x || g.Head.y == g.Tail.y) {
		//fmt.Println("straight move")
		g.Tail.Move(xStep, yStep)
	} else if dist > 2 && (g.Head.x != g.Tail.x && g.Head.y != g.Tail.y) {
		//fmt.Println("diag move")
		g.Tail.Move(xStep, yStep)

		if xStep != 0 {
			if g.Head.y > g.Tail.y {
				g.Tail.Move(0, 1)
			} else {
				g.Tail.Move(0, -1)
			}
		} else {
			if g.Head.x > g.Tail.x {
				g.Tail.Move(1, 0)
			} else {
				g.Tail.Move(-1, 0)
			}
		}
	}

	strRep := g.String()
	if !strings.Contains(strRep, "H") && !strings.Contains(strRep, "T") && g.Head.Dist(*g.Tail) != 0 {
		panic("whaaat")
	}

	g.visited.Add(*g.Tail)
}

func (g Grid) String2() string {
	var s strings.Builder
	for i := -4; i <= 0; i++ {
		for j := 0; j <= 5; j++ {
			jj := j
			ii := i
			if jj == 0 && ii == 0 {
				s.WriteString("s")
			} else if g.Head.At(jj, ii) {
				s.WriteString("H")
			} else if g.Tail.At(jj, ii) {
				s.WriteString("T")
			} else {
				s.WriteString(".")
			}
		}
		s.WriteString("\n")
	}

	return s.String()
}

func (g Grid) String() string {
	var s strings.Builder
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			jj := g.Head.x + j
			ii := g.Head.y + i
			if g.Head.At(jj, ii) {
				s.WriteString("H")
			} else if g.Tail.At(jj, ii) {
				s.WriteString("T")
			} else {
				s.WriteString(".")
			}
		}
		s.WriteString("\n")
	}

	return s.String()
}

func vis(l lib.Set[Pos]) {
	var min, max Pos
	for _, item := range l.Items() {
		if item.x < min.x {
			min.x = item.x
		}

		if item.x > max.x {
			max.x = item.x
		}

		if item.y < min.y {
			min.y = item.y
		}

		if item.y > max.y {
			max.y = item.y
		}
	}

	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else if l.Contains(Pos{x: x, y: y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func stepTest(lines []Motion) string {
	g := Grid{
		Head:    &Pos{},
		Tail:    &Pos{},
		visited: &lib.Set[Pos]{},
	}

	for _, line := range lines {
		//fmt.Printf("Before: %+v\n", line)
		for i := 0; i < line.amount; i++ {
			g.Step(line.dir)
			//fmt.Println(g.String2())
		}

		//fmt.Printf("^After: %+v\n", line)
	}

	return strings.TrimSpace(g.String())
}

func solve(lines []Motion) int {
	g := Grid{
		Head:    &Pos{},
		Tail:    &Pos{},
		visited: &lib.Set[Pos]{},
	}

	//g.visited.Add(*g.Tail)

	fmt.Println(g.String2())

	for _, line := range lines {
		fmt.Printf("Before: %+v\n", line)
		for i := 0; i < line.amount; i++ {
			g.Step(line.dir)
			fmt.Println(g.String2())
		}

		fmt.Printf("^After: %+v\n", line)
	}

	if g.visited.Size() < 100 {
		vis(*g.visited)
	}

	return g.visited.Size()
}

func main() {
	stepSolver := lib.Solver[[]Motion, string]{
		ParseF: lib.ParseLine(parse),
		SolveF: stepTest,
	}
	stepSolver.Expect("U 1\nR 1\nU 1", "...\n.H.\n.T.")
	stepSolver.Expect("U 1\nR 1\nR 1", "...\nTH.\n...")

	solver := lib.Solver[[]Motion, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2", 13)
	solver.Verify(6311)
}
