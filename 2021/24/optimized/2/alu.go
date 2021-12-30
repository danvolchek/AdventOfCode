package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// This file contains an implementation of the ALU described in the puzzle
// It's used  verify the solver's correctness

type operation struct {
	code       opCode
	arg1, arg2 argument
}

type opCode int

const (
	opInp opCode = iota
	opAdd
	opMul
	opDiv
	opMod
	opEql
)

type argument interface {
	Load(a *alu) int
	Store(a *alu, value int)
}

type register int

const (
	registerW register = iota
	registerX
	registerY
	registerZ
)

var (
	argumentW = registerArgument{reg: registerW}
	argumentX = registerArgument{reg: registerX}
	argumentY = registerArgument{reg: registerY}
	argumentZ = registerArgument{reg: registerZ}
)

type registerArgument struct {
	reg register
}

func (r registerArgument) Load(a *alu) int {
	switch r.reg {
	case registerW:
		return a.w
	case registerX:
		return a.x
	case registerY:
		return a.y
	case registerZ:
		return a.z
	default:
		panic(r.reg)
	}
}

func (r registerArgument) Store(a *alu, value int) {
	switch r.reg {
	case registerW:
		a.w = value
	case registerX:
		a.x = value
	case registerY:
		a.y = value
	case registerZ:
		a.z = value
	default:
		panic(r.reg)
	}
}

type constantArgument struct {
	value int
}

func (c constantArgument) Load(_ *alu) int {
	return c.value
}

func (c constantArgument) Store(_ *alu, _ int) {
	panic("can't store to constant argument")
}

func opCodeFromString(str string) opCode {
	switch str {
	case "add":
		return opAdd
	case "inp":
		return opInp
	case "mul":
		return opMul
	case "div":
		return opDiv
	case "eql":
		return opEql
	case "mod":
		return opMod
	default:
		panic(str)
	}
}

func parseArg(str string) argument {
	switch str {
	case "w":
		return argumentW
	case "x":
		return argumentX
	case "y":
		return argumentY
	case "z":
		return argumentZ
	default:
		intVal, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		return constantArgument{value: intVal}
	}
}

func parseOperations(r io.Reader) []operation {
	var operations []operation

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		code := opCodeFromString(parts[0])
		switch len(parts[1:]) {
		case 1:
			operations = append(operations, operation{
				code: code,
				arg1: parseArg(parts[1]),
			})
		case 2:
			operations = append(operations, operation{
				code: code,
				arg1: parseArg(parts[1]),
				arg2: parseArg(parts[2]),
			})
		default:
			panic(parts)
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return operations
}

type alu struct {
	w, x, y, z int

	inputs []int
}

func (a *alu) run(operations []operation) {
	for _, op := range operations {
		switch op.code {
		case opInp:
			op.arg1.Store(a, a.inputs[0])
			a.inputs = a.inputs[1:]
		case opAdd:
			op.arg1.Store(a, op.arg1.Load(a)+op.arg2.Load(a))
		case opMul:
			op.arg1.Store(a, op.arg1.Load(a)*op.arg2.Load(a))
		case opDiv:
			op.arg1.Store(a, op.arg1.Load(a)/op.arg2.Load(a))
		case opMod:
			op.arg1.Store(a, op.arg1.Load(a)%op.arg2.Load(a))
		case opEql:
			if op.arg1.Load(a) == op.arg2.Load(a) {
				op.arg1.Store(a, 1)
			} else {
				op.arg1.Store(a, 0)
			}
		default:
			panic(op.code)
		}
	}
}
