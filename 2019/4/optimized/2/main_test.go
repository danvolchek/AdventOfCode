package main_test

import (
	main "github.com/danvolchek/AdventOfCode/2019/4/optimized/2"
	"testing"
)

func BenchmarkNumPasswords(b *testing.B) {
	min, max := main.Parse()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		main.NumPasswords(min, max)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/danvolchek/AdventOfCode/2019/4/optimized/2
BenchmarkNumPasswords
BenchmarkNumPasswords-16              	  285454	      4081 ns/op
*/
