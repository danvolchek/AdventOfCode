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
	myTicketHeader       = "your ticket:"
	nearbyTicketHeader   = "nearby tickets:"
	ruleSeparator        = ":"
	ticketSeparator      = ","
	valueRangeSeparator  = "-"
	targetFieldPrefix    = "departure"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "16", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func parse(r io.Reader) ([]rule, []int, [][]int) {
	var rules []rule
	var myTicket []int
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
			rules = append(rules, newRule(row))
		case 1: // my ticket
			// discard header
			if row == myTicketHeader {
				continue
			}

			myTicket = newTicket(row)
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

	return rules, myTicket, nearbyTickets
}

type rule struct {
	name        string
	valueRanges []valueRange
}

func newRule(raw string) rule {
	parts := strings.Split(raw, ruleSeparator)
	if len(parts) != 2 {
		panic(raw)
	}

	rawRanges := strings.Split(strings.TrimSpace(parts[1]), valueRangesSeparator)

	valueRanges := make([]valueRange, len(rawRanges))
	for i, rawRange := range rawRanges {
		valueRanges[i] = newRange(rawRange)
	}

	return rule{
		name:        parts[0],
		valueRanges: valueRanges,
	}
}

func (r rule) matches(value int) bool {
	for _, valueRange := range r.valueRanges {
		if valueRange.matches(value) {
			return true
		}
	}

	return false
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

func valueMatchesAnyRule(value int, rules []rule) bool {
	for _, rule := range rules {
		if rule.matches(value) {
			return true
		}
	}

	return false
}

func findInvalidValues(ticket []int, rules []rule) []int {
	var invalidValues []int

	for _, value := range ticket {
		if !valueMatchesAnyRule(value, rules) {
			invalidValues = append(invalidValues, value)
		}
	}

	return invalidValues
}

func findValidTickets(rules []rule, tickets [][]int) [][]int {
	var validTickets [][]int

	for _, ticket := range tickets {
		if len(findInvalidValues(ticket, rules)) == 0 {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func getPossibleFieldOrderings(rules []rule, tickets [][]int) []map[string]bool {
	possibilities := make([]map[string]bool, len(rules))

	for index := 0; index < len(rules); index += 1 {
		for _, rule := range rules {
			allValid := true

			for _, ticket := range tickets {
				if !rule.matches(ticket[index]) {
					allValid = false
					break
				}
			}

			if allValid {
				if possibilities[index] == nil {
					possibilities[index] = make(map[string]bool)
				}
				possibilities[index][rule.name] = true
			}
		}
	}

	return possibilities
}

func findNewDefiniteField(processedCollapsedFields map[string]bool, possibleOrderings []map[string]bool) (int, string) {
	var collapsed string

	for fieldIndex, orderingsForIndex := range possibleOrderings {
		if len(orderingsForIndex) == 1 {
			for k := range orderingsForIndex {
				collapsed = k
			}

			if processedCollapsedFields[collapsed] {
				continue
			}

			processedCollapsedFields[collapsed] = true
			return fieldIndex, collapsed
		}
	}

	return -1, ""
}

func removePossibility(fieldIndex int, fieldName string, possibleOrderings []map[string]bool) {
	for index, orderingsForIndex := range possibleOrderings {
		if index != fieldIndex {
			delete(orderingsForIndex, fieldName)
		}
	}
}

func getDefiniteFieldOrderings(possibleOrderings []map[string]bool) []string {
	processedCollapsedFields := make(map[string]bool)

	definiteOrdering := make([]string, len(possibleOrderings))

	for {
		fieldIndex, fieldName := findNewDefiniteField(processedCollapsedFields, possibleOrderings)

		if fieldIndex == -1 {
			break
		}

		definiteOrdering[fieldIndex] = fieldName

		removePossibility(fieldIndex, fieldName, possibleOrderings)
	}

	return definiteOrdering
}

func solve(rules []rule, myTicket []int, nearbyTickets [][]int) int {
	validNearbyTickets := findValidTickets(rules, nearbyTickets)

	possibleOrderings := getPossibleFieldOrderings(rules, validNearbyTickets)

	definiteOrderings := getDefiniteFieldOrderings(possibleOrderings)

	solution := 1

	for fieldIndex, fieldName := range definiteOrderings {
		if strings.Index(fieldName, targetFieldPrefix) == 0 {
			solution *= myTicket[fieldIndex]
		}
	}

	return solution
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("class: 0-1 or 4-19\nrow: 0-5 or 8-19\nseat: 0-13 or 16-19\n\nyour ticket:\n11,12,13\n\nnearby tickets:\n3,9,18\n15,1,5\n5,14,9"))))
	fmt.Println(solve(parse(input())))
}
