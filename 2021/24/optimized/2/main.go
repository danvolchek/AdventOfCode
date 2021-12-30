package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "24", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

// the solution for this day is split up into two files:
//  - solver.go: contains a solver to solve the puzzle
//  - alu.go: an alu implementation to check the solver's correctness
// the bulk of the code is found in those two files

func solve(r io.Reader) {
	buf := &bytes.Buffer{}
	tee := io.TeeReader(r, buf)

	operations, blocks := parseOperations(tee), parseBlocks(buf)

	validInput := findValidInput(blocks)
	checkValidInput(operations, validInput)

	fmt.Println(combineIntSlice(validInput))
}

func findValidInput(blocks []block) []int {
	fmt.Println("Finding valid input ...")

	validInputs := findValidInputs(blocks)
	sort.Slice(validInputs, func(i, j int) bool {
		return less(validInputs[i], validInputs[j])
	})

	return validInputs[0]
}

func checkValidInput(operations []operation, validInput []int) {
	fmt.Println("Checking input ...")
	a := &alu{
		inputs: validInput,
	}

	a.run(operations)

	if a.z != 0 {
		panic("incorrect solution")
	}
}

func less(a, b []int) bool {
	for index := 0; index < len(a); index++ {
		if a[index] == b[index] {
			continue
		}

		return a[index] < b[index]
	}

	return false
}

func main() {
	solve(input())
}
