package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"sort"
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

	tileOrder := make([]int, len(tiles))
	i := 0
	for id := range tiles {
		tileOrder[i] = id
		i+=1
	}

	sort.Ints(tileOrder)

	for _, id := range tileOrder {
		tile := tiles[id]
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
			fmt.Print(assignments[x][y].id, " ")
		}
		fmt.Println()
	}*/

	newAssignments := make([][]*assignment, size)
	for i := 0; i < size; i++ {
		newAssignments[i] = make([]*assignment, size)
	}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			newAssignments[size - x -1][size - y - 1] = assignments[y][x]
		}
	}

	//fmt.Println()

	//ret = assignments[0][0].id * assignments[0][size - 1].id * assignments[size - 1][0].id * assignments[size - 1][size - 1].id
	//fmt.Println(assignments[0][0].id,  assignments[0][size - 1].id, assignments[size - 1][0].id, assignments[size - 1][size - 1].id, ret)


	//fmt.Println("----------------")

	//printTile(assignments[0][0].tile)
	//picture := amalgamate(assignments)
	//printTile(picture)

	picture := amalgamate(newAssignments)
	printTile(picture)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			fmt.Print(newAssignments[x][y].id, " ")
		}
		fmt.Println()
	}

	monster := loadVeryScaryMonster()

	removed := findAndRemove(picture, monster)


	printTile(removed)
	fmt.Println(count(removed, '#'))
}

func count(tile [][]byte, b byte) int {
	var sum int

	for x := 0; x < len(tile); x++ {
		for y := 0; y < len(tile[x]); y++ {
			if tile[x][y] == b {
				sum += 1
			}
		}
	}

	return sum
}

func find(haystack [][]byte, needle [][]byte) (bool, int, int) {
	for x := 0; x < len(haystack) - len(needle); x++ {
		for y := 0; y < len(haystack[x]) - len(needle[x % len(needle)]); y++ {

			all := true

			for i:= 0; i < len(needle); i++ {
				for j:=0; j < len(needle[i]); j++ {
					if needle[i][j] == 0 {
						continue
					}

					if haystack[x + i][y + j] != needle[i][j] {
						all = false
						break
					}
				}

				if !all {
					break
				}
			}

			if all {
				return true, x ,y
			}

		}
	}

	return false, 0 ,0
}

func remove(haystack [][]byte, needle [][]byte, x, y int) {
	for i:= 0; i < len(needle); i++ {
		for j:=0; j < len(needle[i]); j++ {
			if needle[i][j] == 0 {
				continue
			}

			haystack[x + i][y + j] ='O'
		}

	}


}

func findAndRemove(haystack [][]byte, needle [][]byte) [][]byte {
	for {
		ways := generateRotatesFlips(haystack)

		var didFind bool
		for _, way := range ways {

			found, x ,y := find(way, needle)

			if found {
				remove(way, needle, x, y)
				haystack = way
				didFind = true
				break
			}
		}

		if !didFind {
			break
		}
	}

	return haystack
}

func loadVeryScaryMonster() [][]byte {
	const raw = "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "

	parts := strings.Split(raw, "\n")

	monster := make([][]byte, len(parts))
	for y := 0; y < len(parts); y++ {
		monster[y] = make([]byte, len(parts[y]))
	}

	for y := 0; y < len(parts); y++ {
		for x := 0; x < len(parts[y]); x++ {
			if parts[y][x] == '#' {
				monster[y][x] = '#'
			}
		}
	}

	return monster
}

