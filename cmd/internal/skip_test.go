package internal

import (
	"strings"
	"testing"
)

func singleValueRange(value int) skipRange {
	return skipRange{
		Min: value,
		Max: value,
	}
}

func TestParseSkips_Valid(t *testing.T) {
	rawSkips := "1\n2/3\n4#foo\n\n\n#bar\n5 # bar\n6/7 # baz\n8 \n9-10\n11/12-13\n14-15/16-17"
	expectedSkips := []skip{
		{
			Year: singleValueRange(1),
		},
		{
			Year: singleValueRange(2),
			Day:  singleValueRange(3),
		},
		{
			Year: singleValueRange(4),
		},
		{
			Year: singleValueRange(5),
		},
		{
			Year: singleValueRange(6),
			Day:  singleValueRange(7),
		},
		{
			Year: singleValueRange(8),
		},
		{
			Year: skipRange{
				Min: 9,
				Max: 10,
			},
		},
		{
			Year: singleValueRange(11),
			Day: skipRange{
				Min: 12,
				Max: 13,
			},
		},
		{
			Year: skipRange{
				Min: 14,
				Max: 15,
			},
			Day: skipRange{
				Min: 16,
				Max: 17,
			},
		},
	}

	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}()

	actualSkips := parseSkips(strings.NewReader(rawSkips))

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
			error:   "bad skip: hello: year: isn't a number",
		},
		{
			rawSkip: "foo/bar",
			error:   "bad skip: foo/bar: year: isn't a number",
		},
		{
			rawSkip: "1/bar",
			error:   "bad skip: 1/bar: day: isn't a number",
		},
		{
			rawSkip: "foo-2/bar",
			error:   "bad skip: foo-2/bar: year: min isn't a number",
		},
		{
			rawSkip: "2-foo/bar",
			error:   "bad skip: 2-foo/bar: year: max isn't a number",
		},
		{
			rawSkip: "1/bar-2",
			error:   "bad skip: 1/bar-2: day: min isn't a number",
		},
		{
			rawSkip: "1/2-bar",
			error:   "bad skip: 1/2-bar: day: max isn't a number",
		},
		{
			rawSkip: "1-2-3/4",
			error:   "bad skip: 1-2-3/4: year: range should have at most one separator",
		},
		{
			rawSkip: "1/2-3-4",
			error:   "bad skip: 1/2-3-4: day: range should have at most one separator",
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

			parseSkips(strings.NewReader(testCase.rawSkip))
		}()
	}
}
