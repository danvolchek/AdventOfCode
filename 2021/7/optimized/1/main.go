package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "7", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) []int {
	line, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	var crabs []int

	rawCrabs := strings.Split(strings.TrimSpace(string(line)), ",")
	for _, rawCrab := range rawCrabs {
		crab, err := strconv.Atoi(rawCrab)
		if err != nil {
			panic(err)
		}

		crabs = append(crabs, crab)
	}

	return crabs
}

func solve(r io.Reader) {
	crabs := parse(r)

	sort.Slice(crabs, func(i, j int) bool {
		return crabs[i] < crabs[j]
	})

	// The median is the optimal position to move to. See proof below.
	optimalPosition := crabs[len(crabs)/2]

	minFuel := totalFuel(optimalPosition, crabs)

	fmt.Println(minFuel, optimalPosition)
}

func totalFuel(position int, crabs []int) int {
	fuel := 0
	for _, crab := range crabs {
		fuel += abs(crab - position)
	}

	return fuel
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}

	return v
}

func main() {
	solve(strings.NewReader("16,1,2,0,4,2,7,1,2,14"))
	solve(input())
}

/*
Why the median position is the one which would spend the least fuel:


For an odd numbered amount of crabs:

Consider the case of moving to the median position, m.
Exactly half the crabs are to the left (at position < m) and half the crabs are to the right (at position > m).

Now consider the case of moving to m-1, with left crabs being at position < m - 1 and right crabs being at position > m - 1.
Because m is the median, there are more right crabs than there are left crabs.

Each left crab has to spend 1 fuel less to get to m - 1 as compared to get to m.

Each right crab has to spend 1 fuel more to get to m - 1 as compared to get to m.

Because there are more right crabs than left crabs, the increase in distance needed to move by the right crabs outweighs
the decrease in distance needed to move by the left crabs.
Therefore, the net fuel used is more than moving to m.

The same logic holds for any where m - k: there will be more right crabs than left crabs, so the additional distance the right crabs have to move will dominate the left crabs.

The argument is the same for m + k: now there are more left crabs than right crabs.


For an even numbered amount of crabs:

Let m1 and m2 be the two middle values, where m1 and m2 where m2 > m1. Consider a point m3 where m3 = (m1 + m2) / 2.

Add this crab to the list between m1 and m2. The list has an odd number of crabs now, and m3 is the median. By the above
proof, m3 is the optimal position. Now consider

<not complete>
*/
