package lib

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type SkipRange struct {
	Min, Max int
}

func (s SkipRange) contains(value int) bool {
	return value >= s.Min && value <= s.Max
}

type SkipSolution struct {
	Year, Day SkipRange
}

func (s SkipSolution) ShouldSkip(sol Solution) bool {
	return s.Year.contains(sol.Year) && (s.Day.Max == 0 || s.Day.contains(sol.Day))
}

func ParseSkips(r io.Reader) []SkipSolution {
	var solutionsToSkip []SkipSolution

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if comment := strings.Index(line, "#"); comment != -1 {
			line = strings.TrimSpace(line[0:comment])
		}

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, "/")
		if len(parts) >= 3 {
			panic(fmt.Sprintf("bad skip: %v: wrong number of parts", line))
		}

		var skip SkipSolution
		var err error

		if len(parts) >= 1 {
			skip.Year, err = parseRange(parts[0])
			if err != nil {
				panic(fmt.Sprintf("bad skip: %v: year: %v", line, err))
			}
		}

		if len(parts) >= 2 {
			skip.Day, err = parseRange(parts[1])
			if err != nil {
				panic(fmt.Sprintf("bad skip: %v: day: %v", line, err))
			}
		}

		solutionsToSkip = append(solutionsToSkip, skip)

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return solutionsToSkip
}

func parseRange(s string) (SkipRange, error) {
	parts := strings.Split(s, "-")
	if len(parts) > 2 {
		return SkipRange{}, errors.New("range should have at most one separator")
	}

	min, err := strconv.Atoi(parts[0])
	if err != nil {
		errMsg := "min isn't a number"
		if len(parts) == 1 {
			errMsg = "isn't a number"
		}

		return SkipRange{}, errors.New(errMsg)
	}

	if len(parts) == 1 {
		return SkipRange{
			Min: min,
			Max: min,
		}, nil
	}

	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return SkipRange{}, errors.New("max isn't a number")
	}

	return SkipRange{
		Min: min,
		Max: max,
	}, nil
}
