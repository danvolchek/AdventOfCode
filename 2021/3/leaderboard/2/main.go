package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

const width = 12

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var vals []string
	for scanner.Scan() {
		row := scanner.Text()

		vals = append(vals, row)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	vals2 := make([]string, len(vals))
	copy(vals2, vals)

	pos := 0
	for len(vals) != 1 {
		common := count(vals)

		if common[pos] >= 0 {
			vals = discard(vals, pos, '0')
		} else {
			vals = discard(vals, pos, '1')
		}

		pos += 1
	}

	pos = 0
	for len(vals2) != 1 {
		common := count(vals2)

		if common[pos] < 0 {
			vals2 = discard(vals2, pos, '0')
		} else {
			vals2 = discard(vals2, pos, '1')
		}

		pos += 1
	}

	fmt.Println(vals, vals2, toInt(vals[0]), toInt(vals2[0]), toInt(vals[0])*toInt(vals2[0]))

}

func toInt(v string) int {
	sum := 0
	for i, vv := range v {
		if vv == '1' {
			sum += (1 << (width - i - 1))
		}
	}

	return sum
}

func discard(vals []string, i int, w uint8) []string {
	var ret []string
	for _, v := range vals {
		if v[i] != w {
			ret = append(ret, v)
		}
	}

	return ret
}

func count(nums []string) []int {
	var sum []int = make([]int, width)

	for _, row := range nums {
		for i, v := range row {
			if v == '0' {
				sum[i] -= 1
			} else {
				sum[i] += 1
			}
		}
	}

	return sum
}

func main() {
	//solve(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"))
	solve(input())
}
