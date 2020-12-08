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
	acc := vm(instructions)

	return acc
}

func vm(instructions []instruction) int {
	acc := 0

	visited := make(map[int]bool, len(instructions))

	for i := 0; i < len(instructions); i++ {
		if visited[i] {
			return acc
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

	panic("no infinite loop")
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"))))
	fmt.Println(solve(parse(input())))
}
