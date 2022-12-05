package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime/debug"
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

// Parse parses input and prints the result.
func (s Solver[T, V]) Parse(input string) {
	actual := s.ParseF(strings.NewReader(input))

	fmt.Printf("parse: \"%v\" -> %+v\n", formatInput(input), actual)
}

// ParseExpect parses input, compares it to expected, and prints the result.
func (s Solver[T, V]) ParseExpect(input string, expected T) {
	actual := s.ParseF(strings.NewReader(input))

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)    parse: \"%v\" -> expected %+v, got %+v\n", formatInput(input), expected, actual)
	} else {
		fmt.Printf("(success) parse: \"%v\" -> got %+v\n", formatInput(input), actual)
	}
}

// Expect runs the solution against input, compares it to expected, and prints the result.
func (s Solver[T, V]) Expect(input string, expected V) {
	actual, dur := s.solve(strings.NewReader(input))

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)     test: \"%v\" -> expected %v, got %v%v\n", formatInput(input), expected, actual, dur)
	} else {
		fmt.Printf("(success)  test: \"%v\" -> got %v%v\n", formatInput(input), actual, dur)
	}
}

// Test runs the solution against input and prints the result.
func (s Solver[T, V]) Test(input string) {
	solution, dur := s.solve(strings.NewReader(input))
	fmt.Printf("test: \"%v\" -> %v%v\n", formatInput(input), solution, dur)
}

// Verify runs the solution against the real input, compares it to expected, and prints the result.
func (s Solver[T, V]) Verify(expected V) {
	actual, dur := s.solve(getRealInput())

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)     real: expected %v, got %v%v\n", expected, actual, dur)
	} else {
		fmt.Printf("(success)  real: got %v%v\n", actual, dur)
	}
}

// Solve runs the solution against the real input and prints the result.
func (s Solver[T, V]) Solve() {
	solution, dur := s.solve(getRealInput())
	fmt.Printf("real: %v%v\n", solution, dur)
}

// solve runs the solution against input and returns the result and elapsed time.
func (s Solver[T, V]) solve(input io.Reader) (V, formatDur) {
	now := time.Now()

	solution := s.SolveF(s.ParseF(input))

	return solution, formatDur(time.Now().Sub(now))
}

// getRealInput returns a reader which reads the input file.
func getRealInput() io.Reader {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		panic(ok)
	}

	// path is of the form github.com/danvolchek/AdventOfCode/2015/16/leaderboard/1
	parts := strings.Split(buildInfo.Path, string(os.PathSeparator))

	return Must(os.Open(path.Join(parts[3], parts[4], "input.txt")))
}

type formatInput string

func (f formatInput) String() string {
	return strings.ReplaceAll(string(f), "\n", "\n"+strings.Repeat(" ", 17))
}

type formatDur time.Duration

func (f formatDur) String() string {
	return " (" + time.Duration(f).String() + ")"
}

// ParseBytes is a top level parse function that returns the raw bytes read.
func ParseBytes(input io.Reader) []byte {
	return Must(io.ReadAll(input))
}

// ParseBytesFunc is a top level parse function that returns a parsed value of the raw bytes read.
func ParseBytesFunc[T any](parse func(input []byte) T) func(r io.Reader) T {
	return func(r io.Reader) T {
		return parse(ParseBytes(r))
	}
}

// ParseString is a top level parse function that returns the string representation of the raw bytes read.
func ParseString(input io.Reader) string {
	return string(Must(io.ReadAll(input)))
}

// ParseStringFunc is a top level parse function that returns a parsed value of the raw bytes read.
func ParseStringFunc[T any](parse func(input string) T) func(r io.Reader) T {
	return func(r io.Reader) T {
		return parse(ParseString(r))
	}
}

// ParseStringChunks is a top level parse function that uses a different parse function for every chunk of the input.
func ParseStringChunks[T any](parsers ...func(chunk string, val *T)) func(r io.Reader) T {
	return func(r io.Reader) T {
		var start T

		raw := ParseString(r)

		chunks := strings.Split(raw, "\n\n")

		for i, chunk := range chunks {
			parsers[i](chunk, &start)
		}

		return start
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

// ParseLineChunked is like ParseLine except when it sees a blank line, it parses all the lines seen previously as a single chunk.
func ParseLineChunked[T, V any](parse func(line string) T, parseChunk func(lines []T) V) func(r io.Reader) []V {
	return func(r io.Reader) []V {
		var chunks []V
		var lines []T

		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				chunks = append(chunks, parseChunk(lines))
				lines = nil
			} else {
				lines = append(lines, parse(line))
			}
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}

		chunks = append(chunks, parseChunk(lines))

		return chunks
	}
}

// AsIs is a parse function helper that leaves the value as is. Useful with ParseLine.
func AsIs(line string) string {
	return line
}

// ToByteSlice is a parse function helper that turns the value to a byte slice. Useful with ParseLine.
func ToByteSlice(line string) []byte {
	return []byte(line)
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

// ToString is a solve function helper that converts a byte slice to a string.
func ToString[T any](solve func(T) []byte) func(T) string {
	return func(input T) string {
		return string(solve(input))
	}
}

var intsReg = regexp.MustCompile(`\d+`)

// Ints returns all the positive integers in line.
func Ints(line string) []int {
	numbers := intsReg.FindAllString(line, -1)

	result := make([]int, len(numbers))

	for i, number := range numbers {
		result[i] = Atoi(number)
	}

	return result
}
