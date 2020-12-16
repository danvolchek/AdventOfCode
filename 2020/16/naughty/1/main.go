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
	parts := strings.Split(r, "-")

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

func parseTicket(row string) []int {
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
	for _, r := range rule {
		if val >= r.min && val <= r.max {
			return true
		}
	}

	return false
}

func isValid(ticket []int, rules map[string][]validRange) (bool, int) {

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

	invalid := 0
	for _, nearby := range nearbyTickets {
		if valid, val := isValid(nearby, fields); !valid {
			invalid += val
		}
	}

	fmt.Println(invalid)
}

func main() {
	solve(strings.NewReader("class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\n\nyour ticket:\n7,1,14\n\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12"))
	solve(input())
}
