package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	CountryID = "cid"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) []map[string]bool {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(raw), "\r\n")

	var passports []map[string]bool

	currentPassport := make(map[string]bool)

	for _, row := range rows {
		if len(row) == 0 {
			passports = append(passports, currentPassport)

			currentPassport = make(map[string]bool)
			continue
		}

		rawFields := strings.Split(row, " ")
		for _, rawField := range rawFields {
			rawFieldParts := strings.Split(rawField, ":")

			currentPassport[rawFieldParts[0]] = true
		}
	}

	return passports
}

func isValid(passport map[string]bool) bool {
	_, hasCid := passport[CountryID]
	return len(passport) == 8 || (len(passport) == 7 && !hasCid)
}

func solve(passports []map[string]bool) int {
	valid := 0

	for _, passport := range passports {
		if isValid(passport) {
			valid += 1
		}
	}

	return valid
}

func main() {
	fmt.Println(solve(parse(input())))
}
