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

	acc := 0

	visited := make(map[int]int)

	for i := 0; i < len(instructions); i++ {
		visited[i]++

		if visited[i] == 2 {
			fmt.Println(acc)
			break
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

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(strings.NewReader("nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"))
	solve(input())
}
