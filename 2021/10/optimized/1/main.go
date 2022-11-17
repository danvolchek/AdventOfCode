package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "10", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type character struct {
	value, match rune
	isOpen       bool
	score        int
}

func (c character) Score() int {
	if c.isOpen {
		panic("only closing characters have scores")
	}

	return c.score
}

var characters = []character{
	{
		value:  '(',
		match:  ')',
		isOpen: true,
		score:  3,
	},
	{
		value:  '[',
		match:  ']',
		isOpen: true,
		score:  57,
	},
	{
		value:  '{',
		match:  '}',
		isOpen: true,
		score:  1197,
	},
	{
		value:  '<',
		match:  '>',
		isOpen: true,
		score:  25137,
	},
}

var characterMap map[rune]character

func fillMap() {
	characterMap = make(map[rune]character)
	for _, c := range characters {
		_, hasVal := characterMap[c.value]
		_, hasMatch := characterMap[c.match]
		if hasVal || hasMatch {
			panic("bad character slice")
		}

		characterMap[c.value] = c
		characterMap[c.match] = character{
			value:  c.match,
			match:  c.value,
			isOpen: !c.isOpen,
			score:  c.score,
		}
	}
}

func getCharacterFromRune(r rune) character {
	c, ok := characterMap[r]
	if !ok {
		panic("rune not found")
	}

	return c
}

func parse(r io.Reader) [][]character {
	var chunks [][]character

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		var chunk []character

		for _, r := range line {
			chunk = append(chunk, getCharacterFromRune(r))
		}

		chunks = append(chunks, chunk)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return chunks
}

func solve(r io.Reader) {
	chunks := parse(r)

	score := 0

	for _, chunk := range chunks {
		c, ok := firstIllegalCharacter(chunk)
		if ok {
			score += c.Score()
		}
	}

	fmt.Println(score)
}

func firstIllegalCharacter(chunk []character) (character, bool) {
	var s stack

	for _, c := range chunk {
		switch c.isOpen {
		case true:
			s.push(c.value)
		case false:
			if s.pop() != c.match {
				return c, true
			}
		}
	}

	return character{}, false
}

func main() {
	fillMap()

	solve(strings.NewReader("[({(<(())[]>[[{[]{<()<>>\n[(()[<>])]({[<{<<[]>>(\n{([(<{}[<>[]}>{[]{[(<()>\n(((({<>}<{<{<>}{[]{[]{}\n[[<[([]))<([[{}[[()]]]\n[{[{({}]{}}([{[{{{}}([]\n{<[[]]>}<{[{[{[]{()[[[]\n[<(<(<(<{}))><([]([]()\n<{([([[(<>()){}]>(<<{{\n<{([{{}}[<[[[<>{}]]]>[]]"))
	solve(input())
}
