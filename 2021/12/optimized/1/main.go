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
	value string
	prev  *pathNode
}

// explore returns the number of paths from currentCave to "end", given an already travelled path of path and caves for
// valid transitions between caves according to the rules of the part.
func explore(caves map[string][]string, currentCave string, path *pathNode) int {
	if currentCave == "end" {
		return 1
	}

	currentPath := &pathNode{
		value: currentCave,
		prev:  path,
	}
	paths := 0

	for _, nextCave := range caves[currentCave] {
		// if we've done this transition before, don't do it again to avoid infinite loops
		// if we're visiting a small cave again, don't - that's not allowed
		if transitionedBefore(currentCave, nextCave, currentPath) || visitingSmallAgain(nextCave, currentPath) {
			continue
		}

		paths += explore(caves, nextCave, currentPath)
	}

	return paths
}

func transitionedBefore(from, to string, path *pathNode) bool {
	var seen bool

	lastCave := ""
	seenTwice := iterate(path, func(cave string) bool {
		if cave == from && lastCave == to {
			if seen {
				return false
			}

			seen = true
		}

		lastCave = cave
		return true
	})

	return seenTwice
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

// iterate traverses through the path ending at path, performing an action at each cave.
// it returns true if it was told to stop early
func iterate(path *pathNode, action func(cave string) bool) bool {
	for path != nil {
		cont := action(path.value)
		if !cont {
			return true
		}
		path = path.prev
	}

	return false
}

func main() {
	solve(strings.NewReader("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end"))
	solve(strings.NewReader("fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW"))
	solve(input())
}
