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

const (
	valueRangesSeparator = " or "
	nearbyTicketHeader   = "nearby tickets:"
	ruleSeparator        = ":"
	ticketSeparator      = ","
	valueRangeSeparator  = "-"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "16", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) (map[string][]valueRange, [][]int) {
	rules := make(map[string][]valueRange)
	var nearbyTickets [][]int

	parsingZone := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Text()

		if len(row) == 0 {
			parsingZone += 1
			continue
		}

		switch parsingZone {
		case 0: // rules
			parts := strings.Split(row, ruleSeparator)
			if len(parts) != 2 {
				panic(row)
			}

			ranges := strings.Split(strings.TrimSpace(parts[1]), valueRangesSeparator)

			rules[parts[0]] = []valueRange{newRange(ranges[0]), newRange(ranges[1])}
			continue
		case 1: // my ticket
			// my ticket isn't needed
		case 2: // nearby tickets
			// discard header
			if row == nearbyTicketHeader {
				continue
			}

			nearbyTickets = append(nearbyTickets, newTicket(row))
		default:
			panic(parsingZone)
		}

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return rules, nearbyTickets
}

type valueRange struct {
	min, max int
}

func (v valueRange) matches(value int) bool {
	return value >= v.min && value <= v.max
}

func newRange(raw string) valueRange {
	parts := strings.Split(raw, valueRangeSeparator)

	min, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return valueRange{
		min: min,
		max: max,
	}

}

func newTicket(raw string) []int {
	parts := strings.Split(raw, ticketSeparator)

	ticket := make([]int, len(parts))

	for i, part := range parts {
		intVal, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		ticket[i] = intVal
	}

	return ticket
}

func valueMatchesAnyRange(value int, valueRanges []valueRange) bool {
	for _, valueRange := range valueRanges {
		if valueRange.matches(value) {
			return true
		}
	}

	return false
}

func valueMatchesAnyRule(value int, rules map[string][]valueRange) bool {
	for _, valueRanges := range rules {
		if valueMatchesAnyRange(value, valueRanges) {
			return true
		}
	}

	return false
}

func findInvalidValues(ticket []int, rules map[string][]valueRange) []int {
	var invalidValues []int

	for _, value := range ticket {
		if !valueMatchesAnyRule(value, rules) {
			invalidValues = append(invalidValues, value)
		}
	}

	return invalidValues
}

func solve(rules map[string][]valueRange, nearbyTickets [][]int) int {
	errorRate := 0

	for _, ticket := range nearbyTickets {
		for _, invalidValue := range findInvalidValues(ticket, rules) {
			errorRate += invalidValue
		}
	}

	return errorRate
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50\n\nyour ticket:\n7,1,14\n\nnearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12"))))
	fmt.Println(solve(parse(input())))
}
