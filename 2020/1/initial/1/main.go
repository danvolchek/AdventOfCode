package main

import (
	"fmt"
	"os"
)

func parse() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(input)
}

func main() {
	parse()
}
