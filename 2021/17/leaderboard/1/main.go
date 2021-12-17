package main

import (
	"bufio"
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

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	l := scanner.Text()

	res := reg.FindStringSubmatch(l)

	x1, x2, y1, y2 := toInt(res[1]), toInt(res[2]), toInt(res[3]), toInt(res[4])

	fmt.Println(x1, x2, y1, y2)

	maxY := 0

	for x := -1000; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			p := &probe{0, 0, x, y}

			if m, hits := p.hits(x1, x2, y1, y2); hits {
				if m > maxY {
					maxY = m
				}
			}
		}
	}

	fmt.Println(maxY)

}

func toInt(v string) int {
	vv, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return vv
}

type probe struct {
	x, y, dx, dy int
}

func (p *probe) hits(x1, x2, y1, y2 int) (int, bool) {
	maxY := p.y
	p.simulate()
	if p.y > maxY {
		maxY = p.y
	}

	for {
		if p.intersects(x1, x2, y1, y2) {
			return maxY, true
		}

		p.simulate()
		if p.y > maxY {
			maxY = p.y
		}

		if p.x > x2 || p.y < y1 {
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

func (p *probe) intersects(x1, x2, y1, y2 int) bool {
	return p.x >= x1 && p.x <= x2 && p.y >= y1 && p.y <= y2
}

func main() {
	solve(strings.NewReader("target area: x=20..30, y=-10..-5"))
	solve(input())
}
