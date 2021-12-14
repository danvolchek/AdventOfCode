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
	pair     string
	newPairs [2]string
}

func parse(r io.Reader) (map[string]int, byte, byte, []rule) {
	polymer := make(map[string]int)
	var rules []rule

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	rawPolymer := scanner.Text()
	for i := 0; i < len(rawPolymer)-1; i++ {
		polymer[rawPolymer[i:i+2]] += 1
	}

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " -> ")

		pair := parts[0]
		newElement := parts[1]

		newPairs := [...]string{pair[0:1] + newElement, newElement + pair[1:2]}

		rules = append(rules, rule{
			pair:     pair,
			newPairs: newPairs,
		})
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return polymer, rawPolymer[0], rawPolymer[len(rawPolymer)-1], rules
}

func solve(r io.Reader) {
	polymer, start, end, rules := parse(r)
	for i := 0; i < steps; i++ {
		simulate(polymer, rules)
	}

	amounts := countElements(polymer, start, end)

	mostCommon := amounts[0]
	leastCommon := amounts[len(amounts)-1]

	fmt.Println(mostCommon - leastCommon)
}

func countElements(polymer map[string]int, start, end byte) []int {
	countsByElement := make(map[byte]int)

	// counting each pair double counts every element, because every pair shares elements with adjacent pairs
	// (except for the first and last element, which are not shared, and so are undercounted by one)
	for pair, amount := range polymer {
		for index := range pair {
			countsByElement[pair[index]] += amount
		}
	}

	// add one to first and last element
	countsByElement[start] += 1
	countsByElement[end] += 1

	// undo double counting of each element
	for element := range countsByElement {
		countsByElement[element] /= 2
	}

	var counts []int
	for _, amount := range countsByElement {
		counts = append(counts, amount)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	return counts
}

func simulate(polymer map[string]int, rules []rule) {
	toAdd := make(map[string]int)

	for _, r := range rules {
		if amount := polymer[r.pair]; amount != 0 {
			// this pair is split, so it's removed from the polymer
			delete(polymer, r.pair)

			// each pair this rule generates is added to the polymer
			for _, pair := range r.newPairs {
				toAdd[pair] += amount
			}
		}
	}

	for pair, amount := range toAdd {
		polymer[pair] += amount
	}
}

const steps = 40

func main() {
	solve(strings.NewReader("NNCB\n\nCH -> B\nHH -> N\nCB -> H\nNH -> C\nHB -> C\nHC -> B\nHN -> C\nNN -> C\nBH -> H\nNC -> B\nNB -> B\nBN -> B\nBB -> N\nBC -> B\nCC -> N\nCN -> C"))
	solve(input())
}
