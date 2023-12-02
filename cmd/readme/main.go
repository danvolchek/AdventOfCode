package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	readMePath = "README.md"
)

var tableSection = []byte("# Completion")

// readme keeps the readme up to date with the currently solved solutions.
// The completion section needs to be the last section in the readme.
func main() {
	err := generateReadme(".")
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}

func generateReadme(root string) error {
	readme, err := os.ReadFile(readMePath)
	if err != nil {
		return err
	}

	before, _, found := bytes.Cut(readme, tableSection)
	if !found {
		return errors.New("could not find where to write table in readme")
	}

	completionBuffer := bytes.NewBuffer(before)
	completionBuffer.Write(tableSection)
	completionBuffer.WriteString("\n")

	solutions := internal.NewSolutionsDirectory(root)

	years := solutions.Years()
	slices.Reverse(years)

	for _, year := range years {
		completionBuffer.WriteString("\n## ")
		completionBuffer.WriteString(year)
		completionBuffer.WriteString("\n\n")

		daysTable := createTable(solutions, year)
		daysTable.ToBuffer(completionBuffer)
	}

	err = os.WriteFile(readMePath, completionBuffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

// createTable creates a table for a year of puzzle solutions.
// The table format is:
//
//		            |  1      | 2   | 3
//		Leaderboard | <links> | ... | ...
//	    Optimized   | <links> | ... | ...
//
// In other words, it's 2 rows (solution type) x 25 columns (for each day).
// The text in each box contains links to both parts, if they exist.
func createTable(solutions internal.SolutionDirectory, year string) *internal.Table {
	yearTable := &internal.Table{
		NumRows: 2,
	}

	// The first column is the left header.
	yearTable.AddColumn("", []string{internal.TypeLeaderboard, internal.TypeOptimized})

	// Each subsequent column is the day number, followed by the links to the main files.
	for dayNum := internal.FirstDayNum; dayNum <= internal.LastDayNum; dayNum++ {
		day := strconv.Itoa(dayNum)
		solution := solutions.Get(year, day)

		yearTable.AddColumn(day, []string{
			createLinks(solution.Leaderboard),
			createLinks(solution.Optimized),
		})
	}

	return yearTable
}

// createLinks creates links for a solution type, in the form of individual links for each part
// separated by a comma.
func createLinks(solType internal.SolutionType) string {
	var links []string

	if solType.PartOne.Exists {
		links = append(links, makeLink(internal.PartOne, solType.PartOne.Main.Path))
	}

	if solType.PartTwo.Exists {
		links = append(links, makeLink(internal.PartTwo, solType.PartTwo.Main.Path))
	}

	return strings.Join(links, ",")
}

// makeLink creates a markdown link of the form [visibleText](path/to/link).
func makeLink(visibleText, path string) string {
	return fmt.Sprintf("[%s](%s)", visibleText, path)
}
