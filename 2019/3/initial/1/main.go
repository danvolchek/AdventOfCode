package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parseItem(item string) (dir uint8, length int) {
	dir = item[0]
	length, err := strconv.Atoi(item[1:])
	if err != nil {
		panic(err)
	}

	return
}

type pos struct {
	x, y int
}

func (p pos) equals(other pos) bool {
	return p.x == other.x && p.y == other.y
}

func (p pos) dist(other pos) int {
	xDist := p.x - other.x
	if xDist < 0 {
		xDist *= -1
	}

	yDist := p.y - other.y
	if yDist < 0 {
		yDist *= -1
	}

	return xDist + yDist
}

func parse(r io.Reader) [][]pos {
	csvReader := csv.NewReader(r)
	csvReader.FieldsPerRecord = -1
	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	wires := make([][]pos, len(rows))
	for wireNum, row := range rows {
		wires[wireNum] = transform(row)
	}

	return wires
}

func transform(wireDesc []string) []pos {
	wire := []pos{{1, 1}}
	for _, item := range wireDesc {
		dir, length := parseItem(item)

		orig := wire[len(wire)-1]
		newX, newY := orig.x, orig.y
		xDir := 1
		yDir := 1

		switch dir {
		case 'R':
			newX += length
		case 'L':
			newX -= length
			xDir = -1
		case 'U':
			newY += length
		case 'D':
			newY -= length
			yDir = -1
		}

		if newX != orig.x {
			for i := orig.x; i != newX; {
				i += xDir
				wire = append(wire, pos{i, orig.y})
			}
		} else if newY != orig.y {
			for i := orig.y; i != newY; {
				i += yDir
				wire = append(wire, pos{orig.x, i})
			}
		} else {
			panic("no x or y change")
		}

	}

	return wire
}

func minIntersect(firstWire []pos, secondWire []pos) int {
	var intersections []pos
	for _, firstWirePos := range firstWire {
		for _, secondWirePos := range secondWire {
			if !firstWirePos.equals(pos{1, 1}) && firstWirePos.equals(secondWirePos) {
				intersections = append(intersections, firstWirePos)
			}
		}
	}

	if len(intersections) == 0 {
		panic("no intersections")
	}

	set := false
	minDist := 0
	for _, intersection := range intersections {
		dist := intersection.dist(pos{1, 1})
		if !set || dist < minDist {
			minDist = dist
			set = true
		}
	}

	return minDist
}

func main() {

	test1 := parse(strings.NewReader("R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"))
	fmt.Println(minIntersect(test1[0], test1[1]))
	test2 := parse(strings.NewReader("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"))
	fmt.Println(minIntersect(test2[0], test2[1]))

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	wires := parse(input)
	fmt.Println(minIntersect(wires[0], wires[1]))
}
