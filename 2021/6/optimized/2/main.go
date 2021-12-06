package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "6", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

const days = 256

func parse(r io.Reader) map[int]int {
	line, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rawFish := strings.Split(strings.TrimSpace(string(line)), ",")

	fish := make(map[int]int)

	for _, f := range rawFish {
		timer, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}

		fish[timer] += 1
	}

	return fish
}

func solve(r io.Reader) {
	fish := parse(r)

	for i := 0; i < days; i++ {
		fishToCreate := fish[0]

		// each fish timer ticks down one day
		for timer := 0; timer < 8; timer++ {
			fish[timer] = fish[timer+1]
		}

		// each new fish starts out with timer 8
		fish[8] = fishToCreate

		// each fish that created a fish resets its timer to 6
		fish[6] += fishToCreate
	}

	totalFish := 0
	for _, numFish := range fish {
		totalFish += numFish
	}

	fmt.Println(totalFish)
}

func main() {
	solve(strings.NewReader("3,4,3,1,2"))
	solve(input())
}
