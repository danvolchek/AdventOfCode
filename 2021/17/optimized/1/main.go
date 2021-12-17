package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "17", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

var reg = regexp.MustCompile(`x=(-?\d+)\.\.(-?\d+).*y=(-?\d+)\.\.(-?\d+)`)

type target struct {
	x1, x2, y1, y2 int
}

func parse(r io.Reader) target {
	bytes, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	parts := reg.FindStringSubmatch(string(bytes))

	return target{
		x1: toInt(parts[1]),
		x2: toInt(parts[2]),
		y1: toInt(parts[3]),
		y2: toInt(parts[4]),
	}
}

func toInt(stringVal string) int {
	intVal, err := strconv.Atoi(stringVal)
	if err != nil {
		panic(err)
	}

	return intVal
}

func solve(r io.Reader) {
	targetArea := parse(r)

	maxY := 0

	for dx := 0; dx < targetArea.x2; dx++ {
		for dy := 0; dy < -targetArea.y1; dy++ {
			p := &probe{
				x:          0,
				y:          0,
				dx:         dx,
				dy:         dy,
				targetArea: targetArea,
			}

			if probeMaxY, intersects := p.intersectsEventually(); intersects {
				if probeMaxY > maxY {
					maxY = probeMaxY
				}
			}
		}
	}

	fmt.Println(maxY)
}

type probe struct {
	x, y, dx, dy int

	targetArea target
}

func (p *probe) intersectsEventually() (int, bool) {
	maxY := p.y

	for {
		if p.intersects() {
			return maxY, true
		}

		p.simulate()
		if p.y > maxY {
			maxY = p.y
		}

		// fallen below/to the right of the target
		if p.x > p.targetArea.x2 || p.y < p.targetArea.y1 {
			return 0, false
		}
	}
}

func (p *probe) simulate() {
	p.x += p.dx
	p.y += p.dy

	if p.dx > 0 {
		p.dx -= 1
	} else if p.dx < 0 {
		p.dx += 1
	}

	p.dy -= 1
}

func (p *probe) intersects() bool {
	return p.x >= p.targetArea.x1 && p.x <= p.targetArea.x2 && p.y >= p.targetArea.y1 && p.y <= p.targetArea.y2
}

func main() {
	solve(strings.NewReader("target area: x=20..30, y=-10..-5"))
	solve(input())
}
