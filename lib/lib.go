package lib

import "strconv"

// Must return value if err is non-nil and panics otherwise.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	return value
}

// NoPanic panics if err is non-nil.
func NoPanic(err error) {
	if err != nil {
		panic(err)
	}
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

// Permutations returns all possible permutations of items.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Permutations[T any](items []T) [][]T {
	var results [][]T

	scratch := make([]T, len(items))
	copy(scratch, items)

	savePerm := func() {
		scratch2 := make([]T, len(items))
		copy(scratch2, scratch)

		results = append(results, scratch2)
	}

	stack := make([]int, len(items))

	savePerm()

	for i := 1; i < len(items); {
		if stack[i] < i {
			if i%2 == 0 {
				scratch[0], scratch[i] = scratch[i], scratch[0]
			} else {
				scratch[stack[i]], scratch[i] = scratch[i], scratch[stack[i]]
			}

			savePerm()

			stack[i] += 1
			i = 1
		} else {
			stack[i] = 0
			i += 1
		}
	}

	return results
}
