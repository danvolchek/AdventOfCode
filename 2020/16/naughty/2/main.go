package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "16", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type validRange struct {
	min, max int
}

func parseRange(r string) validRange {
	parts :=  strings.Split(r, "-")

	min, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return validRange{
		min: min,
		max: max,
	}

}

func parseTicket (row string)[]int {
	parts := strings.Split(row, ",")

	ret := make([]int, len(parts))

	for i, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		ret[i] = val
	}

	return ret
}

func isValidVal(val int, rule []validRange) bool {
	for _, r := range rule{
		if val >= r.min && val <= r.max {
			return true
		}
	}

	return false
}

func isValid (ticket []int, rules map[string][]validRange) (bool, int) {

	for _, value := range ticket {
		foundValid := false
		for _, rule := range rules {
			if isValidVal(value, rule) {
				foundValid = true
				break
			}
		}

		if !foundValid {
			return false, value
		}
	}

	return true, 0

}

func getOptions(tickets [][]int, rules map[string][]validRange) []map[string]bool {
	ret := make([]map[string]bool, len(rules))


	for index := 0; index < len(rules); index += 1{
		for name, ranges := range rules {
			allValid := true

			for _, ticket := range tickets {
				if !isValidVal(ticket[index], ranges) {
					allValid = false
					break
				}
			}

			if allValid {
				if ret[index] == nil {
					ret[index] = make(map[string]bool)
				}
				ret[index][name] = true
			}
		}
	}

	return ret

}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	zone := 0

	fields := make(map[string][]validRange)
	var myTicket []int
	var nearbyTickets [][]int

	for scanner.Scan() {
		row := scanner.Text()

		if len(row) == 0 {
			zone += 1
			continue
		}

		switch zone {
		case 0:
			parts := strings.Split(row, ":")
			if len(parts) != 2 {
				panic(row)
			}

			ranges := strings.Split(strings.TrimSpace(parts[1]), " or ")

			fields[parts[0]] = []validRange{parseRange(ranges[0]), parseRange(ranges[1])}
			continue
		case 1:
			if row == "your ticket:" {
				continue
			}

			myTicket = parseTicket(row)
		case 2:
			if row == "nearby tickets:" {
				continue
			}

			nearbyTickets = append(nearbyTickets, parseTicket(row))
		default:
			panic(zone)
		}



	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(fields, myTicket, nearbyTickets)

	var validTickets [][]int


	for _, nearby := range nearbyTickets {
		if valid, _ := isValid(nearby, fields); valid {
			validTickets = append(validTickets, nearby)
		}
	}

	opts := getOptions(validTickets, fields)
	fmt.Println(opts)

	seenSingulars := make(map[string]bool)
	for {
		var singular string

		for _, opt := range opts {
			if len(opt) == 1 {
				for k := range opt {
					singular = k
				}

				if seenSingulars[singular] {
					continue
				}

				seenSingulars[singular] = true
				break
			}
		}

		if singular == "" {
			panic("hee hoo")
		}

		for _, opt := range opts {
			if len(opt) != 1 {
				delete(opt, singular)
			}
		}

		all1 := true
		for _, opt := range opts {
			if len(opt) != 1 {
				all1 = false
				break
			}
		}

		if all1 {
			break
		}
	}

	fmt.Println(opts)

	multi := 1
	for index, opt := range opts {
		for name := range opt {
			if strings.Index(name, "departure") == 0 {
				multi *= myTicket[index]
			}
		}
	}

	fmt.Println(multi)

}

func main() {
	solve(strings.NewReader("class: 0-1 or 4-19\nrow: 0-5 or 8-19\nseat: 0-13 or 16-19\n\nyour ticket:\n11,12,13\n\nnearby tickets:\n3,9,18\n15,1,5\n5,14,9"))
	solve(input())
}
