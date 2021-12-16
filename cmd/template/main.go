package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/internal/parse"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

const (
	inputFileName = "input.txt"

	cmdDirectory         = "cmd"
	templateDirectory    = "template"
	inferenceTypesToSkip = "skip.txt"
	templateStubName     = "main.txt"
	stubTargetName       = "main.go"
)

var (
	solutionTypes = []string{"leaderboard", "optimized"}
	puzzleParts   = []string{"1", "2"}
)

func main() {
	if err := create(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Done!")
	}
}

func create() error {
	args, err := parseArgs()
	if err != nil {
		flag.PrintDefaults()
		fmt.Println()
		return fmt.Errorf("args are invalid: %s", err)
	}

	solutionFolder := path.Join(args.Year, args.Day)

	err = createSolutionFolder(solutionFolder)
	if err != nil {
		return fmt.Errorf("couldn't create solution folder: %s", err)
	}

	if !exists(path.Join(solutionFolder, inputFileName)) {
		input, err := os.Create(path.Join(solutionFolder, inputFileName))
		defer warn(input.Close)
		if err != nil {
			return fmt.Errorf("couldn't create input file: %s", err)
		}
	}

	tmpl, err := loadTemplate(path.Join(cmdDirectory, templateDirectory, templateStubName))
	if err != nil {
		return fmt.Errorf("couldn't load template: %s", err)
	}

	anyCreated := false
	stubsWriter := &multiWriteCloser{}
	defer warn(stubsWriter.Close)

	for _, solutionType := range args.Types {
		for _, puzzleType := range puzzleParts {
			stubDir := path.Join(solutionFolder, solutionType, puzzleType)

			if exists(path.Join(stubDir, stubTargetName)) {
				continue
			}

			stubFile, err := createFileAndDirectories(stubDir, stubTargetName)
			if err != nil {
				return fmt.Errorf("couldn't create stub file %s: %s", path.Join(stubDir, stubTargetName), err)
			}

			anyCreated = true

			stubsWriter.Add(stubFile)
		}
	}
	if !anyCreated {
		return fmt.Errorf("all files already exist, couldn't create anything")
	}

	err = tmpl.Execute(stubsWriter, args)
	if err != nil {
		return fmt.Errorf("couldn't write template to stubs: %s", err)
	}

	return nil
}

type args struct {
	Year, Day string
	Types     []string
}

func (a args) valid() error {
	if a.Year == "" {
		return errors.New("year must be provided")
	}

	if a.Day == "" {
		return errors.New("day must be provided")
	}

	if len(a.Types) == 0 {
		return errors.New("types must be provided")
	}

	for _, solType := range a.Types {
		if solType != solutionTypes[0] && solType != solutionTypes[1] {
			return fmt.Errorf("invalid types %s", solType)
		}
	}

	return nil
}

func parseArgs() (args, error) {
	var parsed args

	var rawTypes string
	flag.StringVar(&parsed.Year, "year", "", "the year to add")
	flag.StringVar(&parsed.Day, "day", "", "the day to add")
	flag.StringVar(&rawTypes, "types", "", "the types to add")
	flag.Parse()
	parsed.Types = strings.Split(rawTypes, ",")

	if parsed.Year == "" && parsed.Day == "" && rawTypes == "" {
		inferred, err := inferArgs()
		if err != nil {
			return args{}, err
		}

		if err := inferred.valid(); err != nil {
			panic(fmt.Sprintf("infered args were invalid: %v:  %s", inferred, err))
		}

		fmt.Printf("arguments not provided so inferred %+v\n", inferred)

		return inferred, nil
	}

	return parsed, parsed.valid()
}

type skipIndicator struct {
	year, day              string
	leaderboard, optimized bool
}

func (s skipIndicator) Skip(y parse.Year, d parse.Day, leaderboard bool) bool {
	if y.Num != s.year {
		return false
	}

	if s.day == "" {
		return true
	}

	if s.day != d.Num {
		return false
	}

	if (leaderboard && s.leaderboard) || (!leaderboard && s.optimized) {
		return true
	}

	return false
}

