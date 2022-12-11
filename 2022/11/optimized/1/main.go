package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"strings"
)

type Monkey struct {
	// the items the monkey is currently holding
	items []int

	// the monkey operation to change your worry
	operation func(int) int

	// the monkey test operation to determine where to throw next
	test func(int) bool

	// the target monkeys based on the test outcome
	target map[bool]int

	// the total inspections this monkey has done
	inspections int
}

func (m *Monkey) Inspections() int {
	return m.inspections
}

func parse(chunk string) *Monkey {
	lines := strings.Split(chunk, "\n")

	var monkey Monkey

	// Line 1:
	// Starting items: 74, 87
	monkey.items = lib.Ints(lines[1])

	// Line 2:
	// Operation: new = old + 2
	// Operation: new = old * 2
	// Operation: new = old * old
	nums := lib.Ints(lines[2])
	mul := strings.Contains(lines[2], "*")
	monkey.operation = func(i int) int {
		var arg int
		if len(nums) == 0 {
			arg = i
		} else {
			arg = nums[0]
		}

		if mul {
			return i * arg
		} else {
			return i + arg
		}
	}

	// Line 3:
	// Test: divisible by 5
	arg := lib.Int(lines[3])
	monkey.test = func(i int) bool {
		return i%arg == 0
	}

	// Line 4:
	// If true: throw to monkey 7
	// Line 5:
	// If false: throw to monkey 4
	monkey.target = map[bool]int{
		true:  lib.Int(lines[4]),
		false: lib.Int(lines[5]),
	}

	return &monkey
}

// Add adds an item to this monkey's items.
func (m *Monkey) Add(item int) {
	m.items = append(m.items, item)
}

// Inspect runs the monkey's inspect and throw algorithm.
func (m *Monkey) Inspect(monkeys []*Monkey) {
	for _, item := range m.items {
		item = m.operation(item) / 3

		targetMonkeyIndex := m.target[m.test(item)]

		monkeys[targetMonkeyIndex].Add(item)
	}

	m.inspections += len(m.items)
	m.items = []int{}
}

const (
	rounds = 20
)

func solve(monkeys []*Monkey) int {
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.Inspect(monkeys)
		}
	}

	inspections := lib.Map(monkeys, (*Monkey).Inspections)

	slices.Sort(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func main() {
	solver := lib.Solver[[]*Monkey, int]{
		ParseF: lib.ParseChunks(parse),
		SolveF: solve,
	}

	solver.Expect("Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1", 10605)
	solver.Verify(56350)
}
