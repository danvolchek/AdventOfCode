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

func parse(r io.Reader) []instr {
	var instructions []instr

	scanner := bufio.NewScanner(r)
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

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return instructions
}

const (
	north   = 'N'
	south   = 'S'
	east    = 'E'
	west    = 'W'
	left    = 'L'
	right   = 'R'
	forward = 'F'
)

var directions = []byte{north, east, south, west}

type instructionHandler func(ship *ferryState, instruction instr)

func handleCardinalMovement(ship *ferryState, instruction instr) {
	switch instruction.action {
	case north:
		ship.y += instruction.value
	case south:
		ship.y -= instruction.value
	case east:
		ship.x += instruction.value
	case west:
		ship.x -= instruction.value
	default:
		panic(instruction.action)
	}
}

func handleTurn(ship *ferryState, instruction instr) {
	if instruction.value%90 != 0 {
		panic(instruction.value)
	}

	turns := instruction.value / 90

	for i := 0; i < turns; i++ {
		switch instruction.action {
		case left:
			ship.directionIndex -= 1
		case right:
			ship.directionIndex += 1
		default:
			panic(instruction.action)
		}

		if ship.directionIndex == -1 {
			ship.directionIndex = len(directions) - 1
		} else if ship.directionIndex == len(directions) {
			ship.directionIndex = 0
		}
	}
}

func handleForward(ship *ferryState, instruction instr) {
	newInstruction := instr{
		action: directions[ship.directionIndex],
		value:  instruction.value,
	}

	handleCardinalMovement(ship, newInstruction)
}

var instructionHandlers = map[byte]instructionHandler{
	north:   handleCardinalMovement,
	south:   handleCardinalMovement,
	east:    handleCardinalMovement,
	west:    handleCardinalMovement,
	left:    handleTurn,
	right:   handleTurn,
	forward: handleForward,
}

type ferryState struct {
	x, y, directionIndex int
}

func solve(instructions []instr) int {
	ship := &ferryState{
		x:              0,
		y:              0,
		directionIndex: 1,
	}

	for _, instruction := range instructions {
		handler, ok := instructionHandlers[instruction.action]
		if !ok {
			panic(instruction.action)
		}

		handler(ship, instruction)
	}

	if ship.x < 0 {
		ship.x *= -1
	}

	if ship.y < 0 {
		ship.y *= -1
	}

	return ship.x + ship.y
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("F10\nN3\nF7\nR90\nF11"))))
	fmt.Println(solve(parse(input())))
}
