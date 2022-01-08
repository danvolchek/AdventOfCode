package lib_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/lib"
	"path/filepath"
	"testing"
)

func TestYearsWithSolutions(t *testing.T) {
	actual := lib.YearsWithSolutions(filepath.Join("testData", "testYears"))
	expected := []int{35, 82, 97}

	if len(actual) != len(expected) {
		t.Fatalf("incorrect length: expected %v, actual %v", len(expected), len(actual))
	}

	for i, actualYear := range actual {
		expectedYear := expected[i]

		if actualYear != expectedYear {
			t.Errorf("entry %v: inccorrect value: expected %v, actual %v", i, expectedYear, actualYear)
		}
	}
}
