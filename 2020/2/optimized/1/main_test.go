package main

import (
	"testing"
)

func BenchmarkSolution(b *testing.B) {
	input := parseFile()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		solve(input)
	}
}

/*
BenchmarkSolution
BenchmarkSolution-16    	  115280	     10037 ns/op
*/
