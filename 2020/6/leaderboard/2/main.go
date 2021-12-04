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

	counter := 0
	answered := make(map[byte]int)
	for _, row := range rows {
		if len(row) == 0 {
			for _, v := range answered {
				if v == counter {
					total += 1
				}
			}
			counter = 0
			answered = make(map[byte]int)
			continue
		}

		counter += 1
		for i := 0; i < len(row); i++ {
			answered[row[i]] += 1
		}
	}
	for _, v := range answered {
		if v == counter {
			total += 1
		}
	}
	fmt.Println(total)

}

func main() {
	solve(input())
}
