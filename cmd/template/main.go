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

	cmdDirectory      = "cmd"
	templateDirectory = "template"
	templateStubName  = "_main.txt"
	stubTargetName    = "main.go"
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

func parseArgs() (args, error) {
	var parsed args

	var rawTypes string
	flag.StringVar(&parsed.Year, "year", "", "the year to add")
	flag.StringVar(&parsed.Day, "day", "", "the day to add")
	flag.StringVar(&rawTypes, "types", "", "the types to add")
	flag.Parse()

	if parsed.Year == "" && parsed.Day == "" && rawTypes == "" {
		return inferArgs()
	}

	if parsed.Year == "" {
		return args{}, errors.New("year must be provided")
	}

	if parsed.Day == "" {
		return args{}, errors.New("day must be provided")
	}

	if rawTypes == "" {
		return args{}, errors.New("types must be provided")
	}

	parsed.Types = strings.Split(rawTypes, ",")
	for _, part := range parsed.Types {
		if part != solutionTypes[0] && part != solutionTypes[1] {
			return args{}, fmt.Errorf("invalid types %s", part)
		}
	}

	return parsed, nil
}

func inferArgs() (args, error) {
	info := parse.SolutionInformation(".")

	year := info[0]
	day := 0
	leaderboard := true
	for day < len(year.Days) && year.Days[day].HasSolution(leaderboard) {
		if leaderboard {
			leaderboard = false
		} else {
			leaderboard = true
			day += 1
		}
	}

	var yearStr, dayStr, typeStr string
	if day == len(year.Days) {
		yearStr = strconv.Itoa(parse.ToInt(year.Num) + 1)
		dayStr = "0"
		typeStr = "leaderboard"
	} else {
		yearStr = year.Num
		dayStr = year.Days[day].Num
		typeStr = "leaderboard"
		if !leaderboard {
			typeStr = "optimized"
		}
	}

	fmt.Printf("Arguments not provided, inferring year %s day %s type %s\n", yearStr, dayStr, typeStr)

	return args{
		Year:  yearStr,
		Day:   dayStr,
		Types: []string{typeStr},
	}, nil
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
