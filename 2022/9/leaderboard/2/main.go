package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
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

type Segment struct {
	Head, Tail *Pos
}

type Grid struct {
	Segments []*Segment

	visited *lib.Set[Pos]
}

func (g *Grid) Step(xStep, yStep int) {
	for _, segment := range g.Segments {
		xStep, yStep = segment.Step(xStep, yStep)
	}

	last := g.Segments[len(g.Segments)-1]
	last.Tail.Move(xStep, yStep)

	g.visited.Add(*last.Tail)
}

func touching(a, b *Pos) bool {
	wasSameRowOrCol := a.x == b.x || a.y == b.y
	dist := a.Dist(*b)

	if wasSameRowOrCol {
		return dist <= 1
	}

	return dist <= 2
}

func sgn(a int) int {
	if a == 0 {
		return 0
	}
	return a / lib.Abs(a)
}
func (s *Segment) Step(xStep, yStep int) (int, int) {
	s.Head.Move(xStep, yStep)

	// already touching, no movement needed
	if touching(s.Head, s.Tail) {
		return 0, 0
	}

	// actually just this, part 1 isn't updated
	return sgn(s.Head.x - s.Tail.x), sgn(s.Head.y - s.Tail.y)
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

func solve(lines []Motion) int {
	var knots []*Pos
	for i := 0; i < 10; i++ {
		knots = append(knots, &Pos{})
	}

	var segments []*Segment
	for i := 0; i < 9; i++ {
		segments = append(segments, &Segment{
			Head: knots[i],
			Tail: knots[i+1],
		})
	}

	g := Grid{
		Segments: segments,
		visited:  &lib.Set[Pos]{},
	}

	//g.visited.Add(*g.Tail)

	for _, line := range lines {
		//fmt.Printf("Before: %+v\n", line)

		xStep, yStep := 0, 0
		switch line.dir {
		case "R":
			xStep = 1
		case "D":
			yStep = 1
		case "L":
			xStep = -1
		case "U":
			yStep = -1
		default:
			panic(line.dir)
		}

		for i := 0; i < line.amount; i++ {

			g.Step(xStep, yStep)
		}

		//fmt.Printf("^After: %+v\n", line)
	}

	if g.visited.Size() < 100 {
		vis(*g.visited)
	}

	return g.visited.Size()
}

func main() {
	solver := lib.Solver[[]Motion, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2", 1)
	solver.Expect("R 5\nU 8\nL 8\nD 3\nR 17\nD 10\nL 25\nU 20", 36)
	solver.Verify(2482)
}
