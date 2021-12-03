package main

import (
	"bufio"
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

func solve(r io.Reader, width int) {
	scanner := bufio.NewScanner(r)

	var nums []string
	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	oxygenGeneratorBits := filterRepeatedly(nums, mostCommonBit)
	epsilonGeneratorBits := filterRepeatedly(nums, leastCommonBit)

	oxygenGeneratorRating := constructIntFromBits(oxygenGeneratorBits, width)
	epsilonGeneratorRating := constructIntFromBits(epsilonGeneratorBits, width)

	fmt.Println(oxygenGeneratorRating, epsilonGeneratorRating, oxygenGeneratorRating*epsilonGeneratorRating)

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
	solve(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"), 5)
	solve(input(), 12)
}
