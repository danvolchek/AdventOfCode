package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
)

const (
	rowEntryPattern = `(\d+)-(\d+) (.): (.+)`
)

var rowEntry = regexp.MustCompile(rowEntryPattern)

type dbEntry struct {
	policy   policy
	password string
}

type policy struct {
	min, max int
	char     byte
}

func input() *os.File {
	input, err := os.Open(path.Join("2020", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) []dbEntry {
	var dbEntries []dbEntry

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		result := rowEntry.FindAllStringSubmatch(scanner.Text(), -1)[0]

		min, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}

		max, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}

		dbEntries = append(dbEntries, dbEntry{
			policy: policy{
				min:  min,
				max:  max,
				char: result[3][0],
			},
			password: result[4],
		})
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return dbEntries
}

func solve(dbEntries []dbEntry) int {
	valid := 0

	for _, dbEntry := range dbEntries {
		occurrences := countOccurrences(dbEntry.policy.char, dbEntry.password)
		if occurrences >= dbEntry.policy.min && occurrences <= dbEntry.policy.max {
			valid += 1
		}
	}

	return valid
}

func countOccurrences(needle byte, haystack string) int {
	num := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			num++
		}
	}

	return num
}

func main() {
	fmt.Println(solve(parse(input())))
}
