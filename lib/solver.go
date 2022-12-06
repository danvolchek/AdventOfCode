package lib

import (
	"bufio"
	"bytes"
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

	expectsRun, expectsCorrect int
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
func (s *Solver[T, V]) Expect(input string, expected V) {
	actual, dur := s.solve(strings.NewReader(input))

	s.expectsRun += 1

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("(fail)     test: \"%v\" -> expected %v, got %v%v\n", formatInput(input), expected, actual, dur)
	} else {
		s.expectsCorrect += 1
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
	input, ok := getRealInput()
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
func (s Solver[T, V]) Solve() {
	input, ok := getRealInput()
	if !ok {
		fmt.Printf("(fail)     real: empty input file\n")
		return
	}

	solution, dur := s.solve(input)
	fmt.Printf("real: %v%v\n", solution, dur)

	if client.sessionCookieErr != nil {
		fmt.Printf("note: can't submit solution: %s\n", client.sessionCookieErr)
		return
	}

	if s.shouldSubmit() {
		fmt.Println("submitting...")
		output, err := client.submitSolution(fmt.Sprint(solution))
		if err != nil {
			fmt.Printf("note: failed to submit solution: %s", err)
		} else {
			fmt.Printf("output:\n%s\n", output)
		}
	}
}

// shouldSubmit returns whether the solution should be automatically submitted.
func (s Solver[T, V]) shouldSubmit() bool {
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
func (s Solver[T, V]) solve(input io.Reader) (V, formatDur) {
	now := time.Now()

	solution := s.SolveF(s.ParseF(input))

	return solution, formatDur(time.Now().Sub(now))
}

// getRealInput returns a reader which reads the input file.
func getRealInput() (io.Reader, bool) {
	year, day, _ := getSolutionMetadata()

	fileBytes := readInputFile(year, day)

	if len(fileBytes) != 0 {
		return bytes.NewReader(fileBytes), true
	}

	fmt.Println("retrieving input because it's not found...")

	input, err := client.retrieveInput()
	if err != nil {
		fmt.Printf("note: tried to retrieve input, failed: %s\n", err)
		return nil, false
	}

	writeInputFile(year, day, input)

	return bytes.NewReader(input), true
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
