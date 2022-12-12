package internal

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"io"
	"os"
	"strconv"
	"strings"
)

type Skipper struct {
	skips []skip
}

func (s *Skipper) Skip(typ *Type) bool {
	for _, skip := range s.skips {
		if skip.shouldSkip(typ) {
			return true
		}
	}

	return false
}

func NewSkipper(path string) *Skipper {
	if !exists(path) {
		return &Skipper{}
	}

	return &Skipper{skips: parseSkips(lib.Must(os.Open(path)))}
}

type skipRange struct {
	Min, Max int
}

func (s skipRange) contains(value int) bool {
	return value >= s.Min && value <= s.Max
}

type skip struct {
	Year, Day skipRange
}

func (s skip) shouldSkip(t *Type) bool {
	return s.Year.contains(t.Day.Year.Number) && (s.Day.Max == 0 || s.Day.contains(t.Day.Number))
}

func parseSkips(r io.Reader) []skip {
	var solutionsToSkip []skip

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

		var skip skip
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

func parseRange(s string) (skipRange, error) {
	parts := strings.Split(s, "-")
	if len(parts) > 2 {
		return skipRange{}, errors.New("range should have at most one separator")
	}

	min, err := strconv.Atoi(parts[0])
	if err != nil {
		errMsg := "min isn't a number"
		if len(parts) == 1 {
			errMsg = "isn't a number"
		}

		return skipRange{}, errors.New(errMsg)
	}

	if len(parts) == 1 {
		return skipRange{
			Min: min,
			Max: min,
		}, nil
	}

	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return skipRange{}, errors.New("max isn't a number")
	}

	return skipRange{
		Min: min,
		Max: max,
	}, nil
}
