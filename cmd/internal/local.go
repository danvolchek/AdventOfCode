package internal

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"os"
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

type hasPath interface {
	Path() string
}

type rootPath struct {
	path string
}

func (r rootPath) Path() string {
	return r.path
}

type pathed struct {
	part   string
	parent hasPath
}

func (p pathed) Path() string {
	return p.parent.Path() + string(os.PathSeparator) + p.part
}

func (p pathed) Exists() bool {
	return exists(p.Path())
}

type Year struct {
	pathed

	Name   string
	Number int
	Days   []*Day
}

func (y Year) String() string {
	return fmt.Sprintf("Year %s", y.Name)
}

func newYear(root, name string) *Year {
	return &Year{
		pathed: pathed{
			part:   name,
			parent: rootPath{path: root},
		},
		Name:   name,
		Number: lib.Atoi(name),
	}
}

type Day struct {
	pathed
	*Year

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
		pathed: pathed{
			part:   name,
			parent: year,
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
	pathed
}

func (i Input) String() string {
	return "input file"
}

func newInput(day *Day) *Input {
	input := Input{
		pathed: pathed{
			part:   "input.txt",
			parent: day,
		},
	}

	return &input
}

type Type struct {
	pathed
	*Day

	Name  string
	Parts []*Part
}

func (t Type) String() string {
	return fmt.Sprintf("%s, Type %s", t.Day.String(), t.Name)
}

func newType(day *Day, name string) *Type {
	typ := Type{
		pathed: pathed{
			part:   name,
			parent: day,
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
	pathed
	*Type

	Main *Main
	Name string
}

func (p Part) String() string {
	return fmt.Sprintf("%s, Part %s", p.Type.String(), p.Name)
}

func newPart(typ *Type, name string) *Part {
	part := Part{
		pathed: pathed{
			part:   name,
			parent: typ,
		},
		Type: typ,
		Name: name,
	}
	part.Main = newMain(&part)

	typ.Parts = append(typ.Parts, &part)

	return &part
}

type Main struct {
	pathed
}

func newMain(part *Part) *Main {
	main := Main{
		pathed: pathed{
			part:   "main.go",
			parent: part,
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
