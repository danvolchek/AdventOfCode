package internal

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"os"
	"path/filepath"
	"strconv"
)

type SolutionDirectory struct {
	root string

	Solutions map[string]map[string]Solution
}

func NewSolutionsDirectory(root string) SolutionDirectory {
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
		Solutions: yearMap,
	}
}

func (s SolutionDirectory) FirstNonexistentType(skipper Skipper) (Solution, string) {
	years := s.Years()

	if len(years) == 0 {
		return newSolution(s.root, FirstYear, FirstDay, false), TypeLeaderboard
	}

	for _, year := range years {
		for dayNum := FirstDayNum; dayNum <= LastDayNum; dayNum++ {
			day := strconv.Itoa(dayNum)
			if skipper.Skip(year, day) {
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

func (s SolutionDirectory) Years() []string {
	years := lib.Keys(s.Solutions)
	slices.Sort(years)
	return years
}

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

type Solution struct {
	Root string
	Year string
	Day  string

	Exists bool

	Leaderboard TTType
	Optimized   TTType
}

func newSolution(root, year, day string, check bool) Solution {
	base := filepath.Join(root, year, day)

	createAndCheck := func(p string) PPPart {
		return PPPart{
			root:   p,
			Exists: check && exists(p),
		}
	}

	createAndCheck2 := func(typ string) TTType {
		p1 := createAndCheck(filepath.Join(base, typ, PartOne))
		p2 := createAndCheck(filepath.Join(base, typ, PartTwo))

		return TTType{
			PartOne: p1,
			PartTwo: p2,
			Exists:  p1.Exists || p2.Exists,
		}
	}

	return Solution{
		Root: root,
		Year: year,
		Day:  day,

		Exists: check && exists(base),

		Leaderboard: createAndCheck2(TypeLeaderboard),
		Optimized:   createAndCheck2(TypeOptimized),
	}
}

type PPPart struct {
	root string

	Exists bool
}

type TTType struct {
	PartOne PPPart
	PartTwo PPPart

	Exists bool
}

func (p PPPart) Main() string {
	return filepath.Join(p.root, "main.go")
}

func (s Solution) Input() string {
	return filepath.Join(s.Root, s.Year, s.Day, "input.txt")
}
