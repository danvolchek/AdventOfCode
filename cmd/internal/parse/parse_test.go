package parse_test

import (
	"encoding/json"
	"github.com/danvolchek/AdventOfCode/cmd/internal/parse"
	"reflect"
	"strconv"
	"testing"
)

var expected = []parse.Year{
	{
		Num: "4567",
		Days: []parse.Day{
			{
				Num: "4",
				PartOne: parse.Part{
					OptimizedSolutionPath: "testData/4567/4/main.go",
				},
				PartTwo: parse.Part{
					OptimizedSolutionPath: "testData/4567/4/main.go",
				},
			},
			{
				Num: "8",
				PartTwo: parse.Part{
					LeaderboardSolutionPath: "testData/4567/8/leaderboard/2/main.go",
				},
			},
			{
				Num: "16",
				PartOne: parse.Part{
					OptimizedSolutionPath: "testData/4567/16/optimized/1/main.go",
				},
			},
		},
	},
	{
		Num: "1234",
		Days: []parse.Day{
			{
				Num: "1",
				PartOne: parse.Part{
					LeaderboardSolutionPath: "testData/1234/1/leaderboard/1/main.go",
					OptimizedSolutionPath:   "testData/1234/1/optimized/1/main.go",
				},
				PartTwo: parse.Part{
					LeaderboardSolutionPath: "testData/1234/1/leaderboard/2/main.go",
					OptimizedSolutionPath:   "testData/1234/1/optimized/2/main.go",
				},
			},
			{
				Num: "2",
				PartOne: parse.Part{
					LeaderboardSolutionPath: "testData/1234/2/leaderboard/1/main.go",
				},
				PartTwo: parse.Part{
					OptimizedSolutionPath: "testData/1234/2/optimized/2/main.go",
				},
			},
		},
	},
}

func TestSolutionInformation(t *testing.T) {
	fill(expected)

	result := parse.SolutionInformation("testData")

	if len(result) != len(expected) {
		t.Fatalf("got %v years, wanted %v years", len(result), len(expected))
	}

	for i := 0; i < len(expected); i++ {
		gotYear := result[i]
		wantYear := expected[i]

		if gotYear.Num != wantYear.Num {
			t.Errorf("year num: got %v, wanted %v", gotYear.Num, wantYear.Num)
			continue
		}

		if len(gotYear.Days) != len(wantYear.Days) {
			t.Errorf("year %v: got %v days, wanted %v days", gotYear.Num, len(gotYear.Days), len(wantYear.Days))
			continue
		}

		for i := 0; i < 25; i++ {
			gotDay := gotYear.Days[i]
			wantDay := wantYear.Days[i]

			if gotDay.Num != wantDay.Num {
				t.Errorf("day num: got %v, wanted %v", gotDay.Num, wantDay.Num)
				continue
			}

			if !reflect.DeepEqual(gotDay, wantDay) {
				t.Errorf("year %v: day %v: got\n%s\nwant\n%s\n", wantYear.Num, wantDay.Num, marshal(gotDay), marshal(wantDay))
			}
		}
	}
}

func fill(expected []parse.Year) {
	for yearNum, year := range expected {
		newDays := make([]parse.Day, 25)

		for i := 1; i <= 25; i++ {
			stringNum := strconv.Itoa(i)

			day, found := findDay(stringNum, year.Days)
			if !found {
				day = parse.Day{
					Num: stringNum,
				}
			}

			newDays[i-1] = day
		}

		expected[yearNum].Days = newDays
	}
}

func findDay(num string, days []parse.Day) (parse.Day, bool) {
	for _, day := range days {
		if day.Num == num {
			return day, true
		}
	}

	return parse.Day{}, false
}

func marshal(v interface{}) []byte {
	result, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return result
}
