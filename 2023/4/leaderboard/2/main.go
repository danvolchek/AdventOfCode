package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type card struct {
	num            int
	winners, yours lib.Set[int]
}

func parse(line string) card {
	winners, yours, _ := strings.Cut(line, "|")
	return card{
		num:     lib.Ints(winners)[0],
		winners: lib.NewSet(lib.Ints(winners)[1:]),
		yours:   lib.NewSet(lib.Ints(yours)),
	}
}

func (c card) winningNumbers() int {
	return len(lib.Filter(c.yours.Items(), c.winners.Contains))
}

func (c card) totalCards(cardRef []card, memo map[int]int) int {
	if result, ok := memo[c.num]; ok {
		return result
	}

	numWinners := c.winningNumbers()

	result := numWinners

	for i := c.num; i < len(cardRef) && i < c.num+numWinners; i++ {
		result += cardRef[i].totalCards(cardRef, memo)
	}

	memo[c.num] = result
	return result
}

func solve(cards []card) int {
	memo := make(map[int]int)
	total := len(cards)

	for i := len(cards) - 1; i >= 0; i-- {
		total += cards[i].totalCards(cards, memo)
	}

	fmt.Println(memo)

	return total
}

func main() {
	solver := lib.Solver[[]card, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 30)
	solver.Verify(9425061)
}
