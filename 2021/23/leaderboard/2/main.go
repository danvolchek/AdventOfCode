package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "23", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type state struct {
	hallway [11]byte
	rooms   [4][4]byte
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var s state
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	first := scanner.Bytes()
	scanner.Scan()
	second := scanner.Bytes()

	i := 0
	for _, v := range bytes.Split(first, []byte{'#'}) {
		if len(v) == 0 || v[0] == '#' || len(bytes.TrimSpace(v)) == 0 {
			continue
		}

		s.rooms[i][0] = v[0]
		i += 1

	}
	i = 0
	for _, v := range bytes.Split(second, []byte{'#'}) {
		if len(v) == 0 || v[0] == '#' || len(bytes.TrimSpace(v)) == 0 {
			continue
		}

		s.rooms[i][3] = v[0]
		i += 1

	}

	s.rooms[0][1] = 'D'
	s.rooms[1][1] = 'C'
	s.rooms[2][1] = 'B'
	s.rooms[3][1] = 'A'
	s.rooms[0][2] = 'D'
	s.rooms[1][2] = 'B'
	s.rooms[2][2] = 'A'
	s.rooms[3][2] = 'C'

	/*seen := map[string]int{}
	for i := 0; i < len(s.rooms); i++ {
		for spot := 0; spot < 2; spot ++ {
			seen[s.rooms[i][spot]]++
			s.rooms[i][spot] += strconv.Itoa(seen[s.rooms[i][spot]])
		}
	}*/

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	stop := make(chan bool)
	go func() {
		s := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-s.C:
				fmt.Println("Solutions", solutions)
			case <-stop:
				s.Stop()
				fmt.Println("Total solutions", solutions)
				stop <- true
				return
			}
		}
	}()
	solutions = 0
	c := minCost(s)
	stop <- true
	<-stop
	fmt.Println(c)
}

var solutions = 0

func minCost(s state) int {
	if done(s) {
		solutions += 1
		return 0
	}

	min := 2 << 32
	for _, m := range validMoves(s) {
		if m.cost > min {
			continue
		}

		cost := minCost(m.s)
		if cost == -1 {
			continue
		}

		vc := m.cost + cost
		if vc < min {
			min = vc
		}
	}

	return min
}

type move struct {
	cost int
	s    state
}

func validMoves(s state) []move {
	return concat(pathFindRooms(s), pathFindHallway(s))
}

func concat(moves ...[]move) []move {
	var ret []move
	for _, m := range moves {
		ret = append(ret, m...)
	}

	return ret
}

func pathFindHallway(s state) []move {
	var ret []move

	for i, hallwaySpot := range s.hallway {
		if hallwaySpot != 0 {
			for _, v := range hallwayMoves(s, i) {
				ret = append(ret, v)
			}
		}
	}

	return ret
}

func pathFindRooms(s state) []move {
	var ret []move

	for i, room := range s.rooms {
		for j, spot := range room {
			if spot != 0 {
				for _, v := range roomMove(s, i, j) {
					ret = append(ret, v)
				}
			}
		}
	}

	return ret
}

func hallwayMoves(s state, hallwayPos int) []move {
	var ret []move

	which := s.hallway[hallwayPos]
	//if which == 0 {
	//	panic(which)
	//}

	//            0        11
	var spots []int
	for j := hallwayPos + 1; j < len(s.hallway); j++ {
		if s.hallway[j] != 0 {
			break
		}

		if j%2 == 0 && j > 0 && j < 10 {
			spots = append(spots, j)
		}
	}

	for j := hallwayPos - 1; j >= 0; j-- {
		if s.hallway[j] != 0 {
			break
		}

		if j%2 == 0 && j > 0 && j < 10 {
			spots = append(spots, j)
		}
	}

	for _, j := range spots {
		possibleRoom := (j / 2) - 1

		for possibleSpot := len(s.rooms[possibleRoom]) - 1; possibleSpot >= 0; possibleSpot-- {
			if roomFits(s, possibleRoom, which, possibleSpot) {
				newState := s
				newState.hallway[hallwayPos] = 0
				newState.rooms[possibleRoom][possibleSpot] = which
				ret = append(ret, move{
					cost: cost(which, abs(hallwayPos-j)+possibleSpot+1),
					s:    newState,
				})

				break // if can go into second spot, just go there; no reason to stop one early
			}
		}
	}

	return ret
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}

	return v
}

func roomMove(s state, room, spot int) []move {
	// blocked to leave room
	for i := spot - 1; i >= 0; i-- {
		if s.rooms[room][i] != 0 {
			return nil
		}
	}

	which := s.rooms[room][spot]
	//if which == 0 {
	//	panic(which)
	//}

	// don't leave if you're in the right spot and everyone behind of you is in the right spot
	if destRoom(which) == room {
		all := true
		for i := spot + 1; i < len(s.rooms[room]); i++ {
			if destRoom(s.rooms[room][i]) != room {
				all = false
				break
			}
		}

		if all {
			return nil
		}
	}

	var ret []move

	// 0, 1, 2, 3 -> 2, 4, 6, 8
	hallwayPos := 2 * (room + 1)

	var spots []int
	for j := hallwayPos + 1; j < len(s.hallway); j++ {
		if s.hallway[j] != 0 {
			break
		}

		if j == 2 || j == 4 || j == 6 || j == 8 {
			continue
		}

		spots = append(spots, j)
	}

	for j := hallwayPos - 1; j >= 0; j-- {
		if s.hallway[j] != 0 {
			break
		}

		if j == 2 || j == 4 || j == 6 || j == 8 {
			continue
		}

		spots = append(spots, j)
	}

	for _, j := range spots {
		newState := s
		newState.hallway[j] = which
		newState.rooms[room][spot] = 0

		c := cost(which, abs(hallwayPos-j)+spot+1)
		ret = append(ret, move{
			cost: c,
			s:    newState,
		})

		for _, v := range hallwayMoves(newState, j) {
			return []move{
				{
					cost: c + v.cost,
					s:    v.s,
				},
			} // if we found a hallway it's the best possible one, don't try others
		}
	}

	return ret
}

func cost(which byte, steps int) int {
	switch which {
	case 'A':
		return steps
	case 'B':
		return 10 * steps
	case 'C':
		return 100 * steps
	case 'D':
		return 1000 * steps
	default:
		panic(which)
	}
}

func roomFits(s state, room int, which byte, spot int) bool {
	if destRoom(which) != room {
		return false
	}

	// can't enter if blocked
	for i := 0; i <= spot; i++ {
		if s.rooms[room][i] != 0 {
			return false
		}
	}

	// can't enter if rest wrong
	for i := spot + 1; i < len(s.rooms[room]); i++ {
		if destRoom(s.rooms[room][i]) != room {
			return false
		}
	}

	return true
}

func destRoom(which byte) int {
	return int(which - 65)

	/*switch which {
	case 'A' :
		return 0
	case 'B':
		return 1
	case 'C':
		return 2
	case 'D':
		return 3
	default:
		panic(which)
	}*/
}

func done(s state) bool {
	for i := 0; i < len(s.rooms); i++ {
		for j := 0; j < len(s.rooms[i]); j++ {
			if int(s.rooms[i][j]) != 65+i {
				return false
			}
		}
	}

	return true
}

func main() {
	solve(strings.NewReader("#############\n#...........#\n###B#C#B#D###\n  #A#D#C#A#\n  #########"))
	solve(input())
}
