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

const (
	opAcc = "acc"
	opJmp = "jmp"
	opNop = "nop"
)

type instruction struct {
	code string
	arg  int
}

func parse(r io.Reader) []instruction {
	var instructions []instruction

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		arg, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		instruction := instruction{
			code: parts[0],
			arg: arg,
		}

		instructions = append(instructions, instruction)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return instructions
}

func solve(instructions []instruction) int {
	swap := func(instruction *instruction) {
		switch instruction.code {
		case opNop:
			instruction.code = opJmp
		case opJmp:
			instruction.code = opNop
		}
	}

	for i := 0; i < len(instructions); i++ {
		if instructions[i].code == opAcc {
			continue
		}

		swap(&instructions[i])
		acc, ok := vm(instructions)
		if ok {
			return acc
		}
		swap(&instructions[i])
	}

	panic("no solution")
}

func vm(instructions []instruction) (int, bool) {
	acc := 0

	visited := make(map[int]bool, len(instructions))

	for i := 0; i < len(instructions); i++ {
		if visited[i] {
			return 0, false
		}

		visited[i] = true

		switch instructions[i].code {
		case opNop:
			continue
		case opJmp:
			i += instructions[i].arg - 1
		case opAcc:
			acc += instructions[i].arg
		default:
			panic(fmt.Sprintf("can't handle %s", instructions[i].code))
		}
	}

	return acc, true
}

func main() {
	fmt.Println(solve(parse(input())))
}
