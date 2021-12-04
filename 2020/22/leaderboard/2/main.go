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

type deck struct {
	cards []int
}

func (d *deck) Pop() int {
	ret := d.cards[0]

	d.cards = d.cards[1:]

	return ret
}

func (d *deck) Push(card int) {
	d.cards = append(d.cards, card)
}

func (d *deck) Length() int {
	return len(d.cards)
}

func (d *deck) Cards() []int {
	return d.cards
}

func (d *deck) String() string {
	return fmt.Sprintf("%v", d.cards)
}

func (d *deck) PushMany(cards []int) {
	for _, card := range cards {
		d.Push(card)
	}
}

/*type deck struct {
	cards      [50]int
	start, end int
}

func (d *deck) Pop() int {
	if d.start == d.end {
		panic("empty deck")
	}

	ret := d.cards[d.start]

	d.start += 1

	if d.start == len(d.cards) {
		d.start = 0
	}

	return ret
}

func (d *deck) PushMany(cards []int) {
	for _, card := range cards {
		d.Push(card)
	}
}

func (d *deck) Push(card int) {
	if d.end == d.start {
		panic("full deck")
	}

	//if d.end == len(d.cards) {
	//	fmt.Println("shift")
	//	d.shift()
	//}

	d.cards[d.end] = card

	d.end += 1

	if d.end == len(d.cards) {
		d.end = 0
	}
}

func (d *deck) Length() int {
	return d.end - d.start
}

func (d *deck) Cards() []int {

	if d.start <= d.end {
		return d.cards[d.start:d.end]
	}

	return append(d.cards[d.start:], d.cards[0:d.end]...)
}

func (d *deck) CardsToHash() [50]int {
	var ret [50]int
	for index, card := range d.Cards() {
		ret[index] = card
	}

	return ret
}

func (d *deck) String() string {
	return fmt.Sprintf("[%v, %v]: %v", d.start, d.end, d.cards)
}

func (d *deck) shift() {
	if d.start == 0 {
		return
	}

	for i := 0; i < (d.end - d.start); i++ {
		d.cards[i] = d.cards[i + d.start]
	}

	for i := d.end - d.start; i < len(d.cards); i++ {
		d.cards[i] = 0
	}

	d.end -= d.start
	d.start = 0
}

func newDeck() *deck {
	return &deck {
		end: 1,
	}
}
*/

/*type gamePool struct {
	workers chan worker
}

type worker struct {
	result chan player
}

func (w worker) run(g *game) {
	go func() {
		g.runToCompletion()
		w.result <- g.winner
	}()
}

func (g *gamePool) RunToCompletion(gg *game) player {
	worker := <-g.workers

	worker.run(gg)

	g.workers <- worker

	return <-worker.result
}

func newGamePool(workers int) gamePool {
	pool := gamePool{
		workers: make(chan worker, workers),
	}

	for i := 0; i < workers; i++ {
		pool.workers <- worker{result: make(chan player)}
	}

	return pool
}*/

/*func hash2(a, b uint64) uint64 {
	return ((a+b)*(a+b+1))/2 + b
}

func hash(a, b []int) uint64 {
	var first, second uint64
	var curr uint64

	if len(a) == 1 {
		first = uint64(a[0])
		second = uint64(b[0])

		curr = hash2(first, second)

		for _, val := range b[1:] {
			curr = hash2(curr, uint64(val))
		}

	} else {
		first = uint64(a[0])
		second = uint64(a[1])

		curr = hash2(first, second)

		for _, val := range a[2:] {
			curr = hash2(curr, uint64(val))
		}

		for _, val := range b {
			curr = hash2(curr, uint64(val))
		}
	}

	return curr
}*/

//var pool = newGamePool(16)

//var history = make(map[uint64]bool)
//var gameHistory = make(map[uint64]player)

func newDeck() *deck {
	return &deck{}
}

type deckSnapshot struct {
	decka, deckb [50]int
}

type player int

const (
	player1 player = iota
	player2
)

type game struct {
	player1Deck, player2Deck *deck
	history                  map[deckSnapshot]bool

	over   bool
	winner player

	r int
}

func newGame(player1Deck, player2Deck *deck) *game {
	return &game{
		player1Deck: player1Deck,
		player2Deck: player2Deck,
		history:     make(map[deckSnapshot]bool),
		r:           1,
	}
}