func amalgamate(assignments [][]*assignment) [][]byte {
	// 1, 0 for biggy testing and -2, 1 for real
	extraRoom := -2
	borderSize := 1

	// assume square
	targetTileSize := len(assignments[0][0].tile) + extraRoom
	size := targetTileSize * len(assignments)

	ret := make([][]byte, size)
	for i := 0; i < size; i++ {
		ret[i] = make([]byte, size)
	}


	for x := 0; x < len(assignments); x++ {
		for y := 0;y < len(assignments[x]); y++ {
			tile := assignments[x][y].tile

			targetX := targetTileSize * x //(len(assignments) - x - 1)
			targetY := targetTileSize * y //(len(assignments[x]) - y - 1)

			tile = flipX(rotate(tile))


			for sourceX := borderSize; sourceX <  len(tile) - borderSize; sourceX ++ {
				for sourceY := borderSize; sourceY < len(tile[sourceX]) - borderSize; sourceY++ {
					ret[targetX + sourceX - borderSize][targetY + sourceY - borderSize] = tile[sourceY][sourceX]
				}
			}

			/*for x := 0; x < targetTileSize; x++ {
				for y :=0; y < targetTileSize; y++ {

					actX := targetTileSize - x - 1 + targetX
					actY := targetTileSize - y - 1 + targetY

					regX := x + targetX
					regY := y + targetY

					//fmt.Println("--", ret[x][y], ret[y][x])
					tmp := ret[actX][actY]
					ret[actX][actY] = ret[regY][regX]
					ret[regY][regX] = tmp
					//fmt.Println(ret[x][y], ret[y][x], "--")
				}
			}*/

			//printTile(ret)
		}
	}

	return ret
}

func printTile(tile [][]byte) {
	for x := 0; x < len(tile); x++ {
		for y := 0; y < len(tile[x]); y++ {
			if tile[x][y] == 0 {
				fmt.Print(" ")
				continue
			}
			fmt.Printf("%s", []byte{tile[x][y]})
		}

		fmt.Println()
	}

	fmt.Println()
}

func main() {

	test := [][]byte{
		{'A', 'B', 'C'},
		{'D', 'E', 'F'},
		{'G', 'H', 'I'}}

	printTile(test)
	/*printTile(rotate(test))
	printTile(flipY(test))
	printTile(flipX(test))*/

	//printTile(loadVeryScaryMonster())

	solve(strings.NewReader("Tile 2311:\n..##.#..#.\n##..#.....\n#...##..#.\n####.#...#\n##.##.###.\n##...#.###\n.#.#.#..##\n..#....#..\n###...#.#.\n..###..###\n\nTile 1951:\n#.##...##.\n#.####...#\n.....#..##\n#...######\n.##.#....#\n.###.#####\n###.##.##.\n.###....#.\n..#.#..#.#\n#...##.#..\n\nTile 1171:\n####...##.\n#..##.#..#\n##.#..#.#.\n.###.####.\n..###.####\n.##....##.\n.#...####.\n#.##.####.\n####..#...\n.....##...\n\nTile 1427:\n###.##.#..\n.#..#.##..\n.#.##.#..#\n#.#.#.##.#\n....#...##\n...##..##.\n...#.#####\n.#.####.#.\n..#..###.#\n..##.#..#.\n\nTile 1489:\n##.#.#....\n..##...#..\n.##..##...\n..#...#...\n#####...#.\n#..#.#.#.#\n...#.#.#..\n##.#...##.\n..##.##.##\n###.##.#..\n\nTile 2473:\n#....####.\n#..#.##...\n#.##..#...\n######.#.#\n.#...#.#.#\n.#########\n.###.#..#.\n########.#\n##...##.#.\n..###.#.#.\n\nTile 2971:\n..#.#....#\n#...###...\n#.#.###...\n##.##..#..\n.#####..##\n.#..####.#\n#..#.#..#.\n..####.###\n..#.#.###.\n...#.#.#.#\n\nTile 2729:\n...#.#.#.#\n####.#....\n..#.#.....\n....#..#.#\n.##..##.#.\n.#.####...\n####.#.#..\n##.####...\n##..#.##..\n#.##...##.\n\nTile 3079:\n#.#.#####.\n.#..######\n..#.......\n######....\n####.#..#.\n.#...#.##.\n#.#####.##\n..#.###...\n..#.......\n..#.###..."))
	solve(input())
}
