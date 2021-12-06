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

	fish := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		fishStr := strings.Split(line, ",")
		for _, f := range fishStr {
			v, err := strconv.Atoi(f)
			if err != nil {
				panic(err)
			}

			fish[v] += 1
		}
	}

	//fmt.Println(fish)
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	for i := 0; i < 256; i++ {

		fish[9] = fish[0]
		fish[7] += fish[0]

		for i := 0; i < 9; i++ {
			fish[i] = fish[i+1]
		}
		//fmt.Println(fish)
	}

	sum := 0

	for i := 0; i < 9; i++ {
		sum += fish[i]
	}

	fmt.Println(sum)

}

func main() {
	solve(strings.NewReader("3,4,3,1,2"))
	solve(input())
}
