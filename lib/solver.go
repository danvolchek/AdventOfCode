package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

	incorrectAnswers []V

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
	metadata := getSolutionMetadata()

	input, ok := s.getRealInput(metadata)
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
	metadata := getSolutionMetadata()

	input, ok := s.getRealInput(metadata)
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

	if s.shouldSubmit(solution) {
		fmt.Println("submitting...")
		output, err := s.client.submitSolution(fmt.Sprint(solution))
		if err != nil {
			fmt.Printf("note: failed to submit solution: %s", err)
		} else {
			fmt.Printf("output:\n%s\n", output)

			s.modifySolutionFile(metadata, solution, output)
		}
	}
}

// Incorrect marks a solution as being not right; this makes it not be automatically submitted through the API.
func (s *Solver[T, V]) Incorrect(solutions ...V) {
	for _, solution := range solutions {
		s.incorrectAnswers = append(s.incorrectAnswers, solution)
	}
}

// shouldSubmit returns whether the solution should be automatically submitted.
func (s *Solver[T, V]) shouldSubmit(solution V) bool {
	// don't submit if it's incorrect
	for _, sol := range s.incorrectAnswers {
		if reflect.DeepEqual(sol, solution) {
			fmt.Println("note: not auto submitting because solution is incorrect")
			return false
		}
	}

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
func (s *Solver[T, V]) getRealInput(metadata solutionMetadata) (string, bool) {
	if s.client == nil {
		s.client = newAocClient(metadata)
	}

	inputPath := filepath.Join(metadata.year, metadata.day, "input.txt")

	fileBytes := readFile(inputPath)

	if len(fileBytes) != 0 {
		return string(fileBytes), true
	}

	fmt.Println("retrieving input because it's not found...")

	input, err := s.client.retrieveInput()
	if err != nil {
		fmt.Printf("note: tried to retrieve input, failed: %s\n", err)
		return "", false
	}

	writeFile(inputPath, input)

	return string(input), true
}

// modifySolutionFile updates the solution file with a call to [Solver.Verify] or
// [Solver.Incorrect] as appropriate based on whether the solution is right or not.
func (s *Solver[T, V]) modifySolutionFile(metadata solutionMetadata, solution V, output string) {
	solFilePath := filepath.Join(metadata.year, metadata.day, metadata.solType, metadata.part, "main.go")

	if strings.Contains(output, "That's the right answer!") {
		solFile := readFile(solFilePath)
		solFile = bytes.ReplaceAll(solFile, []byte("solver.Solve()"), []byte(fmt.Sprintf("solver.Verify(%v)", solution)))
		writeFile(solFilePath, solFile)
		fmt.Println("note: modified solution file to add Verify call")
	} else if strings.Contains(output, "That's not the right answer.") {
		solFile := readFile(solFilePath)
		solFile = bytes.ReplaceAll(solFile, []byte("solver.Solve()"), []byte(fmt.Sprintf("solver.Incorrect(%v)\n\tsolver.Solve()", solution)))
		writeFile(solFilePath, solFile)
		fmt.Println("note: modified solution file to add Incorrect call")
	}
}

// readFile returns the contents of the file at path.
func readFile(path string) []byte {
	return Must(io.ReadAll(Must(os.Open(path))))
}

// writeFile writes data to the file at path.
func writeFile(path string, data []byte) {
	file := Must(os.Create(path))
	Must(file.Write(data))
}

// solutionMetadata describes the metadata the solution is for.
type solutionMetadata struct {
	year, day, solType, part string
}

// getSolutionMetadata returns the year, day, and part the current execution is for.
func getSolutionMetadata() solutionMetadata {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		panic(ok)
	}

	// path is of the form github.com/danvolchek/AdventOfCode/2015/16/leaderboard/1
	parts := strings.Split(buildInfo.Path, string(os.PathSeparator))

	return solutionMetadata{
		year:    parts[3],
		day:     parts[4],
		solType: parts[5],
		part:    parts[6],
	}
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

		if len(matches) == 0 {
			panic("can't parse line: " + line)
		}
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
