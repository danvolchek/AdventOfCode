package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"regexp"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "16", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

var targetProperties = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

type constraint struct {
	num int

	properties map[string]int
}

func (c constraint) matchesTarget() bool {
	for key, val := range c.properties {
		if targetProperties[key] != val {
			return false
		}
	}

	return true
}

var parseRegexp = regexp.MustCompile(`Sue (\d+): (.+)`)

func parse(matches []string) constraint {
	properties := make(map[string]int)

	rawProps := strings.Split(matches[1], ", ")
	for _, property := range rawProps {
		parsedProp := strings.Split(property, ": ")

		properties[parsedProp[0]] = lib.Atoi(parsedProp[1])
	}

	return constraint{
		num:        lib.Atoi(matches[0]),
		properties: properties,
	}
}

func solve(sues []constraint) int {
	for _, sue := range sues {
		if sue.matchesTarget() {
			return sue.num
		}
	}

	return -1
}

func main() {
	solver := lib.Solver[[]constraint, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseRegexp, parse)),
		SolveF: solve,
	}

	solver.Verify(input(), 103)
}
