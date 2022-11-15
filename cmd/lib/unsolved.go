package lib

const (
	firstYear = 2015
)

func FirstUnsolvedSolution(root string, solutionsToSkip []SkipSolution) Solution {
	for year := firstYear; ; year++ {
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
