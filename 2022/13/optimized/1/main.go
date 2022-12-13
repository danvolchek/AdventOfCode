package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strconv"
)

type Ordering int

const (
	Smaller Ordering = iota
	Equal
	Larger
)

type Value struct {
	Val   int
	Items []Value
}

// elements returns the items rawList represents
func elements(rawList string) []string {
	if len(rawList) == 0 {
		return nil
	}

	var result []string

	var last int
	var nesting int

	for i := range rawList {
		switch rawList[i] {
		case '[':
			nesting += 1
		case ']':
			nesting -= 1
		case ',':
			if nesting == 0 {
				result = append(result, rawList[last:i])
				last = i + 1
			}
		}
	}
	result = append(result, rawList[last:])

	return result
}

func parseValue(raw string) Value {
	// If raw is an int, then it is a single value
	if v, err := strconv.Atoi(raw); err == nil {
		return Value{Val: v}
	}

	// Otherwise, strip the list markers and parse the item contents recursively
	return Value{
		Items: lib.Map(elements(raw[1:len(raw)-1]), parseValue),
	}
}

func (v Value) Compare(o Value) Ordering {
	// If both are ints, compare their values
	if v.Items == nil && o.Items == nil {
		return compare(v.Val, o.Val)
	}

	toList := func(v Value) []Value {
		if v.Items != nil {
			return v.Items
		}

		return []Value{v}
	}

	// Otherwise, compare the elements pairwise, replacing an int with a list of that int
	vList, oList := toList(v), toList(o)

	for i := 0; i < len(vList) && i < len(oList); i += 1 {
		result := vList[i].Compare(oList[i])
		if result != Equal {
			return result
		}
	}

	// If one list ran out if items, the ordering is based on list size
	return compare(len(vList), len(oList))
}

func compare(a, b int) Ordering {
	if a == b {
		return Equal
	}

	if a < b {
		return Smaller
	}

	return Larger
}

func solve(packets [][]Value) int {
	sum := 0

	for i, pair := range packets {
		if pair[0].Compare(pair[1]) == Smaller {
			sum += i + 1
		}
	}

	return sum
}

func main() {
	solver := lib.Solver[[][]Value, int]{
		ParseF: lib.ParseChunks(lib.ParseLine(parseValue)),
		SolveF: solve,
	}

	solver.Expect("[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]", 13)
	solver.Verify(5340)
}
