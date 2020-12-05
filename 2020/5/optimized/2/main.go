package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

type boardingPass struct {
	row []bool
	col []bool
}

func (b boardingPass) id() int {
	return decode(b.row)*8 + decode(b.col)
}

func decode(indicator []bool) int {
	value := 0
	step := 1 << (len(indicator) - 1)

	for _, frontHalf := range indicator {
		if !frontHalf {
			value += step
		}

		step /= 2
	}

	return value
}

func elementsMatch(input []byte, target byte) []bool {
	result := make([]bool, len(input))

	for i := 0; i < len(input); i++ {
		result[i] = input[i] == target
	}

	return result
}

func parse(r io.Reader) []boardingPass {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := bytes.Split(bytes.TrimSpace(raw), []byte{'\r', '\n'})

	boardingPasses := make([]boardingPass, len(rows))
	for i, row := range rows {
		boardingPasses[i] = boardingPass{
			row: elementsMatch(row[:7], 'F'),
			col: elementsMatch(row[7:], 'L'),
		}
	}

	return boardingPasses
}

func contains(m map[int]bool, val int) bool {
	_, ok := m[val]
	return ok
}

func solve(boardingPasses []boardingPass) int {
	ids := make(map[int]bool, len(boardingPasses))

	for _, boardingPass := range boardingPasses {
		ids[boardingPass.id()] = true
	}

	for i := range ids {
		if !contains(ids, i-1) && contains(ids, i-2) {
			return i - 1
		}

		if !contains(ids, i+1) && contains(ids, i+2) {
			return i + 1
		}
	}

	panic("not found")
}

func main() {
	input, err := os.Open(path.Join("2020", "5", "input.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(parse(input)))
}
