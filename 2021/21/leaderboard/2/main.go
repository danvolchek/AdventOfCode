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
	input, err := os.Open(path.Join("2021", "21", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type stats struct {
	score, position int
}

type gameState struct {
	p1, p2 stats
	turn bool
}

type result struct {
	p1Wins, p2Wins int
}

var wins = make(map[gameState]result)


func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()

	first := scanner.Text()

	scanner.Scan()

	second := scanner.Text()

	first = strings.Split(first, ": ")[1]
	second = strings.Split(second, ": ")[1]

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	firstPos, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}

	secondPos, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}

	g := gameState{
		p1: stats{
			position: firstPos,
		},
		p2: stats{
			position: secondPos,
		},
		turn: true,
	}

	res := numWins(g)

	if res.p1Wins > res.p2Wins {
		fmt.Println(res.p1Wins)
	} else {
		fmt.Println(res.p2Wins)
	}
}

func numWins(g gameState) result {
	if g.p1.score >= 21 {
		return result{
			p1Wins: 1,
		}
	} else if g.p2.score >= 21 {
		return result {
			p2Wins: 1,
		}
	}

	if r, ok := wins[g]; ok {
		return r
	}

	var results []result

	for _, rolled := range possibleRolls {
		gCopy := g
		roll(&gCopy, rolled)

		results = append(results, numWins(gCopy))
	}

	var r result
	for _, res := range results {
		r.p1Wins += res.p1Wins
		r.p2Wins += res.p2Wins
	}

	wins[g] = r

	return r
}

func roll(g *gameState, num int) {
	var p *stats
	if g.turn {
		p = &g.p1
	} else {
		p = &g.p2
	}

	p.position += num
	if p.position > 10 {
		p.position -= 10
	}

	p.score += p.position

	g.turn = !g.turn
}

var possibleRolls []int = nil

func fillRolls() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				possibleRolls = append(possibleRolls, i + j + k)
			}
		}
	}
}


func main() {
	fillRolls()

	solve(strings.NewReader("Player 1 starting position: 4\nPlayer 2 starting position: 8\n"))
	solve(input())
}
