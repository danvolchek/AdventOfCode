package main

import (
	"flag"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/lib"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

var infer bool
var dryRun bool

var argSolTypes [2]bool
var argSolution lib.Solution

func init() {
	flag.IntVar(&argSolution.Year, "year", 0, "solution year")
	flag.IntVar(&argSolution.Day, "day", 0, "solution day")
	flag.BoolVar(&argSolTypes[0], "l", false, "create leaderboard")
	flag.BoolVar(&argSolTypes[1], "o", false, "create optimized")
	flag.BoolVar(&dryRun, "dryrun", false, "print solution to create and exit")
	flag.Parse()

	if !(argSolTypes[0] || argSolTypes[1] || argSolution.Year != 0 || argSolution.Day != 0) {
		infer = true
		fmt.Println("Inferring solution to create")
		return
	}

	fail := func(message string) {
		fmt.Println(message)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if argSolution.Year == 0 {
		fail("Year needed")
	}

	if argSolution.Day == 0 {
		fail("Day needed")
	}

	if !argSolTypes[0] && !argSolTypes[1] {
		fail("Specify at least leaderboard or optimized")
	}
}

func main() {
	sols := getSolutionsToCreate()

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

func getSolutionsToCreate() []lib.Solution {
	if infer {
		skipFile, err := os.Open("skip.txt")
		if err != nil {
			panic(err)
		}

		return []lib.Solution{lib.FirstUnsolvedSolution(".", lib.ParseSkips(skipFile))}
	}

	var solutions []lib.Solution

	for _, i := range []int{0, 1} {
		if argSolTypes[i] {
			clone := argSolution
			clone.Leaderboard = i == 0
			solutions = append(solutions, clone)
		}
	}

	return solutions
}

func createSolution(sol lib.Solution) error {
	solutionFolder := filepath.Join(strconv.Itoa(sol.Year), strconv.Itoa(sol.Day))

	err := os.MkdirAll(solutionFolder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("couldn't create solution folder: %s", err)
	}

	inputFile := filepath.Join(solutionFolder, "input.txt")

	if !fileExists(inputFile) {
		input, err := os.Create(inputFile)
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
	stubsWriter := &lib.MultiWriteCloser{}
	defer stubsWriter.Close()

	for _, partOne := range []bool{true, false} {
		// day 25 only has 1 part
		if sol.Day == 25 && !partOne {
			continue
		}

		var solPath string
		var exists bool
		if partOne {
			solPath, exists = sol.PartOne(".")
		} else {
			solPath, exists = sol.PartTwo(".")
		}

		if exists {
			continue
		}

		stubFile, err := createFileAndDirectories(solPath)
		if err != nil {
			return fmt.Errorf("couldn't create stub file %s: %s", solPath, err)
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

func fileExists(path string) bool {
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
