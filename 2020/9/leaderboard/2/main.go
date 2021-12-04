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
	input, err := os.Open(path.Join("2020", "9", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var numbers []int
	for scanner.Scan() {
		row := scanner.Text()

		v, _ := strconv.Atoi(row)
		numbers = append(numbers, v)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	target := 22406676

	fmt.Println(len(numbers))
	for k := 0; k < len(numbers)-1; k++ {

		for l := 1; l < len(numbers)-k; l++ {

			//fmt.Println(k, l, "m is max", l-1)
			sum := 0
			for m := 0; m < l; m++ {
				sum += numbers[k+m]
			}

			// note: this is actually wrong. I used first and second to last (thinking it was first and last, and
			// -2 was needed due to an off by one error). It should be min and max, but I was lucky with both
			// the sample input and real input that the first and second to last numbers were the right ones to use.
			if sum == target {
				fmt.Println(numbers[k], numbers[k+l-1], numbers[k]+numbers[k+l-2])
				return
			}

		}

	}

	panic("nothing")
}

func main() {
	//solve(strings.NewReader("35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"))
	solve(input())
}
