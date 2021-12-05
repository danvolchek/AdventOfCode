package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/internal/parse"
	"github.com/danvolchek/AdventOfCode/cmd/readme/table"
	"os"
	"strings"
)

const (
	readMePath = "README.md"
)

var tableSection = []byte("# Completion")

func main() {
	err := generateReadme()
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}

func generateReadme() error {
	readme, err := os.ReadFile(readMePath)
	if err != nil {
		return err
	}

	before, _, found := Cut(readme, tableSection)
	if !found {
		return errors.New("could not find where to write table in readme")
	}

	buf := bytes.NewBuffer(before)
	buf.Write(tableSection)
	buf.WriteString("\n")

	solutions := parse.SolutionInformation(".")
	for _, year := range solutions {
		buf.WriteString("\n## ")
		buf.WriteString(year.Num)
		buf.WriteString("\n\n")

		daysTable := createTable(year.Days)
		daysTable.ToBuffer(buf)
	}

	err = os.WriteFile(readMePath, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func createTable(days []parse.Day) *table.Table {
	yearTable := &table.Table{
		NumRows: 2,
	}

	yearTable.AddColumn("", []string{"leaderboard", "optimized"})

	for _, day := range days {
		yearTable.AddColumn(day.Num, []string{
			createLink(day, func(part parse.Part) string {
				return part.LeaderboardSolutionPath
			}),
			createLink(day, func(part parse.Part) string {
				return part.OptimizedSolutionPath
			}),
		})
	}

	return yearTable
}

func createLink(day parse.Day, selector func(part parse.Part) string) string {
	var parts []string

	partOnePath := selector(day.PartOne)
	if partOnePath != "" {
		parts = append(parts, makeLink("1", partOnePath))
	}

	partTwoPath := selector(day.PartTwo)
	if partTwoPath != "" {
		parts = append(parts, makeLink("2", partTwoPath))
	}

	return strings.Join(parts, ",")
}

func makeLink(visibleText, path string) string {
	return fmt.Sprintf("[%s](%s)", visibleText, path)
}

// Cut cuts s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
//
// Taken from https://github.com/golang/go/issues/46336 as 1.18 hasn't released yet
func Cut(s, sep []byte) (before, after []byte, found bool) {
	if i := bytes.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, []byte{}, false
}
