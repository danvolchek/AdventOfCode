package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

var numberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type stream struct {
	pos    int
	source string
}

func (s *stream) done() bool {
	return s.pos >= len(s.source)
}

func (s *stream) peek() int {
	return int(s.source[s.pos])
}

func (s *stream) consume() {
	s.pos += 1
}

func (s *stream) match(val string) bool {
	if s.pos+len(val) <= len(s.source) && s.source[s.pos:s.pos+len(val)] == val {
		s.pos += len(val) - 1 // -1 so that 'eightwo' only consumes 'eigh' and allows matching 'two'
		return true
	}

	return false
}

func isDigit(val int) (int, bool) {
	if val < '0' || val > '9' {
		return -1, false
	}

	return val - '0', true
}

func parse(line string) int {
	var nums []int

	s := stream{source: line}

outer:
	for !s.done() {
		char := s.peek()
		if digit, ok := isDigit(char); ok {
			nums = append(nums, digit)
			s.consume()
			continue
		}

		for i, word := range numberWords {
			if s.match(word) {
				nums = append(nums, i+1)
				continue outer
			}
		}

		s.consume()
	}

	return nums[0]*10 + nums[len(nums)-1]
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: lib.SumSlice[int],
	}

	solver.Expect("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n", 281)
	solver.Verify(56324)
}
