package lib

import (
	"os"
	"strconv"
)

func YearsWithSolutions(root string) []int {
	var years []int

	entries, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		year, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}

		years = append(years, year)
	}

	return years
}
