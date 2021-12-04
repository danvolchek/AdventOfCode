package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func solve(r io.Reader) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(raw), "\r\n")

	valid := 0

	curr := make(map[string]bool)

	for _, row := range rows {
		if len(row) == 0 {
			_, hasCid := curr["cid"]
			if len(curr) == 8 || (len(curr) == 7 && !hasCid) {
				valid += 1
			}

			curr = make(map[string]bool)
			continue
		}

		items := strings.Split(row, " ")
		for _, item := range items {
			parts := strings.Split(item, ":")

			curr[parts[0]] = true
		}
	}

	fmt.Println(valid)
}

func main() {
	input, err := os.Open(path.Join("2020", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	solve(input)
}
