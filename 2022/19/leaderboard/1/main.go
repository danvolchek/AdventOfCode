package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

type Inventory struct {
	Ore, Clay, Obsidian, Geodes                        int
	OreRobots, ClayRobots, ObsidianRobots, GeodeRobots int

	Time int

	b *Blueprint
}

func (i Inventory) Adjacent() []Inventory {
	if i.Time == 0 {
		return nil
	}

	var results []Inventory

	i.Time -= 1
	i.Ore += i.OreRobots
	i.Clay += i.ClayRobots
	i.Obsidian += i.ObsidianRobots
	i.Geodes += i.GeodeRobots

	if i.Ore >= i.b.OreRobotOreCost {
		newInv := i
		newInv.Ore -= i.b.OreRobotOreCost
		newInv.OreRobots++
		results = append(results, newInv)
	}

	if i.Ore >= i.b.ClayRobotOreCost {
		newInv := i
		newInv.Ore -= i.b.ClayRobotOreCost
		newInv.ClayRobots++
		results = append(results, newInv)
	}

	if i.Ore >= i.b.ObsidianRobotOreCost && i.Clay >= i.b.ObsidianRobotClayCost {
		newInv := i
		newInv.Ore -= i.b.ObsidianRobotOreCost
		newInv.Clay -= i.b.ObsidianRobotClayCost
		newInv.ObsidianRobots++
		results = append(results, newInv)
	}

	if i.Ore >= i.b.GeodeRobotOreCost && i.Clay >= i.b.GeodeRobotObsidianCost {
		newInv := i
		newInv.Ore -= i.b.GeodeRobotOreCost
		newInv.Obsidian -= i.b.GeodeRobotObsidianCost
		newInv.GeodeRobots++
		results = append(results, newInv)
	}

	return results
}

type Blueprint struct {
	Id                                          int
	OreRobotOreCost                             int
	ClayRobotOreCost                            int
	ObsidianRobotOreCost, ObsidianRobotClayCost int
	GeodeRobotOreCost, GeodeRobotObsidianCost   int
}

func parse(line string) *Blueprint {
	ints := lib.Ints(line)
	return &Blueprint{
		Id:                     ints[0],
		OreRobotOreCost:        ints[1],
		ClayRobotOreCost:       ints[2],
		ObsidianRobotOreCost:   ints[3],
		ObsidianRobotClayCost:  ints[4],
		GeodeRobotOreCost:      ints[5],
		GeodeRobotObsidianCost: ints[6],
	}
}

func (b *Blueprint) maxGeodes(inv Inventory, time int) int {
	if time == 0 {
		return inv.Geodes //can't make any more
	}

	inv.Ore += inv.OreRobots
	inv.Clay += inv.ClayRobots
	inv.Obsidian += inv.ObsidianRobots
	inv.Geodes += inv.GeodeRobots

	var max int

	if inv.Ore >= b.OreRobotOreCost {
		newInv := inv
		newInv.Ore -= b.OreRobotOreCost
		newInv.OreRobots++
		max = lib.Max(max, b.maxGeodes(newInv, time-1))
	}

	if inv.Ore >= b.ClayRobotOreCost {
		newInv := inv
		newInv.Ore -= b.ClayRobotOreCost
		newInv.ClayRobots++
		max = lib.Max(max, b.maxGeodes(newInv, time-1))
	}

	if inv.Ore >= b.ObsidianRobotOreCost && inv.Clay >= b.ObsidianRobotClayCost {
		newInv := inv
		newInv.Ore -= b.ObsidianRobotOreCost
		newInv.Clay -= b.ObsidianRobotClayCost
		newInv.ObsidianRobots++
		max = lib.Max(max, b.maxGeodes(newInv, time-1))
	}

	if inv.Ore >= b.GeodeRobotOreCost && inv.Clay >= b.GeodeRobotObsidianCost {
		newInv := inv
		newInv.Ore -= b.GeodeRobotOreCost
		newInv.Obsidian -= b.GeodeRobotObsidianCost
		newInv.GeodeRobots++
		max = lib.Max(max, b.maxGeodes(newInv, time-1))
	}

	return lib.Max(max, b.maxGeodes(inv, time-1))
}

func solve(lines []*Blueprint) int {
	return lib.SumSlice(lib.Map(lines, func(b *Blueprint) int {
		return b.Id * b.maxGeodes(Inventory{OreRobots: 1}, 24)
	}))
}

func main() {
	solver := lib.Solver[[]*Blueprint, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	solver.Expect("Blueprint 1: Each ore robot costs 4 ore.  Each clay robot costs 2 ore.  Each obsidian robot costs 3 ore and 14 clay.  Each geode robot costs 2 ore and 7 obsidian.\nBlueprint 2:  Each ore robot costs 2 ore.  Each clay robot costs 3 ore.  Each obsidian robot costs 3 ore and 8 clay.  Each geode robot costs 3 ore and 12 obsidian.", 33)
	solver.Solve()
}
