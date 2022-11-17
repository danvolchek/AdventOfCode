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
}

var parseRegexp = regexp.MustCompile(`(.+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)

func parse(parts []string) reindeer {
	return reindeer{
		name:       parts[0],
		speed:      lib.Atoi(parts[1]),
		duration:   lib.Atoi(parts[2]),
		restPeriod: lib.Atoi(parts[3]),
	}
}

func (r reindeer) distance(time int) int {
	distance := 0
	flying := true
	flyTimeLeft := r.duration
	restTimeLeft := 0

	for time > 0 {
		if flying {
			flyTime := lib.Min(time, flyTimeLeft)
			distance += r.speed * flyTime

			flyTimeLeft -= flyTime
			time -= flyTime

			if flyTimeLeft == 0 {
				restTimeLeft = r.restPeriod
				flying = false
			}
		} else {
			restTime := lib.Min(time, restTimeLeft)
			restTimeLeft -= restTime
			time -= restTime

			if restTimeLeft == 0 {
				flyTimeLeft = r.duration
				flying = true
			}
		}
	}

	return distance
}

var totalTime int

func solve(reindeers []reindeer) int {
	distances := lib.Map(reindeers, func(reindeer reindeer) int {
		return reindeer.distance(totalTime)
	})

	return lib.Max(distances...)
}

func main() {
	solver := lib.Solver[[]reindeer, int]{
		ParseF: lib.ParseLine(lib.ParseRegexp(parseRegexp, parse)),
		SolveF: solve,
	}

	totalTime = 1000
	solver.Expect("Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.", 1120)
	solver.Expect("Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.", 1056)

	totalTime = 2503
	solver.Verify(input(), 2640)
}
