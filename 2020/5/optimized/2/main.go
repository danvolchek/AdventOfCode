package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func parseFile() [][]byte {
	input, err := os.Open(path.Join("2020", "5", "input.txt"))
	if err != nil {
		panic(err)
	}

	return parse(input)
}

func parse(r io.Reader) [][]byte {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return bytes.Split(bytes.TrimSpace(raw), []byte{'\r', '\n'})
}

func decode(encoded []byte) int {
	decoded := 0

	for i := 0; i < len(encoded); i++ {
		if encoded[i] == 'B' || encoded[i] == 'R' {
			decoded |= 1 << (len(encoded) - i - 1)
		}
	}

	return decoded
}

func contains(m map[int]bool, val int) bool {
	_, ok := m[val]
	return ok
}

func solve(boardingPasses [][]byte) int {
	ids := make(map[int]bool, len(boardingPasses))

	for _, boardingPass := range boardingPasses {
		ids[decode(boardingPass)] = true
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
	fmt.Println(solve(parseFile()))
}
