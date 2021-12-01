package main

import (
	"encoding/csv"
	"fmt"
	"github.com/danvolchek/AdventOfCode/2019/5/vm"
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

func runPrint(tape, input []int) {
	v := vm.VM{
		Tape:  tape,
		In:    input,
		Debug: false,
	}

	v.Run()

	fmt.Println(v)
}

func one(tape []int) {
	runPrint(tape, []int{1})
}

func two(tape []int) {
	runPrint(tape, []int{5})
}

func main() {
	//runPrint(parse(strings.NewReader("1,9,10,3,2,3,11,0,99,30,40,50")), nil)

	//runPrint(parse(strings.NewReader("1,0,0,0,99")), nil)
	//runPrint(parse(strings.NewReader("2,3,0,3,99")), nil)
	//runPrint(parse(strings.NewReader("2,4,4,5,99,0")), nil)
	//runPrint(parse(strings.NewReader("1,1,1,4,99,5,6,0,99")), nil)

	//runPrint(parse(strings.NewReader("1,9,10,3,2,3,11,0,99,30,40,50")), nil)

	//runPrint(parse(strings.NewReader("3,0,4,0,99")), []int{5})
	//runPrint(parse(strings.NewReader("1101,100,-1,4,0,99")), []int{5})

	/*runPrint(parse(strings.NewReader("3,9,8,9,10,9,4,9,99,-1,8")), []int{5})
	runPrint(parse(strings.NewReader("3,9,8,9,10,9,4,9,99,-1,8")), []int{8})

	runPrint(parse(strings.NewReader("3,9,7,9,10,9,4,9,99,-1,8")), []int{5})
	runPrint(parse(strings.NewReader("3,9,7,9,10,9,4,9,99,-1,8")), []int{8})

	runPrint(parse(strings.NewReader("3,3,1108,-1,8,3,4,3,99")), []int{5})
	runPrint(parse(strings.NewReader("3,3,1108,-1,8,3,4,3,99")), []int{8})

	runPrint(parse(strings.NewReader("3,3,1107,-1,8,3,4,3,99")), []int{5})
	runPrint(parse(strings.NewReader("3,3,1107,-1,8,3,4,3,99")), []int{8})*/

	/*runPrint(parse(strings.NewReader("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")), []int{5})
	runPrint(parse(strings.NewReader("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")), []int{0})

	runPrint(parse(strings.NewReader("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")), []int{5})
	runPrint(parse(strings.NewReader("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")), []int{0})*/

	/*const larger = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"

	runPrint(parse(strings.NewReader(larger)), []int{5})
	runPrint(parse(strings.NewReader(larger)), []int{8})
	runPrint(parse(strings.NewReader(larger)), []int{11})*/

	input, err := os.Open(path.Join("2019", "5", "input.txt"))
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
