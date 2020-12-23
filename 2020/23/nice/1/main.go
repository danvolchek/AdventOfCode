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

const (
	numMoves = 100
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "23", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

// a linked list of cups
type cup struct {
	next *cup

	label int
}

func (c *cup) forAll(action func(*cup)) {
	tmp := c
	for {
		action(tmp)

		tmp = tmp.next

		if tmp == c {
			break
		}
	}
}

// returns a linked list of cups and a map of labels to cups for efficient access
func parse(r io.Reader) (*cup, map[int]*cup) {
	var currentCup *cup
	labelsToCups := make(map[int]*cup)

	scanner := bufio.NewScanner(r)
	chomp := func() string {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
		}
		return scanner.Text()
	}
	labels := chomp()


	var lastCup *cup
	addCup := func(label int) {
		newCup := &cup{
			label: label,
		}

		labelsToCups[label] = newCup

		if lastCup != nil {
			lastCup.next = newCup
		}
		lastCup = newCup
	}

	for i := 0; i < len(labels); i++ {
		label, err := strconv.Atoi(labels[i : i+1])
		if err != nil {
			panic(err)
		}

		addCup(label)

		if currentCup == nil {
			currentCup = lastCup
		}
	}

	// connect linked list
	lastCup.next = currentCup

	return currentCup, labelsToCups
}

func move(curr *cup, labelToCup map[int]*cup) *cup {
	// first action: pick up cups (by finding the cups after the current cup, and then removing them from the linked list)
	first, second, third := curr.next, curr.next.next, curr.next.next.next

	curr.next = third.next
	third.next = nil

	// second action: select destination cup (by finding the right destination value, and then the cup associated with it)
	decrement := func(val, total int) int {
		if val-1 == 0 {
			return total
		}

		return val - 1
	}

	destinationValue := decrement(curr.label, len(labelToCup))

	for i := 0; i < 3; i++ {
		if destinationValue == first.label || destinationValue == second.label || destinationValue == third.label {
			destinationValue = decrement(destinationValue, len(labelToCup))
		}
	}

	destination := labelToCup[destinationValue]

	// third action: put cups back (by inserting them after the destination cup)
	third.next = destination.next
	destination.next = first

	// fourth action: select new current cup
	return curr.next
}

func solve(currentCup *cup, labelsToCups map[int]*cup) string {
	for i := 0; i < numMoves; i++ {
		currentCup = move(currentCup, labelsToCups)
	}

	// the ordering starting from 1, with 1 removed
	return ordering(labelsToCups[1])[1:]
}

func ordering(startingCup *cup) string {
	var sb strings.Builder

	startingCup.forAll(func(cup *cup) {
		sb.WriteString(fmt.Sprintf("%d", cup.label))
	})

	return sb.String()
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("389125467"))))
	fmt.Println(solve(parse(input())))
}
