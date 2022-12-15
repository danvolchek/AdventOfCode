package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"golang.org/x/exp/slices"
)

type Pos struct {
	x, y int
}

type Reading struct {
	Sensor        Pos
	ClosestBeacon Pos
}

func parse(line string) Reading {
	nums := lib.Ints(line)
	return Reading{
		Sensor: Pos{
			x: nums[0],
			y: nums[1],
		},
		ClosestBeacon: Pos{
			x: nums[2],
			y: nums[3],
		},
	}
}

func dist(a, b Pos) int {
	return lib.Abs(a.x-b.x) + lib.Abs(a.y-b.y)
}

func coverageInRow(source Pos, maxD int, targRow int) (int, int, bool) {
	distToRow := dist(source, Pos{x: source.x, y: targRow})

	if distToRow > maxD {
		return 0, 0, false
	}

	remaining := maxD - distToRow

	return source.x - remaining, source.x + remaining, true
}

var maxRow int

func solve(lines []Reading) int {
	var covered []func(targRow int) (int, int, bool)

	for _, line := range lines {
		d := dist(line.Sensor, line.ClosestBeacon)

		q := line.Sensor
		covered = append(covered, func(targRow int) (int, int, bool) {
			return coverageInRow(q, d, targRow)
		})
	}

	//s := time.Now()

	for y := 0; y <= maxRow; y++ {
		rc := coverageForRow(covered, y)
		//fmt.Println(rc)
		rc = simplify(rc)
		//fmt.Println(rc, sz(rc, 0, maxRow))

		//fmt.Println("to find right row", time.Now().Sub(s))

		//if sz(rc, 0, maxRow) == maxRow+1 {
		//	continue
		//}

		// faster than size check, always true?
		if len(rc) != 2 {
			continue
		}

		slices.SortFunc(rc, func(r1, r2 rangey) bool {
			return r1.s < r2.s
		})

		return (rc[0].e+1)*4000000 + y

		/*for i, rr := range rc[1:] {
			if rc[i].e != rr.s && rc[i].e < maxRow && rc[i].e >= -1 {
				fmt.Println(y)
				return (rc[i].e+1)*4000000 + y
			}
		}*/

	}

	panic("not found")
}

func (r rangey) intersect(o rangey) bool {
	return (r.e >= o.s && r.s <= o.e) || (o.e >= r.s && o.s <= r.e)
}

func simplify(ranges []rangey) []rangey {
	for start := 0; start < len(ranges)-1; start++ {
		if len(ranges) == 1 {
			return ranges
		}

		toMerge := ranges[start]

		for i := start + 1; i < len(ranges); i++ {
			if toMerge.intersect(ranges[i]) {
				//fmt.Println("intersect", toMerge, ranges[i])
				newR := rangey{lib.Min(toMerge.s, ranges[i].s), lib.Max(toMerge.e, ranges[i].e)}

				ranges = lib.Remove(ranges, i)
				ranges[start] = newR
				start = -1
				break
			}
		}
	}

	return ranges
}

func sz(ranges []rangey, min, max int) int {
	return lib.SumSlice(lib.Map(ranges, func(r rangey) int {
		s, e := r.s, r.e
		if s < min {
			s = min
		}
		if e > max {
			e = max
		}
		return e - s + 1
	}))
}

type rangey struct{ s, e int }

func coverageForRow(covered []func(targRow int) (int, int, bool), row int) []rangey {
	var result []rangey

	for _, c := range covered {
		minX, maxX, ok := c(row)
		if !ok {
			continue
		}

		result = append(result, rangey{minX, maxX})

	}

	return result
}

func main() {
	solver := lib.Solver[[]Reading, int]{
		ParseF: lib.ParseLine(parse),
		SolveF: solve,
	}

	maxRow = 20
	solver.Expect("Sensor at x=2, y=18: closest beacon is at x=-2, y=15\nSensor at x=9, y=16: closest beacon is at x=10, y=16\nSensor at x=13, y=2: closest beacon is at x=15, y=3\nSensor at x=12, y=14: closest beacon is at x=10, y=16\nSensor at x=10, y=20: closest beacon is at x=10, y=16\nSensor at x=14, y=17: closest beacon is at x=10, y=16\nSensor at x=8, y=7: closest beacon is at x=2, y=10\nSensor at x=2, y=0: closest beacon is at x=2, y=10\nSensor at x=0, y=11: closest beacon is at x=2, y=10\nSensor at x=20, y=14: closest beacon is at x=25, y=17\nSensor at x=17, y=20: closest beacon is at x=21, y=22\nSensor at x=16, y=7: closest beacon is at x=15, y=3\nSensor at x=14, y=3: closest beacon is at x=15, y=3\nSensor at x=20, y=1: closest beacon is at x=15, y=3", 56000011)
	maxRow = 4000000
	solver.Verify(11016575214126)
}
