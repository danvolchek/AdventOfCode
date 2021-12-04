package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "23", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func remove(cups []int, index int) []int {
	return append(cups[0:index], cups[index+1:]...)
}

func insert(cups []int, index int, value int) []int {
	//fmt.Println("beggining", cups[0:index])
	//fmt.Println("index", index)
	//fmt.Println("value", value)
	//fmt.Println("end", cups[index:])
	ret := make([]int, len(cups)+1)
	copy(ret, cups[0:index])
	ret[index] = value
	copy(ret[index+1:], cups[index:])
	return ret
}

func removeByValue(cups []int, val int) []int {
	index := getIndex(cups, val)
	if index == -1 {
		panic(val)
	}
	return remove(cups, index)
}

func getIndex(cups []int, value int) int {
	for index, cup := range cups {
		if cup == value {
			return index
		}
	}

	return -1
}

func round(cups []int, current int) ([]int, int) {
	fmt.Println("cups: ", cups)
	currentIndex := getIndex(cups, current)
	fmt.Println("current value: ", current, "current index: ", currentIndex)

	if currentIndex == -1 {
		panic("asd")
	}

	first, second, third := cups[(currentIndex+1)%len(cups)], cups[(currentIndex+2)%len(cups)], cups[(currentIndex+3)%len(cups)]

	fmt.Println("picked up", first, second, third)
	cups = removeByValue(removeByValue(removeByValue(cups, first), second), third)

	//fmt.Println("removed:", cups)

	destination := current - 1

	for getIndex(cups, destination) == -1 {
		destination -= 1

		if destination < 1 {
			destination = 9
		}
	}

	destinationIndex := (getIndex(cups, destination) + 1) % len(cups)

	fmt.Println("destination value: ", destination, "destination index: ", destinationIndex)

	cups = insert(insert(insert(cups, destinationIndex, third), destinationIndex, second), destinationIndex, first)

	return cups, cups[(getIndex(cups, current)+1)%len(cups)]
}

func solve(r io.Reader) {
	var cups []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := scanner.Text()

		for i := 0; i < len(row); i++ {
			val, err := strconv.Atoi(row[i : i+1])
			if err != nil {
				panic(err)
			}

			cups = append(cups, val)

		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	current := cups[0]

	//printy(cups, current)
	for i := 0; i < 100; i++ {
		cups, current = round(cups, current)
		fmt.Println()
		//printy(cups, current)
	}

	fmt.Println(cups)

	one := getIndex(cups, 1)

	curr := one + 1

	var ret strings.Builder
	for curr != one {
		ret.WriteString(fmt.Sprintf("%d", cups[curr]))
		curr = (curr + 1) % len(cups)
	}

	fmt.Println(ret.String())

}

func printy(cups []int, curr int) {
	fmt.Println("current: ", curr)
	fmt.Println(cups)
	fmt.Println()
}

func main() {
	//solve(strings.NewReader("389125467"))
	solve(input())
}
