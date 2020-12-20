package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "20", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

const (
	edgeTop = 0
	edgeRight = 1
	edgeBottom = 2
	edgeLeft = 3
)

func getEdge(t [][]byte, which int) []byte {
	if t == nil {
		return nil
	}

	switch which {
	case edgeTop:
		// top
		return t[0]
	case edgeRight:
		// right
		newContainer := make([]byte, len(t))
		for i := 0; i < len(t); i++ {
			newContainer[i] = t[i][len(t[i]) - 1]
		}

		return newContainer
	case edgeBottom:
		// bottom
		return t[len(t) - 1]
	case edgeLeft:
		// left
		newContainer := make([]byte, len(t))
		for i := 0; i < len(t); i++ {
			newContainer[i] = t[i][0]
		}

		return newContainer
	default:
		panic(which)
	}
}

func match(edgeA, edgeB []byte) bool {
	if edgeB == nil {
		return true
	}

	if len(edgeA) != len(edgeB) {
		return false
	}

	for i := 0; i < len(edgeA); i++ {
		if edgeA[i] != edgeB[i] {
			return false
		}
	}

	return true
}

func rotate(tile [][]byte) [][]byte {
	// assumes tile is a square
	newTile := make([][]byte, len(tile))
	for x := 0; x < len(tile); x++ {
		newTile[x] = make([]byte, len(tile[x]))
	}

	for x := 0; x < len(tile); x++ {
		for y := 0; y < len(tile); y++ {
			newTile[len(tile) - y - 1][x] = tile[x][y]
		}
	}

	return newTile
}

func flipY(tile [][]byte) [][]byte {
	newTile := make([][]byte, len(tile))
	for x := 0; x < len(tile); x++ {
		newTile[x] = make([]byte, len(tile[x]))
	}

	for x := 0; x < len(tile); x++ {
		for y := 0; y < len(tile); y++ {
			newTile[x][y] = tile[len(tile) - x - 1][y]
		}
	}

	return newTile
}

func flipX(tile [][]byte) [][]byte {
	newTile := make([][]byte, len(tile))
	for x := 0; x < len(tile); x++ {
		newTile[x] = make([]byte, len(tile[x]))
	}

	for x := 0; x < len(tile); x++ {
		for y := 0; y < len(tile); y++ {
			newTile[x][y] = tile[x][len(tile[x]) - y - 1]
		}
	}

	return newTile
}


func eq(tileA, tileB [][]byte) bool {
	for x := 0; x < len(tileA); x++ {
		for y := 0; y < len(tileB); y++ {
			if tileA[x][y] != tileB[x][y] {
				return false
			}
		}
	}

	return true
}

func contains(tiles [][][]byte, tile [][]byte) bool {
	for _, tileA := range tiles {
		if eq(tileA, tile) {
			return true
		}
	}

	return false
}

func generateRotatesFlips(tile [][]byte) [][][]byte {
	var retTiles [][][]byte
	for rotation := 0; rotation < 4; rotation += 1 {
		for xFlip := 0; xFlip < 2; xFlip += 1 {
			for yFlip := 0; yFlip < 2; yFlip += 1 {

				newTile := tile
				for r := 0; r < rotation; r+= 1 {
					newTile = rotate(tile)
				}

				if xFlip == 1{
					newTile = flipY(newTile)
				}

				if yFlip == 1{
					newTile = flipX(newTile)
				}

				if !contains(retTiles, newTile) {
					retTiles = append(retTiles, newTile)
				}

			}
		}
	}

	return retTiles
}


type assignment struct {
	id int

	tile [][]byte
}

func (a assignment) getEdge(which int) []byte{
	return getEdge(a.tile, which)
}

func possibilities(tiles map[int][][]byte, used map[int]bool, assignments [][]*assignment, x, y int) []*assignment {
	var topMatch []byte
	var rightMatch []byte
	var bottomMatch []byte
	var leftMatch []byte

	if y != 0 && assignments[x][y-1] != nil{
		topMatch = assignments[x][y-1].getEdge(edgeBottom)
	}

	if x != len(assignments) - 1 && assignments[x + 1][y] != nil{
		rightMatch = assignments[x + 1][y].getEdge(edgeLeft)
	}

	if y != len(assignments[x]) - 1 && assignments[x][y + 1] != nil {
		bottomMatch = assignments[x][y + 1].getEdge(edgeTop)
	}

	if x != 0 && assignments[x - 1][y] != nil {
		leftMatch = assignments[x - 1][y].getEdge(edgeRight)
	}

	var possib []*assignment

	for id, tile := range tiles {
		if used[id] {
			continue
		}

		ways := generateRotatesFlips(tile)

		for _, way := range ways {
			if match(getEdge(way, edgeTop), topMatch) &&
				match(getEdge(way, edgeRight), rightMatch) &&
				match(getEdge(way, edgeBottom), bottomMatch) &&
				match(getEdge(way, edgeLeft), leftMatch) {
				possib = append(possib, &assignment{
					id:   id,
					tile: way,
				})
			}
		}
	}

	return possib
}

