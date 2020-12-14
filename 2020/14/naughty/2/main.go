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

func addressToInt(addressSpace string) int {
	val := 0
	for i := 0; i < 36; i++ {
		if addressSpace[i] == '1' {
			val |= 1 << (36 - i - 1)
		}

	}

	return val
}

func intToAddress(memIndex int) string {
	val := ""

	for i := 0; i < 36; i++ {
		if (memIndex>>i)&1 == 1 {
			val = "1" + val
		} else {
			val = "0" + val
		}

	}

	return val
}

func setR(memory map[int]int, index int, addressSpace string, memValue int) {
	for i := index; i < 36; i++ {
		if addressSpace[i] != 'X' {
			continue
		}

		for _, val := range []string{"0", "1"} {
			newAddr := addressSpace[:i] + val + addressSpace[i+1:]
			//fmt.Println("recurse on ", newAddr)
			setR(memory, i+1, newAddr, memValue)
		}
		return

	}

	//fmt.Println("setting", addressSpace, addressToInt(addressSpace))

	addr := addressToInt(addressSpace)
	memory[addr] = memValue

}

func set(memory map[int]int, mask string, memIndex int, memValue int) {
	indexAsStr := intToAddress(memIndex)

	for i := 0; i < 36; i++ {
		if mask[i] == '0' {
			continue
		} else if mask[i] == '1' {
			indexAsStr = indexAsStr[:i] + "1" + indexAsStr[i+1:]
		} else {
			indexAsStr = indexAsStr[:i] + "X" + indexAsStr[i+1:]
		}
	}

	setR(memory, 0, indexAsStr, memValue)
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
	solve(strings.NewReader("mask = 000000000000000000000000000000X1001X\nmem[42] = 100\nmask = 00000000000000000000000000000000X0XX\nmem[26] = 1"))
	solve(input())
}
