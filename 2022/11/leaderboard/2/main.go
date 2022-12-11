package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"strings"
)

type Monkey struct {
	items       []int
	operation   func(int) int
	test        func(int) bool
	true, false int

	inspects int
}

func parse(lines string) []*Monkey {
	chunks := strings.Split(lines, "\n\n")

	var monkeys []*Monkey

	var divisors []int
	for _, chunk := range chunks {
		divisors = append(divisors, lib.Ints(strings.Split(chunk, "\n")[3])[0])
	}

	divisor := 1
	for _, d := range divisors {
		divisor *= d
	}

	for _, chunk := range chunks {
		clines := strings.Split(chunk, "\n")

		var monkey Monkey

		monkey.items = lib.Ints(clines[1])
		monkey.operation = func(i int) int {

			num := lib.Ints(clines[2])

			var arg int
			if len(num) == 0 {
				arg = i
			} else {
				arg = num[0]
			}

			if strings.Contains(clines[2], "*") {
				return (i * arg) % divisor
			} else if strings.Contains(clines[2], "+") {
				return (i + arg) % divisor
			} else {
				panic(clines[2])
			}
		}
		monkey.test = func(i int) bool {
			arg := lib.Ints(clines[3])[0]

			return i%arg == 0
		}
		monkey.true = lib.Ints(clines[4])[0]
		monkey.false = lib.Ints(clines[5])[0]

		monkeys = append(monkeys, &monkey)
	}

	return monkeys
}

func solve(monkeys []*Monkey) int {
	for i := 0; i < 10000; i++ {
		round(monkeys)
	}

	inspections := lib.Map(monkeys, func(m *Monkey) int {
		return m.inspects
	})

	slices.Sort(inspections)

	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func round(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		for i, item := range monkey.items {
			monkey.items[i] = monkey.operation(item)

			monkey.inspects += 1

			var targ *Monkey
			if monkey.test(monkey.items[i]) {
				targ = monkeys[monkey.true]
			} else {
				targ = monkeys[monkey.false]
			}

			targ.items = append(targ.items, monkey.items[i])
		}

		monkey.items = []int{}
	}
}

func main() {
	solver := lib.Solver[[]*Monkey, int]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1", 2713310158)
	solver.Verify(13954061248)
}
