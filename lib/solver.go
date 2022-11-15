package lib

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Solver is a wrapper around running a solution, providing helper methods to simplify boilerplate.
type Solver[T, V any] struct {
	// ParseF is a top level parse function that accepts an [io.Reader] and returns the parsed value.
	// See below for common top level parse functions.
	ParseF func(input io.Reader) T

	// SolveF is the function which solves the puzzle using the parsed input.
	SolveF func(parsed T) V
}

// Expect runs the solution against input, compares it to expected, and prints the result.
func (s Solver[T, V]) Expect(input string, expected V) {
	actual, dur := s.solve(strings.NewReader(input))

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)    test: \"%v\" -> expected %v, got %v%v\n", input, expected, actual, dur)
	} else {
		fmt.Printf("(success) test: \"%v\" -> got %v%v\n", input, actual, dur)
	}
}

// Test runs the solution against input and prints the result.
func (s Solver[T, V]) Test(input string) {
	solution, dur := s.solve(strings.NewReader(input))
	fmt.Printf("test: \"%v\" -> %v%v\n", input, solution, dur)
}

// Verify runs the solution against the real input, compares it to expected, and prints the result.
func (s Solver[T, V]) Verify(input io.Reader, expected V) {
	actual, dur := s.solve(input)

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)    real: expected %v, got %v%v\n", expected, actual, dur)
	} else {
		fmt.Printf("(success) real: got %v%v\n", actual, dur)
	}
}

// Solve runs the solution against the real input and prints the result.
func (s Solver[T, V]) Solve(input io.Reader) {
	solution, dur := s.solve(input)
	fmt.Printf("real: %v%v\n", solution, dur)
}

// solve runs the solution against input and returns the result and elapsed time.
func (s Solver[T, V]) solve(input io.Reader) (V, formatDur) {
	now := time.Now()

	solution := s.SolveF(s.ParseF(input))

	return solution, formatDur{dur: time.Now().Sub(now)}
}

type formatDur struct {
	dur time.Duration
}

func (f formatDur) String() string {
	return " (" + f.dur.String() + ")"
}

// ParseBytes is a top level parse function that returns the raw bytes read.
func ParseBytes() func(input io.Reader) []byte {
	return func(input io.Reader) []byte {
		return Must(io.ReadAll(input))
	}
}

// ParseLine is a top level function helper that splits parsing into one line at a time, returning a slice of items.
// It accepts a parse function to parse each line seen.
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

// AsIs is a parse function helper that leaves the value as is. Useful with ParseLine.
func AsIs[T any]() func(value T) T {
	return func(line T) T {
		return line
	}
}

// ParseRegexp is a parse function helper that returns substring matches from a string. Useful with ParseLine.
// There should only be one match of reg in the string; others will be ignored.
func ParseRegexp[T any](reg *regexp.Regexp, parse func(parts []string) T) func(line string) T {
	return func(line string) T {
		matches := reg.FindAllStringSubmatch(line, -1)

		firstMatchSubmatches := matches[0][1:]

		return parse(firstMatchSubmatches)
	}
}