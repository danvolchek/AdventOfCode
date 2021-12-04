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

const (
	divisor = 20201227
	subject = 7
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
		value %= divisor
	}

	return value
}

func unTransform(pk int) int {
	loopSize := 0

	calcKey := 1

	for {
		if calcKey == pk {
			return loopSize
		}

		calcKey *= subject
		calcKey %= divisor

		loopSize += 1
	}
}

func parse(r io.Reader) (int, int) {
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

	return toInt(chomp()), toInt(chomp())
}

func solve(pk1, pk2 int) int {
	loopSize2 := unTransform(pk2)

	return transform(pk1, loopSize2)
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("5764801\n17807724"))))
	fmt.Println(solve(parse(input())))
}
