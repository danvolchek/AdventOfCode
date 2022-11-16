package lib

import "strconv"

// Must return value if err is non-nil and panics otherwise.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

// Min returns the smallest value in values.
func Min(values ...int) int {
	min := 0

	for i, value := range values {
		if i == 0 || value < min {
			min = value
		}
	}

	return min
}

// Max returns the largest value in values.
func Max(values ...int) int {
	max := 0

	for i, value := range values {
		if i == 0 || value > max {
			max = value
		}
	}

	return max
}

// Atoi is a convenience wrapper on [strconv.Atoi] that panics if it fails.
func Atoi(s string) int {
	return Must(strconv.Atoi(s))
}

// Remove returns a new slice that has the item at index removed from items.
func Remove[T any](items []T, index int) []T {
	result := make([]T, len(items)-1)

	copy(result, items[:index])
	copy(result[index:], items[index+1:])

	return result
}
