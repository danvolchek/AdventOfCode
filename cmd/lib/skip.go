package lib

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type SkipSolution struct {
	Year, Day int
}

func (s SkipSolution) ShouldSkip(sol Solution) bool {
	return s.Year == sol.Year && (s.Day == 0 || s.Day == sol.Day)
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
			skip.Year, err = strconv.Atoi(parts[0])
			if err != nil {
				panic(fmt.Sprintf("bad skip: %v: year isn't a number", line))
			}
		}

		if len(parts) >= 2 {
			skip.Day, err = strconv.Atoi(parts[1])
			if err != nil {
				panic(fmt.Sprintf("bad skip: %v: day isn't a number", line))
			}
		}

		solutionsToSkip = append(solutionsToSkip, skip)

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return solutionsToSkip
}
