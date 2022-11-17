package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
	"os"
	"path"
	"regexp"
)

func input() *os.File {
	input, err := os.Open(path.Join("2015", "14", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type reindeer struct {
	name                        string
	speed, duration, restPeriod int

	prog *progress
}

type progress struct {
	distance, flyTimeLeft, restTimeLeft int
	flying                              bool
}

var parseRegexp = regexp.MustCompile(`(.+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

func parse(parts []string) reindeer {
	return reindeer{
		name:       parts[0],
		speed:      lib.Atoi(parts[1]),
		duration:   lib.Atoi(parts[2]),
		restPeriod: lib.Atoi(parts[3]),
		prog: &progress{
			flyTimeLeft: lib.Atoi(parts[2]),
			flying:      true,
		},
	}
}

func (r reindeer) tick(time int) {
	for time > 0 {
		if r.prog.flying {
			flyTime := lib.Min(time, r.prog.flyTimeLeft)
			r.prog.distance += r.speed * flyTime

			r.prog.flyTimeLeft -= flyTime
			time -= flyTime

			if r.prog.flyTimeLeft == 0 {
				r.prog.restTimeLeft = r.restPeriod
				r.prog.flying = false
			}
		} else {
			restTime := lib.Min(time, r.prog.restTimeLeft)
			r.prog.restTimeLeft -= restTime
			time -= restTime

			if r.prog.restTimeLeft == 0 {
				r.prog.flyTimeLeft = r.duration
				r.prog.flying = true
			}
		}
	}
}

var totalTime int

func solve(reindeers []reindeer) int {
	points := make([]int, len(reindeers))
	for i := 0; i < totalTime; i++ {
		for _, r := range reindeers {
			r.tick(1)
		}

		maxDistance := lib.Max(lib.Map(reindeers, func(r reindeer) int {
			return r.prog.distance
		})...)

		for i, r := range reindeers {
			if r.prog.distance == maxDistance {
				points[i] += 1
			}
		}

	}

	return lib.Max(points...)
}

func main() {
	solver := lib.Solver[[]reindeer, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseRegexp, parse)),
		SolveF: solve,
	}

	totalTime = 1000
	solver.Expect("Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.\nDancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.", 689)

	totalTime = 2503
	solver.Verify(input(), 1102)
}
