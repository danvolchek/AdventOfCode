package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"regexp"
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
	index       int
	name        string
	flowRate    int
	neighbors   []*Node
	travelTimes []int
}

func (n *Node) Adjacent() []*Node {
	return n.neighbors
}

func buildGraph(valves []Valve) map[string]*Node {
	nodeMap := make(map[string]*Node)

	for i, valve := range valves {
		nodeMap[valve.name] = &Node{
			index:    i,
			name:     valve.name,
			flowRate: valve.flowRate,
		}
	}

	for _, valve := range valves {
		for _, neighbor := range valve.neighbors {
			nodeMap[valve.name].neighbors = append(nodeMap[valve.name].neighbors, nodeMap[neighbor])
		}
	}

	return nodeMap
}

func buildGraphZeroesRemoved(nodeMap map[string]*Node) *Node {
	newMap := make(map[string]*Node)
	for name, node := range nodeMap {
		if node.flowRate == 0 && name != "AA" {
			continue
		}

		newMap[name] = &Node{
			index:    node.index,
			name:     node.name,
			flowRate: node.flowRate,
		}
	}

	for startName, start := range nodeMap {
		if start.flowRate == 0 && startName != "AA" {
			continue
		}

		for endName, end := range nodeMap {
			if start == end {
				continue
			}

			if end.flowRate == 0 && endName != "AA" {
				continue
			}

			path, ok := lib.BFS(start, func(n *Node) bool { return n == end })
			if !ok {
				panic("disconnected valve")
			}

			for i := 0; i < len(path)-1; i++ {
				curr := newMap[path[i].name]

				counter := 1
				nextNonZero := path[i+counter]
				for nextNonZero.flowRate == 0 && nextNonZero.name != "AA" {
					counter++
					nextNonZero = path[i+counter]
				}

				nextNonZero = newMap[nextNonZero.name]

				if curr != nil {
					hasNeighbor := false
					for _, n := range curr.neighbors {
						if n == nextNonZero {
							hasNeighbor = true
							break
						}
					}

					if hasNeighbor {
						continue
					}

					curr.neighbors = append(curr.neighbors, nextNonZero)
					curr.travelTimes = append(curr.travelTimes, counter)
				}

				i += counter - 1
			}
		}
	}

	return newMap["AA"]
}

type cacheArgs struct {
	position string
	timeLeft int
}

type cacheResult struct {
	flowRate         int
	pressureReleased int
}

func turnValveOn(index int, on uint64) uint64 {
	return on + 1<<index
}

func isValveOn(index int, on uint64) bool {
	return (on>>index)&1 == 1
}

func numValvesOn(on uint64) int {
	sum := 0
	for on > 0 {
		sum += int(on & 1)

		on = on >> 1
	}

	return sum
}

var totalValves int
var totalTime int
var cache map[cacheArgs]cacheResult

func findBestPressure(position *Node, pressureReleased int, timeLeft int, flowRate int, valves uint64) int {
	if timeLeft <= 0 {
		return 0
	}

	if numValvesOn(valves) == totalValves {
		return 0
	}

	for time := timeLeft + 1; time < totalTime; time++ {
		arg := cacheArgs{
			position: position.name,
			timeLeft: time,
		}

		if v, ok := cache[arg]; ok && pressureReleased < v.pressureReleased && flowRate < v.flowRate {
			return 0
		}
	}

	cache[cacheArgs{
		position: position.name,
		timeLeft: timeLeft,
	}] = cacheResult{
		flowRate:         flowRate,
		pressureReleased: pressureReleased,
	}

	var max int
	if !isValveOn(position.index, valves) && position.flowRate != 0 {
		newPressureReleased := pressureReleased + flowRate
		newFlowRate := flowRate + position.flowRate
		newValves := turnValveOn(position.index, valves)

		max = lib.Max(max, newPressureReleased+findBestPressure(position, newPressureReleased, timeLeft-1, newFlowRate, newValves))
	}

	for i, neighbor := range position.neighbors {
		travelTime := position.travelTimes[i]
		newPressureReleased := pressureReleased + travelTime*flowRate

		max = lib.Max(max, findBestPressure(neighbor, newPressureReleased, timeLeft-travelTime, flowRate, valves))
	}

	return max
}

func solve(valves []Valve) int {
	originalGraph := buildGraph(valves)
	newGraph := buildGraphZeroesRemoved(originalGraph)

	totalValves = len(lib.Filter(valves, func(v Valve) bool { return v.flowRate != 0 }))
	totalTime = 30
	cache = make(map[cacheArgs]cacheResult)

	result := findBestPressure(newGraph, 0, totalTime, 0, 0)

	return result
}

func main() {
	solver := lib.Solver[[]Valve, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(reg, parse)),
		SolveF: solve,
	}

	solver.Expect("Valve AA has flow rate=0; tunnels lead to valves DD, II, BB\nValve BB has flow rate=13; tunnels lead to valves CC, AA\nValve CC has flow rate=2; tunnels lead to valves DD, BB\nValve DD has flow rate=20; tunnels lead to valves CC, AA, EE\nValve EE has flow rate=3; tunnels lead to valves FF, DD\nValve FF has flow rate=0; tunnels lead to valves EE, GG\nValve GG has flow rate=0; tunnels lead to valves FF, HH\nValve HH has flow rate=22; tunnel leads to valve GG\nValve II has flow rate=0; tunnels lead to valves AA, JJ\nValve JJ has flow rate=21; tunnel leads to valve II", 1651)
	solver.Verify(1724)
}
