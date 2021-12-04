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

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	numBlack := 0

	tiles := make(map[tile]bool)

	for scanner.Scan() {
		row := scanner.Text()

		tile := traverse(row)

		tiles[tile] = !tiles[tile]

		if tiles[tile] {
			numBlack += 1
		} else {
			numBlack -= 1
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(numBlack)
}

func main() {
	solve(strings.NewReader("sesenwnenenewseeswwswswwnenewsewsw\nneeenesenwnwwswnenewnwwsewnenwseswesw\nseswneswswsenwwnwse\nnwnwneseeswswnenewneswwnewseswneseene\nswweswneswnenwsewnwneneseenw\neesenwseswswnenwswnwnwsewwnwsene\nsewnenenenesenwsewnenwwwse\nwenwwweseeeweswwwnwwe\nwsweesenenewnwwnwsenewsenwwsesesenwne\nneeswseenwwswnwswswnw\nnenwswwsewswnenenewsenwsenwnesesenew\nenewnwewneswsewnwswenweswnenwsenwsw\nsweneswneswneneenwnewenewwneswswnese\nswwesenesewenwneswnwwneseswwne\nenesenwswwswneneswsenwnewswseenwsese\nwnwnesenesenenwwnenwsewesewsesesew\nnenewswnwewswnenesenwnesewesw\neneswnwswnwsenenwnwnwwseeswneewsenese\nneswnwewnwnwseenwseesewsenwsweewe\nwseweeenwnesenwwwswnew"))
	solve(input())
}
