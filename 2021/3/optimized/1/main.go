package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "3", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	raw, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	nums := strings.Split(string(bytes.TrimSpace(raw)), "\n")
	count := bitCount(nums)
	width := len(nums[0])

	gammaBits := bitsByIndexFunc(count, mostCommonBit, width)
	gammaRate := constructIntFromBits(gammaBits, width)

	// epsilon is always the negation of gamma because the least common bit
	// is by definition the opposite of the most common bit
	epsilonRate := constructIntFromBits(negate(gammaBits), width)

	fmt.Println(gammaRate, epsilonRate, gammaRate*epsilonRate)
}

func negate(bits []int) []int {
	result := make([]int, len(bits))
	for i, b := range bits {
		if b == 0 {
			result[i] = 1
		}
	}

	return result
}

func bitsByIndexFunc(count []int, bitAtIndex func(count int) int, width int) []int {
	result := make([]int, width)

	for i := 0; i < width; i++ {
		result[i] = bitAtIndex(count[i])
	}

	return result
}

func constructIntFromBits(bits []int, width int) int {
	value := 0
	for i := 0; i < width; i++ {
		value += bits[i] << (width - i - 1)
	}

	return value
}

// bitCount returns a slice of bits where the value at position i is
// - positive if the most common bit in nums at position i is 1
// - 0 if 0s and 1s are equally common in nums at position i
// - negative if the most common bit in nums at position i is 0
func bitCount(nums []string) []int {
	var count []int
	for i := 0; i < len(nums[0]); i++ {
		count = append(count, 0)
	}

	for _, num := range nums {
		for index, character := range num {
			if character == '0' {
				count[index] -= 1
			} else {
				count[index] += 1
			}
		}
	}

	return count
}

func mostCommonBit(count int) int {
	if count >= 0 {
		return 1
	}

	return 0
}

func main() {
	solve(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"))
	solve(input())
}
