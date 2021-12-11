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

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	score := 0
	for scanner.Scan() {
		line := scanner.Text()

		if c, ok := isIllegal(line); !ok {
			fmt.Printf("%s %c\n", line, c)
			switch c {
			case ')':
				score += 3
			case ']':
				score += 57
			case '}':
				score += 1197
			case '>':
				score += 25137
			}
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(score)
}

func isIllegal(chunk string) (byte, bool) {
	s := newStack(nil)

	handleOpen := func(c byte) {
		s.push(c)
		//fmt.Println(s)
	}
	handleClose := func(c byte) bool {
		//fmt.Println(s)
		v := s.pop()
		switch v {
		case '[':
			return c == ']'
		case '{':
			return c == '}'
		case '<':
			return c == '>'
		case '(':
			return c == ')'
		default:
			panic("bad char")
		}
	}

	for i := 0; i < len(chunk); i++ {
		c := chunk[i]

		switch c {
		case '[':
			fallthrough
		case '{':
			fallthrough
		case '<':
			fallthrough
		case '(':
			handleOpen(c)
			continue
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			fallthrough
		case ')':
			ok := handleClose(c)
			if !ok {
				return c, false
			}
		default:
			panic("bad char")
		}
	}

	return 0, true
}

func main() {
	solve(strings.NewReader("[({(<(())[]>[[{[]{<()<>>\n[(()[<>])]({[<{<<[]>>(\n{([(<{}[<>[]}>{[]{[(<()>\n(((({<>}<{<{<>}{[]{[]{}\n[[<[([]))<([[{}[[()]]]\n[{[{({}]{}}([{[{{{}}([]\n{<[[]]>}<{[{[{[]{()[[[]\n[<(<(<(<{}))><([]([]()\n<{([([[(<>()){}]>(<<{{\n<{([{{}}[<[[[<>{}]]]>[]]"))
	solve(input())
}
