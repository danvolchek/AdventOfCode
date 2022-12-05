package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

// Instruction represents a crane instruction to move amount crates from stack from to stack to.
type Instruction struct {
	amount, from, to int
}

type Puzzle struct {
	stacks []lib.Stack[string]

	instructions []Instruction
}

func parseCrates(crates string, p *Puzzle) {
	cratesByLine := strings.Split(crates, "\n")
	for lineIndex := range cratesByLine {
		// parse in reverse order so stack is built properly
		line := cratesByLine[len(cratesByLine)-1-lineIndex]

		for i := 1; i < len(line); i += 4 {
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
}

func parseInstructions(instructions string, p *Puzzle) {
	for _, instruction := range strings.Split(strings.TrimSpace(instructions), "\n") {
		nums := lib.Ints(instruction)

		p.instructions = append(p.instructions, Instruction{
			amount: nums[0],
			from:   nums[1] - 1,
			to:     nums[2] - 1,
		})
	}
}

func solve(puzzle Puzzle) string {
	for _, instr := range puzzle.instructions {
		for i := 0; i < instr.amount; i++ {
			puzzle.stacks[instr.to].Push(puzzle.stacks[instr.from].Pop())
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
		ParseF: lib.ParseStringChunks(parseCrates, parseInstructions),
		SolveF: solve,
	}

	solver.Expect("    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2", "CMZ")
	solver.Verify("VJSFHWGFT")
}
