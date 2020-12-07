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
	targetBag = "shiny gold"
)

var (
	innerBag      = regexp.MustCompile(`(\d+?) (.+?) bags?`)
	hasNoContents = regexp.MustCompile(`^(.+) bags contain no other bags.$`)
	hasContents   = regexp.MustCompile(`^(.+?) bags contain (.+?).$`)
)

func parseRule(raw string) (string, map[string]int) {
	noContents := hasNoContents.FindStringSubmatch(raw)
	if noContents != nil {
		return noContents[1], nil
	}

	contents := hasContents.FindStringSubmatch(raw)
	if contents == nil {
		panic(fmt.Sprintf("don't know how to parse %s", raw))
	}

	inner := make(map[string]int)

	rawContents := strings.Split(contents[2], ",")

	for _, rawBag := range rawContents {
		innerBag := innerBag.FindStringSubmatch(rawBag)

		numInner, err := strconv.Atoi(innerBag[1])
		if err != nil {
			panic(err)
		}
		inner[innerBag[2]] = numInner
	}

	return contents[1], inner

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

func solve(rules map[string]map[string]int) int {
	return numContents(targetBag, make(map[string]int), rules) - 1
}

func main() {
	fmt.Println(solve(parse(input())))
}
