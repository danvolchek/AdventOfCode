package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(prefix []byte) int {
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
			return num
		}

		hash.Reset()
		num += 1
	}
}

func main() {
	lib.TestSolveBytes("abcdef", solve)
	lib.SolveBytes(input(), solve)
}
