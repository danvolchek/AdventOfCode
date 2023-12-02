package internal

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"os"
	"path/filepath"
	"strconv"
)

const (
	FirstDayNum = 1
	FirstDay    = "1"

	LastDayNum = 25
	LastDay    = "25"

	TypeLeaderboard = "leaderboard"
	TypeOptimized   = "optimized"

	PartOne = "1"
	PartTwo = "2"

	FirstYearNum = 2015
	FirstYear    = "2015"
)

// SolutionDirectory represents a directory holding advent of code solutions for zero or more years.
// It's a read only, parsed metadata view of the filesystem. Don't edit any of the values.
type SolutionDirectory struct {
	// root is the path to the directory.
	root string

	// skipper handles skipping solutions when looking for the first unsolved solution.
	skipper Skipper

	// Solutions is a map from year -> day -> solution information for that day in that year.
	Solutions map[string]map[string]Solution
}

// NewSolutionsDirectory creates a new solutions directory using the data at root.
func NewSolutionsDirectory(root string) SolutionDirectory {
	// The algorithm is:
	//  - Iterate over every top level directory whose name only contains digits, assuming them to be years
	//  - Iterate over every subdirectory of those whose names also only contain digits, assume them to be days
	//  - Create a solution for each day
	//  - Store everything indexed by year and day

	yearDirs, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}

	yearMap := make(map[string]map[string]Solution)

	for _, yearDir := range yearDirs {
		if !yearDir.IsDir() {
			continue
		}

		year := yearDir.Name()

		if _, err := strconv.Atoi(year); err != nil {
			continue
		}

		dayDirs, err := os.ReadDir(filepath.Join(root, year))
		if err != nil {
			panic(err)
		}

		dayMap := make(map[string]Solution)

		for _, dayDir := range dayDirs {
			if !dayDir.IsDir() {
				continue
			}

			day := dayDir.Name()

			if _, err := strconv.Atoi(day); err != nil {
				continue
			}

			dayMap[day] = newSolution(root, year, day, true)
		}

		// Cache days that don't exist in years that do
		for dayNum := FirstDayNum; dayNum <= LastDayNum; dayNum++ {
			day := strconv.Itoa(dayNum)
			if _, ok := dayMap[day]; !ok {
				dayMap[day] = newSolution(root, year, day, false)
			}
		}

		yearMap[year] = dayMap
	}

	return SolutionDirectory{
		root:      root,
		skipper:   NewSkipper(filepath.Join(root, "skip.txt")),
		Solutions: yearMap,
	}
}

// FirstUnsolvedSolutionType returns the first solution + type of the solution that isn't solved - i.e.
// that doesn't exist yet.
func (s SolutionDirectory) FirstUnsolvedSolutionType() (Solution, string) {
	years := s.Years()

	if len(years) == 0 {
		return newSolution(s.root, FirstYear, FirstDay, false), TypeLeaderboard
	}

	for _, year := range years {
		for dayNum := FirstDayNum; dayNum <= LastDayNum; dayNum++ {
			day := strconv.Itoa(dayNum)
			if s.skipper.Skip(year, day) {
				continue
			}

			solution := s.Solutions[year][day]

			if !solution.Leaderboard.Exists {
				return solution, TypeLeaderboard
			} else if !solution.Optimized.Exists {
				return solution, TypeOptimized
			}
		}
	}

	lastYear := years[len(years)-1]
	return newSolution(s.root, strconv.Itoa(lib.Atoi(lastYear)+1), FirstDay, false), TypeLeaderboard
}

// Years returns all years that contains solutions in the directory. This doesn't necessarily mean every day in the year
// has a solution.
func (s SolutionDirectory) Years() []string {
	years := lib.Keys(s.Solutions)
	slices.Sort(years)
	return years
}

// Get returns the solution for day in year. It (or parts of it) might not exist - use the Exists field to check.
func (s SolutionDirectory) Get(year, day string) Solution {
	solutions, ok := s.Solutions[year]
	if !ok {
		return newSolution(s.root, year, day, false) // Not an existing year in the repo
	}

	solution, ok := solutions[day]
	if !ok {
		return newSolution(s.root, year, day, false) // Not an existing day in the repo
	}

	return solution
}

// Solution represents metadata about a puzzle solution.
type Solution struct {
	// The directory path of the solution.
	isPath

	// Year is the year the solution is for.
	Year string

	// Day is the day the solution is for.
	Day string

	// Input is the input file for the day.
	Input isPath

	// Leaderboard is the leaderboard solution.
	Leaderboard SolutionType

	// Optimized is the optimized solution.
	Optimized SolutionType
}

// newSolution creates a new solution. When check is false, all file paths are assumed not to exist.
func newSolution(root, year, day string, check bool) Solution {
	base := filepath.Join(root, year, day)

	newSolutionPart := func(path string) SolutionPart {
		return SolutionPart{
			isPath: newPath(path, check),
			Main:   newPath(filepath.Join(path, "main.go"), check),
		}
	}

	newSolutionType := func(path string) SolutionType {
		p1 := newSolutionPart(filepath.Join(path, PartOne))
		p2 := newSolutionPart(filepath.Join(path, PartTwo))

		return SolutionType{
			PartOne: p1,
			PartTwo: p2,
			isPath:  newPath(path, check),
		}
	}

	return Solution{
		isPath: newPath(base, check),

		Year: year,
		Day:  day,

		Input:       newPath(filepath.Join(base, "input.txt"), check),
		Leaderboard: newSolutionType(filepath.Join(base, TypeLeaderboard)),
		Optimized:   newSolutionType(filepath.Join(base, TypeOptimized)),
	}
}

// SolutionType represents metadata about solutions of a certain type for a puzzle.
type SolutionType struct {
	// The directory path to the solution type.
	isPath

	// PartOne is the solution to part one.
	PartOne SolutionPart

	// PartTwo is the solution to part two.
	PartTwo SolutionPart
}

// SolutionPart represents metadata about a solution to the part of a puzzle.
//
// Note: The path of the solution part itself isn't needed currently, and this could just be replaced with an
// [.isPath] pointing to the main.go file in [.SolutionType], but it's left as is in case I want files other than main
// in the future.
type SolutionPart struct {
	// The directory path to the solution part.
	isPath

	// Main is the path to the main.go source file that can be run to solve the puzzle.
	Main isPath
}

// isPath represents a file or folder path.
type isPath struct {
	// The path to the file or folder.
	Path string

	// Whether the file or folder exists.
	Exists bool
}

// newPath creates a new path, and checks whether it exists if check is true.
func newPath(path string, check bool) isPath {
	return isPath{
		Path:   path,
		Exists: check && exists(path),
	}
}

// exists checks if path exists, handling errors if possible and panicking if not.
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err)
	}

	return true
}
