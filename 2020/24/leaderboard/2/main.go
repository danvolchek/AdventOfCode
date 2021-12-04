package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "24", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func traverse(path string) tile {
	x, y := 0, 0
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'e':
			x += 1
		case 'w':
			x -= 1
		case 'n':
			switch path[i+1] {
			case 'w':
				x += 1
				y += 1
			case 'e':
				y += 1
			default:
				panic(path[i:])
			}
		case 's':
			switch path[i+1] {
			case 'w':
				y -= 1
			case 'e':
				x -= 1
				y -= 1
			default:
				panic(path[i:])
			}
		default:
			panic(path[i:])
		}
	}

	return tile{x: x, y: y}
}

type tile struct {
	x, y int
}

func adjacentBlackTiles(tiles map[tile]bool, curr tile) int {
	num := 0

	for _, offset := range []tile{{1, 0}, {-1, 0}, {1, 1}, {0, 1}, {0, -1}, {-1, -1}} {
		if tiles[tile{curr.x + offset.x, curr.y + offset.y}] {
			num += 1
		}
	}

	return num
}

func day(tiles map[tile]bool) {
	const limit = 100

	changes := make(map[tile]bool)

	for x := -limit; x < limit; x += 1 {
		for y := -limit; y < limit; y += 1 {
			tile := tile{x, y}
			isBlack := tiles[tile]
			adjacent := adjacentBlackTiles(tiles, tile)

			if isBlack && (adjacent == 0 || adjacent > 2) {
				changes[tile] = false
			} else if !isBlack && adjacent == 2 {
				changes[tile] = true
			}
		}
	}

	for k, v := range changes {
		tiles[k] = v
	}
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	tiles := make(map[tile]bool)

	for scanner.Scan() {
		row := scanner.Text()

		tile := traverse(row)

		tiles[tile] = !tiles[tile]
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	for i := 0; i < 100; i += 1 {
		day(tiles)
	}

	numBlack := 0
	for _, isBlack := range tiles {
		if isBlack {
			numBlack += 1
		}
	}
	fmt.Println(numBlack)
}

func main() {
	solve(strings.NewReader("sesenwnenenewseeswwswswwnenewsewsw\nneeenesenwnwwswnenewnwwsewnenwseswesw\nseswneswswsenwwnwse\nnwnwneseeswswnenewneswwnewseswneseene\nswweswneswnenwsewnwneneseenw\neesenwseswswnenwswnwnwsewwnwsene\nsewnenenenesenwsewnenwwwse\nwenwwweseeeweswwwnwwe\nwsweesenenewnwwnwsenewsenwwsesesenwne\nneeswseenwwswnwswswnw\nnenwswwsewswnenenewsenwsenwnesesenew\nenewnwewneswsewnwswenweswnenwsenwsw\nsweneswneswneneenwnewenewwneswswnese\nswwesenesewenwneswnwwneseswwne\nenesenwswwswneneswsenwnewswseenwsese\nwnwnesenesenenwwnenwsewesewsesesew\nnenewswnwewswnenesenwnesewesw\neneswnwswnwsenenwnwnwwseeswneewsenese\nneswnwewnwnwseenwseesewsenwsweewe\nwseweeenwnesenwwwswnew"))
	solve(input())
}
