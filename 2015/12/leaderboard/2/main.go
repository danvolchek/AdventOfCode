package main

import (
	"encoding/json"

	"github.com/danvolchek/AdventOfCode/lib"
)

func sum(value any) int {
	switch value.(type) {
	case map[string]any:
		valueMap := value.(map[string]any)
		for _, val := range valueMap {
			if val == "red" {
				return 0
			}
		}

		total := 0
		for _, val := range valueMap {
			total += sum(val)
		}

		return total
	case []any:
		valueSlice := value.([]any)
		total := 0
		for _, val := range valueSlice {
			total += sum(val)
		}

		return total
	case float64:
		return int(value.(float64))
	default: // string, bool, null
		return 0
	}
}

func solve(input []byte) int {
	var value any

	lib.NoPanic(json.Unmarshal(input, &value))

	return sum(value)
}

func main() {
	solver := lib.Solver[[]byte, int]{
		ParseF: lib.ParseBytes,
		SolveF: solve,
	}

	solver.Expect(`[1,2,3]`, 6)
	solver.Expect(`{"a":2,"b":4}`, 6)
	solver.Expect(`[[[3]]]`, 3)
	solver.Expect(`{"a":{"b":4},"c":-1}`, 3)
	solver.Expect(`{"a":[-1,1]}`, 0)
	solver.Expect(`[-1,{"a":1}]`, 0)
	solver.Expect(`[]`, 0)
	solver.Expect(`{}`, 0)
	solver.Expect(`[1,{"c":"red","b":2},3]`, 4)
	solver.Expect(`{"d":"red","e":[1,2,3,4],"f":5}`, 0)
	solver.Expect(`[1,"red",5]`, 6)

	solver.Verify(68466)
}
