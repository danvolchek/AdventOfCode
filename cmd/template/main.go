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

var argCreateLeaderboard bool
var argCreateOptimized bool
var argYear string
var argDay string

func init() {
	flag.StringVar(&argYear, "year", "0", "solution year")
	flag.StringVar(&argDay, "day", "0", "solution day")
	flag.BoolVar(&argCreateLeaderboard, "l", false, "create leaderboard")
	flag.BoolVar(&argCreateOptimized, "o", false, "create optimized")
	flag.BoolVar(&dryRun, "dryrun", false, "print solution to create and exit")
	flag.Parse()

	if !(argCreateLeaderboard || argCreateOptimized || argYear != "0" || argDay != "0") {
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

	if !argCreateLeaderboard && !argCreateOptimized {
		fail("Specify at least leaderboard or optimized")
	}
}

func main() {
	solution, types := getSolutionsToCreate(".")

	for _, solutionType := range types {
		if dryRun {
			fmt.Printf("Would create: Year %s Day %s Type %s\n", solution.Year, solution.Day, solutionType)
			continue
		}
		fmt.Printf("Creating: Year %s Day %s Type %s\n", solution.Year, solution.Day, solutionType)

		err := createSolution(solution, solutionType)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Done!")
}

func getSolutionsToCreate(root string) (internal.Solution, []string) {
	solutions := internal.NewSolutionsDirectory(root)

	if infer {
		solution, solutionType := solutions.FirstUnsolvedSolutionType()
		return solution, []string{solutionType}
	}

	solution := solutions.Get(argYear, argDay)

	var solutionTypes []string
	if argCreateLeaderboard {
		solutionTypes = append(solutionTypes, internal.TypeLeaderboard)
	}

	if argCreateOptimized {
		solutionTypes = append(solutionTypes, internal.TypeOptimized)
	}

	return solution, solutionTypes
}

func createSolution(solution internal.Solution, solutionType string) error {
	err := os.MkdirAll(solution.pathy, os.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't create solution folder: %s", err)
	}

	if !solution.Input.Exists {
		input, err := os.Create(solution.Input.Path)
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

	var parts []internal.SolutionPart
	switch solutionType {
	case internal.TypeLeaderboard:
		parts = append(parts, solution.Leaderboard.PartOne, solution.Leaderboard.PartTwo)
	case internal.TypeOptimized:
		parts = append(parts, solution.Optimized.PartOne, solution.Optimized.PartTwo)
	}

	stubsWriter := &internal.MultiWriteCloser{}
	defer stubsWriter.Close()

	for _, part := range parts {
		if part.Main.Exists {
			continue
		}

		stubFile, err := createFileAndDirectories(part.Main.Path)
		if err != nil {
			return fmt.Errorf("couldn't create stub file %s: %s", part.Main.Path, err)
		}

		anyCreated = true

		stubsWriter.Add(stubFile)
	}

	if !anyCreated {
		fmt.Println("all files already exist, couldn't create anything")
		return nil
	}

	err = tmpl.Execute(stubsWriter, solution)
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
