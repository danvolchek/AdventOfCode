package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "8", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type op struct {
	instr string
	arg   int
}

func run(instructions []op) (int, bool) {
	acc := 0

	visited := make(map[int]int)

	for i := 0; i < len(instructions); i++ {
		visited[i]++

		if visited[i] == 10000 {
			return 0, false
		}

		switch instructions[i].instr {
		case "nop":
			continue
		case "jmp":
			i += instructions[i].arg
			i--
		case "acc":
			acc += instructions[i].arg
		}
	}

	return acc, true
}

func solve(r io.Reader) {

	scanner := bufio.NewScanner(r)

	var instructions []op

	for scanner.Scan() {
		row := scanner.Text()

		parts := strings.Split(row, " ")

		op := op{
			instr: parts[0],
		}

		arg, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		op.arg = arg

		instructions = append(instructions, op)

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	bob := make([]op, len(instructions))
	for i := 0; i < len(instructions); i++ {
		if instructions[i].instr == "acc" {
			continue
		}

		copy(bob, instructions)

		if bob[i].instr == "nop" {
			bob[i].instr = "jmp"
		} else {
			bob[i].instr = "nop"
		}

		res, ok := run(bob)
		if ok {
			fmt.Println(res)
			return
		}
	}

	panic("nothing")
}

func main() {
	solve(input())
}
