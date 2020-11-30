package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
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

	fmt.Println(vm(parse(strings.NewReader("1,0,0,0,99"))))
	fmt.Println(vm(parse(strings.NewReader("2,1,0,1,99"))))
	fmt.Println(vm(parse(strings.NewReader("2,2,2,5,99,0"))))
	fmt.Println(vm(parse(strings.NewReader("1,1,1,2,99,5,6,0,99"))))

	input, err := os.Open(path.Join("2019", "2", "input.txt"))
	if err != nil {
		panic(err)
	}

	tape := parse(input)
	tape[1] = 12
	tape[2] = 2
	fmt.Println(vm(tape))
}
