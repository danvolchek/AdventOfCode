package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/lib"
	"os"
	"sort"
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

	before, _, found := Cut(readme, tableSection)
	if !found {
		return errors.New("could not find where to write table in readme")
	}

	completionBuffer := bytes.NewBuffer(before)
	completionBuffer.Write(tableSection)
	completionBuffer.WriteString("\n")

	years := lib.YearsWithSolutions(root)
	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	for _, year := range years {
		completionBuffer.WriteString("\n## ")
		completionBuffer.WriteString(strconv.Itoa(year))
		completionBuffer.WriteString("\n\n")

		daysTable := createTable(root, year)
		daysTable.ToBuffer(completionBuffer)
	}

	err = os.WriteFile(readMePath, completionBuffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func createTable(root string, year int) *lib.Table {
	yearTable := &lib.Table{
		NumRows: 2,
	}

	yearTable.AddColumn("", []string{"leaderboard", "optimized"})

	for day := 1; day <= 25; day++ {
		yearTable.AddColumn(strconv.Itoa(day), []string{
			createLink(root, year, day, true),
			createLink(root, year, day, false),
		})
	}

	return yearTable
}

func createLink(root string, year, day int, leaderboard bool) string {
	var parts []string

	sol := lib.Solution{
		Year:        year,
		Day:         day,
		Leaderboard: leaderboard,
	}

	partOnePath, ok := sol.PartOne(root)
	if ok {
		parts = append(parts, makeLink("1", partOnePath))
	}

	partTwoPath, ok := sol.PartTwo(root)
	if ok {
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
