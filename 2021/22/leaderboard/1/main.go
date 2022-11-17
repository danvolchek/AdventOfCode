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
	input, err := os.Open(path.Join("2021", "22", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type action struct {
	on       bool
	from, to pos
}

type pos struct {
	x, y, z int
}

var reg = regexp.MustCompile(`(...?) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var actions []action
	for scanner.Scan() {
		line := scanner.Text()

		parts := reg.FindStringSubmatch(line)
		actions = append(actions, action{
			on: parts[1] == "on",
			from: pos{
				x: toInt(parts[2]),
				y: toInt(parts[4]),
				z: toInt(parts[6]),
			},
			to: pos{
				x: toInt(parts[3]),
				y: toInt(parts[5]),
				z: toInt(parts[7]),
			},
		})

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	//fmt.Println(actions)

	cubes := make(map[pos]bool)

	for _, action := range actions {
		if oob(action.from) || oob(action.to) {
			continue
		}

		for x := action.from.x; x <= action.to.x; x++ {
			for y := action.from.y; y <= action.to.y; y++ {
				for z := action.from.z; z <= action.to.z; z++ {
					if action.on {
						cubes[pos{x: x, y: y, z: z}] = true
					} else {
						delete(cubes, pos{x: x, y: y, z: z})
					}
				}
			}
		}
	}

	fmt.Println(len(cubes))
}

func oob(p pos) bool {
	oobs := func(i int) bool { return i < -50 || i > 50 }
	return oobs(p.x) || oobs(p.y) || oobs(p.z)
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return v
}

func main() {
	solve(strings.NewReader("on x=10..12,y=10..12,z=10..12\non x=11..13,y=11..13,z=11..13\noff x=9..11,y=9..11,z=9..11\non x=10..10,y=10..10,z=10..10"))
	solve(strings.NewReader("on x=-20..26,y=-36..17,z=-47..7\non x=-20..33,y=-21..23,z=-26..28\non x=-22..28,y=-29..23,z=-38..16\non x=-46..7,y=-6..46,z=-50..-1\non x=-49..1,y=-3..46,z=-24..28\non x=2..47,y=-22..22,z=-23..27\non x=-27..23,y=-28..26,z=-21..29\non x=-39..5,y=-6..47,z=-3..44\non x=-30..21,y=-8..43,z=-13..34\non x=-22..26,y=-27..20,z=-29..19\noff x=-48..-32,y=26..41,z=-47..-37\non x=-12..35,y=6..50,z=-50..-2\noff x=-48..-32,y=-32..-16,z=-15..-5\non x=-18..26,y=-33..15,z=-7..46\noff x=-40..-22,y=-38..-28,z=23..41\non x=-16..35,y=-41..10,z=-47..6\noff x=-32..-23,y=11..30,z=-14..3\non x=-49..-5,y=-3..45,z=-29..18\noff x=18..30,y=-20..-8,z=-3..13\non x=-41..9,y=-7..43,z=-33..15\non x=-54112..-39298,y=-85059..-49293,z=-27449..7877\non x=967..23432,y=45373..81175,z=27513..53682"))
	solve(input())
}
