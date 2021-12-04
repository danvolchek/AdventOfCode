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

type dbEntry struct {
	policy   policy
	password string
}

type policy struct {
	firstPosition, secondPosition int
	char                          byte
}

const (
	rowEntryPattern = `(\d+)-(\d+) (.): (.+)`
)

var rowEntry = regexp.MustCompile(rowEntryPattern)

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

		firstPosition, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}

		secondPosition, err := strconv.Atoi(result[2])
		if err != nil {
			panic(err)
		}

		dbEntries = append(dbEntries, dbEntry{
			policy: policy{
				firstPosition:  firstPosition - 1,
				secondPosition: secondPosition - 1,
				char:           result[3][0],
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
		atFirstPosition := dbEntry.password[dbEntry.policy.firstPosition] == dbEntry.policy.char
		atSecondPosition := dbEntry.password[dbEntry.policy.secondPosition] == dbEntry.policy.char
		if atFirstPosition != atSecondPosition {
			valid += 1
		}
	}

	return valid
}

func main() {
	fmt.Println(solve(parse(input())))
}
