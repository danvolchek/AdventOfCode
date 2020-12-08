package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
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

func parse(r io.Reader) map[string]map[string]bool {
	scanner := bufio.NewScanner(r)

	rules := make(map[string]map[string]bool)

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

func parseRule(raw string) (string, map[string]bool) {
	ruleMatches := ruleRegex.FindStringSubmatch(raw)
	if ruleMatches == nil {
		panic(fmt.Sprintf("don't know how to parse rule %s", raw))
	}

	outer := ruleMatches[1]

	if ruleMatches[2] == noContents {
		return outer, nil
	}

	inner := make(map[string]bool)
	rawContents := strings.Split(ruleMatches[2], ",")

	for _, rawBag := range rawContents {
		innerBagMatches := innerBag.FindStringSubmatch(rawBag)
		if innerBagMatches == nil {
			panic(fmt.Sprintf("don't know how to parse inner bag %s", rawBag))
		}

		inner[innerBagMatches[2]] = true
	}

	return outer, inner
}

func solve(rules map[string]map[string]bool) int {
	sum := 0

	cache := make(map[string]bool)
	for color := range rules {
		if containsTargetBag(color, cache, rules) {
			sum += 1
		}
	}

	return sum
}

func containsTargetBag(color string, cache map[string]bool, rules map[string]map[string]bool) bool {
	if containsBag, ok := cache[color]; ok {
		return containsBag
	}

	result := func(result bool) bool {
		cache[color] = result
		return result
	}

	// search direct contents
	for newColor := range rules[color] {
		if newColor == targetBag {
			return result(true)
		}
	}

	// search indirect contents
	for newColor := range rules[color] {
		if containsTargetBag(newColor, cache, rules) {
			return result(true)
		}
	}

	return result(false)
}

func main() {
	fmt.Println(parseRule("muted white bags contain 3 muted tomato bags, 5 light black bags, 4 pale black bags, 5 shiny gold bags."))
	fmt.Println(parseRule("light black bags contain 1 striped yellow bag."))
	fmt.Println(parseRule("dotted tomato bags contain no other bags."))

	fmt.Println(solve(parse(input())))
}
