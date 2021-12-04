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

func isValid(name, value string) bool {
	switch name {
	case "byr":
		if len(value) != 4 {
			return false
		}

		iv, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		return iv >= 1920 && iv <= 2002
	case "iyr":
		if len(value) != 4 {
			return false
		}

		iv, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		return iv >= 2010 && iv <= 2020
	case "eyr":
		if len(value) != 4 {
			return false
		}

		iv, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		return iv >= 2020 && iv <= 2030
	case "hgt":
		if len(value) <= 2 {
			return false
		}

		iv, err := strconv.Atoi(value[:len(value)-2])
		if err != nil {
			return false
		}

		switch value[len(value)-2:] {
		case "cm":
			return iv >= 150 && iv <= 193
		case "in":
			return iv >= 59 && iv <= 76
		default:
			return false
		}
	case "hcl":
		reg := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return reg.MatchString(value)
	case "ecl":
		switch value {
		case "amb":
			fallthrough
		case "blu":
			fallthrough
		case "brn":
			fallthrough
		case "gry":
			fallthrough
		case "grn":
			fallthrough
		case "hzl":
			fallthrough
		case "oth":
			return true
		default:
			return false
		}
	case "pid":
		reg := regexp.MustCompile(`^\d{9}$`)
		return reg.MatchString(value)
	case "cid":
		return true
	}

	panic(name)
}

func explain(name, value string) string {
	switch name {
	case "byr":
		return "" //fmt.Sprintf("1920 <= %s <= 2002", value)

	case "iyr":
		return "" //fmt.Sprintf("2010 <= %s <= 2020", value)
	case "eyr":
		return "" //fmt.Sprintf("2020 <= %s <= 2030", value)
	case "hgt":
		switch value[len(value)-2:] {
		case "cm":
			return "" //fmt.Sprintf("150 <= %s <= 193", value)
		case "in":
			return "" //fmt.Sprintf("59 <= %s <= 76", value)
		default:
			panic("should be valid")
		}
	case "hcl":
		return "" //fmt.Sprintf("%s matches regexp #[0-9a-f]{6}", value)
	case "ecl":
		return "" //fmt.Sprintf("%s is one of amb blu brn gry grn hzl oth", value)
	case "pid":
		return fmt.Sprintf("%s is 9 length digit", value)
	case "cid":
		return "" //fmt.Sprintf("%s always", value)
	}

	panic(name)
}

func solve(r io.Reader) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(raw), "\r\n")
	if len(rows) == 1 {
		rows = strings.Split(string(raw), "\n")
	}

	valid := 0

	curr := make(map[string]string)

	for _, row := range rows {
		if len(row) == 0 {
			if entriesValid(curr) {
				valid += 1
			}

			curr = make(map[string]string)
			continue
		}

		items := strings.Split(row, " ")
		for _, item := range items {
			parts := strings.Split(item, ":")

			curr[parts[0]] = parts[1]
		}
	}

	fmt.Println(valid)
}

func entriesValid(curr map[string]string) bool {
	_, hasCid := curr["cid"]
	if len(curr) == 8 || (len(curr) == 7 && !hasCid) {

		allValid := true

		for k, v := range curr {
			if !isValid(k, v) {
				allValid = false
				break
			}
		}

		if allValid {
			for k, v := range curr {
				exp := explain(k, v)
				if exp != "" {
					fmt.Println(k, exp)
				}

			}

			fmt.Println()
		}

		return allValid
	}

	return false
}

func main() {
	//solve(strings.NewReader("eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n\niyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946\n\nhcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n\nhgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007"))
	//solve(strings.NewReader("pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f\n\neyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n\nhcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022\n\niyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"))

	input, err := os.Open(path.Join("2020", "4", "input.txt"))
	if err != nil {
		panic(err)
	}

	solve(input)
}
