package main_test

import (
	main "github.com/danvolchek/AdventOfCode/2019/4/final/1"
	"testing"
)

func BenchmarkNumPasswords(b *testing.B) {
	min, max := main.Parse()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		main.NumPasswords(min, max)
	}
}

func BenchmarkNumPasswordsGoroutines(b *testing.B) {
	min, max := main.Parse()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		main.NumPasswordsGoroutines(min, max, 2)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/danvolchek/AdventOfCode/2019/4/final/1
BenchmarkNumPasswords
BenchmarkNumPasswords-16              	  428180	      2709 ns/op
BenchmarkNumPasswordsGoroutines
BenchmarkNumPasswordsGoroutines-16    	  217981	      5599 ns/op
 */
