package lib_test

import (
	"github.com/danvolchek/AdventOfCode/cmd/lib"
	"strings"
	"testing"
)

func TestParseSkips_Valid(t *testing.T) {
	rawSkips := "1\n2/3\n4#foo\n\n\n#bar\n5 # bar\n6/7 # baz\n8 "
	expectedSkips := []lib.SkipSolution{
		{
			Year: 1,
			Day:  0,
		},
		{
			Year: 2,
			Day:  3,
		},
		{
			Year: 4,
			Day:  0,
		},
		{
			Year: 5,
			Day:  0,
		},
		{
			Year: 6,
			Day:  7,
		},
		{
			Year: 8,
			Day:  0,
		},
	}

	actualSkips := lib.ParseSkips(strings.NewReader(rawSkips))

	if len(expectedSkips) != len(actualSkips) {
		t.Fatalf("expected: len %v, actual: len %v", len(expectedSkips), len(actualSkips))
	}

	for i, expected := range expectedSkips {
		actual := actualSkips[i]

		if actual.Year != expected.Year {
			t.Errorf("skip %v: year: expected %v, actual %v", i, expected.Year, actual.Year)
		}

		if actual.Day != expected.Day {
			t.Errorf("skip %v: day: expected %v, actual %v", i, expected.Day, actual.Day)
		}
	}
}

func TestParseSkips_Invalid(t *testing.T) {
	for _, testCase := range []struct {
		rawSkip string
		error   string
	}{
		{
			rawSkip: "a/b/c",
			error:   "bad skip: a/b/c: wrong number of parts",
		},
		{
			rawSkip: "hello",
			error:   "bad skip: hello: year isn't a number",
		},
		{
			rawSkip: "foo/bar",
			error:   "bad skip: foo/bar: year isn't a number",
		},
		{
			rawSkip: "1/bar",
			error:   "bad skip: 1/bar: day isn't a number",
		},
	} {
		func() {
			defer func() {
				err := recover()
				if err == nil {
					t.Errorf("skip %v: expected an error", testCase.rawSkip)
				}

				if !strings.Contains(err.(string), testCase.error) {
					t.Errorf("skip %v: expected error '%v', actual '%v'", testCase.rawSkip, testCase.error, err.(string))
				}
			}()

			lib.ParseSkips(strings.NewReader(testCase.rawSkip))
		}()
	}
}
