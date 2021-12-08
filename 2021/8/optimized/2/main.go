package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "8", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type record struct {
	patterns []string
	output   []string
}

// map alphabetically sorted segment display to the number it represents
var segmentsToDigit = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

// alphabetical list of segments
var segments = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

func parse(r io.Reader) []record {
	scanner := bufio.NewScanner(r)

	var records []record
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " | ")

		records = append(records, record{
			patterns: strings.Split(parts[0], " "),
			output:   strings.Split(parts[1], " "),
		})
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return records
}

func solve(r io.Reader) {
	records := parse(r)

	sum := 0
	now := time.Now()
	for i, record := range records {
		nowNow := time.Now()
		mapping := findSignalSegmentMapping(record.patterns)

		sum += translateDigits(record.output, mapping)
		fmt.Println(i, time.Now().Sub(nowNow))
	}

	fmt.Println(sum, time.Now().Sub(now))
}

// findSignalSegmentMapping finds the mapping which satisfies all digits in patterns, or in other words
// the mapping which results in all translated patterns being valid digits.
// It does so by trying every single mapping. This is the place to optimize.
func findSignalSegmentMapping(patterns []string) map[byte]byte {
	mapping := map[byte]byte{
		'a': 'a',
		'b': 'a',
		'c': 'a',
		'd': 'a',
		'e': 'a',
		'f': 'a',
		'g': 'a',
	}

	for !areAllDigitsValid(patterns, mapping) {
		index := 0
		for mapping[segments[index]] == 'g' {
			mapping[segments[index]] = 'a'
			index++
		}

		mapping[segments[index]]++
	}

	return mapping
}

// translateDigits translates every digit in mapping to an integer, and then assembles the digits into a single integer,
// treading digits as a single seven-segment display.
func translateDigits(digits []string, mapping map[byte]byte) int {
	sum := 0
	for i, digit := range digits {
		v, ok := translateDigit(digit, mapping)
		if !ok {
			panic("should be a valid mapping")
		}

		sum += int(math.Pow10(len(digits)-i-1)) * v
	}

	return sum
}

// areAllDigitsValid translates every digit in digits using mapping. If all result in a digit, it returns true.
func areAllDigitsValid(digits []string, mapping map[byte]byte) bool {
	for _, digit := range digits {
		_, ok := translateDigit(digit, mapping)
		if !ok {
			return false
		}
	}

	return true
}

// translateDigit takes a digit composed of mixed up segments and returns the integer digit it represents.
// It does so by passing digit through mapping to un-scramble the segments and then comparing against segmentsToDigit.
// If mapping is invalid, and results in  the second return value will be false.
func translateDigit(digit string, mapping map[byte]byte) (int, bool) {
	bytes := make([]byte, len(digit))
	for i := 0; i < len(digit); i++ {
		bytes[i] = mapping[digit[i]]
	}

	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	val, ok := segmentsToDigit[string(bytes)]
	return val, ok
}

func main() {
	solve(strings.NewReader("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\nedbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc\nfgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg\nfbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb\naecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea\nfgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb\ndbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe\nbdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef\negadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb\ngcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"))
	solve(input())
}
