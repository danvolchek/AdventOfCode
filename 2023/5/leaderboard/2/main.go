package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"strings"
)

type Almanac struct {
	seeds []Range

	maps map[string]Map
}

type Range struct {
	start, length int
}

// subtract returns the numbers in r that are not in others - the logical XOR.
// Each range in others must be a subset of r.
// Each range in others may overlap with other ranges in others - others need not all be unique subsets of r.
//
// For example:
//
//	  (2 other ranges)               (3 other ranges)
//	r  other        result         r  other        result
//	|   |                          |   ||
//	|   |     ->                   |   |     ->
//	|                 |            |   ||
//	|   |                          |                 |
func (r Range) subtract(others []Range) []Range {
	missing := []Range{r}

	for _, cover := range others {
	outer:
		if len(missing) == 0 {
			break
		}

		for i, miss := range missing {
			intersection, ok := miss.intersect(cover)
			if !ok {
				continue
			}

			parts := miss.remove(intersection)
			newMissing := append(lib.Remove(missing, i), parts...)
			missing = newMissing
			goto outer
		}
	}

	return missing
}

// remove returns the numbers in r that are not in other - the logical XOR. other must be a subset of r.
//
// The cases are:
//
//	 (r starts with other)           (r contains other)            (r ends with other)
//	r  other        result         r  other        result         r  other        result
//	|   |                          |                 |            |                 |
//	|   |     ->                   |   |      ->                  |         ->      |
//	|                |             |                 |            |   |
//	|                |             |                 |            |   |
func (r Range) remove(other Range) []Range {
	var result []Range
	if other.start > r.start {
		//before

		result = append(result, Range{
			start:  r.start,
			length: other.start - r.start,
		})
	}

	if other.start+other.length < r.start+r.length {
		// after

		result = append(result, Range{
			start:  other.start + other.length,
			length: (r.start + r.length) - (other.start + other.length),
		})
	}

	return result
}

// intersect returns a range representing the numbers r and other share - the logical AND.
// other need not be a subset of r.
// True is returned if the intersection is non-empty, otherwise false.
//
// The non-empty intersecting cases are:
//
//	  (other is before r)            (other is in r)                (other is after r)
//	r  other        result         r  other        result         r  other        result
//	     |
//	|    |            |            |
//	|    |            |            |   |             |            |
//	|          ->                  |   |      ->     |            |         ->
//	|                              |                              |
//	|                              |                              |   |             |
//	                                                              |   |             |
//	                                                                  |
func (r Range) intersect(other Range) (Range, bool) {
	if other.start > r.start+r.length || other.start+other.length < r.start {
		return Range{}, false
	}

	startPoint := max(other.start, r.start)
	endPoint := min(other.start+other.length, r.start+r.length)
	length := endPoint - startPoint

	return Range{
		start:  startPoint,
		length: length,
	}, length > 0
}

type Map struct {
	source, dest string

	ranges []MapRange
}

// mapping maps r from source space to destination space. The result can be multiple ranges if r matches
// multiple mapping directives in the Map.
func (m Map) mapping(r Range) []Range {
	var destSpaceRanges []Range
	var sourceSpaceRanges []Range

	for _, mapRange := range m.ranges {
		intersection, ok := mapRange.sourceRange.intersect(r)
		if !ok {
			continue
		}

		destSpaceRanges = append(destSpaceRanges, Range{
			start:  mapRange.destStart + (intersection.start - mapRange.sourceRange.start),
			length: intersection.length,
		})

		sourceSpaceRanges = append(sourceSpaceRanges, intersection)

	}

	destSpaceRanges = append(destSpaceRanges, r.subtract(sourceSpaceRanges)...)

	return destSpaceRanges
}

type MapRange struct {
	destStart   int
	sourceRange Range
}

func parseSeeds(chunk string, almanac *Almanac) {
	rawSeeds := lib.Ints(chunk)
	seedRanges := make([]Range, len(rawSeeds)/2)

	for i := 0; i < len(rawSeeds); i += 2 {
		seedRanges[i/2] = Range{
			start:  rawSeeds[i],
			length: rawSeeds[i+1],
		}
	}
	almanac.seeds = seedRanges
	almanac.maps = make(map[string]Map)
}

func parseMap(chunk string, almanac *Almanac) {
	lines := strings.Split(chunk, "\n")

	names := strings.Split(strings.ReplaceAll(lines[0], " map:", ""), "-to-")

	var ranges []MapRange

	for _, rawRange := range lines[1:] {
		nums := lib.Ints(rawRange)
		ranges = append(ranges, MapRange{
			destStart: nums[0],
			sourceRange: Range{
				start:  nums[1],
				length: nums[2],
			},
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

	for _, seedRange := range almanac.seeds {

		value := []Range{seedRange}

		current := "seed"
		target := "location"

		for current != target {
			newMap := almanac.maps[current]

			var newValue []Range
			for _, currValue := range value {
				newValue = append(newValue, newMap.mapping(currValue)...)
			}

			value = newValue
			current = newMap.dest
		}

		newLowest := lib.Min(lib.Map(value, func(v Range) int {
			return v.start
		})...)

		lowestLocation = min(lowestLocation, newLowest)
	}

	return lowestLocation
}

func main() {
	solver := lib.Solver[Almanac, int]{
		ParseF: lib.ParseChunksUnique[Almanac](parseSeeds, parseMap),
		SolveF: solve,
	}

	solver.Expect("seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4", 46)
	solver.Verify(60568880)
}
