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

func parse(r io.Reader) []passport {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(raw), "\r\n")

	var passports []passport

	currentPassport := newPassport()

	for _, row := range rows {
		if len(row) == 0 {
			passports = append(passports, currentPassport)

			currentPassport = newPassport()
			continue
		}

		rawFields := strings.Split(row, " ")
		for _, rawField := range rawFields {
			rawFieldParts := strings.Split(rawField, ":")

			currentPassport.AddField(rawFieldParts[0])
		}
	}

	return passports
}

func parseFile() []passport {
	input, err := os.Open(path.Join("2020", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	return parse(input)
}

type passport struct {
	fields map[string]bool
}

func newPassport() passport {
	return passport{
		fields: make(map[string]bool),
	}
}

func (p passport) AddField(name string) {
	p.fields[name] = true
}

func (p passport) IsValid() bool {
	_, hasCid := p.fields[CountryID]
	return len(p.fields) == 8 || (len(p.fields) == 7 && !hasCid)
}

func solve(passports []passport) int {
	valid := 0

	for _, passport := range passports {
		if passport.IsValid() {
			valid += 1
		}
	}

	return valid
}

func main() {
	fmt.Println(solve(parseFile()))
}
