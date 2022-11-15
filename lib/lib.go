package lib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Must return value if err is non-nil and panics otherwise.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

// Min returns the smallest value in values.
func Min(values ...int) int {
	min := 0

	for i, value := range values {
		if i == 0 || value < min {
			min = value
		}
	}

	return min
}

// SolveBytes runs solution using the raw result of reading r.
func SolveBytes[T any](r io.Reader, solution func(input []byte) T) {
	input := Must(io.ReadAll(r))
	result := solution(input)
	fmt.Println(result)
}

// TestSolveBytes is SolveBytes with a string input (for test cases).
func TestSolveBytes[T any](input string, solution func(input []byte) T) {
	SolveBytes(strings.NewReader(input), solution)
}

// SolveLines runs solution interpreting r as a newline delimited list of strings.
func SolveLines[T any](r io.Reader, solution func(lines []string) T) {
	var lines []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	result := solution(lines)
	fmt.Println(result)
}

// TestSolveLines is SolveLines with a string input (for test cases).
func TestSolveLines[T any](input string, solution func(lines []string) T) {
	SolveLines(strings.NewReader(input), solution)
}
