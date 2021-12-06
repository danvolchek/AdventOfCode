package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "6", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var fish []int
	for scanner.Scan() {
		line := scanner.Text()

		fishStr := strings.Split(line, ",")
		for _, f := range fishStr {
			v, err := strconv.Atoi(f)
			if err != nil {
				panic(err)
			}

			fish = append(fish, v)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	for i := 0; i < 80; i++ {
		var additions []int

		for i, f := range fish {
			if f != 0 {
				fish[i]--
			} else {
				fish[i] = 6
				additions = append(additions, 8)
			}
		}

		for _, add := range additions {
			fish = append(fish, add)
		}
	}

	fmt.Println(len(fish))

}

func main() {
	solve(strings.NewReader("3,4,3,1,2"))
	solve(input())
}
