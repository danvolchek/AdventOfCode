package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
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

func (c constMatcher) matches(value string) []matchResult {
	if len(value) < len(c.value) || c.value != value[:len(c.value)]{
		return []matchResult{}
	}

	return []matchResult{{
		matchLength: len(c.value),
	}}
}

type referenceMatcher struct {
	name  string
	rules map[string]rule
}

func (r referenceMatcher) matches(value string) []matchResult {
	return r.rules[r.name].matches(value)
}

type sequentialMatcher struct {
	rules []rule
}

func (s sequentialMatcher) matches(value string) []matchResult {
	var newResults []matchResult

	lastResults := s.rules[0].matches(value)


	for i:= 1; i < len(s.rules) && len(lastResults) > 0; i++{

		var thisIterationResults []matchResult
		for _, lastResult := range lastResults {
			results := s.rules[i].matches(value[lastResult.matchLength:])

			for _, result := range results {
				fixed := matchResult{matchLength: lastResult.matchLength + result.matchLength}

				thisIterationResults = append(thisIterationResults, fixed)
			}


		}

		lastResults = thisIterationResults

	}

	newResults = append(newResults, lastResults...)

	return newResults
}

type orMatcher struct {
	first  rule
	second rule
}

func (o orMatcher) matches(value string) []matchResult {
	results := o.first.matches(value)

	results = append(results, o.second.matches(value)...)

	return results
}

type matchResult struct {
	matchLength int
}

type rule interface {
	matches(value string) []matchResult
}

type repeaterMatcher struct {
	repeated rule
}

func (r repeaterMatcher) matches(value string) []matchResult {
	newResults := make(map[matchResult]struct{})


	currentRule := sequentialMatcher{rules: []rule{r.repeated}}

	for {
		matchResults := currentRule.matches(value)

		if len(matchResults) == 0 {
			break
		}

		for _, result := range matchResults {
			newResults[result] = struct{}{}
		}

		currentRule = sequentialMatcher{rules: append(currentRule.rules, r.repeated)}
	}

	container := make([]matchResult, len(newResults))

	i := 0
	for newResult := range newResults {
		container[i] = newResult
		i+=1
	}
	return container
}

type sandwichRepeaterMatcher struct {
	left rule
	right rule
}

func (s sandwichRepeaterMatcher) matches(value string) []matchResult {
	newResults := make(map[matchResult]struct{})


	currentRule := sequentialMatcher{rules: []rule{s.left, s.right}}

	for i := 0; i < 4; i++{
		matchResults := currentRule.matches(value)

		newRules := []rule{s.left}
		newRules = append(newRules, currentRule.rules...)
		newRules = append(newRules, s.right)

		currentRule = sequentialMatcher{rules: newRules}

		for _, result := range matchResults {
			newResults[result] = struct{}{}
		}
	}

	container := make([]matchResult, len(newResults))

	i := 0
	for newResult := range newResults {
		container[i] = newResult
		i+=1
	}
	return container
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

func parseRuleWithoutName(raw string, rules map[string]rule) rule {
	var rule rule
	if strings.Index(raw, "|") != -1 {

		parts := strings.Split(raw, " | ")

		rule = orMatcher{
			first:  parseRuleWithoutPipe(parts[0], rules),
			second: parseRuleWithoutPipe(parts[1], rules),
		}

	} else {
		rule = parseRuleWithoutPipe(raw, rules)
	}

	return rule
}

func parseRule(raw string, rules map[string]rule) (string, rule) {
	parts := strings.Split(raw, ": ")

	return parts[0], parseRuleWithoutName(parts[1], rules)
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
			name, rule := parseRule(row, rules)
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

	rules["8"] = repeaterMatcher{repeated: parseRuleWithoutName("42", rules)}

	rules["11"] = sandwichRepeaterMatcher{
		left: parseRuleWithoutName("42", rules),
		right: parseRuleWithoutName("31", rules),
	}

	a := time.Now()
	sum := 0
	for _, input := range inputs {
		results := rules["0"].matches(input)

		matched := false
		for _, result := range results {
			if result.matchLength == len(input) {
				matched = true
				break
			}
		}

		//fmt.Println(input, matched)
		if matched {
			sum += 1
		}
	}
	fmt.Println(sum, time.Now().Sub(a))
}

func main() {
	solve(strings.NewReader("0: 8 45\n42: \"a\"\n45: \"b\"\n\nab\naab\naaab\naaaab"))
	solve(strings.NewReader("0: 11 31\n42: \"a\"\n31: \"b\"\n\nabb\naabbb\naaabbbb\naaaabbbbb"))
	solve(strings.NewReader("0: 1 3 1 3\n1: \"a\"\n2: \"b\"\n3: 1 2\n\naabaab"))
	solve(strings.NewReader("0: 1 2\n1: \"a\"\n2: 1 3 | 3 1\n3: \"b\"\n\naab\naba"))
	solve(strings.NewReader("0: 4 1 5\n1: 2 3 | 3 2\n2: 4 4 | 5 5\n3: 4 5 | 5 4\n4: \"a\"\n5: \"b\"\n\nababbb\nbababa\nabbbab\naaabbb\naaaabbb"))
	solve(strings.NewReader("42: 9 14 | 10 1\n9: 14 27 | 1 26\n10: 23 14 | 28 1\n1: \"a\"\n11: 42 31\n5: 1 14 | 15 1\n19: 14 1 | 14 14\n12: 24 14 | 19 1\n16: 15 1 | 14 14\n31: 14 17 | 1 13\n6: 14 14 | 1 14\n2: 1 24 | 14 4\n0: 8 11\n13: 14 3 | 1 12\n15: 1 | 14\n17: 14 2 | 1 7\n23: 25 1 | 22 14\n28: 16 1\n4: 1 1\n20: 14 14 | 1 15\n3: 5 14 | 16 1\n27: 1 6 | 14 18\n14: \"b\"\n21: 14 1 | 1 14\n25: 1 1 | 1 14\n22: 14 14\n8: 42\n26: 14 22 | 1 20\n18: 15 15\n7: 14 5 | 1 21\n24: 14 1\n\nabbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa\nbbabbbbaabaabba\nbabbbbaabbbbbabbbbbbaabaaabaaa\naaabbbbbbaaaabaababaabababbabaaabbababababaaa\nbbbbbbbaaaabbbbaaabbabaaa\nbbbababbbbaaaaaaaabbababaaababaabab\nababaaaaaabaaab\nababaaaaabbbaba\nbaabbaaaabbaaaababbaababb\nabbbbabbbbaaaababbbbbbaaaababb\naaaaabbaabaaaaababaa\naaaabbaaaabbaaa\naaaabbaabbaaaaaaabbbabbbaaabbaabaaa\nbabaaabbbaaabaababbaabababaaab\naabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba"))
	solve(input())
}
