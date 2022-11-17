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
	on bool
	r  cuboid
}

type pos struct {
	x, y, z int
}

type cuboid struct {
	from, to pos
}

func (c cuboid) volume() int {
	return (c.to.x - c.from.x + 1) * (c.to.y - c.from.y + 1) * (c.to.z - c.from.z + 1)
}

type cutOutCuboid struct {
	area cuboid
	cuts []*cutOutCuboid
}

func (c cutOutCuboid) volume() int {
	cutVolume := 0
	for _, cut := range c.cuts {
		cutVolume += cut.volume()
	}

	vol := c.area.volume() - cutVolume
	//if vol < 0 {
	//	return 0
	//}

	return vol
}

var reg = regexp.MustCompile(`(...?) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)

func parseCube(line string) action {
	parts := reg.FindStringSubmatch(line)

	return action{
		on: parts[1] == "on",
		r: cuboid{
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
		},
	}
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var actions []action
	for scanner.Scan() {
		line := scanner.Text()
		actions = append(actions, parseCube(line))
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	var onCuboids []*cutOutCuboid
	for _, act := range actions {
		onCuboids = addCuboid(onCuboids, act)
	}

	fmt.Println(volume(onCuboids))
}

func volume(cubes []*cutOutCuboid) int {
	ret := 0
	for _, cube := range cubes {
		ret += cube.volume()
	}

	return ret
}

func (c *cutOutCuboid) checkIntersect(cub cuboid) {
	inter, ok := intersection(c.area, cub)
	if !ok {
		return
	}

	for _, cut := range c.cuts {
		cut.checkIntersect(inter)
	}

	c.cuts = append(c.cuts, &cutOutCuboid{area: inter})
}

func addCuboid(onCuboids []*cutOutCuboid, act action) []*cutOutCuboid {
	for _, cub := range onCuboids {
		cub.checkIntersect(act.r)
	}

	if act.on {
		onCuboids = append(onCuboids, &cutOutCuboid{
			area: act.r,
			cuts: nil,
		})
	}

	return onCuboids
}

func intersection(c1, c2 cuboid) (cuboid, bool) {
	if c1.from.x > c2.to.x || c1.to.x < c2.from.x { // c1 right of c2 || c1 left of c2
		return cuboid{}, false
	}

	if c1.from.y > c2.to.y || c1.to.y < c2.from.y { // c1 right of c2 || c1 left of c2
		return cuboid{}, false
	}

	if c1.from.z > c2.to.z || c1.to.z < c2.from.z { // c1 right of c2 || c1 left of c2
		return cuboid{}, false
	}

	return cuboid{
		from: pos{
			x: max(c1.from.x, c2.from.x),
			y: max(c1.from.y, c2.from.y),
			z: max(c1.from.z, c2.from.z),
		},
		to: pos{
			x: min(c1.to.x, c2.to.x),
			y: min(c1.to.y, c2.to.y),
			z: min(c1.to.z, c2.to.z),
		},
	}, true

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return v
}

func testIntersect() {
	for _, tc := range []struct{ a, b, c cuboid }{
		{
			a: parseCube("on x=0..10,y=0..10,z=0..10").r,
			b: parseCube("on x=0..5,y=0..5,z=0..5").r,
			c: parseCube("on x=0..5,y=0..5,z=0..5").r,
		},
		{
			a: parseCube("on x=0..10,y=0..10,z=0..10").r,
			b: parseCube("on x=5..10,y=5..10,z=5..10").r,
			c: parseCube("on x=5..10,y=5..10,z=5..10").r,
		},
		{
			a: parseCube("on x=0..10,y=0..10,z=0..10").r,
			b: parseCube("on x=-5..5,y=-5..5,z=-5..5").r,
			c: parseCube("on x=0..5,y=0..5,z=0..5").r,
		},
		{
			a: parseCube("on x=0..10,y=0..10,z=0..10").r,
			b: parseCube("on x=5..15,y=5..15,z=5..15").r,
			c: parseCube("on x=5..10,y=5..10,z=5..10").r,
		},
		{
			a: parseCube("on x=0..10,y=0..10,z=0..10").r,
			b: parseCube("on x=15..20,y=5..15,z=5..15").r,
			c: cuboid{},
		},
		{
			a: parseCube("on x=0..10,y=0..10,z=0..10").r,
			b: parseCube("on x=5..15,y=15..20,z=5..15").r,
			c: cuboid{},
		},
		{
			a: parseCube("on x=0..10,y=0..10,z=0..10").r,
			b: parseCube("on x=5..15,y=5..15,z=15..20").r,
			c: cuboid{},
		},
	} {
		inter, _ := intersection(tc.a, tc.b)
		if inter != tc.c {
			fmt.Println(tc.a, tc.b)
			fmt.Println("want", tc.c)
			fmt.Println("got", inter)
			panic("fail")
		}
	}
}

func main() {
	//testIntersect()
	solve(strings.NewReader("on x=0..10,y=0..10,z=0..10\noff x=0..5,y=0..10,z=0..10\noff x=0..5,y=0..10,z=0..10"))
	solve(strings.NewReader("on x=10..12,y=10..12,z=10..12\non x=11..13,y=11..13,z=11..13\noff x=9..11,y=9..11,z=9..11\non x=10..10,y=10..10,z=10..10"))
	solve(strings.NewReader("on x=-20..26,y=-36..17,z=-47..7\non x=-20..33,y=-21..23,z=-26..28\non x=-22..28,y=-29..23,z=-38..16\non x=-46..7,y=-6..46,z=-50..-1\non x=-49..1,y=-3..46,z=-24..28\non x=2..47,y=-22..22,z=-23..27\non x=-27..23,y=-28..26,z=-21..29\non x=-39..5,y=-6..47,z=-3..44\non x=-30..21,y=-8..43,z=-13..34\non x=-22..26,y=-27..20,z=-29..19\noff x=-48..-32,y=26..41,z=-47..-37\non x=-12..35,y=6..50,z=-50..-2\noff x=-48..-32,y=-32..-16,z=-15..-5\non x=-18..26,y=-33..15,z=-7..46\noff x=-40..-22,y=-38..-28,z=23..41\non x=-16..35,y=-41..10,z=-47..6\noff x=-32..-23,y=11..30,z=-14..3\non x=-49..-5,y=-3..45,z=-29..18\noff x=18..30,y=-20..-8,z=-3..13\non x=-41..9,y=-7..43,z=-33..15\non x=-54112..-39298,y=-85059..-49293,z=-27449..7877\non x=967..23432,y=45373..81175,z=27513..53682"))
	solve(input())
}
