package lib

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