package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/danvolchek/AdventOfCode/lib"
)

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
	solver := lib.Solver[[]byte, int]{
		ParseF: lib.ParseBytes,
		SolveF: solve,
	}

	solver.Expect("abcdef", 609043)
	solver.Expect("pqrstuv", 1048970)
	solver.Verify(117946)
}
