package lib

import (
	"fmt"
	"sort"
)

func FirstUnsolvedSolution(root string, solutionsToSkip []SkipSolution) Solution {
	years := YearsWithSolutions(root)
	if len(years) == 0 {
		panic(fmt.Sprintf("No years found in %v", root))
	}

	sort.Ints(years)
	smallestYear := years[0]

	for year := smallestYear; ; year++ {
		for day := 1; day <= 25; day++ {
			for _, leaderboard := range []bool{true, false} {
				solution := Solution{
					Year:        year,
					Day:         day,
					Leaderboard: leaderboard,
				}

				_, solExists := solution.Path(root)

				if solExists {
					continue
				}

				doSkip := false

				for _, skip := range solutionsToSkip {
					if skip.ShouldSkip(solution) {
						doSkip = true
						break
					}
				}

				if doSkip {
					continue
				}

				return solution
			}
		}
	}
}
