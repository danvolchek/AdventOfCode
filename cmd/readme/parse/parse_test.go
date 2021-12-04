package parse_test

import (
	"encoding/json"
	"github.com/danvolchek/AdventOfCode/cmd/readme/parse"
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
		t.Fatalf("missing years entirely")
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
		dayNum := 0
		newDays := make([]parse.Day, 25)
		for i := 1; i <= 25; i++ {
			stringNum := strconv.Itoa(i)
			if dayNum >= len(year.Days) || year.Days[dayNum].Num != stringNum {
				newDays[i-1] = parse.Day{
					Num: stringNum,
				}
			} else {
				newDays[i-1] = year.Days[dayNum]
				dayNum++
			}
		}

		expected[yearNum].Days = newDays
	}
}

func marshal(v interface{}) []byte {
	result, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}

	return result
}