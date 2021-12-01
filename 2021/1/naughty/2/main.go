package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	items := make([]int, 3)
	which := 0

	inc := 0

	i := 0
	for scanner.Scan() {
		row := scanner.Text()

		v, err := strconv.ParseInt(row, 10, 32)
		vv := int(v)
		if err != nil {
			panic(err)
		}

		old := sum(items)

		items[which] = vv
		which = (which + 1) % len(items)

		if i > 2 {
			if sum(items) > old {
				inc++
			}
		}

		fmt.Println(items)
		i++
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(inc)

}

func sum(v []int) int {
	var s int
	for _, vv := range v {
		s += vv
	}

	return s
}

func main() {
	solve(input())
}
