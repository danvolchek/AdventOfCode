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
	for _, record := range records {
		for _, digit := range record.output {
			switch len(digit) {
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 7:
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	solve(strings.NewReader("be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\nedbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc\nfgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg\nfbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb\naecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea\nfgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb\ndbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe\nbdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef\negadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb\ngcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"))
	solve(input())
}
