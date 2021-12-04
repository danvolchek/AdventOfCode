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

func numContents(color string, rules map[string]map[string]int) int {

	contents := rules[color]

	sum := 1

	for newColor, amount := range contents {
		sum += amount * numContents(newColor, rules)
	}

	return sum
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

	return numContents("shiny gold", rules) - 1
}

func main() {
	fmt.Println(solve(input()))
}
