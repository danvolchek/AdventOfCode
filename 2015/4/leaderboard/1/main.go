package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	prefix := lib.Must(io.ReadAll(r))

	hash := md5.New()
	num := 1
	for {
		hash.Write(prefix)
		hash.Write([]byte(strconv.Itoa(num)))

		hashVal := hex.EncodeToString(hash.Sum(hash.Sum(nil)))

		found := true
		for index := 0; index < 5; index += 1 {
			if hashVal[index] != '0' {
				found = false
				break
			}
		}

		if found {
			fmt.Println(num)
			return
		}

		hash.Reset()
		num += 1
	}

}

func main() {
	solve(strings.NewReader("abcdef"))
	solve(input())
}
