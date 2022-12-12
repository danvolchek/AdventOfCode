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
	// ParseF is a parse function that accepts a string and returns the parsed value.
	// See below for common parse functions.
	ParseF func(input string) T

	// SolveF is the function which solves the puzzle using the parsed input.
	SolveF func(parsed T) V

	expectsRun, expectsCorrect int

	client *aocClient
}

// Parse parses input and prints the result.
func (s *Solver[T, V]) Parse(input string) {
	actual := s.ParseF(input)

	fmt.Printf("parse: \"%v\" -> %+v\n", formatInput(input), actual)
}

// ParseExpect parses input, compares it to expected, and prints the result.
func (s *Solver[T, V]) ParseExpect(input string, expected T) {
	actual := s.ParseF(input)

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)    parse: \"%v\" -> expected %+v, got %+v\n", formatInput(input), expected, actual)
	} else {
		fmt.Printf("(success) parse: \"%v\" -> got %+v\n", formatInput(input), actual)
	}
}

// Expect runs the solution against input, compares it to expected, and prints the result.
func (s *Solver[T, V]) Expect(input string, expected V) {
	actual, dur := s.solve(input)

	s.expectsRun += 1

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)     test: \"%v\" -> expected %v, got %v%v\n", formatInput(input), expected, actual, dur)
	} else {
		s.expectsCorrect += 1
		fmt.Printf("(success)  test: \"%v\" -> got %v%v\n", formatInput(input), actual, dur)
	}
}

// Test runs the solution against input and prints the result.
func (s *Solver[T, V]) Test(input string) {
	solution, dur := s.solve(input)
	fmt.Printf("test: \"%v\" -> %v%v\n", formatInput(input), solution, dur)
}

// Verify runs the solution against the real input, compares it to expected, and prints the result.
func (s *Solver[T, V]) Verify(expected V) {
	input, ok := s.getRealInput()
	if !ok {
		fmt.Printf("(fail)     real: empty input file\n")
		return
	}

	actual, dur := s.solve(input)

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)     real: expected %v, got %v%v\n", expected, actual, dur)
	} else {
		fmt.Printf("(success)  real: got %v%v\n", actual, dur)
	}
}

// Solve runs the solution against the real input and prints the result.
func (s *Solver[T, V]) Solve() {
	input, ok := s.getRealInput()
	if !ok {
		fmt.Printf("(fail)     real: empty input file\n")
		return
	}

	solution, dur := s.solve(input)
	fmt.Printf("real: %v%v\n", solution, dur)

	if s.client.sessionCookieErr != nil {
		fmt.Printf("note: can't submit solution: %s\n", s.client.sessionCookieErr)
		return
	}

	if s.shouldSubmit() {
		fmt.Println("submitting...")
		output, err := s.client.submitSolution(fmt.Sprint(solution))
		if err != nil {
			fmt.Printf("note: failed to submit solution: %s", err)
		} else {
			fmt.Printf("output:\n%s\n", output)
		}
	}
}

// shouldSubmit returns whether the solution should be automatically submitted.
func (s *Solver[T, V]) shouldSubmit() bool {
	// auto submit if all expects passed
	if s.expectsRun > 0 && s.expectsCorrect == s.expectsRun {
		fmt.Println("note: auto submitting because all expects passed")
		return true
	}

	// don't submit if some expects failed
	if s.expectsRun != 0 {
		fmt.Println("note: not auto submitting because some expects failed")
		return false
	}

	// otherwise ask user
	fmt.Println("submit? y/n")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("note: failed to read response: %s\n", err)
		return false
	}

	return strings.ToLower(strings.TrimSpace(text)) == "y"
}

// solve runs the solution against input and returns the result and elapsed time.
func (s *Solver[T, V]) solve(input string) (V, formatDur) {
	now := time.Now()

	solution := s.SolveF(s.ParseF(input))

	return solution, formatDur(time.Now().Sub(now))
}

