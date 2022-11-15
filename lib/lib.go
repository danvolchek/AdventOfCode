package lib

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"regexp"
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

// Solver is a wrapper around running a solution, providing helper methods to simplify boilerplate.
type Solver[T, V any] struct {
	ParseF func(input io.Reader) T
	SolveF func(parsed T) V
}

// Expect runs the solution against input, compares it to expected, and prints the result.
func (s Solver[T, V]) Expect(input string, expected V) {
	actual := s.solve(strings.NewReader(input))

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)    test: \"%v\" -> expected %v, got %v\n", input, expected, actual)
	} else {
		fmt.Printf("(success) test: \"%v\" -> got %v\n", input, actual)
	}
}

// Test runs the solution against input and prints the result.
func (s Solver[T, V]) Test(input string) {
	fmt.Printf("test: \"%v\" -> %v\n", input, s.solve(strings.NewReader(input)))
}

// Solve runs the solution against the real input and prints the result.
func (s Solver[T, V]) Solve(input io.Reader) {
	fmt.Printf("real: %v\n", s.solve(input))
}

// solve runs the solution against input and returns the result.
func (s Solver[T, V]) solve(input io.Reader) V {
	return s.SolveF(s.ParseF(input))
}

// ParseBytes is a parse function helper that returns the raw bytes read.
func ParseBytes() func(input io.Reader) []byte {
	return func(input io.Reader) []byte {
		return Must(io.ReadAll(input))
	}
}

// ParseLine is a parse function helper that splits parsing into one line at a time, returning a slice of items.
func ParseLine[T any](parse func(line string) T) func(r io.Reader) []T {
	return func(r io.Reader) []T {
		var lines []T

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			lines = append(lines, parse(scanner.Text()))
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		return lines
	}
}

// AsIs is a parse function helper that leaves the value as is.
func AsIs[T any]() func(value T) T {
	return func(line T) T {
		return line
	}
}

// ParseRegexp is a parse function helper that returns substring matches from a string. Useful with ParseLine.
func ParseRegexp[T any](reg *regexp.Regexp, parse func(parts []string) T) func(line string) T {
	return func(line string) T {
		result := reg.FindAllStringSubmatch(line, -1)

		parts := result[0][1:]

		return parse(parts)
	}
}
