package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

func parse(r io.Reader) []dbEntry {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := bytes.Split(bytes.TrimSpace(raw), []byte{'\r', '\n'})
	dbEntries := make([]dbEntry, len(rows))

	for i, row := range rows {
		result := rowEntry.FindAllSubmatch(row, -1)[0]

		firstPosition, err := strconv.Atoi(string(result[1]))
		if err != nil {
			panic(err)
		}

		secondPosition, err := strconv.Atoi(string(result[2]))
		if err != nil {
			panic(err)
		}

		dbEntries[i] = dbEntry{
			policy: policy{
				firstPosition:  firstPosition - 1,
				secondPosition: secondPosition - 1,
				char:           result[3][0],
			},
			password: string(result[4]),
		}

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

func parseFile() []dbEntry {
	input, err := os.Open(path.Join("2020", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	return parse(input)
}

func main() {
	fmt.Println(solve(parseFile()))
}
