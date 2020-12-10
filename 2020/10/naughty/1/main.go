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

	diff1 := 0
	diff3 := 0
	for i := 1; i < len(adapters); i++ {
		if adapters[i]-adapters[i-1] == 1 {
			diff1 += 1
		} else if adapters[i]-adapters[i-1] == 3 {
			diff3 += 1
		} else {
			panic(fmt.Sprintf("%d %d", adapters[i], adapters[i-1]))
		}
	}

	diff3 += 1
	fmt.Println(diff1, diff3, diff1*diff3)
}

func main() {
	solve(strings.NewReader("16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4"))
	solve(strings.NewReader("28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3"))
	solve(input())
}
