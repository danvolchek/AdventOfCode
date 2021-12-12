package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "12", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) map[string][]string {
	scanner := bufio.NewScanner(r)

	caves := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "-")
		first := parts[0]
		second := parts[1]

		caves[first] = append(caves[first], second)
		caves[second] = append(caves[second], first)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return caves
}

func solve(r io.Reader) {
	caves := parse(r)
	t := time.Now()

	numPaths := explore(caves, "start", nil)

	fmt.Println(numPaths)
	fmt.Println(time.Now().Sub(t))
}

// pathNode represents the current cave being visited along some path. By storing a pointer to the previous pathNode, it
// builds a chain of caves in reverse order.
type pathNode struct {
	cave string
	prev *pathNode
}

// explore returns the number of paths from currentCave to "end", given an already travelled path of path and caves for
// valid transitions between caves according to the rules of the part.
func explore(caves map[string][]string, currentCave string, path *pathNode) int {
	if currentCave == "end" {
		return 1
	}

	currentPath := &pathNode{
		cave: currentCave,
		prev: path,
	}
	paths := 0

	for _, nextCave := range caves[currentCave] {
		// if we're visiting start again, don't - that's not allowed
		// else, if all small caves have been visited exactly once, any cave is a valid cave
		// else, if we've done this transition two or more times before, don't do it again to avoid infinite loops
		// else, if we're visiting a small cave again, don't - that's not allowed (the double small cave has already been used)
		if nextCave == "start" || (!smallCavesAllVisitedOnce(currentPath) && (transitionedBefore(currentCave, nextCave, currentPath) || visitingSmallAgain(nextCave, currentPath))) {
			continue
		}

		paths += explore(caves, nextCave, currentPath)
	}

	return paths
}

// transitionedBefore returns whether from -> to has happened two or more times already.
// this is up one from part 1, because being able to visit a small cave twice means we may need to repeat transitions twice
func transitionedBefore(from, to string, path *pathNode) bool {
	seen := 0
	lastCave := ""

	seenTooManyTimes := iterate(path, func(cave string) bool {
		if cave == from && lastCave == to {
			if seen == 2 {
				return false
			}

			seen += 1
		}

		lastCave = cave
		return true
	})

	return seenTooManyTimes
}

func visitingSmallAgain(cave string, path *pathNode) bool {
	if strings.ToLower(cave) != cave {
		return false
	}

	foundSmall := iterate(path, func(visitedCave string) bool {
		if visitedCave == cave {
			return false
		}

		return true
	})

	return foundSmall
}

func smallCavesAllVisitedOnce(path *pathNode) bool {
	smalls := make(map[string]int)

	foundDoubleSmall := iterate(path, func(cave string) bool {
		if strings.ToLower(cave) != cave {
			return true
		}

		if smalls[cave] == 1 {
			return false
		}

		smalls[cave] = 1
		return true
	})

	return !foundDoubleSmall
}

// iterate traverses through the path ending at path, performing an action at each cave.
// it returns true if it was told to stop early
func iterate(path *pathNode, action func(cave string) bool) bool {
	for path != nil {
		cont := action(path.cave)
		if !cont {
			return true
		}
		path = path.prev
	}

	return false
}

func printCave(path *pathNode) {
	var s string
	iterate(path, func(cave string) bool {
		s = cave + " " + s

		return true
	})

	fmt.Println(s)
}

func main() {
	solve(strings.NewReader("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end"))
	solve(strings.NewReader("fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW"))
	solve(input())
}
