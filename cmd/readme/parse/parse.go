package parse

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

type Year struct {
	Num  string
	Days []Day
}

type Day struct {
	Num string

	PartOne Part
	PartTwo Part
}

type Part struct {
	LeaderboardSolutionPath string
	OptimizedSolutionPath   string
}

const (
	leaderboardFolderName = "leaderboard"
	optimizedFolderName   = "optimized"
	solutionFileName      = "main.go"
)

var numberRegexp = regexp.MustCompile(`^\d+$`)

func SolutionInformation(root string) []Year {
	var years []Year

	files, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		yearNum := file.Name()

		if !numberRegexp.MatchString(yearNum) {
			continue
		}

		year := parseYear(root, yearNum)
		years = append(years, year)
	}

	sort.Slice(years, func(i, j int) bool {
		return toInt(years[i].Num) > toInt(years[j].Num)
	})

	return years
}

func parseYear(root, yearNum string) Year {
	var days []Day

	for i := 1; i <= 25; i++ {
		dayNum := strconv.Itoa(i)
		if !exists(path.Join(root, yearNum, dayNum)) {
			days = append(days, Day{
				Num: dayNum,
			})
			continue
		}

		day := parseDay(root, yearNum, dayNum)
		days = append(days, day)
	}

	sort.Slice(days, func(i, j int) bool {
		return toInt(days[i].Num) < toInt(days[j].Num)
	})

	return Year{
		Num:  yearNum,
		Days: days,
	}
}

func parseDay(root, yearNum, dayNum string) Day {
	partOne := parsePart(root, yearNum, dayNum, "1")
	partTwo := parsePart(root, yearNum, dayNum, "2")

	return Day{
		Num:     dayNum,
		PartOne: partOne,
		PartTwo: partTwo,
	}
}

func parsePart(root, yearNum, dayNum, partNum string) Part {
	// If there are no solution types, assume an optimized solution for both parts exists
	// This currently applies to a few days (e.g. 2019 5, 9, and 13)
	notSplitSolutionPath := filepath.Join(root, yearNum, dayNum, solutionFileName)
	if exists(notSplitSolutionPath) {
		return Part{
			OptimizedSolutionPath: notSplitSolutionPath,
		}
	}
	leaderboardSolutionPath := filepath.Join(root, yearNum, dayNum, leaderboardFolderName, partNum, solutionFileName)
	if !exists(leaderboardSolutionPath) {
		leaderboardSolutionPath = ""
	}

	optimizedSolutionPath := filepath.Join(root, yearNum, dayNum, optimizedFolderName, partNum, solutionFileName)
	if !exists(optimizedSolutionPath) {
		optimizedSolutionPath = ""
	}

	return Part{
		LeaderboardSolutionPath: leaderboardSolutionPath,
		OptimizedSolutionPath:   optimizedSolutionPath,
	}
}

func toInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err)
	}

	return true
}
