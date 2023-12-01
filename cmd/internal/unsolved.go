package internal

import (
	"strconv"
)

func FirstUnsolvedSolution(root string, years []*Year, skipper *Skipper) *Type {
	for _, year := range years {
		for _, day := range year.Days {
			for _, typ := range day.Types {
				if !typ.Exists() && !skipper.Skip(typ.Day.Year.Name, typ.Day.Name) {
					return typ
				}
			}
		}
	}

	if len(years) == 0 {
		return NewType(root, FirstYear, FirstDay, TypeLeaderboard)
	}

	lastYear := years[len(years)-1]
	return NewType(root, strconv.Itoa(lastYear.Number+1), FirstDay, TypeLeaderboard)
}
