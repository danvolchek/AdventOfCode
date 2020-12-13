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

	currX := 0
	currY := 0

	wayPointOffsetX := 10
	wayPointOffsetY := 1

	turn := func(direction, amount int) {
		if amount%90 != 0 {
			panic(amount)
		}
		for amount != 0 {
			amount -= 90
			switch direction {
			case 1:
				/* clockwise

				top right
				1,  10 -> 10, -1
				10,   1 -> 1, -10

				x, y -> y, -x

				bottom right
				1, -10 -> -10, -1
				10, -1 -> -1, -10

				x, y -> y, -x

				bottom left
				-1, -10 -> -10, 1
				-10, -1 -> -1, 10

				x, y - > y, -x

				top left
				-1,  10 -> 10, 1
				-10, 1  -> 1, 10

				x, y -> y, -x

				north
				0, 10 -> 10, 0

				x, y -> y, -x

				east
				10, 0 -> 0, -10

				x, y -> y, -x

				south
				0, -10 -> -10, 0

				x, y -> y, -x

				west
				-10, 0 -> 0, 10

				x, y -> y, -x
				*/

				tmp := wayPointOffsetX
				wayPointOffsetX = wayPointOffsetY
				wayPointOffsetY = -tmp

			case -1:
				/* counter clockwise

				top right
				1,  10 -> -10, 1
				10,   1 -> -1, 10

				x, y -> -y, x

				*/
				tmp := wayPointOffsetX
				wayPointOffsetX = -wayPointOffsetY
				wayPointOffsetY = tmp
			default:
				panic(direction)
			}
		}
	}

	handle := func(instruction instr) {
		switch instruction.action {
		case 'N':
			wayPointOffsetY += instruction.value
		case 'S':
			wayPointOffsetY -= instruction.value
		case 'E':
			wayPointOffsetX += instruction.value
		case 'W':
			wayPointOffsetX -= instruction.value
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
			for i := 0; i < instruction.value; i++ {
				currX += wayPointOffsetX
				currY += wayPointOffsetY
			}
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
