package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
	"strings"
)

type stack[T any] struct {
	items []T
}

func (s *stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *stack[T]) Empty() bool {
	return len(s.items) == 0
}

func (s *stack[T]) Pop() T {
	if s.Empty() {
		panic("empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *stack[T]) Top() T {
	return s.items[len(s.items)-1]
}

func (s *stack[T]) Reverse() {
	var items []T

	for !s.Empty() {
		items = append(items, s.Pop())
	}

	s.items = items
}

type instruction struct {
	amount, from, to int
}

type puzzle struct {
	stacks []*stack[string]

	instructions []instruction
}

func parse(inp string) puzzle {
	doingInstructs := false

	stacks := []*stack[string]{{}, {}, {}}
	var p puzzle

	for _, line := range strings.Split(inp, "\n") {
		if line == "" {
			doingInstructs = true
			continue
		}

		switch doingInstructs {
		case true:
			reg := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
			matches := reg.FindAllStringSubmatch(line, -1)

			firstMatchSubmatches := matches[0][1:]

			p.instructions = append(p.instructions, instruction{
				amount: lib.Atoi(firstMatchSubmatches[0]),
				from:   lib.Atoi(firstMatchSubmatches[1]) - 1,
				to:     lib.Atoi(firstMatchSubmatches[2]) - 1,
			})
		case false:
			for i := 0; i < len(line); i++ {
				if line[i] >= 'A' && line[i] <= 'Z' {
					targ := (i - 1) / 4
					for len(stacks) < targ+1 {
						stacks = append(stacks, &stack[string]{})
					}
					stacks[(i-1)/4].Push(string(line[i]))
				}
			}
		}
	}

	for _, stack := range stacks {
		stack.Reverse()
	}

	return puzzle{
		stacks:       stacks,
		instructions: p.instructions,
	}
}

func solve(puzz puzzle) string {
	for _, instr := range puzz.instructions {
		for i := 0; i < instr.amount; i++ {
			puzz.stacks[instr.to].Push(puzz.stacks[instr.from].Pop())
		}
	}

	ret := ""
	for _, s := range puzz.stacks {
		ret += s.Top()
	}
	return ret
}

func main() {
	solver := lib.Solver[puzzle, string]{
		ParseF: parse,
		SolveF: solve,
	}

	solver.Expect("    [D]    \n[N] [C]    \n[Z] [M] [P]\n 1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2", "CMZ")
	solver.Verify("VJSFHWGFT")
}
