package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

// this has the solver for both parts
// it works by recognizing the instructions are split into a block per input, each which modify z according to
// z = (p_z % 26) + a != in_i ? 26 * (p_z / b) + in11 + c : (p_z / b)
// where a, b, and c are numbers that vary in the instructions for the block
// the equation can be applied for inputs 1-14, with every possible input value (1-9) to figure out
// all the possible z values. Once we have those, we trace back input choices by z value that end with
// z = 0 after input 14 (the target z) - this yields every valid input. Then we pick the min/max


func input() *os.File {
	input, err := os.Open(path.Join("2021", "24", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

// z = (p_z % 26) + a != in_i ? 26 * (p_z / b) + in11 + c : (p_z / b)
type block struct {
	a, b, c uint64
}

func toInt(v string) uint64 {
	vv, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(err)
	}

	return uint64(vv)
}

func parse(r io.Reader) []block {
	scanner := bufio.NewScanner(r)

	move := func(num int) bool {
		last := true

		for i := 0; i < num; i++ {
			last = scanner.Scan()
		}

		return last
	}

	arg2 := func() uint64 {
		return toInt(strings.Split(scanner.Text(), " ")[2])
	}

	move(1)

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

		for !strings.Contains(scanner.Text(), "inp") {
			if !move(1) {
				return blocks
			}
		}
	}
}

// this is a space efficient representation of the result of the z equation above for all possible inputs
// see solveResult.zValues for how to reconstruct the z value for a given input value
// a map from input Value to Z value takes too much space and causes my PC to run out RAM
type solveResult struct {
	// The minimum z value, corresponding to input = 1. Because of the form of the equation, if there was no
	// equality in the z equation, input = 2 is Min + 1, input = 3 is Min + 2, and so on.
	Min     uint64 `json:"m,omitempty"`

	// If there was an equality in the z equation, this is the corresponding z value.
	MatchZ  uint64 `json:"z"`

	// If there was an equality in the z equation, it happened at only one input. This is that input.
	MatchIn byte	`json:"i,omitempty"`
}

func (s solveResult) hasEqualityMatch() bool {
	return s.MatchIn != 0
}

func (s solveResult) zValues() []uint64 {
	ret := make([]uint64, 9)

	var i byte
	for i = 1; i <= 9; i++ {
		zToMatch := uint64(i - 1) + s.Min

		if i == s.MatchIn {
			zToMatch = s.MatchZ
		}

		ret[i - 1] = zToMatch
	}

	return ret
}

// -> map[inVal][resulting z value]
func (b block) solve(prevZ uint64) solveResult {

	ret := solveResult{
		Min: 26*(prevZ/b.b) + 1 + b.c,
		//max:    26*(prevZ/b.b) + 9 + b.c,
		//MatchZ: prevZ / b.b,
	}

	var in uint64
	for in = 1; in <= 9; in++ {
		cond := (prevZ % 26) + b.a
		if cond == in {
			ret.MatchIn = byte(in)
			ret.MatchZ = prevZ/b.b
			break
		}
	}

	return ret
}

func solve(r io.Reader) {
	blocks := parse(r)

	// block -> map[prevZ] -> map[inpNum] -> resulting Z
	//var lastInValues map[uint64]solveResult

	var lastInValuesMem []map[uint64]solveResult

	var blockNum int

	//if read(0, &lastInValuesMem) {
	//	blockNum = len(blocks)
	//}
	/*for blockNum = len(blocks) - 1; blockNum >= 0; blockNum-- {
		if q, ok := read(blockNum); ok {
			lastInValues = q
			blockNum++
			break
		}
	}*/

	if lastInValuesMem == nil {
		lastInValuesMem = append(lastInValuesMem, map[uint64]solveResult{
			0: blocks[0].solve(0),
		})
		//write(0, lastInValues)
		blockNum = 1
	}

	fmt.Println("Starting from", blockNum)
	for ; blockNum < len(blocks); blockNum++ {
		fmt.Println(blockNum)
		b := blocks[blockNum]

		possibForBlock := make(map[uint64]solveResult)

		lastInValues := lastInValuesMem[len(lastInValuesMem) - 1]

		for _, lastResult := range lastInValues {
			for _, zToMatch := range lastResult.zValues() {
				sol := b.solve(zToMatch)
				if blockNum == len(blocks) - 1 {
					if !sol.hasEqualityMatch() {
						continue
					}
				}

				possibForBlock[zToMatch] = sol
			}
		}

		//lastInValues = possibForBlock

		//write(blockNum, lastInValues)

		lastInValuesMem = append(lastInValuesMem, possibForBlock)
	}

	//write(0, lastInValuesMem)

	fmt.Println("reconstructing....")


	fmt.Println("chasing 0")
	stop := make(chan int)
	go func() {
		vv := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-vv.C:
				fmt.Println("Found", found, "nums")
			case <-stop:
				fmt.Println("Found", found, "nums")
				vv.Stop()
				break
			}
		}
	}()

	var numsInt []int
	for _, num := range chase(lastInValuesMem, 0) {
		numsInt = append(numsInt, int(toInt(combiner(num).String())))
	}
	stop <- 0

	sort.Slice(numsInt, func(i, j int) bool {
		return numsInt[i] < numsInt[j]
	})

	fmt.Println(numsInt[0], numsInt[len(numsInt) - 1])
}

type combiner []int

func (s combiner) String() string {
	var ss strings.Builder
	for _, v := range s {
		ss.WriteString(strconv.Itoa(v))
	}

	return ss.String()
}

var found = 0

func chase(zss []map[uint64]solveResult, target uint64) [][]int {
	if len(zss) == 0 {
		found++
		return [][]int{{}}
	}


	var ret [][]int

	zs := zss[len(zss) - 1]

	for prevZ, result := range zs {
		for i, nextZ := range result.zValues() {
			inpNum := i + 1

			if nextZ == target {

				//fmt.Println("found", target, "with num", inpNum, "now chasing", prevZ)

				for _, v := range ap(chase(zss[:len(zss) - 1], prevZ), inpNum) {
					//return [][]int{v}

					ret = append(ret, v)
				}
			}
		}
	}


	return ret
}

func ap(rets [][]int, v int) [][]int {
	n := make([][]int, len(rets))
	for i, ret := range rets {
		n[i] = make([]int, len(ret) + 1)
		copy(n[i], ret)
		n[i][len(n[i]) - 1] = v
	}

	return n
}

func read(i int, q interface{}) bool {
	f, err := os.Open(fmt.Sprintf("24_%d.json", i))
	if err != nil {
		return  false
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return false
	}

	err = json.Unmarshal(b, q)
	if err != nil {
		return false
	}

	return true
}

func write(i int, q interface{}) {
	f, err := os.Create(fmt.Sprintf("24_%d.json", i))
	if err != nil {
		panic(err)
	}

	v, err := json.Marshal(q)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(v)
	if err != nil {
		panic(err)
	}
}

func main() {
	solve(input())
}
