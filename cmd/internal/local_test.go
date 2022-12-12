package internal_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/internal"
	"path/filepath"
	"strconv"
	"testing"
)

func TestGetLocalSolutionInfo_Years(t *testing.T) {
	root := filepath.Join("testData", "testLocal", "testYears")
	actual := internal.GetLocalSolutionInfo(root)
	expected := []int{35, 82, 97}

	if len(actual) != len(expected) {
		t.Fatalf("incorrect length: expected %v, actual %v", len(expected), len(actual))
	}

	for i, actualYear := range actual {
		expectedYear := expected[i]

		if actualYear.Number != expectedYear {
			t.Errorf("year %d: inccorrect number: expected %v, actual %v", i, expectedYear, actualYear)
		}

		if actualYear.Name != strconv.Itoa(expectedYear) {
			t.Errorf("year %d: inccorrect name: expected %s, actual %s", i, strconv.Itoa(expectedYear), actualYear.Name)
		}

		if actualYear.Path != filepath.Join(root, strconv.Itoa(expectedYear)) {
			t.Errorf("year %d: inccorrect path: expected %s, actual %s", i, filepath.Join(root, strconv.Itoa(expectedYear)), actualYear.Path)
		}
	}
}

func TestGetLocalSolutionInfo_Days(t *testing.T) {
	root := filepath.Join("testData", "testLocal", "testDays")
	year := internal.GetLocalSolutionInfo(root)[0]

	if len(year.Days) != internal.LastDayNum {
		t.Fatalf("year %s: incorrect number of days: expected %d, got %d", year.Name, internal.LastDayNum, len(year.Days))
	}

	for i := 1; i <= internal.LastDayNum; i++ {
		day := year.Days[i-1]

		if day.Number != i {
			t.Errorf("day %d: incorrect number: expected %d, got %d", i, i, day.Number)
		}

		if day.Name != strconv.Itoa(i) {
			t.Errorf("day %d: incorrect number: expected %s, got %s", i, strconv.Itoa(i), day.Name)
		}

		if day.Path != filepath.Join(root, "2015", strconv.Itoa(i)) {
			t.Errorf("day %d: incorrect path: expected %s, actual %s", i, filepath.Join(root, "2015", strconv.Itoa(i)), day.Path)
		}
	}
}
