package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
	"strconv"
)

type Value struct {
	Val   int
	Items []Value
}

func (v Value) Int() bool {
	return v.Items == nil
}

type Packet struct {
	val Value
}

func parts(line string) []string {
	nesting := 0

	var parts []string

	var last int

	for i := range line {
		switch line[i] {
		case '[':
			nesting += 1
		case ']':
			nesting -= 1
		case ',':
			if nesting == 0 {
				parts = append(parts, line[last:i])
				last = i + 1
			}
		}
	}
	parts = append(parts, line[last:])

	return parts
}

func parseValue(line string) Value {
	if v, err := strconv.Atoi(line); err == nil {
		return Value{Val: v}
	}

	stripped := line[1 : len(line)-1]

	p := parts(stripped)

	return Value{
		Items: lib.Map(lib.Filter(p, func(p string) bool {
			return len(p) > 0
		}), parseValue),
	}

}

func parsePacket(line string) Packet {
	return Packet{val: parseValue(line)}
}

func parseChunk(chunk string) []Packet {
	return lib.ParseLine(parsePacket)(chunk)
}

func smallerLists(p1, p2 []Value) (bool, bool) {
	i := 0

	for {
		if i >= len(p1) {
			return true, len(p1) != len(p2)
		}

		if i >= len(p2) {
			return false, true
		}

		sm, ok := smaller(p1[i], p2[i])
		if ok {
			return sm, true
		}

		i += 1
	}
}

func smaller(p1, p2 Value) (bool, bool) {
	if p1.Int() && p2.Int() {
		return p1.Val < p2.Val, p1.Val != p2.Val
	}

	if !p1.Int() && !p2.Int() {
		return smallerLists(p1.Items, p2.Items)
	}

	if p1.Int() {
		return smallerLists([]Value{p1}, p2.Items)
	}

	return smallerLists(p1.Items, []Value{p2})
}

func solve(packets [][]Packet) int {

	var allPackets []Packet

	for _, o := range packets {
		for _, i := range o {
			allPackets = append(allPackets, i)
		}
	}

	allPackets = append(allPackets, Packet{val: Value{Items: []Value{
		{Items: []Value{{Val: 2}}},
	}}})
	allPackets = append(allPackets, Packet{val: Value{Items: []Value{
		{Items: []Value{{Val: 6}}},
	}}})

	slices.SortFunc(allPackets, func(p1, p2 Packet) bool {
		sm, ok := smaller(p1.val, p2.val)
		return sm && ok
	})

	mul := 1

	for i, packet := range allPackets {
		if len(packet.val.Items) == 1 && len(packet.val.Items[0].Items) == 1 && (packet.val.Items[0].Items[0].Val == 2 || packet.val.Items[0].Items[0].Val == 6) {
			mul *= i + 1
		}
	}

	return mul
}

func main() {
	solver := lib.Solver[[][]Packet, int]{
		ParseF: lib.ParseChunks(parseChunk),
		SolveF: solve,
	}

	solver.Expect("[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]", 140)
	solver.Solve()
}
