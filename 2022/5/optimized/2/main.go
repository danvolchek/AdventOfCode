package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

// Instruction represents a crane instruction to move amount crates from stack from to stack to.
type Instruction struct {
	amount, from, to int
}

// Puzzle represents the puzzle input, which is a set of stacks of crates and a set of movement instructions.
type Puzzle struct {
	stacks []lib.Stack[string]

	instructions []Instruction
}

func parse(inp string) Puzzle {
	crates, instructions, found := strings.Cut(inp, "\n\n")
	if !found {
		panic("bad input")
	}

	var p Puzzle

	cratesByLine := strings.Split(crates, "\n")
	for lineIndex := range cratesByLine {
		// parse in reverse order so stack is built properly
		line := cratesByLine[len(cratesByLine)-1-lineIndex]

		for i := 0; i < len(line); i++ {
			if line[i] >= 'A' && line[i] <= 'Z' {
				stackNumber := (i - 1) / 4

				// grow stacks if needed
				for len(p.stacks) < stackNumber+1 {
					p.stacks = append(p.stacks, lib.Stack[string]{})
				}

				// add crate to stack
				p.stacks[stackNumber].Push(string(line[i]))
			}
		}
	}

	for _, instruction := range strings.Split(strings.TrimSpace(instructions), "\n") {
		nums := lib.Ints(instruction)

		p.instructions = append(p.instructions, Instruction{
			amount: nums[0],
			from:   nums[1],
			to:     nums[2],
		})
	}

	return p
}

func solve(puzzle Puzzle) string {
	for _, instr := range puzzle.instructions {
		// simulate lifting crates with order preserved by putting them into a temporary stack and popping that onto to
		var temp lib.Stack[string]
		for i := 0; i < instr.amount; i++ {
			temp.Push(puzzle.stacks[instr.from-1].Pop())
		}

		for i := 0; i < instr.amount; i++ {
			puzzle.stacks[instr.to-1].Push(temp.Pop())
		}
	}

	result := ""
	for _, stack := range puzzle.stacks {
		result += stack.Pop()
	}
	return result
}

func main() {
	solver := lib.Solver[Puzzle, string]{
		ParseF: lib.ParseStringFunc(parse),
		SolveF: solve,
	}

	solver.Expect("    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2", "MCD")
	solver.Verify("LCTQFBVZV")
}
