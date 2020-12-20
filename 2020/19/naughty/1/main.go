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
	input, err := os.Open(path.Join("2020", "19", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type constMatcher struct {
	value string
}

func (c constMatcher) matches(value string) (bool, int) {
	if len(value) < len(c.value) {
		return false, 0
	}

	return c.value == value[:len(c.value)], len(c.value)
}

type referenceMatcher struct {
	name string
	rules map[string]rule
}

func (r referenceMatcher) matches(value string) (bool, int) {
	return r.rules[r.name].matches(value)
}

type sequentialMatcher struct {
	rules []rule
}


func (s sequentialMatcher) matches(value string) (bool, int) {
	// how much to first/second?

	currLen := 0
	for _, rule := range s.rules {
		matches, length := rule.matches(value[currLen:])
		if !matches {
			return false, 0
		}

		currLen += length
	}

	return true, currLen
}

type orMatcher struct {
	first rule
	second rule
}

func (o orMatcher) matches(value string) (bool, int) {
	matches, length := o.first.matches(value)
	if matches {
		return true, length
	}

	return o.second.matches(value)
}

type rule interface {
	matches(value string) (bool, int)
}

func parseRuleWithoutPipe(raw string, rules map[string]rule) rule {
	if raw[0] == '"' {
		return constMatcher{value: raw[1:2]}
	}

	if strings.Index(raw, " ") == -1 {
		return referenceMatcher{
			name:  raw,
			rules: rules,
		}
	}

	parts := strings.Split(raw, " ")

	components := make([]rule, len(parts))

	for i, part := range parts {
		components[i] = referenceMatcher{
			name:  part,
			rules: rules,
		}
	}

	return sequentialMatcher{
		rules: components,
	}
}

func parseRule(raw string, rules map[string]rule) (string, rule) {
	parts := strings.Split(raw, ": ")

	name := parts[0]
	body := parts[1]

	var rule rule
	if strings.Index(body, "|") != -1 {

		parts := strings.Split(body, " | ")

		rule = orMatcher{
			first: parseRuleWithoutPipe(parts[0], rules),
			second: parseRuleWithoutPipe(parts[1], rules),
		}

	} else {
		rule = parseRuleWithoutPipe(body, rules)
	}

	return name, rule
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	rules := make(map[string]rule)
	var inputs []string

	zone := 0
	for scanner.Scan() {
		row := scanner.Text()

		if len(row) == 0 {
			zone++
			continue
		}

		switch zone {
		case 0:
			name, rule  := parseRule(row, rules)
			rules[name] = rule
		case 1:
			inputs = append(inputs, row)
		default:
			panic(zone)
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}


	sum := 0
	for _, input := range inputs {
		matches, length := rules["0"].matches(input)

		//fmt.Println(input, matches && length == len(input))
		if matches && length == len(input) {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func main() {
	solve(strings.NewReader("0: 1 2\n1: \"a\"\n2: 1 3 | 3 1\n3: \"b\"\n\naab\naba"))
	solve(strings.NewReader("0: 4 1 5\n1: 2 3 | 3 2\n2: 4 4 | 5 5\n3: 4 5 | 5 4\n4: \"a\"\n5: \"b\"\n\nababbb\nbababa\nabbbab\naaabbb\naaaabbb"))
	solve(input())
}
