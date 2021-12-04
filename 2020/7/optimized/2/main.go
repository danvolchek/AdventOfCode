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
	input, err := os.Open(path.Join("2020", "7", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

const (
	targetBag  = "shiny gold"
	noContents = "no other bags"
)

var (
	innerBag  = regexp.MustCompile(`(\d+?) (.+?) bags?`)
	ruleRegex = regexp.MustCompile(`^(.+?) bags contain (.+?)\.$`)
)

func parse(r io.Reader) map[string]map[string]int {
	scanner := bufio.NewScanner(r)

	rules := make(map[string]map[string]int)

	for scanner.Scan() {
		row := scanner.Text()

		outer, inner := parseRule(row)

		rules[outer] = inner
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rules
}

func parseRule(raw string) (string, map[string]int) {
	ruleMatches := ruleRegex.FindStringSubmatch(raw)
	if ruleMatches == nil {
		panic(fmt.Sprintf("don't know how to parse rule %s", raw))
	}

	outer := ruleMatches[1]

	if ruleMatches[2] == noContents {
		return outer, nil
	}

	inner := make(map[string]int)
	rawContents := strings.Split(ruleMatches[2], ",")

	for _, rawBag := range rawContents {
		innerBagMatches := innerBag.FindStringSubmatch(rawBag)
		if innerBagMatches == nil {
			panic(fmt.Sprintf("don't know how to parse inner bag %s", rawBag))
		}

		numInner, err := strconv.Atoi(innerBagMatches[1])
		if err != nil {
			panic(err)
		}

		inner[innerBagMatches[2]] = numInner
	}

	return outer, inner
}

func numContents(color string, cache map[string]int, rules map[string]map[string]int) int {
	if num, ok := cache[color]; ok {
		return num
	}

	sum := 1

	for newColor, amount := range rules[color] {
		sum += amount * numContents(newColor, cache, rules)
	}

	cache[color] = sum

	return sum
}

func solve(rules map[string]map[string]int) int {
	return numContents(targetBag, make(map[string]int), rules) - 1
}

func main() {
	fmt.Println(solve(parse(input())))
}
