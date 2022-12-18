package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
	"sort"
	"strings"
)

type Valve struct {
	name      string
	flowRate  int
	neighbors []string
}

var reg = regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func parse(parts []string) Valve {
	rate := lib.Atoi(parts[1])
	connections := strings.Split(parts[2], ", ")

	return Valve{
		name:      parts[0],
		flowRate:  rate,
		neighbors: connections,
	}
}

type Node struct {
	name      int
	flowRate  int
	neighbors []*Node
}

func (n *Node) Adjacent() []*Node {
	return n.neighbors
}

func buildGraph(valves []Valve) *Node {
	nodeMap := make(map[string]*Node)

	for i, valve := range valves {
		nodeMap[valve.name] = &Node{
			name:     i,
			flowRate: valve.flowRate,
		}
	}

	for _, valve := range valves {
		for _, neighbor := range valve.neighbors {
			nodeMap[valve.name].neighbors = append(nodeMap[valve.name].neighbors, nodeMap[neighbor])
		}
	}

	return nodeMap["AA"]
}

func copyMap[T comparable, V any](m map[T]V) map[T]V {
	r := make(map[T]V, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

type args struct {
	you      int
	elephant int
	time     int
	on       uint64
}

func hash(ints []int) uint64 {
	var result uint64

	for _, v := range ints {
		result += 1 << v
	}

	return result
}

var cache map[args]int
var total int
var maxTime int

var cacheSeenAtAll map[args]bool

func onToStr(a map[string]bool) string {
	v := lib.Keys(a)
	sort.Strings(v)
	return strings.Join(v, "")
}

func set2(index1 int, index2 int, on uint64) uint64 {
	return on + 1<<index1 + 1<<index2
}

func set(index int, on uint64) uint64 {
	return on + 1<<index
}

func isSet(index int, on uint64) bool {
	return (on>>index)&1 == 1
}

func size(on uint64) int {
	sum := 0
	for on > 0 {
		sum += int(on & 1)

		on = on >> 1
	}

	return sum
}

func exScore(you *Node, elephant *Node, time int, on uint64) int {
	// nothing more we can do if there's no time left
	if time <= 0 {
		return 0
	}

	onSize := size(on)
	// if all the valves are on there's no point moving around more
	if onSize == total {
		return 0
	}

	aaargs := args{you.name, elephant.name, time, on}

	if val, ok := cache[aaargs]; ok {
		return val
	}

	// if all the valves are open at a previous time, this path must be worse
	for i := time + 1; i < maxTime-onSize; i++ {
		if _, ok := cacheSeenAtAll[args{you.name, elephant.name, i, on}]; ok {
			return 0
		}
	}

	cacheSeenAtAll[aaargs] = true

	var max int

	var scoreFromYouOn int
	if !isSet(you.name, on) && you.name != elephant.name {
		// we can try turning your valve on
		scoreFromYouOn = you.flowRate * (time - 1)
	}

	var scoreFromElephantOn int
	if !isSet(elephant.name, on) {
		// we can try turning the elephant's valve on (if we're not at the same spot)
		scoreFromElephantOn = elephant.flowRate * (time - 1)
	}

	for _, youNeighbor := range you.neighbors {
		if scoreFromElephantOn != 0 {
			// mr elephant opens his valve but I move on
			turnedOn := set(elephant.name, on)
			turnOnAndMove := scoreFromElephantOn + exScore(youNeighbor, elephant, time-1, turnedOn)
			max = lib.Max(max, turnOnAndMove)
		}

		for _, elephantNeighbor := range elephant.neighbors {
			// we both can open our valve
			if scoreFromYouOn != 0 && scoreFromElephantOn != 0 {
				turnedOn := set2(elephant.name, you.name, on)
				turnOnAndMove := scoreFromYouOn + scoreFromElephantOn + exScore(you, elephant, time-1, turnedOn)
				max = lib.Max(max, turnOnAndMove)
			}

			if scoreFromYouOn != 0 {
				// I open my valve but mr elephant moves on
				turnedOn := set(you.name, on)
				turnOnAndMove := scoreFromYouOn + exScore(you, elephantNeighbor, time-1, turnedOn)
				max = lib.Max(max, turnOnAndMove)
			}

			// neither of us open our valve
			justMove := exScore(youNeighbor, elephantNeighbor, time-1, on)
			max = lib.Max(max, justMove)
		}
	}

	cache[aaargs] = max
	return max
}

func solve(valves []Valve) int {
	graph := buildGraph(valves)

	var on uint64
	for i, valve := range valves {
		if valve.flowRate == 0 {
			on = set(i, on)
		}
	}

	total = len(valves)
	cache = make(map[args]int)
	cacheSeenAtAll = make(map[args]bool)
	maxTime = 26

	r := exScore(graph, graph, 26, 0)

	if r == 3389 {
		panic("that's not right")
	}

	return r
}

func main() {
	solver := lib.Solver[[]Valve, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(reg, parse)),
		SolveF: solve,
	}

	solver.Expect("Valve AA has flow rate=0; tunnels lead to valves DD, II, BB\nValve BB has flow rate=13; tunnels lead to valves CC, AA\nValve CC has flow rate=2; tunnels lead to valves DD, BB\nValve DD has flow rate=20; tunnels lead to valves CC, AA, EE\nValve EE has flow rate=3; tunnels lead to valves FF, DD\nValve FF has flow rate=0; tunnels lead to valves EE, GG\nValve GG has flow rate=0; tunnels lead to valves FF, HH\nValve HH has flow rate=22; tunnel leads to valve GG\nValve II has flow rate=0; tunnels lead to valves AA, JJ\nValve JJ has flow rate=21; tunnel leads to valve II", 1707)
	solver.Solve()
}
