package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func decode(encoded []byte) int {
	decoded := 0

	for i := 0; i < len(encoded); i++ {
		if encoded[i] == 'B' || encoded[i] == 'R' {
			decoded |= 1 << (len(encoded) - i - 1)
		}
	}

	return decoded
}

func parse(r io.Reader) [][]byte {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return bytes.Split(bytes.TrimSpace(raw), []byte{'\r', '\n'})
}

func solve(boardingPasses [][]byte) int {
	maxId := 0

	for _, boardingPass := range boardingPasses {
		id := decode(boardingPass)

		if id > maxId {
			maxId = id
		}
	}

	return maxId
}

func main() {
	input, err := os.Open(path.Join("2020", "5", "input.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(solve(parse(input)))
}
