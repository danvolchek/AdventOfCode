package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parse(reader io.Reader) []int {

	csvReader := csv.NewReader(reader)

	items, err := csvReader.Read()
	if err != nil {
		panic(err)
	}

	ret := make([]int, len(items))
	for i, item := range items {

		val, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		ret[i] = val
	}

	return ret
}

func vm(tape []int) []int {
	for i := 0; i < len(tape); i += 4 {
		switch tape[i] {
		case 1:
			tape[tape[i+3]] = tape[tape[i+1]] + tape[tape[i+2]]
		case 2:
			tape[tape[i+3]] = tape[tape[i+1]] * tape[tape[i+2]]
		case 99:
			return tape
		}
	}

	panic("no ret")
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	tape := parse(input)
loop:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			newTape := make([]int, len(tape))
			for k, v := range tape {
				newTape[k] = v
			}

			newTape[1] = i
			newTape[2] = j

			res := vm(newTape)[0]
			if res == 19690720 {
				fmt.Println(100*i + j)
				break loop
			}
		}
	}
}
