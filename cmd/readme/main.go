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

func createTable(solutions internal.SolutionDirectory, year string) *internal.Table {
	yearTable := &internal.Table{
		NumRows: 2,
	}

	yearTable.AddColumn("", []string{internal.TypeLeaderboard, internal.TypeOptimized})

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

func makeLink(visibleText, path string) string {
	return fmt.Sprintf("[%s](%s)", visibleText, path)
}
