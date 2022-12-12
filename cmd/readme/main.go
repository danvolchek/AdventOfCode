package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
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

	years := internal.GetLocalSolutionInfo(root)

	for i := range years {
		year := years[len(years)-i-1]

		completionBuffer.WriteString("\n## ")
		completionBuffer.WriteString(year.Name)
		completionBuffer.WriteString("\n\n")

		daysTable := createTable(year)
		daysTable.ToBuffer(completionBuffer)
	}

	err = os.WriteFile(readMePath, completionBuffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func createTable(year *internal.Year) *internal.Table {
	yearTable := &internal.Table{
		NumRows: len(internal.SolutionTypes),
	}

	yearTable.AddColumn("", internal.SolutionTypes)

	for _, day := range year.Days {
		yearTable.AddColumn(day.Name, lib.Map(day.Types, createLink))
	}

	return yearTable
}

func createLink(typ *internal.Type) string {
	var parts []string

	for _, part := range typ.Parts {
		if !part.Main.Exists() {
			continue
		}

		parts = append(parts, makeLink(part.Name, part.Main.Path()))
	}

	return strings.Join(parts, ",")
}

func makeLink(visibleText, path string) string {
	return fmt.Sprintf("[%s](%s)", visibleText, path)
}
