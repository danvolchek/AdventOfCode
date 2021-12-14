package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2021", "14", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type rule struct {
	pair                            string
	newLeftProduct, newRightProduct string

	addition string
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var rules []rule

	scanner.Scan()
	start := scanner.Text()

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		parts := strings.Split(line, " -> ")

		input := parts[0]
		output := parts[1]

		rules = append(rules, rule{
			pair:            input,
			newLeftProduct:  string(input[0]) + output,
			newRightProduct: output + string(input[1]),
			addition:        output,
		})
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	counts := make(map[string]int)
	polymer := make(map[string]int)
	for i := 0; i < len(start)-1; i++ {
		polymer[start[i:i+2]] += 1
	}

	for i := 0; i < len(start); i++ {
		counts[start[i:i+1]]++
	}

	for i := 0; i < 10; i++ {
		simulate(polymer, counts, rules)
	}

	amnts := amounts(counts)

	fmt.Println(amnts[0] - amnts[len(amnts)-1])
}

func amounts(counts map[string]int) []int {
	ret := make([]int, len(counts))

	i := 0
	for _, v := range counts {
		ret[i] = v
		i++
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i] > ret[j]
	})

	return ret
}
func simulate(polymer map[string]int, counts map[string]int, rules []rule) {
	toAdd := make(map[string]int)

	for _, r := range rules {
		if amnt := polymer[r.pair]; amnt != 0 {
			delete(polymer, r.pair)
			toAdd[r.newLeftProduct] += amnt
			toAdd[r.newRightProduct] += amnt

			counts[r.addition] += amnt
		}
	}

	for k, v := range toAdd {
		polymer[k] += v
	}
}

func main() {
	solve(strings.NewReader("NNCB\n\nCH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C"))
	solve(input())
}
