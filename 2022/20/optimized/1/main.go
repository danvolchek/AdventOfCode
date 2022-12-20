package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func buildLinkedList(numbers []int) ([]*lib.LinkedList[int], *lib.LinkedList[int]) {
	links := make([]*lib.LinkedList[int], len(numbers))

	var zero *lib.LinkedList[int]
	var current *lib.LinkedList[int]

	for i, number := range numbers {
		link := &lib.LinkedList[int]{
			Item: number,
		}

		if current != nil {
			current.Append(link)
		}

		current = link

		if number == 0 {
			zero = link
		}

		links[i] = link
	}
	current.Append(links[0])

	return links, zero
}

func mix(links []*lib.LinkedList[int], totalSize int) {
	for _, link := range links {
		linkNum := link.Item

		if linkNum == 0 {
			continue
		}

		// totalSteps is the total number of steps needed to get to the target
		totalSteps := lib.Abs(linkNum) % (totalSize - 1)
		// step gets one step closer to the target link
		var step func(*lib.LinkedList[int]) *lib.LinkedList[int]
		// action is the action to take once we reach the target
		var action func(*lib.LinkedList[int], *lib.LinkedList[int])

		// note: IntelliJ says these function references are wrong, but they're actually correct and this code
		// runs properly. Generics confuse the IDE - without generics it sees the types as what they are,
		// but with generics it sees the type as missing the struct argument, e.g. it thinks step is
		// func() *lib.LinkedList[int] and action is func(*lib.LinkedList[int])
		// annoyingly, the type mismatch inspection can't be disabled

		if linkNum < 0 {
			// if link is a negative number we step backwards the link number of steps and then prepend
			step = (*lib.LinkedList[int]).Prev
			action = (*lib.LinkedList[int]).Prepend
		} else {
			// if link is a positive number we step forwards the link number of steps and then append
			step = (*lib.LinkedList[int]).Next
			action = (*lib.LinkedList[int]).Append
		}

		target := link
		for i := 0; i < totalSteps; i++ {
			target = step(target)
		}

		link.Remove()
		action(target, link)
	}
}

func groveCoordinates[T any](zero *lib.LinkedList[T], targetCoords []int) []T {
	result := make([]T, len(targetCoords))

	i := 0

	for targetIndex, target := range targetCoords {
		for i != target {
			i += 1
			zero = zero.Next()
		}

		result[targetIndex] = zero.Item
	}

	return result
}

func solve(numbers []int) int {
	links, zero := buildLinkedList(numbers)

	mix(links, len(numbers))

	coords := groveCoordinates(zero, []int{1000, 2000, 3000})

	return lib.SumSlice(coords)
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseLine(lib.Atoi),
		SolveF: solve,
	}

	solver.Expect("1\n2\n-3\n3\n-2\n0\n4", 3)

	solver.Verify(27726)
}