// getRealInput returns a reader which reads the input file.
func (s *Solver[T, V]) getRealInput() (string, bool) {
	year, day, part := getSolutionMetadata()

	if s.client == nil {
		s.client = newAocClient(year, day, part)
	}

	fileBytes := readInputFile(year, day)

	if len(fileBytes) != 0 {
		return string(fileBytes), true
	}

	fmt.Println("retrieving input because it's not found...")

	input, err := s.client.retrieveInput()
	if err != nil {
		fmt.Printf("note: tried to retrieve input, failed: %s\n", err)
		return "", false
	}

	writeInputFile(year, day, input)

	return string(input), true
}

// readInputFile returns the contents of the input file for year and day.
func readInputFile(year, day string) []byte {
	file := Must(os.Open(path.Join(year, day, "input.txt")))

	return Must(io.ReadAll(file))
}

// writeInputFile writes data to the input file for year and day.
func writeInputFile(year, day string, data []byte) {
	file, err := os.Create(path.Join(year, day, "input.txt"))
	if err != nil {
		fmt.Printf("note: tried to create input file, failed: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("note: tried to write input to file, failed: %s\n", err)
		return
	}
}

// getSolutionMetadata returns the year, day, and part the current execution is for.
func getSolutionMetadata() (string, string, string) {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		panic(ok)
	}

	// path is of the form github.com/danvolchek/AdventOfCode/2015/16/leaderboard/1
	parts := strings.Split(buildInfo.Path, string(os.PathSeparator))

	return parts[3], parts[4], parts[6]
}

type formatInput string

func (f formatInput) String() string {
	return strings.ReplaceAll(string(f), "\n", "\n"+strings.Repeat(" ", 17))
}

type formatDur time.Duration

func (f formatDur) String() string {
	return " (" + time.Duration(f).String() + ")"
}

// ParseBytes is a parse function that returns the raw bytes read.
func ParseBytes(input string) []byte {
	return []byte(input)
}

// ParseBytesFunc is a parse function that returns a parsed value of the raw bytes read.
func ParseBytesFunc[T any](parse func(input []byte) T) func(input string) T {
	return func(input string) T {
		return parse(ParseBytes(input))
	}
}

// ParseGrid parses a grid of characters.
func ParseGrid[T any](parse func(s string) T) func(input string) [][]T {
	return func(input string) [][]T {
		rawGrid := strings.Split(strings.TrimSpace(input), "\n")

		var grid [][]T

		for i, row := range rawGrid {
			grid = append(grid, make([]T, len(row)))
			for j := range row {
				grid[i][j] = parse(string(row[j]))
			}
		}

		return grid
	}
}

// ParseLine is a parse function that splits parsing into one line at a time, returning a slice of items.
// It accepts a parse function to parse each line seen.
func ParseLine[T any](parse func(line string) T) func(input string) []T {
	return func(input string) []T {
		rawLines := strings.Split(strings.TrimSpace(input), "\n")

		return Map(rawLines, parse)
	}
}

// ParseChunks is like ParseLine but parses lines delimited by two new lines, not one
func ParseChunks[T any](parse func(chunk string) T) func(input string) []T {
	return func(input string) []T {
		rawChunks := strings.Split(strings.TrimSpace(input), "\n\n")

		return Map(rawChunks, parse)
	}
}

// ParseChunksUnique is like ParseChunks but uses a unique parser for every chunk.
// It's different from all the other functions in that it passes in a T to the parse func to modify as needed.
// The intended use is for each parser to be able to return a different type safe type - this couldn't work with generics -
// so T is intended a container for the result of each parser.
func ParseChunksUnique[T any](parsers ...func(chunk string, val *T)) func(input string) T {
	return func(input string) T {
		var start T

		rawChunks := strings.Split(input, "\n\n")

		for i, chunk := range rawChunks {
			parsers[i](chunk, &start)
		}

		return start
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
