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
	inner expression
}

func (g grouping) evaluate() int {
	return g.inner.evaluate()
}

const (
	tokenConstant = iota
	tokenPlus
	tokenMul
	tokenLeftParenthesis
	tokenRightParenthesis
)

type token struct {
	tokenType int
	value     int
}

func tokenize(raw string) []token {
	var tokens []token

	for i := 0; i < len(raw); i++ {
		curr := raw[i]

		switch curr {
		case ' ':
			continue
		case '+':
			tokens = append(tokens, token{
				tokenType: tokenPlus,
			})
			continue
		case '*':
			tokens = append(tokens, token{
				tokenType: tokenMul,
			})
			continue
		case '(':
			tokens = append(tokens, token{
				tokenType: tokenLeftParenthesis,
			})
			continue
		case ')':
			tokens = append(tokens, token{
				tokenType: tokenRightParenthesis,
			})
			continue
		}

		if curr >= '0' && curr <= '9' {
			end := i + 1

			for end < len(raw) && raw[end] >= '0' && raw[end] <= '9' {
				end++
			}

			intVal, err := strconv.Atoi(raw[i:end])
			if err != nil {
				panic(err)
			}

			tokens = append(tokens, token{
				tokenType: tokenConstant,
				value:     intVal,
			})
			continue
		}

		panic(curr)
	}

	return tokens
}

func parse(tokens []token) expression {
	if len(tokens) == 1 {
		if tokens[0].tokenType != tokenConstant {
			panic(tokens)
		}

		return constant{value: tokens[0].value}
	}

	var firstExpr expression
	secondTokenIndex := 1

	if tokens[0].tokenType == tokenRightParenthesis {
		end := 1

		nestingLevel := 1
		for nestingLevel != 0 {
			switch tokens[end].tokenType {
			case tokenRightParenthesis:
				nestingLevel += 1
			case tokenLeftParenthesis:
				nestingLevel -= 1
			}

			end += 1

			if end >= len(tokens) && nestingLevel != 0 {
				panic("unbalanced parens")
			}
		}

		firstExpr = grouping{
			inner: parse(tokens[1 : end-1]),
		}
		secondTokenIndex = end

		if end == len(tokens) {
			return firstExpr
		}

		//fmt.Println("Parsed group", firstExpr)
	} else {
		firstExpr = parse(tokens[:1])
	}
	//fmt.Println("Parsed first expr as", firstExpr)

	switch tokens[secondTokenIndex].tokenType {
	case tokenPlus:
		return binary{
			operation: "+",
			left:      firstExpr,
			right:     parse(tokens[secondTokenIndex+1:]),
		}
	case tokenMul:
		return binary{
			operation: "*",
			left:      firstExpr,
			right:     parse(tokens[secondTokenIndex+1:]),
		}
	default:
		fmt.Println("don't know what to do")
		panic(tokens[1])
	}

}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	sum := 0
	for scanner.Scan() {
		row := scanner.Text()

		tokens := tokenize(row)

		newTokens := make([]token, len(tokens))

		for i := 0; i < len(tokens); i++ {
			newTokens[i] = tokens[len(tokens)-i-1]
		}

		expr := parse(newTokens)

		fmt.Println(expr, expr.evaluate())
		sum += expr.evaluate()
	}

	fmt.Println("SUM", sum)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	solve(strings.NewReader("2 * 3 + (4 * 5)"))
	solve(strings.NewReader("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	solve(strings.NewReader("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	solve(strings.NewReader("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 "))
	solve(input())
}
