package main

import (
	"fmt"
	"github.com/danvolchek/AdventOfCode/lib"
)

type linkedList[T any] struct {
	item T

	prev *linkedList[T]
	next *linkedList[T]
}

func solve(lines []int) int {
	start := &linkedList[int]{
		item: lines[0],
	}

	moveOrder := make([]*linkedList[int], len(lines))

	var zero *linkedList[int]

	moveOrder[0] = start
	curr := start
	for i := 1; i < len(lines); i++ {
		next := &linkedList[int]{
			item: lines[i],
			prev: curr,
		}
		curr.next = next
		curr = next

		if lines[i] == 0 {
			zero = next
		}

		moveOrder[i] = curr
	}
	curr.next = start
	start.prev = curr

	for _, curr = range moveOrder {
		targNum := curr.item

		if targNum == 0 {
			continue
		}

		remove(curr)

		target := curr
		if targNum < 0 {
			for j := 0; j < ((lib.Abs(targNum) + 1) % (len(lines) - 1)); j++ {
				target = target.prev
			}
		} else {
			for j := 0; j < (targNum % (len(lines) - 1)); j++ {
				target = target.next
			}
		}

		if target == curr {
			panic("omg")
		}

		insertAfter(curr, target)

	}

	return get(zero, 1000) + get(zero, 2000) + get(zero, 3000)
}

func get[T any](start *linkedList[T], index int) T {
	for i := 0; i < index; i++ {
		start = start.next
	}

	fmt.Println(start.item)
	return start.item
}

func remove[T any](item *linkedList[T]) {
	item.prev.next = item.next
	item.next.prev = item.prev

	//item.prev = nil
	//item.next = nil
}

func insertAfter[T any](item, target *linkedList[T]) {
	targNext := target.next

	targNext.prev = item
	target.next = item

	item.next = targNext
	item.prev = target
}

func main() {
	solver := lib.Solver[[]int, int]{
		ParseF: lib.ParseLine(lib.Atoi),
		SolveF: solve,
	}

	solver.Expect("1\n2\n-3\n3\n-2\n0\n4", 3)
	solver.Expect("1\n2\n-3\n10\n-2\n0\n4", 3)

	solver.Verify(27726)
}
