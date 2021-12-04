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
	width := len(nums[0])

	oxygenGeneratorBits := filterRepeatedly(nums, mostCommonBit)
	co2ScrubberBits := filterRepeatedly(nums, leastCommonBit)

	oxygenGeneratorRating := constructIntFromBits(oxygenGeneratorBits, width)
	co2ScrubberRating := constructIntFromBits(co2ScrubberBits, width)

	fmt.Println(oxygenGeneratorRating, co2ScrubberRating, oxygenGeneratorRating*co2ScrubberRating)
}

func filterRepeatedly(values []string, bitAtIndex func(count int) byte) string {
	index := 0

	for len(values) != 1 {
		count := bitCount(values)

		values = filter(values, func(value string) bool {
			return value[index] == bitAtIndex(count[index])
		})

		index += 1
	}

	return values[0]
}

func filter(values []string, keep func(value string) bool) []string {
	var result []string
	for _, value := range values {
		if keep(value) {
			result = append(result, value)
		}
	}

	return result
}

func constructIntFromBits(num string, width int) int {
	value := 0
	for i, v := range num {
		bit := 0
		if v == '1' {
			bit = 1
		}
		value += bit << (width - i - 1)
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

func mostCommonBit(count int) byte {
	if count >= 0 {
		return '1'
	}

	return '0'
}

func leastCommonBit(count int) byte {
	if count < 0 {
		return '1'
	}

	return '0'
}

func main() {
	solve(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"))
	solve(input())
}
