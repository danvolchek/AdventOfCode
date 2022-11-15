package lib

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
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

// Atoi is a convenience wrapper on [strconv.Atoi] that panics if it fails.
func Atoi(s string) int {
	return Must(strconv.Atoi(s))
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
	SolveParseLines(r, func(line string) string { return line }, solution)
}

// TestSolveLines is SolveLines with a string input (for test cases).
func TestSolveLines[T any](input string, solution func(lines []string) T) {
	SolveLines(strings.NewReader(input), solution)
}

// SolveParseLines runs solution interpreting r as a newline delimited list of strings, parsed according to a function.
func SolveParseLines[T, K any](r io.Reader, parse func(line string) K, solution func(lines []K) T) {
	var lines []K

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, parse(scanner.Text()))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	result := solution(lines)
	fmt.Println(result)
}

// TestSolveParseLines is SolveParseLines with a string input (for test cases).
func TestSolveParseLines[T, K any](input string, parse func(line string) K, solution func(lines []K) T) {
	SolveParseLines(strings.NewReader(input), parse, solution)
}