package internal_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSolutionsDirectory_Create(t *testing.T) {
	root := filepath.Join("testData", "testLocal")
	solutions := internal.NewSolutionsDirectory(root)
	expectedYears := []string{"35", "82", "97"}

	if len(solutions.Years()) != len(expectedYears) {
		t.Fatalf("years: incorrect length: expected %v, actual %v", len(expectedYears), len(solutions.Years()))
	}

	for _, expectedYear := range expectedYears {
		solutions, ok := solutions.Solutions[expectedYear]
		if !ok {
			t.Errorf("year %s: not found in solution directory", expectedYear)
			continue
		}

		if expectedYear != "35" {
			continue
		}

		expectedDays := []string{"5", "6", "7", "8"}

		for _, expectedDay := range expectedDays {

			checkError := func(descriptor string, actual, expected any) {
				if !reflect.DeepEqual(actual, expected) {
					t.Helper()
					t.Errorf("year %s: day %s: %s: actual %v, expected %v", expectedYear, expectedDay, descriptor, actual, expected)
				}
			}

			solution, ok := solutions[expectedDay]
			checkError("presence in map", ok, true)
			if !ok {
				continue
			}

			checkError("Exists", solution.Exists, true)
			checkError("Path", solution.Path, filepath.Join(root, expectedYear, expectedDay))
			checkError("Year", solution.Year, expectedYear)
			checkError("Day", solution.Day, expectedDay)

			checkError("Input exists", solution.Input.Exists, expectedDay == "5")
			checkError("Input path", solution.Input.Path, filepath.Join(root, expectedYear, expectedDay, "input.txt"))

			checkError("Lb exists", solution.Leaderboard.Exists, expectedDay == "5" || expectedDay == "7")
			checkError("Opt exists", solution.Optimized.Exists, expectedDay == "6" || expectedDay == "8")

			checkError("Lb path", solution.Leaderboard.Path, filepath.Join(root, expectedYear, expectedDay, "leaderboard"))
			checkError("Opt path", solution.Optimized.Path, filepath.Join(root, expectedYear, expectedDay, "optimized"))

			checkError("Lb Part 1 exists", solution.Leaderboard.PartOne.Exists, expectedDay == "5")
			checkError("Lb Part 2 exists", solution.Leaderboard.PartTwo.Exists, expectedDay == "7")
			checkError("Opt Part 1 exists", solution.Optimized.PartOne.Exists, expectedDay == "6")
			checkError("Opt Part 2 exists", solution.Optimized.PartTwo.Exists, expectedDay == "8")

			checkError("Lb Part 1 path", solution.Leaderboard.PartOne.Path, filepath.Join(root, expectedYear, expectedDay, "leaderboard", "1"))
			checkError("Lb Part 2 path", solution.Leaderboard.PartTwo.Path, filepath.Join(root, expectedYear, expectedDay, "leaderboard", "2"))
			checkError("Opt Part 1 path", solution.Optimized.PartOne.Path, filepath.Join(root, expectedYear, expectedDay, "optimized", "1"))
			checkError("Opt Part 2 path", solution.Optimized.PartTwo.Path, filepath.Join(root, expectedYear, expectedDay, "optimized", "2"))

			checkError("Lb Part 1 main exists", solution.Leaderboard.PartOne.Main.Exists, expectedDay == "5")
			checkError("Lb Part 2 main exists", solution.Leaderboard.PartTwo.Main.Exists, expectedDay == "7")
			checkError("Opt Part 1 main exists", solution.Optimized.PartOne.Main.Exists, expectedDay == "6")
			checkError("Opt Part 2 main exists", solution.Optimized.PartTwo.Main.Exists, expectedDay == "8")

			checkError("Lb Part 1 main path", solution.Leaderboard.PartOne.Main.Path, filepath.Join(root, expectedYear, expectedDay, "leaderboard", "1", "main.go"))
			checkError("Lb Part 2 main path", solution.Leaderboard.PartTwo.Main.Path, filepath.Join(root, expectedYear, expectedDay, "leaderboard", "2", "main.go"))
			checkError("Opt Part 1 main path", solution.Optimized.PartOne.Main.Path, filepath.Join(root, expectedYear, expectedDay, "optimized", "1", "main.go"))
			checkError("Opt Part 2 main path", solution.Optimized.PartTwo.Main.Path, filepath.Join(root, expectedYear, expectedDay, "optimized", "2", "main.go"))
		}
	}
}

func TestSolutionsDirectory_FirstUnsolvedSolutionType(t *testing.T) {
	type Expected struct {
		Year, Day, Type string
	}
	for _, testCase := range []struct {
		root     string
		expected Expected
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
			root: filepath.Join("testData", "testUnsolved", "firstYear", "dayTwentyFiveLeaderboard"),
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
			root: filepath.Join("testData", "testUnsolved", "secondYearSkips", "dayTwentyFiveLeaderboard"),
			expected: Expected{
				Year: "2015",
				Day:  "25",
				Type: "leaderboard",
			},
		},
	} {
		solutions := internal.NewSolutionsDirectory(testCase.root)
		solution, solutionType := solutions.FirstUnsolvedSolutionType()
		expected := testCase.expected

		if solution.Year != expected.Year {
			t.Errorf("root %v: year: expected %v, actual %v", testCase.root, expected.Year, solution.Year)
		}

		if solution.Day != expected.Day {
			t.Errorf("root %v: day: expected %v, actual %v", testCase.root, expected.Day, solution.Day)
		}

		if solutionType != expected.Type {
			t.Errorf("root %v: type: expected %v, actual %v", testCase.root, expected.Type, solutionType)
		}
	}
}
