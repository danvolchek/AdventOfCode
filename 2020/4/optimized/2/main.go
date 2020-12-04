package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type fieldValueValidator interface {
	IsValid(value string) bool
}

type length struct {
	length int
}

func (l length) IsValid(value string) bool {
	return len(value) == l.length
}

type between struct {
	min, max int
}

func (i between) IsValid(value string) bool {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return intVal >= i.min && intVal <= i.max
}

type conditional struct {
	keyFunc       func(value string) (string, string)
	keyValidators map[string]fieldValueValidator
}

func (m conditional) IsValid(value string) bool {
	key, newValue := m.keyFunc(value)

	return m.keyValidators[key].IsValid(newValue)
}

type matches struct {
	reg *regexp.Regexp
}

func (r matches) IsValid(value string) bool {
	return r.reg.MatchString(value)
}

type oneOf struct {
	values []string
}

func (o oneOf) IsValid(value string) bool {
	for _, allowed := range o.values {
		if allowed == value {
			return true
		}
	}

	return false
}

type alwaysValid struct{}

func (a alwaysValid) IsValid(_ string) bool {
	return true
}

type both struct {
	first, second fieldValueValidator
}

func (s both) IsValid(value string) bool {
	return s.first.IsValid(value) && s.second.IsValid(value)
}

const (
	BirthYear      = "byr"
	IssueYear      = "iyr"
	ExpirationYear = "eyr"
	Height         = "hgt"
	HairColor      = "hcl"
	EyeColor       = "ecl"
	PassportID     = "pid"
	CountryID      = "cid"
)

var validatorForField = map[string]fieldValueValidator{
	BirthYear: both{
		first:  length{4},
		second: between{min: 1920, max: 2002},
	},
	IssueYear: both{
		first:  length{4},
		second: between{min: 2010, max: 2020},
	},
	ExpirationYear: both{
		first:  length{4},
		second: between{min: 2020, max: 2030},
	},
	Height: both{
		first: matches{regexp.MustCompile(`^\d+(cm|in)$`)},
		second: conditional{
			keyFunc: func(value string) (string, string) {
				num, unit := value[:len(value)-2], value[len(value)-2:]

				return unit, num
			},
			keyValidators: map[string]fieldValueValidator{
				"cm": between{min: 150, max: 193},
				"in": between{min: 59, max: 76},
			},
		},
	},
	HairColor:  matches{regexp.MustCompile(`^#[0-9a-f]{6}$`)},
	EyeColor:   oneOf{[]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}},
	PassportID: matches{regexp.MustCompile(`^\d{9}$`)},
	CountryID:  alwaysValid{},
}

type passport struct {
	fields map[string]string
}

func newPassport() passport {
	return passport{
		fields: make(map[string]string),
	}
}

func (p passport) AddField(name, value string) {
	p.fields[name] = value
}

func (p passport) IsValid() bool {
	if !p.hasRequiredFields() {
		return false
	}

	for fieldName, fieldValue := range p.fields {
		if !validatorForField[fieldName].IsValid(fieldValue) {
			return false
		}
	}

	return true
}

func (p passport) hasRequiredFields() bool {
	_, hasCid := p.fields[CountryID]
	return len(p.fields) == 8 || (len(p.fields) == 7 && !hasCid)
}

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

			currentPassport.AddField(rawFieldParts[0], rawFieldParts[1])
		}
	}

	return passports
}

func getValidPassports(passports []passport) int {
	valid := 0

	for _, passport := range passports {
		if passport.IsValid() {
			valid += 1
		}
	}

	return valid
}

func main() {
	input, err := os.Open(path.Join("2020", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(getValidPassports(parse(input)))
}
