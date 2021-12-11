package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
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
	isOpen bool
	score int
}

func (c character) Score() int {
	if c.isOpen {
		panic("only closing characters have scores")
	}

	return c.score
}

var characters = []character{
	{
		value: '(',
		match: ')',
		isOpen: true,
		score: 1,
	},
	{
		value: '[',
		match: ']',
		isOpen: true,
		score: 2,
	},
	{
		value: '{',
		match: '}',
		isOpen: true,
		score: 3,
	},
	{
		value: '<',
		match: '>',
		isOpen: true,
		score: 4,
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

	return chunks
}

func solve(r io.Reader) {
	chunks := parse(r)

	var scores []int

	for _, chunk := range chunks {
		rest, ok := finish(chunk)
		if ok {
			scores = append(scores, calculateScore(rest))
		}
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	fmt.Println(scores[len(scores)/2])
}

func calculateScore(rest []character) int {
	score := 0
	for _, c := range rest {
		score = score * 5 + c.Score()
	}

	return score
}

// finish completes the chunk. It returns false if the chunk is illegal.
func finish(chunk []character) ([]character, bool) {
	var s stack

	for _, c := range chunk {
		switch c.isOpen {
		case true:
			s.push(c)
		case false:
			if s.pop().value != c.match {
				return nil, false
			}
		}
	}

	var rest []character
	for !s.empty() {
		rest = append(rest, getCharacterFromRune(s.pop().match))
	}
	return rest, true
}

func main() {
	fillMap()

	solve(strings.NewReader("[({(<(())[]>[[{[]{<()<>>\n[(()[<>])]({[<{<<[]>>(\n{([(<{}[<>[]}>{[]{[(<()>\n(((({<>}<{<{<>}{[]{[]{}\n[[<[([]))<([[{}[[()]]]\n[{[{({}]{}}([{[{{{}}([]\n{<[[]]>}<{[{[{[]{()[[[]\n[<(<(<(<{}))><([]([]()\n<{([([[(<>()){}]>(<<{{\n<{([{{}}[<[[[<>{}]]]>[]]"))
	solve(input())
}
