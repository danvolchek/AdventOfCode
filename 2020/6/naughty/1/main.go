package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "6", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := bytes.Split(raw, []byte{'\r', '\n'})
	total := 0

	answered := make(map[byte]bool)
	for _, row := range rows {
		if len(row) == 0 {
			total += len(answered)
			answered = make(map[byte]bool)
			continue
		}

		for i := 0; i < len(row); i++ {
			answered[row[i]] = true
		}
	}
	total += len(answered)
	fmt.Println(total)

}

func main() {
	solve(input())
}
