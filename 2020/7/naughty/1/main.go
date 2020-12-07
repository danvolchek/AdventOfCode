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

var innerRegex = regexp.MustCompile(`(\d+?) (.+?) bags?`)

var ruleRegex = regexp.MustCompile(`^(.+?) bags contain (.+?).$`)

var noneRegex = regexp.MustCompile(`(.+) bags contain no other bags.`)

type rule struct {
	outer string
	inner map[string]int
}

func parseRule(def string) rule {
	if noneRegex.MatchString(def) {
		matches := noneRegex.FindAllStringSubmatch(def, -1)

		return rule{
			outer: matches[0][1],
		}
	}

	matches := ruleRegex.FindAllStringSubmatch(def, -1)

	rule := rule{
		inner: make(map[string]int),
	}

	rule.outer = matches[0][1]

	rawInner := strings.Split(matches[0][2], ",")

	for _, item := range rawInner {
		items := innerRegex.FindStringSubmatch(item)

		numInner, _ := strconv.Atoi(items[1])
		rule.inner[items[2]] = numInner
	}

	return rule

}

func containsShinyGold(color string, rules map[string]map[string]int) bool {

	contents := rules[color]

	for newColor, _ := range contents {
		if newColor == "shiny gold" {
			return true
		}

		if containsShinyGold(newColor, rules) {
			return true
		}
	}

	return false
}

func solve(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	rules := make(map[string]map[string]int)

	for scanner.Scan() {
		row := scanner.Text()

		parsed := parseRule(row)

		rules[parsed.outer] = parsed.inner
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	hasShiny := make(map[string]bool)

	for color := range rules {
		if containsShinyGold(color, rules) {
			hasShiny[color] = true
		}
	}

	return len(hasShiny)
}

func main() {
	fmt.Println(parseRule("muted white bags contain 3 muted tomato bags, 5 light black bags, 4 pale black bags, 5 shiny gold bags."))
	fmt.Println(parseRule("light black bags contain 1 striped yellow bag."))
	fmt.Println(parseRule("dotted tomato bags contain no other bags."))

	fmt.Println(solve(input()))

}
