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
	input, err := os.Open(path.Join("2020", "12", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type instr struct {
	action byte
	value  int
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var instructions []instr

	for scanner.Scan() {
		row := scanner.Text()

		intVal, err := strconv.Atoi(row[1:])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, instr{
			action: row[0],
			value:  intVal,
		})
	}

	directions := []byte{'N', 'E', 'S', 'W'}

	currDirection := 1
	currX := 0
	currY := 0

	turn := func(direction, amount int) {
		if amount%90 != 0 {
			panic(amount)
		}
		for amount != 0 {
			currDirection += direction
			amount -= 90

			if currDirection == -1 {
				currDirection = len(directions) - 1
			}

			if currDirection == len(directions) {
				currDirection = 0
			}
		}

	}

	handle := func(instruction instr) {
		switch instruction.action {
		case 'N':
			currY += instruction.value
		case 'S':
			currY -= instruction.value
		case 'E':
			currX += instruction.value
		case 'W':
			currX -= instruction.value
		case 'L':
			turn(-1, instruction.value)
		case 'R':
			turn(1, instruction.value)
		default:
			panic(instruction)
		}
	}

	for _, instruction := range instructions {

		switch instruction.action {
		case 'F':
			val := currDirection
			//if val == 0 || val == 2 {
			//	val = (currDirection + 2) % 4
			//}
			handle(instr{
				action: directions[val],
				value:  instruction.value,
			})
		default:
			handle(instruction)
		}
	}

	fmt.Println(currX, currY)

	if currX < 0 {
		currX *= -1
	}

	if currY < 0 {
		currY *= -1
	}

	fmt.Println(currX + currY)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(strings.NewReader("F10\nN3\nF7\nR90\nF11"))
	solve(input())
}