func (g *game) createSnapshot() *deckSnapshot {

	//return hash(g.player1Deck.Cards(), g.player2Deck.Cards())

	var ret deckSnapshot

	for i, card := range g.player1Deck.Cards() {
		ret.decka[i] = card
	}

	for i, card := range g.player2Deck.Cards() {
		ret.deckb[i] = card
	}

	return &ret
}

func (g *game) round() {

	snapshot := *g.createSnapshot()

	if g.history[snapshot] {
		//fmt.Println("bailing out")
		g.over = true
		g.winner = player1
		return
	}

	g.history[snapshot] = true

	player1Card, player2Card := g.player1Deck.Pop(), g.player2Deck.Pop()

	var winner player
	if g.player1Deck.Length() >= player1Card && g.player2Deck.Length() >= player2Card {
		winner = g.determineWinnerRecursive(g.player1Deck, player1Card, g.player2Deck, player2Card)
	} else {
		winner = g.determineWinnerBaseCase(player1Card, player2Card)
	}

	g.updateDecks(winner, player1Card, player2Card)

	g.checkOver()

	g.r += 1
}

func (g *game) determineWinnerRecursive(deck11 *deck, toDraw1 int, deck22 *deck, toDraw2 int) player {

	d := newDeck()
	d.PushMany(deck11.cards[:toDraw1])
	d2 := newDeck()
	d2.PushMany(deck22.cards[:toDraw2])
	recursiveGame := newGame(d, d2)

	//fmt.Println("new game")

	//snapshot := recursiveGame.createSnapshot()

	//if winner, ok := gameHistory[snapshot]; ok {
	//	//fmt.Println("seen this entire game before")
	//	return winner
	//}

	recursiveGame.runToCompletion()
	//pool.RunToCompletion(recursiveGame)

	//gameHistory[snapshot] = recursiveGame.winner

	return recursiveGame.winner
}

func (g *game) determineWinnerBaseCase(player1Card, player2Card int) player {
	if player1Card > player2Card {
		return player1
	}

	return player2
}

func (g *game) updateDecks(winner player, player1Card, player2Card int) {
	if winner == player1 {
		g.player1Deck.Push(player1Card)
		g.player1Deck.Push(player2Card)
	} else {
		g.player2Deck.Push(player2Card)
		g.player2Deck.Push(player1Card)
	}
}

func (g *game) runToCompletion() {
	for !g.over {
		//fmt.Println(g)
		g.round()
	}

	//fmt.Println(g)

	//fmt.Println("done")

	//if g.r == 100 {
	//	panic(123)
	//}

}

func (g *game) checkOver() {
	if g.player1Deck.Length() == 0 {
		g.over = true
		g.winner = player2
	} else if g.player2Deck.Length() == 0 {
		g.over = true
		g.winner = player1
	}
}

func (g *game) String() string {
	return fmt.Sprintf("Round %v\n%v\n%v\n", g.r, g.player1Deck, g.player2Deck)
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var player1Cards []int
	var player2Cards []int

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
			player1Cards = append(player1Cards, val)
		case 1:
			player2Cards = append(player2Cards, val)
		default:
			panic(zone)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	player1Deck := newDeck()
	player1Deck.PushMany(player1Cards)

	player2Deck := newDeck()
	player2Deck.PushMany(player2Cards)

	g := newGame(player1Deck, player2Deck)

	g.runToCompletion()

	fmt.Println(g, "winner", g.winner, "p1 score", score(player1Deck.Cards()), "p2 score", score(player2Deck.Cards()))
}

func score(deck []int) int {
	score := 0

	multiplier := 1
	for i := len(deck) - 1; i >= 0; i-- {
		score += multiplier * deck[i]
		multiplier += 1
	}

	return score
}

func main() {
	//solve(strings.NewReader("Player 1:\n43\n19\n\nPlayer 2:\n2\n29\n14"))
	//solve(strings.NewReader("Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"))
	//solve(strings.NewReader("Player 1:\n48\n23\n9\n34\n37\n36\n40\n26\n49\n7\n12\n20\n6\n45\n14\n42\n18\n31\n39\n47\n44\n15\n43\n10\n35\n\nPlayer 2:\n13\n19\n21\n32\n27\n16\n11\n29\n41\n46\n33\n1\n30\n22\n38\n5\n17\n4\n50\n2\n3\n28\n8\n25\n24"))
	solve(input())
}
