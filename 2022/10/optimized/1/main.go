package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type PendingCommand struct {
	countdown int
	action    func(*CPU)
}

type CPU struct {
	x int

	cycle int

	commands []*PendingCommand
}

type Command interface {
	Cycles() int
	Run(*CPU)
}

type NoOpCommand struct {
}

func (n NoOpCommand) Cycles() int {
	return 1
}

func (n NoOpCommand) Run(cpu *CPU) {
}

type AddXCommand struct {
	arg int
}

func (a AddXCommand) Cycles() int {
	return 2
}

func (a AddXCommand) Run(cpu *CPU) {
	cpu.x += a.arg
}

func (c *CPU) addCommand(command Command) {
	c.commands = append(c.commands, &PendingCommand{
		countdown: command.Cycles(),
		action:    command.Run,
	})
}

func (c *CPU) step() {
	if len(c.commands) > 0 {
		command := c.commands[0]
		command.countdown -= 1
		if command.countdown == 0 {
			command.action(c)
			c.commands = c.commands[1:]
		}
	}

	c.cycle += 1
}

func parse(line string) Command {
	if line == "noop" {
		return NoOpCommand{}
	} else if strings.Contains(line, "addx") {
		arg := lib.Atoi(line[len("addx "):])
		return AddXCommand{arg: arg}
	}

	panic(line)
}

func solve(commands []Command) int {
	strength := 0

	cpu := CPU{
		x:     1,
		cycle: 1,
	}

	for _, line := range commands {
		cpu.addCommand(line)
	}

	for len(cpu.commands) != 0 {
		//fmt.Printf("During cycle %v, x is %v\n", i, cpu.x)
		if cpu.cycle >= 20 && (cpu.cycle-20)%40 == 0 {
			strength += cpu.cycle * cpu.x
		}

		cpu.step()
		//fmt.Printf("After cycle %v, x is %v\n", i, cpu.x)
	}

	return strength
}

func main() {
	solver := lib.Solver[[]Command, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Test("noop\naddx 3\naddx -5")
	solver.Expect("addx 15\naddx -11\naddx 6\naddx -3\naddx 5\naddx -1\naddx -8\naddx 13\naddx 4\nnoop\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx -35\naddx 1\naddx 24\naddx -19\naddx 1\naddx 16\naddx -11\nnoop\nnoop\naddx 21\naddx -15\nnoop\nnoop\naddx -3\naddx 9\naddx 1\naddx -3\naddx 8\naddx 1\naddx 5\nnoop\nnoop\nnoop\nnoop\nnoop\naddx -36\nnoop\naddx 1\naddx 7\nnoop\nnoop\nnoop\naddx 2\naddx 6\nnoop\nnoop\nnoop\nnoop\nnoop\naddx 1\nnoop\nnoop\naddx 7\naddx 1\nnoop\naddx -13\naddx 13\naddx 7\nnoop\naddx 1\naddx -33\nnoop\nnoop\nnoop\naddx 2\nnoop\nnoop\nnoop\naddx 8\nnoop\naddx -1\naddx 2\naddx 1\nnoop\naddx 17\naddx -9\naddx 1\naddx 1\naddx -3\naddx 11\nnoop\nnoop\naddx 1\nnoop\naddx 1\nnoop\nnoop\naddx -13\naddx -19\naddx 1\naddx 3\naddx 26\naddx -30\naddx 12\naddx -1\naddx 3\naddx 1\nnoop\nnoop\nnoop\naddx -9\naddx 18\naddx 1\naddx 2\nnoop\nnoop\naddx 9\nnoop\nnoop\nnoop\naddx -1\naddx 2\naddx -37\naddx 1\naddx 3\nnoop\naddx 15\naddx -21\naddx 22\naddx -6\naddx 1\nnoop\naddx 2\naddx 1\nnoop\naddx -10\nnoop\nnoop\naddx 20\naddx 1\naddx 2\naddx 2\naddx -6\naddx -11\nnoop\nnoop\nnoop", 13140)
	solver.Verify(13680)
}
