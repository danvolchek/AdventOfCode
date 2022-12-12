package main

import (
	"flag"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"os"
	"path/filepath"
	"text/template"
)

var infer bool
var dryRun bool

var argSolTypes [2]bool
var argYear string
var argDay string

func init() {
	flag.StringVar(&argYear, "year", "0", "solution year")
	flag.StringVar(&argDay, "day", "0", "solution day")
	flag.BoolVar(&argSolTypes[0], "l", false, "create leaderboard")
	flag.BoolVar(&argSolTypes[1], "o", false, "create optimized")
	flag.BoolVar(&dryRun, "dryrun", false, "print solution to create and exit")
	flag.Parse()

	if !(argSolTypes[0] || argSolTypes[1] || argYear != "0" || argDay != "0") {
		infer = true
		fmt.Println("Inferring solution to create")
		return
	}

	fail := func(message string) {
		fmt.Println(message)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if argYear == "0" {
		fail("Year needed")
	}

	if argDay == "0" {
		fail("Day needed")
	}

	if !argSolTypes[0] && !argSolTypes[1] {
		fail("Specify at least leaderboard or optimized")
	}
}

func main() {
	sols := getSolutionsToCreate(".")

	for _, sol := range sols {
		if dryRun {
			fmt.Println("Would create:", sol)
			continue
		}
		fmt.Println("Creating solution:", sol)

		err := createSolution(sol)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Done!")
}

func getSolutionsToCreate(root string) []*internal.Type {
	if infer {
		years := internal.GetLocalSolutionInfo(root)
		skipper := internal.NewSkipper(filepath.Join(root, "skips.txt"))

		return []*internal.Type{internal.FirstUnsolvedSolution(root, years, skipper)}
	}

	var solutions []*internal.Type

	for _, i := range []int{0, 1} {
		if argSolTypes[i] {

			typ := internal.TypeLeaderboard
			if i == 1 {
				typ = internal.TypeOptimized
			}

			solutions = append(solutions, internal.NewType(root, argYear, argDay, typ))
		}
	}

	return solutions
}

func createSolution(sol *internal.Type) error {
	err := os.MkdirAll(sol.Day.Path(), os.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't create solution folder: %s", err)
	}

	if !sol.Input.Exists() {
		input, err := os.Create(sol.Input.Path())
		defer input.Close()
		if err != nil {
			return fmt.Errorf("couldn't create input file: %s", err)
		}
	}

	tmpl, err := loadTemplate(filepath.Join("cmd", "template", "main.txt"))
	if err != nil {
		return fmt.Errorf("couldn't load template: %s", err)
	}

	anyCreated := false
	stubsWriter := &internal.MultiWriteCloser{}
	defer stubsWriter.Close()

	for _, part := range sol.Parts {
		if part.Exists() {
			continue
		}

		stubFile, err := createFileAndDirectories(part.Main.Path())
		if err != nil {
			return fmt.Errorf("couldn't create stub file %s: %s", part.Main.Path(), err)
		}

		anyCreated = true

		stubsWriter.Add(stubFile)
	}

	if !anyCreated {
		fmt.Println("all files already exist, couldn't create anything")
		return nil
	}

	err = tmpl.Execute(stubsWriter, sol)
	if err != nil {
		return fmt.Errorf("couldn't write template to stubs: %s", err)
	}

	return nil
}

func loadTemplate(path string) (*template.Template, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't read template file: %s", err)
	}

	tmpl, err := template.New("main").Parse(string(contents))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse template: %s", err)
	}

	return tmpl, nil
}

func createFileAndDirectories(path string) (*os.File, error) {
	parent := filepath.Dir(path)
	child, err := filepath.Rel(parent, path)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(parent, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("couldn't make directories: %s", err)
	}

	file, err := os.Create(filepath.Join(parent, child))
	if err != nil {
		return nil, fmt.Errorf("couldn't create file: %s", err)
	}

	return file, nil
}
