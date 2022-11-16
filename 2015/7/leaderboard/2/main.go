package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "7", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type evaluable interface {
	Ready(context map[string]uint16) bool
	Value(context map[string]uint16) uint16
}

type constValue struct {
	value uint16
}

func (c constValue) Ready(context map[string]uint16) bool {
	return true
}

func (c constValue) Value(context map[string]uint16) uint16 {
	return c.value
}

type referenceValue struct {
	name string
}

func (c referenceValue) Ready(context map[string]uint16) bool {
	_, ok := context[c.name]

	return ok
}

func (c referenceValue) Value(context map[string]uint16) uint16 {
	return context[c.name]
}

type binaryOperator int
type monaryOperator int

const (
	opAnd binaryOperator = iota
	opOr
	opLShift
	opRShift
)

const (
	opNot monaryOperator = iota
)

type binaryExpression struct {
	arg1, arg2 evaluable
	op         binaryOperator
}

func (b binaryExpression) Ready(context map[string]uint16) bool {
	return b.arg1.Ready(context) && b.arg2.Ready(context)
}

func (b binaryExpression) Value(context map[string]uint16) uint16 {
	arg1Val, arg2Val := b.arg1.Value(context), b.arg2.Value(context)

	switch b.op {
	case opAnd:
		return arg1Val & arg2Val
	case opOr:
		return arg1Val | arg2Val
	case opLShift:
		return arg1Val << arg2Val
	case opRShift:
		return arg1Val >> arg2Val
	default:
		panic(b.op)
	}
}

type monaryExpression struct {
	arg evaluable
	op  monaryOperator
}

func (m monaryExpression) Ready(context map[string]uint16) bool {
	return m.arg.Ready(context)
}

func (m monaryExpression) Value(context map[string]uint16) uint16 {
	argVal := m.arg.Value(context)

	switch m.op {
	case opNot:
		return ^argVal
	default:
		panic(m.op)
	}
}

type instruction struct {
	op     evaluable
	target string
}

var valueRegexp = regexp.MustCompile(`^[a-z]+$`)

func parseValue(raw string) (evaluable, bool) {
	num, err := strconv.Atoi(raw)
	if err == nil {
		return constValue{value: uint16(num)}, true
	}

	if valueRegexp.MatchString(raw) {
		return referenceValue{name: raw}, true
	}

	return nil, false
}

var binOpRegexp = regexp.MustCompile(`(.+) (.+) (.+)`)

func parseBinExpr(raw string) (evaluable, bool) {
	matches := binOpRegexp.FindAllStringSubmatch(raw, -1)
	if len(matches) != 1 {
		return nil, false
	}

	left, op, right := matches[0][1], matches[0][2], matches[0][3]

	var parsedOp binaryOperator

	switch op {
	case "AND":
		parsedOp = opAnd
	case "OR":
		parsedOp = opOr
	case "LSHIFT":
		parsedOp = opLShift
	case "RSHIFT":
		parsedOp = opRShift
	default:
		return nil, false
	}

	v1, ok1 := parseValue(left)
	v2, ok2 := parseValue(right)

	return binaryExpression{
		arg1: v1,
		arg2: v2,
		op:   parsedOp,
	}, ok1 && ok2
}

func parseMonExpr(raw string) (evaluable, bool) {
	if strings.Index(raw, "NOT ") == 0 {
		val, ok := parseValue(raw[len("NOT "):])
		if !ok {
			panic(raw)
		}

		return monaryExpression{
			arg: val,
			op:  opNot,
		}, true
	}

	return nil, false
}

func parseExpression(raw string) evaluable {
	binExpr, ok := parseBinExpr(raw)
	if ok {
		return binExpr
	}

	monExpr, ok := parseMonExpr(raw)
	if ok {
		return monExpr
	}

	value, ok := parseValue(raw)
	if ok {
		return value
	}

	panic(raw)
}

func parse(line string) instruction {
	before, target, ok := strings.Cut(line, " -> ")
	if !ok {
		panic(line)
	}

	return instruction{
		op:     parseExpression(before),
		target: target,
	}
}

func run(instructions []instruction, context map[string]uint16) uint16 {
	// note: the instruction dependencies form a DAG, the most efficient solution is to do a topological sort
	// and then evaluate instructions in that order. This is not that.
	for {
		if aVal, ok := context["a"]; ok {
			return aVal
		}

		var notReady []instruction

		for _, instr := range instructions {
			if !instr.op.Ready(context) {
				notReady = append(notReady, instr)
				continue
			}

			if _, ok := context[instr.target]; ok {
				continue
			}

			context[instr.target] = instr.op.Value(context)
		}

		instructions = notReady
	}
}

func solve(instructions []instruction) uint16 {
	aVal := run(instructions, map[string]uint16{})

	return run(instructions, map[string]uint16{"b": aVal})
}

func main() {
	solver := lib.Solver[[]instruction, uint16]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Verify(input(), 2797)
}
