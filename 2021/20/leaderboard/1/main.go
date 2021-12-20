package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "20", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type pos struct {
	x, y int
}

var outsideZero = true

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	algo := strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), ".", "0"), "#", "1")

	picture := make(map[pos]rune)

	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		part := strings.ReplaceAll(strings.ReplaceAll(line, ".", "0"), "#", "1")
		for x, p := range part {
			picture[pos{x: x, y: y}] = p
		}
		y += 1
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	doSwap := false
	switch algo[0] {
	case '0':
		// 3x3 off stays off, nothing to do
	case '1':
		switch algo[len(algo)-1] {
		case '0':
			// 3x3 all on turns off, we need to swap the out of bounds char
			doSwap = true
		case '1':
			panic("infinite pixels lit according to this algo")
		}
	}

	draw(picture)
	for i := 0; i < 2; i++ {
		enhance(algo, picture)
		draw(picture)

		if doSwap {
			outsideZero = !outsideZero
		}
	}

	sum := 0
	for _, v := range picture {
		if v == '1' {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func draw(picture map[pos]rune) {
	minp, maxp := bounds(picture)

	for y := minp.y; y <= maxp.y; y++ {
		for x := minp.x; x <= maxp.x; x++ {
			val := '.'
			v := get(picture, pos{x: x, y: y})
			if v == '1' {
				val = '#'
			}

			fmt.Printf("%c", val)
		}

		fmt.Println()
	}

	fmt.Println()

}

func bounds(picture map[pos]rune) (pos, pos) {
	minp, maxp := pos{x: 9999999, y: 999999}, pos{x: -9999999, y: -9999999}

	for p := range picture {
		if p.x < minp.x {
			minp.x = p.x
		}
		if p.y < minp.y {
			minp.y = p.y
		}

		if p.x > maxp.x {
			maxp.x = p.x
		}
		if p.y > maxp.y {
			maxp.y = p.y
		}
	}

	return minp, maxp
}

func enhance(algo string, picture map[pos]rune) {
	updates := make(map[pos]rune)

	minp, maxp := bounds(picture)

	for y := minp.y - 2; y <= maxp.y+2; y++ {
		for x := minp.x - 2; x <= maxp.x+2; x++ {
			p := pos{x: x, y: y}

			val := combine(
				get(picture, pos{
					x: p.x - 1,
					y: p.y - 1,
				}),
				get(picture, pos{
					x: p.x,
					y: p.y - 1,
				}),
				get(picture, pos{
					x: p.x + 1,
					y: p.y - 1,
				}),
				get(picture, pos{
					x: p.x - 1,
					y: p.y,
				}),
				get(picture, pos{
					x: p.x,
					y: p.y,
				}),
				get(picture, pos{
					x: p.x + 1,
					y: p.y,
				}),
				get(picture, pos{
					x: p.x - 1,
					y: p.y + 1,
				}),
				get(picture, pos{
					x: p.x,
					y: p.y + 1,
				}),
				get(picture, pos{
					x: p.x + 1,
					y: p.y + 1,
				}),
			)

			intVal, err := strconv.ParseInt(val, 2, 32)
			if err != nil {
				panic(err)
			}

			v := rune(algo[intVal])
			updates[p] = v
		}
	}

	for p, v := range updates {
		picture[p] = v
	}
}

func combine(rs ...rune) string {
	var s strings.Builder
	for _, r := range rs {
		s.WriteRune(r)
	}

	return s.String()
}

func get(picture map[pos]rune, p pos) rune {
	v, ok := picture[p]
	if ok {
		//fmt.Printf("get, %+v = %v\n", p, v)
		return v
	}

	if outsideZero {
		return '0'
	}

	return '1'
}

func main() {
	solve(strings.NewReader("..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#\n\n#..#.\n#....\n##..#\n..#..\n..###"))
	solve(input())
}
