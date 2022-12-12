package internal

import (
	"fmt"
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

var (
	SolutionTypes = []string{TypeLeaderboard, TypeOptimized}
	SolutionParts = []string{PartOne, PartTwo}
)

type hasPath struct {
	Path string
}

func (p hasPath) Exists() bool {
	return exists(p.Path)
}

type Year struct {
	hasPath

	Name   string
	Number int
	Days   []*Day
}

func (y Year) String() string {
	return fmt.Sprintf("Year %s", y.Name)
}

func newYear(root, name string) *Year {
	return &Year{
		hasPath: hasPath{
			Path: filepath.Join(root, name),
		},
		Name:   name,
		Number: lib.Atoi(name),
	}
}

type Day struct {
	hasPath
	Year *Year

	Name   string
	Number int
	Input  *Input
	Types  []*Type
}

func (d Day) String() string {
	return fmt.Sprintf("%s, Day %s", d.Year.String(), d.Name)
}

func newDay(year *Year, name string) *Day {
	day := Day{
		hasPath: hasPath{
			Path: filepath.Join(year.Path, name),
		},
		Year:   year,
		Number: lib.Atoi(name),
		Name:   name,
	}
	day.Input = newInput(&day)

	year.Days = append(year.Days, &day)

	return &day
}

type Input struct {
	hasPath
}

func (i Input) String() string {
	return "input file"
}

func newInput(day *Day) *Input {
	input := Input{
		hasPath: hasPath{
			Path: filepath.Join(day.Path, "input.txt"),
		},
	}

	return &input
}

type Type struct {
	hasPath
	Day *Day

	Name  string
	Parts []*Part
}

func (t Type) String() string {
	return fmt.Sprintf("%s, Type %s", t.Day.String(), t.Name)
}

func newType(day *Day, name string) *Type {
	typ := Type{
		hasPath: hasPath{
			Path: filepath.Join(day.Path, name),
		},
		Day:  day,
		Name: name,
	}

	day.Types = append(day.Types, &typ)

	return &typ
}

func NewType(root, yearN, dayN, typN string) *Type {
	year := newYear(root, yearN)
	day := newDay(year, dayN)
	typ := newType(day, typN)
	_ = newPart(typ, "1")
	_ = newPart(typ, "2")

	return typ
}

type Part struct {
	hasPath
	Type *Type

	Main *Main
	Name string
}

func (p Part) String() string {
	return fmt.Sprintf("%s, Part %s", p.Type.String(), p.Name)
}

func newPart(typ *Type, name string) *Part {
	part := Part{
		hasPath: hasPath{
			Path: filepath.Join(typ.Path, name),
		},
		Type: typ,
		Name: name,
	}
	part.Main = newMain(&part)

	typ.Parts = append(typ.Parts, &part)

	return &part
}

type Main struct {
	hasPath
}

func newMain(part *Part) *Main {
	main := Main{
		hasPath: hasPath{
			Path: filepath.Join(part.Path, "main.go"),
		},
	}

	return &main
}

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

func GetLocalSolutionInfo(root string) []*Year {
	var years []*Year

	dirs, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		if _, err := strconv.Atoi(dir.Name()); err != nil {
			continue
		}

		year := newYear(root, dir.Name())

		for dayNum := FirstDayNum; dayNum <= LastDayNum; dayNum++ {
			day := newDay(year, strconv.Itoa(dayNum))

			for _, typName := range SolutionTypes {
				typ := newType(day, typName)

				// Day 25 only has 1 part
				parts := SolutionParts
				if dayNum == LastDayNum {
					parts = []string{"1"}
				}

				for _, partNum := range parts {
					newPart(typ, partNum)
				}
			}
		}

		years = append(years, year)
	}

	slices.SortFunc(years, func(a, b *Year) bool {
		return a.Number < b.Number
	})
	return years
}
