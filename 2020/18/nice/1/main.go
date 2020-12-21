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
	input, err := os.Open(path.Join("2020", "18", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func read(r io.Reader) []string{
	var lines []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return lines
}

type expression interface {
	evaluate() int
}

type constant struct {
	value int
}

func (c constant) evaluate() int {
	return c.value
}

type binary struct {
	left      expression
	operation string
	right     expression
}

func (b binary) evaluate() int {
	switch b.operation {
	case "*":
		return b.left.evaluate() * b.right.evaluate()
	case "+":
		return b.left.evaluate() + b.right.evaluate()
	default:
		panic(b.operation)
	}
}

type grouping struct {
	expr expression
}

func (g grouping) evaluate() int {
	return g.expr.evaluate()
}

type tokenType int

const (
	tokenConstant tokenType = iota
	tokenPlus
	tokenMultiply
	tokenLeftParenthesis
	tokenRightParenthesis
)

type token struct {
	which tokenType
	value int
}

func tokenize(raw string) []token {
	var tokens []token

	for i := 0; i < len(raw); i++ {
		switch raw[i] {
		case ' ':
			continue
		case '+':
			tokens = append(tokens, token{
				which: tokenPlus,
			})
			continue
		case '*':
			tokens = append(tokens, token{
				which: tokenMultiply,
			})
			continue
		case '(':
			tokens = append(tokens, token{
				which: tokenLeftParenthesis,
			})
			continue
		case ')':
			tokens = append(tokens, token{
				which: tokenRightParenthesis,
			})
			continue
		}

		isDigit := func(b byte) bool { return b >= '0' && b <= '9'}

		if !isDigit(raw[i]) {
			panic(raw[i])
		}

		value, err := strconv.Atoi(string(raw[i]))
		if err != nil {
			panic(err)
		}

		tokens = append(tokens, token{
			which: tokenConstant,
			value: value,
		})
	}

	return tokens
}

type tokenStack struct {
	tokens []token
}

func (t *tokenStack) Push(token token) {
	t.tokens = append(t.tokens, token)
}

func (t *tokenStack) Pop() token {
	lastIndex := len(t.tokens) - 1
	token := t.tokens[lastIndex]
	t.tokens = t.tokens[:lastIndex]

	return token
}

func parse(tokens []token) expression {
	//tokenQueue := make(chan token, 10)
	operatorStack := tokenStack{}

	for _, token := range tokens {
		switch token.which {
		case tokenPlus:
			operatorStack.Push(token)
		case tokenMultiply:
			operatorStack.Push(token)
		case tokenConstant:
		default:

		}
	}

	return nil
}

func solve(equations []string) int {
	sum := 0

	for _, equation := range equations {
		tokens := tokenize(equation)

		newTokens := make([]token, len(tokens))

		for i := 0; i < len(tokens); i++ {
			newTokens[i] = tokens[len(tokens)-i-1]
		}

		expr := parse(newTokens)

		fmt.Println(expr, expr.evaluate())
		sum += expr.evaluate()
	}

	return sum
}

func main() {
	fmt.Println(solve(read(strings.NewReader("2 * 3 + (4 * 5)"))))
	fmt.Println(solve(read(strings.NewReader("5 + (8 * 3 + 9 + 3 * 4 * 3)"))))
	fmt.Println(solve(read(strings.NewReader("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))))
	fmt.Println(solve(read(strings.NewReader("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 "))))
	fmt.Println(solve(read(input())))
}
