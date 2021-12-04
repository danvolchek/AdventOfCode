package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "23", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type cup struct {
	next *cup

	label int
}

func round3(curr *cup, labelToCup map[int]*cup) *cup {
	first, second, third := curr.next, curr.next.next, curr.next.next.next

	curr.next = curr.next.next.next.next
	third.next = nil

	destinationValue := curr.label - 1
	if destinationValue == 0 {
		destinationValue = len(labelToCup)
	}

	for i := 0; i < 3; i++ {
		if destinationValue == first.label || destinationValue == second.label || destinationValue == third.label {
			destinationValue -= 1

			if destinationValue == 0 {
				destinationValue = len(labelToCup)
			}
		}
	}

	destination := labelToCup[destinationValue]

	third.next = destination.next
	destination.next = first

	return curr.next
}

func solve(r io.Reader) {
	var curr *cup

	var builderCup *cup

	scanner := bufio.NewScanner(r)

	labelToCup := make(map[int]*cup)

	addCup := func(val int) {
		tempCup := &cup{
			label: val,
		}

		labelToCup[val] = tempCup

		builderCup.next = tempCup

		builderCup = tempCup
	}

	for scanner.Scan() {
		row := scanner.Text()

		for i := 0; i < len(row); i++ {
			val, err := strconv.Atoi(row[i : i+1])
			if err != nil {
				panic(err)
			}

			if curr == nil {
				curr = &cup{
					label: val,
				}

				labelToCup[val] = curr

				builderCup = curr
			} else {
				addCup(val)
			}

		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	const real = true

	if real {
		for i := 10; i <= 1000000; i += 1 {
			addCup(i)
		}
	}

	builderCup.next = curr

	printy := func(curr *cup) {
		forAllCups(curr, func(cup *cup) {
			fmt.Print(cup.label, " ")
		})
		fmt.Println()
	}

	numCups := func(curr *cup) int {
		sum := 0
		forAllCups(curr, func(c *cup) {
			sum += 1
		})

		return sum
	}

	if !real {
		printy(curr)
	}

	totalCups := numCups(curr)

	fmt.Println(totalCups)

	rounds := 100
	if real {
		rounds = 10000000
	}
	for i := 0; i < rounds; i++ {
		curr = round3(curr, labelToCup)
		if !real {
			printy(curr)
		}
	}

	if !real {
		printy(curr)
	} else {
		fmt.Print(labelToCup[1].next.label, labelToCup[1].next.next.label, labelToCup[1].next.label*labelToCup[1].next.next.label)
	}
}

func forAllCups(curr *cup, action func(*cup)) {
	tmp := curr
	for {
		action(tmp)

		tmp = tmp.next

		if tmp == curr {
			break
		}
	}
}

func main() {
	//solve(strings.NewReader("389125467"))
	solve(input())
}
