package lib_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/lib"
	"path/filepath"
	"testing"
)

func TestFirstUnsolvedSolution(t *testing.T) {
	for _, testCase := range []struct {
		root     string
		expected lib.Solution
		skips    []lib.SkipSolution
	}{
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayOneLeaderboard"),
			expected: lib.Solution{
				Year:        2012,
				Day:         1,
				Leaderboard: true,
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayOneLeaderboardWithSecondYear"),
			expected: lib.Solution{
				Year:        2012,
				Day:         1,
				Leaderboard: true,
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayOneOptimized"),
			expected: lib.Solution{
				Year:        2012,
				Day:         1,
				Leaderboard: false,
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayThreeLeaderboard"),
			expected: lib.Solution{
				Year:        2012,
				Day:         3,
				Leaderboard: true,
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayThreeOptimized"),
			expected: lib.Solution{
				Year:        2012,
				Day:         3,
				Leaderboard: false,
			},
		},
		{
			root:  filepath.Join("testData", "testUnsolved", "firstYear", "dayTwentyFiveLeaderboard"),
			skips: []lib.SkipSolution{{Year: singleValueRange(2012), Day: lib.SkipRange{Min: 1, Max: 24}}},
			expected: lib.Solution{
				Year:        2012,
				Day:         25,
				Leaderboard: true,
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "secondYearNoSkips", "noNextYear"),
			expected: lib.Solution{
				Year:        2013,
				Day:         1,
				Leaderboard: true,
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "secondYearNoSkips", "someNextYear"),
			expected: lib.Solution{
				Year:        2013,
				Day:         3,
				Leaderboard: false,
			},
		},
		{
			root:  filepath.Join("testData", "testUnsolved", "secondYearSkips", "dayTwentyFiveLeaderboard"),
			skips: []lib.SkipSolution{{Year: singleValueRange(2011)}, {Year: singleValueRange(2012), Day: lib.SkipRange{Min: 1, Max: 24}}},
			expected: lib.Solution{
				Year:        2012,
				Day:         25,
				Leaderboard: true,
			},
		},
	} {
		actual := lib.FirstUnsolvedSolution(testCase.root, testCase.skips)
		expected := testCase.expected

		if actual.Year != expected.Year {
			t.Errorf("root %v: year: expected %v, actual %v", testCase.root, expected.Year, actual.Year)
		}

		if actual.Day != expected.Day {
			t.Errorf("root %v: day: expected %v, actual %v", testCase.root, expected.Day, actual.Day)
		}

		if actual.Leaderboard != expected.Leaderboard {
			t.Errorf("root %v: leaderboard: expected %v, actual %v", testCase.root, expected.Leaderboard, actual.Leaderboard)
		}
	}
}
