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
	input, err := os.Open(path.Join("2020", "25", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func transform(subject, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i += 1 {
		value *= subject
		value %= 20201227
	}

	return value
}

func unTransform(subject, pk int) int {
	// find l such that 7 * l % 20201227
	loopSize := 1

	calcKey := 1

	for {
		calcKey *= subject
		calcKey %= 20201227

		if calcKey == pk {
			return loopSize
		}

		loopSize += 1
	}
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	chomp := func() string {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
		}
		return scanner.Text()
	}
	toInt := func(val string) int {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return intVal
	}

	pk1 := toInt(chomp())
	pk2 := toInt(chomp())

	//loopSize1 := unTransform(7, pk1)
	loopSize2 := unTransform(7, pk2)

	fmt.Println("loop:", loopSize2)

	encryptionKey := transform(pk1, loopSize2)

	fmt.Println(encryptionKey)
}

func main() {
	solve(strings.NewReader("5764801\n17807724"))
	solve(input())
}
