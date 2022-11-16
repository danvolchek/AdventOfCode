package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"regexp"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "9", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type distance struct {
	start, end string
	length     int
}

var parseRegExp = regexp.MustCompile(`(.*) to (.*) = (\d+)`)

func parse(parts []string) distance {
	return distance{
		start:  parts[0],
		end:    parts[1],
		length: lib.Atoi(parts[2]),
	}
}

// note: this is the travelling salesman problem. There are more efficient solutions for this problem;
// this algorithm is not that - it tries every possible combination without any caching.
func solve(instructions []distance) int {
	distances := map[string]map[string]int{}
	for _, instr := range instructions {
		if _, ok := distances[instr.start]; !ok {
			distances[instr.start] = map[string]int{}
		}
		distances[instr.start][instr.end] = instr.length

		if _, ok := distances[instr.end]; !ok {
			distances[instr.end] = map[string]int{}
		}
		distances[instr.end][instr.start] = instr.length
	}

	calcDistance := func(path []string) int {
		total := 0
		for i := 0; i < len(distances)-1; i++ {
			total += distances[path[i]][path[i+1]]
		}

		return total
	}

	locations := lib.Keys(distances)

	paths := lib.Permutations(locations)

	totalDistances := lib.Map(paths, calcDistance)

	return lib.Max(totalDistances...)
}

func main() {
	solver := lib.Solver[[]distance, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseRegExp, parse)),
		SolveF: solve,
	}

	solver.Expect("London to Dublin = 464\nLondon to Belfast = 518\nDublin to Belfast = 141", 982)
	solver.Verify(input(), 909)
}
