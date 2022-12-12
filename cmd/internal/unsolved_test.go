package internal

import (
	"path/filepath"
	"testing"
)

func TestFirstUnsolvedSolution(t *testing.T) {
	type Expected struct {
		Year, Day, Type string
	}
	for _, testCase := range []struct {
		root     string
		expected Expected
		skips    []skip
	}{
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayOneLeaderboard"),
			expected: Expected{
				Year: "2015",
				Day:  "1",
				Type: "leaderboard",
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayOneLeaderboardWithSecondYear"),
			expected: Expected{
				Year: "2015",
				Day:  "1",
				Type: "leaderboard",
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayOneOptimized"),
			expected: Expected{
				Year: "2015",
				Day:  "1",
				Type: "optimized",
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayThreeLeaderboard"),
			expected: Expected{
				Year: "2015",
				Day:  "3",
				Type: "leaderboard",
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayThreeOptimized"),
			expected: Expected{
				Year: "2015",
				Day:  "3",
				Type: "optimized",
			},
		},
		{
			root:  filepath.Join("testData", "testUnsolved", "firstYear", "dayTwentyFiveLeaderboard"),
			skips: []skip{{Year: singleValueRange(2015), Day: skipRange{Min: 1, Max: 24}}},
			expected: Expected{
				Year: "2015",
				Day:  "25",
				Type: "leaderboard",
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "secondYearNoSkips", "noNextYear"),
			expected: Expected{
				Year: "2016",
				Day:  "1",
				Type: "leaderboard",
			},
		},
		{
			root: filepath.Join("testData", "testUnsolved", "secondYearNoSkips", "someNextYear"),
			expected: Expected{
				Year: "2016",
				Day:  "3",
				Type: "optimized",
			},
		},
		{
			root:  filepath.Join("testData", "testUnsolved", "secondYearSkips", "dayTwentyFiveLeaderboard"),
			skips: []skip{{Year: singleValueRange(2011)}, {Year: singleValueRange(2015), Day: skipRange{Min: 1, Max: 24}}},
			expected: Expected{
				Year: "2015",
				Day:  "25",
				Type: "leaderboard",
			},
		},
	} {
		years := GetLocalSolutionInfo(testCase.root)
		actual := FirstUnsolvedSolution(testCase.root, years, &Skipper{skips: testCase.skips})
		expected := testCase.expected

		if actual.Year.Name != expected.Year {
			t.Errorf("root %v: year: expected %v, actual %v", testCase.root, expected.Year, actual.Year)
		}

		if actual.Day.Name != expected.Day {
			t.Errorf("root %v: day: expected %v, actual %v", testCase.root, expected.Day, actual.Day)
		}

		if actual.Name != expected.Type {
			t.Errorf("root %v: leaderboard: expected %v, actual %v", testCase.root, expected.Type, actual.Name)
		}
	}
}
