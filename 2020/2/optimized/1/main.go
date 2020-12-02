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
	min, max int
	char     uint8
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

	rows := bytes.Split(bytes.TrimSpace(raw), []byte("\r\n"))
	dbEntries := make([]dbEntry, len(rows))

	for i, row := range rows {
		result := rowEntry.FindAllSubmatch(row, -1)[0]

		min, err := strconv.Atoi(string(result[1]))
		if err != nil {
			panic(err)
		}

		max, err := strconv.Atoi(string(result[2]))
		if err != nil {
			panic(err)
		}

		dbEntries[i] = dbEntry{
			policy: policy{
				min: min,
				max: max,
				char: result[3][0],
			},
			password: string(result[4]),
		}

	}

	return dbEntries
}

func getValidPasswords(dbEntries []dbEntry) int {
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
	input, err := os.Open(path.Join("2020", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	entries := parse(input)

	fmt.Println(getValidPasswords(entries))
}
