package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "14", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

var reg = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

func set(memory map[int]int, mask string, memIndex int, memValue int) {
	val := 0

	for i := 0; i < 36; i++ {
		if mask[36-i-1] != 'X' {
			if mask[36-i-1] == '1' {
				val |= 1 << i
			}
			continue
		}

		val |= ((memValue >> i) & 1) << i
	}

	memory[memIndex] = val
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var mask string

	memory := make(map[int]int)

	for scanner.Scan() {
		row := scanner.Text()

		if strings.Index(row, "mask") == 0 {
			mask = strings.Split(row, " = ")[1]
			continue
		} else {
			parts := reg.FindStringSubmatch(row)

			index, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			value, err := strconv.Atoi(parts[2])
			if err != nil {
				panic(err)
			}

			set(memory, mask, index, value)

		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	sum := 0

	for _, val := range memory {
		sum += val
	}

	fmt.Println(sum)
}

func main() {
	solve(strings.NewReader("mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X\nmem[8] = 11\nmem[7] = 101\nmem[8] = 0"))
	solve(input())
}
