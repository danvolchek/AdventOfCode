package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Effect struct {
	countdown int
	action    func(*CPU)
}

type CPU struct {
	x int

	cycle int

	effects []*Effect
}

func (c *CPU) addCommand(command string) {
	if strings.Contains(command, "noop") {
		c.effects = append(c.effects, &Effect{
			countdown: 1,
			action: func(cpu *CPU) {

			},
		})
	} else if strings.Contains(command, "addx") {
		arg := lib.Atoi(command[len("addx "):])
		c.effects = append(c.effects, &Effect{
			countdown: 2,
			action: func(cpu *CPU) {
				cpu.x += arg
			},
		})
	}
}

func (c *CPU) step() {
	defer func() {
		c.cycle += 1
	}()

	if len(c.effects) > 0 {
		c.effects[0].countdown -= 1
		if c.effects[0].countdown == 0 {
			c.effects[0].action(c)
			c.effects = c.effects[1:]
		}
	}
}

func parse(line string) string {
	return line
}

func solve(lines []string) string {

	var cpu CPU
	cpu.x = 1
	cpu.cycle = 1

	for _, line := range lines {
		cpu.addCommand(line)
	}

	var picture strings.Builder

	for len(cpu.effects) != 0 {
		i := cpu.cycle

		xPos := (cpu.cycle - 1) % 40

		if lib.Abs(xPos-cpu.x) <= 1 {
			picture.WriteString("#")
		} else {
			picture.WriteString(".")
		}

		//fmt.Printf("During cycle %v, x is %v\n", i, cpu.x)
		if i >= 20 && (i-20)%40 == 0 {

			//fmt.Printf("strength: %v * %v = %v\n", i, cpu.x, i*cpu.x)
		}

		cpu.step()

		//fmt.Printf("After cycle %v, x is %v\n", i, cpu.x)

		if i%40 == 0 {
			picture.WriteString("\n")
		}

	}

	return "\n" + strings.TrimSpace(picture.String())
}

func main() {
	solver := lib.Solver[[]string, string]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Test("noop\naddx 3\naddx -5")
	solver.Expect("addx 15\naddx -11\naddx 6\naddx -3\naddx 5\naddx -1\naddx -8\naddx 13\naddx 4\nnoop\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx 5\naddx -1\naddx -35\naddx 1\naddx 24\naddx -19\naddx 1\naddx 16\naddx -11\nnoop\nnoop\naddx 21\naddx -15\nnoop\nnoop\naddx -3\naddx 9\naddx 1\naddx -3\naddx 8\naddx 1\naddx 5\nnoop\nnoop\nnoop\nnoop\nnoop\naddx -36\nnoop\naddx 1\naddx 7\nnoop\nnoop\nnoop\naddx 2\naddx 6\nnoop\nnoop\nnoop\nnoop\nnoop\naddx 1\nnoop\nnoop\naddx 7\naddx 1\nnoop\naddx -13\naddx 13\naddx 7\nnoop\naddx 1\naddx -33\nnoop\nnoop\nnoop\naddx 2\nnoop\nnoop\nnoop\naddx 8\nnoop\naddx -1\naddx 2\naddx 1\nnoop\naddx 17\naddx -9\naddx 1\naddx 1\naddx -3\naddx 11\nnoop\nnoop\naddx 1\nnoop\naddx 1\nnoop\nnoop\naddx -13\naddx -19\naddx 1\naddx 3\naddx 26\naddx -30\naddx 12\naddx -1\naddx 3\naddx 1\nnoop\nnoop\nnoop\naddx -9\naddx 18\naddx 1\naddx 2\nnoop\nnoop\naddx 9\nnoop\nnoop\nnoop\naddx -1\naddx 2\naddx -37\naddx 1\naddx 3\nnoop\naddx 15\naddx -21\naddx 22\naddx -6\naddx 1\nnoop\naddx 2\naddx 1\nnoop\naddx -10\nnoop\nnoop\naddx 20\naddx 1\naddx 2\naddx 2\naddx -6\naddx -11\nnoop\nnoop\nnoop", "\n##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....")
	solver.Verify("\n###..####..##..###..#..#.###..####.###..\n#..#....#.#..#.#..#.#.#..#..#.#....#..#.\n#..#...#..#....#..#.##...#..#.###..###..\n###...#...#.##.###..#.#..###..#....#..#.\n#....#....#..#.#....#.#..#....#....#..#.\n#....####..###.#....#..#.#....####.###..")
}
