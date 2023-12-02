package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"os"
	"path/filepath"
)

var infer bool
var dryRun bool

var argCreateLeaderboard bool
var argCreateOptimized bool
var argYear string
var argDay string

var verb string

func init() {
	flag.StringVar(&argYear, "year", "0", "solution year")
	flag.StringVar(&argDay, "day", "0", "solution day")
	flag.BoolVar(&argCreateLeaderboard, "l", false, "create leaderboard")
	flag.BoolVar(&argCreateOptimized, "o", false, "create optimized")
	flag.BoolVar(&dryRun, "dryrun", false, "print solution to create and exit")
	flag.Parse()

	verb = "Creating"
	if dryRun {
		verb = "Would create"
	}

	if !(argCreateLeaderboard || argCreateOptimized || argYear != "0" || argDay != "0") {
		infer = true
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
	solution, types := getSolutionToCreate(".")

	for _, solutionType := range types {
		fmt.Printf("%s: %s %s %s\n", verb, solution.Year, solution.Day, solutionType)

		filesToCreate, err := getFilesToCreate(solution, solutionType)
		if err != nil {
			panic(err)
		}

		err = createFiles(filesToCreate)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Done!")
}

// getSolutionToCreate figures out which solution and which types for it need to be created.
func getSolutionToCreate(root string) (internal.Solution, []string) {
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

// getFilesToCreate figures out which files to create for a solution and a type.
func getFilesToCreate(solution internal.Solution, solutionType string) ([]fileToCreate, error) {
	var files []fileToCreate

	if !solution.Input.Exists {
		files = append(files, fileToCreate{
			Path:     solution.Input.Path,
			Contents: nil,
		})
	}

	template, err := os.ReadFile(filepath.Join("cmd", "template", "main.txt"))
	if err != nil {
		return nil, fmt.Errorf("couldn't read template: %s", err)
	}

	addMainFile := func(part internal.SolutionPart, contents []byte) {
		if !part.Main.Exists {
			files = append(files, fileToCreate{
				Path:     part.Main.Path,
				Contents: contents,
			})
		}
	}

	addMainFileFromPrevious := func(part internal.SolutionPart, previous internal.SolutionPart, fallback []byte) {
		contents := fallback
		if previous.Main.Exists {
			var err error
			contents, err = os.ReadFile(previous.Main.Path)
			if err != nil {
				panic(err)
			}
		}

		addMainFile(part, contents)
	}

	switch solutionType {
	case internal.TypeLeaderboard:
		addMainFile(solution.Leaderboard.PartOne, template)
		addMainFile(solution.Leaderboard.PartTwo, template)
	case internal.TypeOptimized:
		addMainFileFromPrevious(solution.Optimized.PartOne, solution.Leaderboard.PartOne, template)
		addMainFileFromPrevious(solution.Optimized.PartTwo, solution.Leaderboard.PartTwo, template)
	}

	return files, nil
}

// createFiles creates files to be created.
func createFiles(files []fileToCreate) error {
	if len(files) == 0 {
		fmt.Println("all files already exist, nothing to create")
		return nil
	}

	var errs []error
	for _, file := range files {
		fmt.Printf("%s: %s\n", verb, file.Path)

		if dryRun {
			continue
		}

		err := os.MkdirAll(filepath.Dir(file.Path), os.ModePerm)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = os.WriteFile(file.Path, file.Contents, 0666)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return fmt.Errorf("couldn't create files: %s", errors.Join(errs...))
	}

	return nil
}

// fileToCreate represents a file to be created.
type fileToCreate struct {
	// Path is the path to the file.
	Path string

	// Contents is the contents the file will have after creation.
	Contents []byte
}
