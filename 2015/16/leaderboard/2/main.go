package main

import (
	"regexp"
	"strings"

	"github.com/danvolchek/AdventOfCode/lib"
)

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
		switch key {
		case "cats", "trees":
			if val <= targetProperties[key] {
				return false
			}
		case "pomeranians", "goldfish":
			if val >= targetProperties[key] {
				return false
			}
		default:
			if targetProperties[key] != val {
				return false
			}
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

	solver.Verify(405)
}
