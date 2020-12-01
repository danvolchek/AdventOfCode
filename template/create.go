package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	solutionTypeInitial = "initial"
	solutionTypeClean   = "clean"
	puzzleTypeFirst     = "1"
	puzzleTypeSecond    = "2"

	inputFileName = "input.txt"

	templateDirectory = "template"
	templateFileName  = "main.go"
)

var (
	solutionTypes = []string{solutionTypeInitial, solutionTypeClean}
	puzzleTypes   = []string{puzzleTypeFirst, puzzleTypeSecond}
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
		return fmt.Errorf("args are invalid: %s", err)
	}

	solutionFolder := path.Join(args.Year, args.Day)

	err = createSolutionFolder(solutionFolder)
	if err != nil {
		return fmt.Errorf("couldn't create solution folder: %s", err)
	}

	input, err := os.Create(path.Join(solutionFolder, inputFileName))
	defer warn(input.Close)
	if err != nil {
		return fmt.Errorf("couldn't create input file: %s", err)
	}

	tmpl, err := loadTemplate(path.Join(templateDirectory, templateFileName))
	if err != nil {
		return err
	}

	stubsWriter := &multiWriteCloser{}
	defer warn(stubsWriter.Close)

	for _, solutionType := range solutionTypes {
		for _, puzzleType := range puzzleTypes {
			stubDir := path.Join(solutionFolder, solutionType, puzzleType)
			stubFile, err := createStubFile(stubDir, templateFileName)
			if err != nil {
				return fmt.Errorf("couldn't create stub %s: %s", path.Join(stubDir, templateFileName), err)
			}

			stubsWriter.Add(stubFile)
		}
	}

	err = tmpl.Execute(stubsWriter, args)
	if err != nil {
		return fmt.Errorf("couldn't write template to stubs: %s", err)
	}

	return nil
}

type args struct {
	Year, Day string
}

func parseArgs() (args, error) {
	var parsed args

	flag.StringVar(&parsed.Year, "Year", "", "the year to add")
	flag.StringVar(&parsed.Day, "Day", "", "the day to add")
	flag.Parse()

	if parsed.Year == "" {
		return args{}, errors.New("year must be provided")
	}

	if parsed.Day == "" {
		return args{}, errors.New("day must be provided")
	}

	return parsed, nil
}

func createSolutionFolder(path string) error {
	exists, err := exists(path)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("%s already exists", path)
	}

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't make directories: %s", err)
	}

	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func loadTemplate(path string) (*template.Template, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't read stub template: %s", err)
	}

	tmpl, err := template.New("main").Parse(string(contents))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse stub template: %s", err)
	}

	return tmpl, nil
}

func createStubFile(stubDir string, stub string) (*os.File, error) {
	err := os.MkdirAll(stubDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("couldn't make folder for stub: %s", err)
	}

	file, err := os.Create(path.Join(stubDir, stub))
	if err != nil {
		return nil, fmt.Errorf("couldn't create stub file: %s", err)
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
