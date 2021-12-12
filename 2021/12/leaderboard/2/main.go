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

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	caves := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "-")
		a := parts[0]
		b := parts[1]

		caves[a] = append(caves[a], b)
		caves[b] = append(caves[b], a)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	t := time.Now()
	result := explore(caves, "start", nil)
	fmt.Println(len(result))
	fmt.Println(time.Now().Sub(t))

	//for _, path := range result {
	//	fmt.Println(path)
	//}
}

func explore(caves map[string][]string, current string, path []string) [][]string {
	next := append(clone(path), current)
	if current == "end" {
		return [][]string{next}
	}

	var result [][]string
	for _, cave := range caves[current] {
		if cave == "start" || (!allSmallsExactlyOnce(next) && (transitionedBefore(current, cave, next) || visitingSmallAgain(cave, next))) {
			continue
		}

		for _, newPath := range explore(caves, cave, next) {
			result = append(result, newPath)
		}
	}

	return result
}

func transitionedBefore(from, to string, path []string) bool {
	seen := 0

	for i := 0; i < len(path)-1; i++ {
		if path[i] == from && path[i+1] == to {
			seen += 1
		}
	}

	return seen > 1
}

func visitingSmallAgain(cave string, path []string) bool {
	if strings.ToLower(cave) != cave {
		return false
	}

	for _, visitedCave := range path {
		if visitedCave == cave {
			return true
		}
	}

	return false
}

func allSmallsExactlyOnce(path []string) bool {
	smalls := make(map[string]int)
	for _, cave := range path {
		if strings.ToLower(cave) != cave {
			continue
		}

		smalls[cave] += 1
	}

	for _, amount := range smalls {
		if amount != 1 {
			return false
		}
	}

	return true
}

func clone(path []string) []string {
	ret := make([]string, len(path))
	copy(ret, path)
	return ret
}

func main() {
	solve(strings.NewReader("start-A\nstart-b\nA-c\nA-b\nb-d\nA-end\nb-end"))
	solve(strings.NewReader("fs-end\nhe-DX\nfs-he\nstart-DX\npj-DX\nend-zg\nzg-sl\nzg-pj\npj-he\nRW-he\nfs-DX\npj-RW\nzg-RW\nstart-pj\nhe-WI\nzg-he\npj-fs\nstart-RW"))
	solve(input())
}
