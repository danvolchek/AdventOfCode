package lib_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/lib"
	"path/filepath"
	"testing"
)

type solutionTestCase struct {
	s            lib.Solution
	expectedPath string
	expectedOkay bool
}

var solutionTestRoot = filepath.Join("testData", "testSolution")

func TestSolution_Path(t *testing.T) {
	for _, testCase := range []solutionTestCase{
		{
			s: lib.Solution{
				Year:        2015,
				Day:         5,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "5", "leaderboard"),
			expectedOkay: true,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         5,
				Leaderboard: false,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "5", "optimized"),
			expectedOkay: false,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         6,
				Leaderboard: false,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "6", "optimized"),
			expectedOkay: true,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         6,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "6", "leaderboard"),
			expectedOkay: false,
		},
	} {
		actualPath, actualOk := testCase.s.Path(solutionTestRoot)
		if actualOk != testCase.expectedOkay {
			t.Errorf("input %+v: ok: expected %v, actual %v", testCase.s, testCase.expectedOkay, actualOk)
		}

		if actualPath != testCase.expectedPath {
			t.Errorf("input %+v: path: expected %v, actual %v", testCase.s, testCase.expectedPath, actualPath)
		}
	}
}

func TestSolution_PartOne(t *testing.T) {
	for _, testCase := range []solutionTestCase{
		{
			s: lib.Solution{
				Year:        2015,
				Day:         5,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "5", "leaderboard", "1", "main.go"),
			expectedOkay: true,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         7,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "7", "leaderboard", "1", "main.go"),
			expectedOkay: false,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         6,
				Leaderboard: false,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "6", "optimized", "1", "main.go"),
			expectedOkay: true,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         8,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "8", "leaderboard", "1", "main.go"),
			expectedOkay: false,
		},
	} {
		actualPath, actualOk := testCase.s.PartOne(solutionTestRoot)
		if actualOk != testCase.expectedOkay {
			t.Errorf("input %+v: ok: expected %v, actual %v", testCase.s, testCase.expectedOkay, actualOk)
		}

		if actualPath != testCase.expectedPath {
			t.Errorf("input %+v: path: expected %v, actual %v", testCase.s, testCase.expectedPath, actualPath)
		}
	}
}

func TestSolution_PartTwo(t *testing.T) {
	for _, testCase := range []solutionTestCase{
		{
			s: lib.Solution{
				Year:        2015,
				Day:         7,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "7", "leaderboard", "2", "main.go"),
			expectedOkay: true,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         5,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "5", "leaderboard", "2", "main.go"),
			expectedOkay: false,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         8,
				Leaderboard: false,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "8", "optimized", "2", "main.go"),
			expectedOkay: true,
		},
		{
			s: lib.Solution{
				Year:        2015,
				Day:         6,
				Leaderboard: true,
			},
			expectedPath: filepath.Join(solutionTestRoot, "2015", "6", "leaderboard", "2", "main.go"),
			expectedOkay: false,
		},
	} {
		actualPath, actualOk := testCase.s.PartTwo(solutionTestRoot)
		if actualOk != testCase.expectedOkay {
			t.Errorf("input %+v: ok: expected %v, actual %v", testCase.s, testCase.expectedOkay, actualOk)
		}

		if actualPath != testCase.expectedPath {
			t.Errorf("input %+v: path: expected %v, actual %v", testCase.s, testCase.expectedPath, actualPath)
		}
	}
}

func TestSolution_String(t *testing.T) {
	for _, testCase := range []struct {
		s        lib.Solution
		expected string
	}{
		{
			s: lib.Solution{
				Year:        2015,
				Day:         6,
				Leaderboard: true,
			},
			expected: "{Year: 2015, Day: 6, Type: Leaderboard}",
		},
		{
			s: lib.Solution{
				Year:        2016,
				Day:         8,
				Leaderboard: false,
			},
			expected: "{Year: 2016, Day: 8, Type: Optimized}",
		},
	} {
		actual := testCase.s.String()

		if testCase.expected != actual {
			t.Errorf("input year %v, day %v, lb %v: expected %v, actual %v", testCase.s.Year, testCase.s.Day, testCase.s.Leaderboard, testCase.expected, actual)
		}
	}
}