func parseSkipfile() ([]skipIndicator, error) {
	skipFile, err := os.Open(path.Join(cmdDirectory, templateDirectory, inferenceTypesToSkip))
	if err != nil {
		return nil, err
	}

	skipBytes, err := io.ReadAll(skipFile)
	if err != nil {
		return nil, err
	}

	var skipIndicators []skipIndicator
	rawSkips := strings.Split(strings.TrimSpace(string(skipBytes)), "\n")

	for _, skip := range rawSkips {
		skip = strings.TrimSpace(skip)
		if commentIndex := strings.Index(skip, "#"); commentIndex != -1 {
			skip = strings.TrimSpace(skip[0:commentIndex])
		}

		parts := strings.Split(skip, "/")

		year, day := "", ""
		leaderboard, optimized := true, true
		switch len(parts) {
		case 1:
			year = parts[0]
		case 2:
			year, day = parts[0], parts[1]
		case 3:
			year, day = parts[0], parts[1]

			if len(parts[2]) > 0 {
				solutionTypes := strings.Split(parts[2], ",")
				leaderboard, optimized = contains(solutionTypes, "l"), contains(solutionTypes, "o")
			}
		}
		skipIndicators = append(skipIndicators, skipIndicator{
			year:        year,
			day:         day,
			leaderboard: leaderboard,
			optimized:   optimized,
		})
	}

	return skipIndicators, nil
}

func contains(items []string, item string) bool {
	for _, value := range items {
		if value == item {
			return true
		}
	}

	return false
}

func inferArgs() (args, error) {
	info := parse.SolutionInformation(".")

	skipIndicators, err := parseSkipfile()
	if err != nil {
		return args{}, err
	}

	return findIncompleteSolution(info, skipIndicators)
}

func shouldSkip(skipIndicators []skipIndicator, year parse.Year, day parse.Day, leaderboard bool) bool {
	for _, indicator := range skipIndicators {
		if indicator.Skip(year, day, leaderboard) {
			return true
		}
	}

	return false
}

func findIncompleteSolution(info []parse.Year, skipIndicators []skipIndicator) (args, error) {
	for i := len(info) - 1; i >= 0; i-- {
		year := info[i]

		for _, day := range year.Days {
			for _, lb := range []bool{true, false} {
				if !shouldSkip(skipIndicators, year, day, lb) && !day.HasSolution(lb) {
					return createArgs(year, day, lb), nil
				}
			}
		}
	}

	yearInt, err := strconv.Atoi(info[0].Num)
	if err != nil {
		return args{}, err
	}

	return args{
		Year:  strconv.Itoa(yearInt + 1),
		Day:   "1",
		Types: []string{"leaderboard"},
	}, nil

}

func createArgs(y parse.Year, d parse.Day, leaderboard bool) args {
	typeStr := "leaderboard"
	if !leaderboard {
		typeStr = "optimized"
	}

	return args{
		Year:  y.Num,
		Day:   d.Num,
		Types: []string{typeStr},
	}
}

func createSolutionFolder(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't make directories: %s", err)
	}

	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err)
	}

	return true
}

func loadTemplate(path string) (*template.Template, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't read template file: %s", err)
	}

	tmpl, err := template.New("main").Parse(string(contents))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse template: %s", err)
	}

	return tmpl, nil
}

func createFileAndDirectories(parent string, child string) (*os.File, error) {
	err := os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("couldn't make directories: %s", err)
	}

	file, err := os.Create(path.Join(parent, child))
	if err != nil {
		return nil, fmt.Errorf("couldn't create file: %s", err)
	}

	return file, nil
}

func warn(action func() error) {
	err := action()
	if err != nil {
		fmt.Printf("warn: %s", err)
	}
}

// based on io.MultiWriter, but is an io.WriteCloser and can add new io.WriteClosers on the fly
type multiWriteCloser struct {
	writeClosers []io.WriteCloser
}

func (mwc *multiWriteCloser) Write(p []byte) (n int, err error) {
	for _, wc := range mwc.writeClosers {
		n, err = wc.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (mwc *multiWriteCloser) Close() error {
	var errMessages []string

	for _, wc := range mwc.writeClosers {
		err := wc.Close()
		if err != nil {
			errMessages = append(errMessages, err.Error())
		}
	}

	if len(errMessages) == 0 {
		return nil
	}

	return errors.New(strings.Join(errMessages, ","))
}

func (mwc *multiWriteCloser) Add(w io.WriteCloser) {
	mwc.writeClosers = append(mwc.writeClosers, w)
}
