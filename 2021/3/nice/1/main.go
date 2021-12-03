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

	count := bitCount(nums)

	gammaRate := constructIntFromBits(bitsByIndexFunc(count, mostCommonBit, width), width)
	epsilonRate := constructIntFromBits(bitsByIndexFunc(count, leastCommonBit, width), width)

	fmt.Println(gammaRate, epsilonRate, gammaRate*epsilonRate)

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

func leastCommonBit(count int) int {
	if count < 0 {
		return 1
	}

	return 0
}

func main() {
	solve(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"), 5)
	solve(input(), 12)
}
