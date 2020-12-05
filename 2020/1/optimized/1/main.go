package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func parse(r io.Reader) map[int]bool {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rawExpenses := strings.Split(strings.TrimSpace(string(raw)), "\r\n")

	expenses := make(map[int]bool, len(rawExpenses))
	for _, row := range rawExpenses {
		expense, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}

		expenses[expense] = true
	}

	return expenses
}

func parseFile() map[int]bool {
	input, err := os.Open(path.Join("2020", "1", "input.txt"))
	if err != nil {
		panic(err)
	}

	return parse(input)
}

func solve(expenses map[int]bool) int {
	const target = 2020

	for expense := range expenses {
		// look for whether the other needed value exists
		needed := target - expense

		if _, ok := expenses[needed]; ok {
			return expense * (needed)
		}
	}

	panic("no matches")
}

func main() {
	fmt.Println(solve(parseFile()))
}
