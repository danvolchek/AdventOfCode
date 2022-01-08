package main

import (
	"encoding/csv"
	"fmt"
	"github.com/danvolchek/AdventOfCode/2019/13/optimized/2/vm"
	"io"
	"os"
	"path"
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

func one(tape []int) {
	v := vm.VM{
		Tape:  tape,
		In:    nil,
		debug: false,
	}

	v.Run()

	blocks := 0
	for i := 0; i < len(v.Out); i += 3 {
		if v.Out[i+2] == 2 {
			blocks += 1
		}
	}

	fmt.Println(blocks)
}

func two(tape []int) {
	tape[0] = 2

	v := vm.VM{
		Tape:  tape,
		In:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		debug: false,
	}

	v.Run()

	fmt.Println(v.Out)
}

func main() {

	input, err := os.Open(path.Join("2019", "13", "input.txt"))
	if err != nil {
		panic(err)
	}

	tape := parse(input)

	one(cp(tape))

	two(cp(tape))
}

func cp(t []int) []int {
	ret := make([]int, len(t))
	copy(ret, t)
	return ret
}
