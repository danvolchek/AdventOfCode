package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type Monkey struct {
	name string

	num     int
	isConst bool

	dep1, dep2 string
	op         byte
}

func parse(line string) Monkey {
	nums := lib.Ints(line)

	if len(nums) == 1 {
		return Monkey{
			name:    line[0:4],
			num:     nums[0],
			isConst: true,
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

type MonkeyOp interface {
	// HasVariable returns whether this MonkeyOp eventually contains the variable being solved for.
	HasVariable() bool

	// Evaluate evaluates the value of this MonkeyOp. It must not have any variables in it.
	Evaluate() int

	// Reverse computes the value needed for this MonkeyOp such that Evaluate would return target.
	Reverse(target int) int
}

// Used to sanity check our answer. nil if we don't have an answer yet.
var sanityCheckAnswer *int

// The human monkey. It is the variable, can only be evaluated if the problem is solved/being checked,
// and reversing it yields the answer.
type humanMonkey struct {
}

func (h humanMonkey) HasVariable() bool {
	return true
}

func (h humanMonkey) Evaluate() int {
	if sanityCheckAnswer == nil {
		panic("can't evaluate human")
	}

	return *sanityCheckAnswer
}

func (h humanMonkey) Reverse(target int) int {
	return target
}

// A constant monkey. It is not the variable, evaluates to its constant value, and can't be reversed.
type constMonkey struct {
	name     string
	constant int
}

func (c constMonkey) HasVariable() bool {
	return false
}

func (c constMonkey) Evaluate() int {
	return c.constant
}

func (c constMonkey) Reverse(_ int) int {
	panic("can't reverse constant")
}

// An operation monkey, which performs a binary operation on two other monkeys.
// It has a variable if either of its sub operations have a variable, and can be evaluated and reversed.
type operationMonkey struct {
	left, right MonkeyOp
	which       byte
}

func (b operationMonkey) HasVariable() bool {
	return b.left.HasVariable() || b.right.HasVariable()
}

func (b operationMonkey) Evaluate() int {
	left, right := b.left.Evaluate(), b.right.Evaluate()

	var result int
	switch b.which {
	case '+':
		result = left + right
	case '*':
		result = left * right
	case '/':
		result = left / right
	case '-':
		result = left - right
	case '=':
		result = 1
		if left != right {
			result = 0
		}
	default:
		panic(b.which)
	}

	return result
}

func (b operationMonkey) Reverse(target int) int {
	var val int
	var other MonkeyOp

	leftVar, rightVar := b.left.HasVariable(), b.right.HasVariable()

	if leftVar && rightVar || (!leftVar && !rightVar) {
		panic("must be exactly one side constant and one side variable")
	}

	if leftVar {
		other = b.left
		val = b.right.Evaluate()
	} else {
		other = b.right
		val = b.left.Evaluate()
	}

	switch b.which {
	case '+':
		return other.Reverse(target - val)
	case '*':
		return other.Reverse(target / val)
	case '/':
		if other == b.left {
			return other.Reverse(val * target)
		}
		return other.Reverse(val / target)
	case '-':
		if other == b.left {
			return other.Reverse(val + target)
		}

		return other.Reverse(val - target)
	case '=':
		return other.Reverse(val)
	default:
		panic(b.which)
	}
}

// createOp creates a MonkeyOp from a Monkey.
func createOp(curr Monkey, monkeyMap map[string]Monkey) MonkeyOp {
	switch curr.name {
	case "root":
		return operationMonkey{
			left:  createOp(monkeyMap[curr.dep1], monkeyMap),
			right: createOp(monkeyMap[curr.dep2], monkeyMap),
			which: '=',
		}
	case "humn":
		return humanMonkey{}
	default:
		if curr.isConst {
			return constMonkey{
				name:     curr.name,
				constant: curr.num,
			}
		}
		return &operationMonkey{
			left:  createOp(monkeyMap[curr.dep1], monkeyMap),
			right: createOp(monkeyMap[curr.dep2], monkeyMap),
			which: curr.op,
		}
	}
}

func solve(lines []Monkey) int {
	sanityCheckAnswer = nil

	// build monkey operation tree
	monkeyMap := make(map[string]Monkey)
	for _, monkey := range lines {
		monkeyMap[monkey.name] = monkey
	}
	rootMonkey := createOp(monkeyMap["root"], monkeyMap)

	// calculate the answer
	answer := rootMonkey.Reverse(0)

	// sanity check the answer by evaluating with that answer
	sanityCheckAnswer = &answer
	result := rootMonkey.Evaluate()
	if result != 1 {
		panic("answer is wrong somehow")
	}

	return answer
}

func main() {
	solver := lib.Solver[[]Monkey, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("root: pppw + sjmn\ndbpl: 5\ncczh: sllz + lgvd\nzczc: 2\nptdq: humn - dvpt\ndvpt: 3\nlfqf: 4\nhumn: 5\nljgn: 2\nsjmn: drzm * dbpl\nsllz: 4\npppw: cczh / lfqf\nlgvd: ljgn * ptdq\ndrzm: hmdt - zczc\nhmdt: 32", 301)
	solver.Verify(3342154812537)
}