func assign(tiles map[int][][]byte, used map[int]bool, assignments [][]*assignment) bool {
	for x := 0; x < len(assignments); x++ {
		for y := 0; y < len(assignments[x]); y++ {

			if assignments[x][y] != nil {
				continue
			}

			for _, possibAssignment := range possibilities(tiles, used, assignments, x, y) {

				assignments[x][y] = possibAssignment
				used[possibAssignment.id] = true

				if assign(tiles, used, assignments) {
					return true
				}

				assignments[x][y] = nil
				delete(used, possibAssignment.id)
			}

			return false

		}
	}

	return true
}

func solve(r io.Reader) {

	tiles := make(map[int][][]byte)

	scanner := bufio.NewScanner(r)

	var currTile [][]byte
	var currId int
	for scanner.Scan() {
		row := scanner.Text()

		if strings.Index(row, "Tile ") == 0 {
			var err error
			currId, err = strconv.Atoi(row[5:len(row) - 1])
			if err != nil {
				panic(err)
			}
		} else if len(row) != 0 {
			currTile = append(currTile, []byte(row))
		} else {
			if currId == 0 || currTile == nil {
				panic("bad format")
			}

			tiles[currId] = currTile

			currTile = nil
			currId = 0
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	if currTile != nil {
		tiles[currId] = currTile
	}

	//fmt.Println(tiles)

	size := int(math.Sqrt(float64(len(tiles))))

	assignments := make([][]*assignment, size)
	for i := 0; i < size; i++ {
		assignments[i] = make([]*assignment, size)
	}

	used := make(map[int]bool)

	if !assign(tiles, used, assignments) {
		panic("impossible")
	}

	ret := assignments[0][0].id * assignments[0][size - 1].id * assignments[size - 1][0].id * assignments[size - 1][size - 1].id
	fmt.Println(assignments[0][0].id,  assignments[0][size - 1].id, assignments[size - 1][0].id, assignments[size - 1][size - 1].id, ret)

	/*for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			printTile(assignments[x][y].tile)
		}
	}*/
}

func printTile(tile [][]byte) {
	for x := 0; x < len(tile); x++ {
		for y := 0; y < len(tile); y++ {
			fmt.Printf("%s", []byte{tile[x][y]})
		}

		fmt.Println()
	}

	fmt.Println()
}

func main() {

	/*test := [][]byte{
		{'A', 'B', 'C'},
		{'D', 'E', 'F'},
		{'G', 'H', 'I'}}

	printTile(test)
	printTile(rotate(test))
	printTile(flipY(test))
	printTile(flipX(test))*/

	solve(strings.NewReader("Tile 2311:\n..##.#..#.\n##..#.....\n#...##..#.\n####.#...#\n##.##.###.\n##...#.###\n.#.#.#..##\n..#....#..\n###...#.#.\n..###..###\n\nTile 1951:\n#.##...##.\n#.####...#\n.....#..##\n#...######\n.##.#....#\n.###.#####\n###.##.##.\n.###....#.\n..#.#..#.#\n#...##.#..\n\nTile 1171:\n####...##.\n#..##.#..#\n##.#..#.#.\n.###.####.\n..###.####\n.##....##.\n.#...####.\n#.##.####.\n####..#...\n.....##...\n\nTile 1427:\n###.##.#..\n.#..#.##..\n.#.##.#..#\n#.#.#.##.#\n....#...##\n...##..##.\n...#.#####\n.#.####.#.\n..#..###.#\n..##.#..#.\n\nTile 1489:\n##.#.#....\n..##...#..\n.##..##...\n..#...#...\n#####...#.\n#..#.#.#.#\n...#.#.#..\n##.#...##.\n..##.##.##\n###.##.#..\n\nTile 2473:\n#....####.\n#..#.##...\n#.##..#...\n######.#.#\n.#...#.#.#\n.#########\n.###.#..#.\n########.#\n##...##.#.\n..###.#.#.\n\nTile 2971:\n..#.#....#\n#...###...\n#.#.###...\n##.##..#..\n.#####..##\n.#..####.#\n#..#.#..#.\n..####.###\n..#.#.###.\n...#.#.#.#\n\nTile 2729:\n...#.#.#.#\n####.#....\n..#.#.....\n....#..#.#\n.##..##.#.\n.#.####...\n####.#.#..\n##.####...\n##..#.##..\n#.##...##.\n\nTile 3079:\n#.#.#####.\n.#..######\n..#.......\n######....\n####.#..#.\n.#...#.##.\n#.#####.##\n..#.###...\n..#.......\n..#.###..."))
	solve(input())
}
