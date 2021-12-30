package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func parseBlocks(r io.Reader) []block {
	scanner := bufio.NewScanner(r)

	move := func(num int) bool {
		last := true

		for i := 0; i < num; i++ {
			last = scanner.Scan()
		}

		return last
	}

	arg2 := func() int {
		stringVal := strings.Split(scanner.Text(), " ")[2]

		intVal, err := strconv.ParseInt(stringVal, 10, 64)
		if err != nil {
			panic(err)
		}

		return int(intVal)
	}

	moveToNextInput := func() bool {
		for !strings.Contains(scanner.Text(), "inp") {
			if !move(1) {
				return false
			}
		}

		return true
	}

	moveToNextInput()

	var blocks []block

	for {
		move(4)
		b := arg2() // div z d

		move(1)
		a := arg2() // add x a

		move(10)
		c := arg2() // add y c

		blocks = append(blocks, block{
			a: a,
			b: b,
			c: c,
		})

		if !moveToNextInput() {
			break
		}
	}

	return blocks
}

// block represents a sequence of alu instructions that ultimately result in the modification of z according to:
// z = (p_z % 26) + a != in_i ? 26 * (p_z / b) + in_i + c : (p_z / b)
// where p_z is the previous z value, in_i is the current input value, and a, b, c are constants determined by the instructions
type block struct {
	a, b, c int
}

// solve evaluates the z equation the block represents for input values from 1 to 9, returning a zRange that represents all solutions
func (b block) solve(prevZ int) zRange {
	solution := zRange{
		startZ:    26*(prevZ/b.b) + 1 + b.c,
		equalityZ: prevZ / b.b,
	}

	for inputVal := 1; inputVal <= 9; inputVal++ {
		condition := (prevZ % 26) + b.a
		if condition == inputVal {
			solution.equalityInput = byte(inputVal)
			break
		}
	}

	return solution
}

// zRange is a space-efficient representation of a collection of zValues
// it works by noticing that there are two cases in the z equation: equality and inequality
// the equality case can happen only for one input value, and results in one z value
// the inequality case happens for the rest of the input values, and results in linearly increasing z values
type zRange struct {
	// startZ is the z value associated with input value 1
	startZ int

	// equalityZ is the z value associated with the input that triggered the equality case
	equalityZ int

	// equalityInput is the input value that triggered the equality case, or 0 if no input did
	equalityInput byte
}

// zValue represents an input value and it's corresponding z value
type zValue struct {
	input int
	value int
}

func (z zRange) zValues() [9]zValue {
	var zValues [9]zValue

	for i := 1; i <= 9; i++ {
		zVal := z.startZ + i - 1

		if byte(i) == z.equalityInput {
			zVal = z.equalityZ
		}

		zValues[i-1] = zValue{
			input: i,
			value: zVal,
		}
	}

	return zValues
}

func (z zRange) hasEqualityCase() bool {
	return z.equalityInput != 0
}

type progressIndicator struct {
	duration time.Duration
	value    int
	stop     chan bool
}

func (p progressIndicator) Start() *int {
	go func() {
		ticker := time.NewTicker(p.duration)
		for {
			select {
			case <-ticker.C:
				fmt.Println("Found", p.value, "nums")
			case <-p.stop:
				fmt.Println("Found", p.value, "nums")
				ticker.Stop()
				return
			}
		}
	}()

	return &p.value
}

func (p progressIndicator) Stop() {
	p.stop <- true
}

// The solver works by finding every possible z value for all input values at each block,
// and then reconstructing the input values that led to a z value of 0 after the last block
func findValidInputs(blocks []block) [][]int {
	return reconstructValidInputs(buildZValues(blocks))
}

// buildZValues returns a slice of every possible z value for every possible input value at every block.
// importantly, it records the previous z value and input value used to reach the current z value.
// this means that the series of input numbers which resulted in a certain z value after any number of blocks can be reconstructed.
func buildZValues(blocks []block) []map[int]zRange {
	zValuesForAllBlocks := []map[int]zRange{
		{0: blocks[0].solve(0)},
	}

	for blockNum := 1; blockNum < len(blocks); blockNum++ {
		fmt.Println("Building z values for block", blockNum, "...")
		currentBlock := blocks[blockNum]

		zValuesForBlock := make(map[int]zRange)

		zValuesForPreviousBlock := zValuesForAllBlocks[len(zValuesForAllBlocks)-1]

		for _, prevZValues := range zValuesForPreviousBlock {
			for _, prevZValue := range prevZValues.zValues() {
				currentZRange := currentBlock.solve(prevZValue.value)

				// the equality case is the only case which can reduce the z value
				// assuming the z value is not 0 until the last block, the last block must have it to take the z value
				// to 0; so we can throw away ranges which don't have an equality match
				if blockNum == len(blocks)-1 {
					if !currentZRange.hasEqualityCase() {
						continue
					}
				}

				zValuesForBlock[prevZValue.value] = currentZRange
			}
		}

		zValuesForAllBlocks = append(zValuesForAllBlocks, zValuesForBlock)
	}

	return zValuesForAllBlocks
}

func reconstructValidInputs(zValuesForAllBlocks []map[int]zRange) [][]int {
	fmt.Println("Reconstructing inputs from z values ...")

	p := progressIndicator{
		stop:     make(chan bool),
		duration: time.Second * 5,
	}

	r := reconstructor{
		numFound:            p.Start(),
		zValuesForAllBlocks: zValuesForAllBlocks,
	}
	defer p.Stop()

	// the valid inputs according to the puzzle are the ones which leave z 0 after all blocks execute
	return r.reconstruct(len(zValuesForAllBlocks)-1, 0)
}

type reconstructor struct {
	numFound            *int
	zValuesForAllBlocks []map[int]zRange
}

// reconstruct returns all possible series of inputs that result in the target z value after
// the block blockNum was executed. It does so by tracing z values from the desired block to block 1,
// recording the input values used to reach the target z values it finds along the way.
func (r reconstructor) reconstruct(blockNum, targetZValue int) [][]int {
	if blockNum == -1 {
		*r.numFound += 1
		return [][]int{{}}
	}

	var inputs [][]int

	zValuesForCurrentBlock := r.zValuesForAllBlocks[blockNum]

	for prevZ, result := range zValuesForCurrentBlock {
		for _, nextZ := range result.zValues() {
			if nextZ.value == targetZValue {
				for _, inputValue := range r.reconstruct(blockNum-1, prevZ) {
					inputs = append(inputs, append(inputValue, nextZ.input))
				}
			}
		}
	}

	return inputs
}

func combineIntSlice(ints []int) int {
	value := 0
	for index := 0; index < len(ints); index++ {
		value *= 10
		value += ints[index]
	}
	return value
}
