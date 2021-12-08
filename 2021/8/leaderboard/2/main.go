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

func solve(r io.Reader) {
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

	sum := 0
	now := time.Now()
	for i, record := range records {
		fmt.Println(i, time.Now().Sub(now))
		mapping := findSignalSegmentMapping(record.patterns)

		sum += translateDigits(record.output, mapping)
	}

	fmt.Println(sum)
}

var segments = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

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

func areAllDigitsValid(digits []string, mapping map[byte]byte) bool {
	for _, digit := range digits {
		_, ok := translateDigit(digit, mapping)
		if !ok {
			return false
		}
	}

	return true
}

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
