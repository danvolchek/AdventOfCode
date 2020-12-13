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

type instructionHandler func(ship *ferryState, instruction instr)

func handleCardinalMovement(ship *ferryState, instruction instr) {
	switch instruction.action {
	case north:
		ship.waypointOffsetY += instruction.value
	case south:
		ship.waypointOffsetY -= instruction.value
	case east:
		ship.waypointOffsetX += instruction.value
	case west:
		ship.waypointOffsetX -= instruction.value
	default:
		panic(instruction)
	}
}

func handleTurn(ship *ferryState, instruction instr) {
	if instruction.value%90 != 0 {
		panic(instruction)
	}

	turns := instruction.value / 90

	for i := 0; i < turns; i++ {
		switch instruction.action {
		case left:
			tmp := ship.waypointOffsetX
			ship.waypointOffsetX = -ship.waypointOffsetY
			ship.waypointOffsetY = tmp
		case right:
			tmp := ship.waypointOffsetX
			ship.waypointOffsetX = ship.waypointOffsetY
			ship.waypointOffsetY = -tmp
		default:
			panic(instruction)
		}
	}
}

func handleForward(ship *ferryState, instruction instr) {
	for i := 0; i < instruction.value; i++ {
		ship.x += ship.waypointOffsetX
		ship.y += ship.waypointOffsetY
	}
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
	x, y                             int
	waypointOffsetX, waypointOffsetY int
}

func solve(instructions []instr) int {
	ship := &ferryState{
		x:               0,
		y:               0,
		waypointOffsetX: 10,
		waypointOffsetY: 1,
	}

	for _, instruction := range instructions {
		handler := instructionHandlers[instruction.action]
		if handler == nil {
			panic(instruction)
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
