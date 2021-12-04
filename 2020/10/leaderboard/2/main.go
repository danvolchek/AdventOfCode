package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "10", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func contains(a map[int]bool, b int) bool {
	_, ok := a[b]
	return ok
}

func combinations(adapters map[int]bool, ways []int) int {
	ways[0] = 1
	for i := 1; i < len(ways); i++ {
		for j := 1; j < 4; j++ {
			if contains(adapters, i) && contains(adapters, i-j) {
				//fmt.Printf("to get to %d, can use %d\n", i, i - j)
				ways[i] += ways[i-j]
			}
		}

		//fmt.Println("2")
	}

	return ways[len(ways)-1]

}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var adapters []int
	adapters = append(adapters, 0)
	for scanner.Scan() {
		row := scanner.Text()

		i, _ := strconv.Atoi(row)
		adapters = append(adapters, i)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	//currentJoltage := 0

	//adaptersInUse := make(map[int]bool)

	sort.Ints(adapters)

	max := adapters[len(adapters)-1]

	adaptamap := make(map[int]bool, len(adapters))
	for _, ad := range adapters {
		adaptamap[ad] = true
	}

	combos := combinations(adaptamap, make([]int, max+1))

	fmt.Println(combos)
}

func main() {
	solve(strings.NewReader("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"))
	solve(strings.NewReader("28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"))
	solve(input())
}
