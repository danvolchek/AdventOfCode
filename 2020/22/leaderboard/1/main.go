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
	input, err := os.Open(path.Join("2020", "22", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func round(player1Deck, player2Deck []int) ([]int, []int) {
	player1Card := player1Deck[0]
	player2Card := player2Deck[0]

	player1Deck = player1Deck[1:]
	player2Deck = player2Deck[1:]

	if player1Card > player2Card {
		player1Deck = append(player1Deck, player1Card, player2Card)
	} else {
		player2Deck = append(player2Deck, player2Card, player1Card)
	}

	return player1Deck, player2Deck
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var player1Deck []int
	var player2Deck []int

	zone := 0
	for scanner.Scan() {
		row := scanner.Text()

		if row == "" {
			zone += 1
			continue
		}

		if strings.Index(row, "Player") == 0 {
			continue
		}

		val, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}

		switch zone {
		case 0:
			player1Deck = append(player1Deck, val)
		case 1:
			player2Deck = append(player2Deck, val)
		default:
			panic(zone)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	for len(player1Deck) != 0 && len(player2Deck) != 0 {
		player1Deck, player2Deck = round(player1Deck, player2Deck)
	}

	fmt.Println(score(player1Deck) + score(player2Deck))

}

func score(deck []int) int {
	score := 0

	for index, card := range deck {
		score += (len(deck) - index) * card
	}

	return score
}

func main() {
	solve(strings.NewReader("Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"))
	solve(input())
}
