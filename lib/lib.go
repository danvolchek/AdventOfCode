package lib

import (
	"golang.org/x/exp/constraints"
	"math"
	"regexp"
	"strconv"
)

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
func Min[T constraints.Ordered](values ...T) T {
	var min T

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

// Abs returns the absolute value of v.
func Abs[T constraints.Integer | constraints.Float](v T) T {
	if v < 0 {
		return -v
	}

	return v
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

// Clone returns a copy of items.
func Clone[T any](items []T) []T {
	result := make([]T, len(items))
	copy(result, items)
	return result
}

// Keys returns the keys of a map. It does not return the items in a consistent order.
func Keys[T comparable, V any](items map[T]V) []T {
	keys := make([]T, len(items))

	i := 0
	for key := range items {
		keys[i] = key
		i += 1
	}

	return keys
}

// Map converts a slice of items of type T to type V using a mapping function.
func Map[T, V any](items []T, mapper func(T) V) []V {
	result := make([]V, len(items))

	for i, t := range items {
		result[i] = mapper(t)
	}

	return result
}

// Filter returns entries in items that the filter function returns true for.
func Filter[T any](items []T, filter func(T) bool) []T {
	var result []T

	for _, item := range items {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}

// Reverse returns a new slice which is items in reverse order.
func Reverse[T any](items []T) []T {
	result := make([]T, len(items))
	for i, item := range items {
		result[len(result)-i-1] = item
	}
	return result

}

// SumSlice returns the sum of items.
func SumSlice[T constraints.Integer | constraints.Float](items []T) T {
	var sum T
	for _, val := range items {
		sum += val
	}

	return sum
}

// MulSlice returns the multiplication of items.
func MulSlice[T constraints.Integer | constraints.Float](items []T) T {
	var sum T = 1

	for _, val := range items {
		sum *= val
	}

	return sum
}

// MinSlice returns the min of items.
func MinSlice[T constraints.Ordered](items []T) T {
	var min T
	for i, val := range items {
		if i == 0 || val < min {
			min = val
		}
	}

	return min
}

// MaxSlice returns the max of items.
func MaxSlice[T constraints.Ordered](items []T) T {
	var max T
	for i, val := range items {
		if i == 0 || val > max {
			max = val
		}
	}

	return max
}

// Adjacent returns the adjacent items in a 2d grid of items.
// Diag controls whether diagonals are considered as adjacent.
func Adjacent[T any](diag bool, row, col int, grid [][]T) []T {
	return Map(
		AdjacentPosBounds(diag, row, col, grid),
		func(pos Pos) T {
			return grid[pos.Row][pos.Col]
		})
}

// Pos represents a two-dimensional position.
type Pos struct {
	Row, Col int
}

// Add returns a new Pos that's the sum of the two rows and cols.
func (p Pos) Add(o Pos) Pos {
	return Pos{Row: p.Row + o.Row, Col: p.Col + o.Col}
}

// Min returns a new Pos that's the minimum of the two rows and cols.
func (p Pos) Min(o Pos) Pos {
	return Pos{Row: Min(p.Row, o.Row), Col: Min(p.Col, o.Col)}
}

// Max returns a new Pos that's the maximum of the two rows and cols.
func (p Pos) Max(o Pos) Pos {
	return Pos{Row: Max(p.Row, o.Row), Col: Max(p.Col, o.Col)}
}

// AdjacentPos returns the adjacent positions in a 2d grid of items (excluding bounds checks).
// Diag controls whether diagonals are considered as adjacent.
func AdjacentPos(diag bool, row, col int) []Pos {
	var results []Pos

	for di := -1; di <= 1; di += 1 {
		for dj := -1; dj <= 1; dj += 1 {
			if di == 0 && dj == 0 {
				continue
			}

			if !diag && Abs(di)+Abs(dj) == 2 {
				continue
			}

			adjI := row + di
			adjJ := col + dj

			results = append(results, Pos{Row: adjI, Col: adjJ})
		}
	}

	return results
}

// AdjacentPosBounds returns the adjacent positions in a 2d grid of items within the given bounds.
// Diag controls whether diagonals are considered as adjacent.
func AdjacentPosBounds[T any](diag bool, row, col int, grid [][]T) []Pos {
	return Filter(AdjacentPos(diag, row, col), func(pos Pos) bool {
		return !(pos.Row < 0 || pos.Col < 0 || pos.Row >= len(grid) || pos.Col >= len(grid[row]))
	})
}

// Permutations returns all possible permutations of items.
// It uses https://en.wikipedia.org/wiki/Heap%27s_algorithm.
//
// Note: Almost every time, there is a more efficient solution than generating every possible permutation of a list.
// I.e. there's a way to filter out potential options that are known not right.
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

// Subsets returns all subsets of items by enumerating the 2**n-1 possible combinations, using the bits in the
// counter as whether to include an item or not.
func Subsets[T any](items []T) [][]T {
	if len(items) == 0 {
		return nil
	}

	var result [][]T

	n := len(items)

	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		var subset []T

		for bit := 0; bit < n; bit++ {
			if (i>>bit)&1 == 1 {
				subset = append(subset, items[bit])

			}
		}

		result = append(result, subset)
	}

	return result
}

var intsReg = regexp.MustCompile(`-?\d+`)

// Ints returns all the integers in line.
func Ints(line string) []int {
	numbers := intsReg.FindAllString(line, -1)

	result := make([]int, len(numbers))

	for i, number := range numbers {
		result[i] = Atoi(number)
	}

	return result
}

// Int is like Ints except it returns the first int found.
// See Atoi to convert a string representation of a number into an int.
func Int(line string) int {
	ints := Ints(line)
	if len(ints) == 0 {
		panic("no ints in '" + line + "'")
	}

	return ints[0]
}

// Pow is like math.Pow but for integers.
func Pow(num, pow int) int {
	return int(math.Pow(float64(num), float64(pow)))
}
