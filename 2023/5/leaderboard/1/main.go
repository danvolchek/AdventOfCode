package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Almanac struct {
	seeds []int

	maps map[string]Map
}

type Map struct {
	source, dest string

	ranges []Range
}

func (m Map) mapping(value int) int {
	for _, r := range m.ranges {
		if !(value >= r.sourceStart && value <= r.sourceStart+r.length) {
			continue
		}

		return (value - r.sourceStart) + r.destStart
	}

	return value
}

type Range struct {
	destStart, sourceStart, length int
}

func parseSeeds(chunk string, almanac *Almanac) {
	almanac.seeds = lib.Ints(chunk)
	almanac.maps = make(map[string]Map)
}

func parseMap(chunk string, almanac *Almanac) {
	lines := strings.Split(chunk, "\n")

	names := strings.Split(strings.ReplaceAll(lines[0], " map:", ""), "-to-")

	var ranges []Range

	for _, rawRange := range lines[1:] {
		nums := lib.Ints(rawRange)
		ranges = append(ranges, Range{
			destStart:   nums[0],
			sourceStart: nums[1],
			length:      nums[2],
		})
	}

	almanac.maps[names[0]] = Map{
		source: names[0],
		dest:   names[1],
		ranges: ranges,
	}
}

func solve(almanac Almanac) int {
	lowestLocation := 9999999999999

	for _, seed := range almanac.seeds {
		value := seed
		current := "seed"
		target := "location"

		for current != target {
			newMap := almanac.maps[current]

			value = newMap.mapping(value)
			current = newMap.dest
		}

		lowestLocation = min(lowestLocation, value)
	}

	return lowestLocation
}

func main() {
	solver := lib.Solver[Almanac, int]{
		ParseF: lib.ParseChunksUnique[Almanac](parseSeeds, parseMap),
		SolveF: solve,
	}

	solver.Expect("seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4", 35)
	solver.Verify(289863851)
}
