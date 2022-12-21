package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type Monkey struct {
	name string

	num     int
	isReady bool

	dep1, dep2 string
	op         byte
}

func parse(line string) Monkey {
	nums := lib.Ints(line)

	if len(nums) == 1 {
		return Monkey{
			name:    line[0:4],
			num:     nums[0],
			isReady: true,
		}
	}

	lineAfter := line[6:]

	return Monkey{
		name: line[0:4],
		dep1: lineAfter[0:4],
		dep2: lineAfter[7:],
		op:   lineAfter[5],
	}
}

func solve(lines []Monkey) int {
	results := make(map[string]int)

	for _, monkey := range lines {
		if monkey.isReady {
			results[monkey.name] = monkey.num
		}
	}

	for {
		v, ok := results["root"]
		if ok {
			return v
		}

		for _, monkey := range lines {
			_, ok := results[monkey.name]
			if ok {
				continue
			}

			num1, okD1 := results[monkey.dep1]
			num2, okD2 := results[monkey.dep2]
			if okD1 && okD2 {
				switch monkey.op {
				case '+':
					results[monkey.name] = num1 + num2
				case '*':
					results[monkey.name] = num1 * num2
				case '/':
					results[monkey.name] = num1 / num2
				case '-':
					results[monkey.name] = num1 - num2
				default:
					panic(monkey.op)
				}
			}
		}
	}

}

func main() {
	solver := lib.Solver[[]Monkey, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("root: pppw + sjmn\ndbpl: 5\ncczh: sllz + lgvd\nzczc: 2\nptdq: humn - dvpt\ndvpt: 3\nlfqf: 4\nhumn: 5\nljgn: 2\nsjmn: drzm * dbpl\nsllz: 4\npppw: cczh / lfqf\nlgvd: ljgn * ptdq\ndrzm: hmdt - zczc\nhmdt: 32", 152)
	solver.Verify(155708040358220)
}
