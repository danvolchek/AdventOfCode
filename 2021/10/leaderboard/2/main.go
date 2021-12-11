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

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var scores []int
	for scanner.Scan() {
		line := scanner.Text()

		if rest, ok := finish(line); ok {
			score := 0
			for i := 0; i < len(rest); i++ {
				score *= 5

				switch rest[i] {
				case ')':
					score += 1
				case ']':
					score += 2
				case '}':
					score += 3
				case '>':
					score += 4
				}
			}

			scores = append(scores, score)

		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	fmt.Println(scores[len(scores)/2])
}

func finish(chunk string) ([]byte, bool) {
	s := newStack(nil)

	handleOpen := func(c byte) {
		s.push(c)
		//fmt.Println(s)
	}
	want := func(c byte) byte {
		switch c {
		case '[':
			return ']'
		case '{':
			return '}'
		case '<':
			return '>'
		case '(':
			return ')'
		default:
			panic("bad char")
		}
	}
	handleClose := func(c byte) bool {
		//fmt.Println(s)
		v := s.pop()
		return c == want(v)
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
				return nil, false
			}
		default:
			panic("bad char")
		}
	}

	var finish []byte
	for !s.Empty() {
		v := s.pop()
		finish = append(finish, want(v))
	}

	return finish, true
}

func main() {
	solve(strings.NewReader("[({(<(())[]>[[{[]{<()<>>\n[(()[<>])]({[<{<<[]>>(\n{([(<{}[<>[]}>{[]{[(<()>\n(((({<>}<{<{<>}{[]{[]{}\n[[<[([]))<([[{}[[()]]]\n[{[{({}]{}}([{[{{{}}([]\n{<[[]]>}<{[{[{[]{()[[[]\n[<(<(<(<{}))><([]([]()\n<{([([[(<>()){}]>(<<{{\n<{([{{}}[<[[[<>{}]]]>[]]"))
	solve(input())
}
