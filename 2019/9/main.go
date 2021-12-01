package main

import (
	"encoding/csv"
	"fmt"
	"github.com/danvolchek/AdventOfCode/2019/9/vm"
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

func runPrint(tape, input []int) {
	v := vm.VM{
		Tape: tape,
		In:   input,
		Debug: false,
	}

	v.Run()

	fmt.Println(v)
}


func main() {
	runPrint(parse(strings.NewReader("104,1125899906842624,99")), []int{})
	runPrint(parse(strings.NewReader("1102,34915192,34915192,7,4,7,99,0")), []int{})
	runPrint(parse(strings.NewReader("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")), []int{})


	input, err := os.Open(path.Join("2019", "9", "input.txt"))
	if err != nil {
		panic(err)
	}

	tape := parse(input)

	runPrint(cp(tape), []int{1})
	runPrint(cp(tape), []int{2})
}

func cp(t []int) []int {
	ret := make([]int, len(t))
	copy(ret, t)
	return ret
}
