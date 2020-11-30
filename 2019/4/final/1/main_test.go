package main_test

import (
	main "github.com/danvolchek/AdventOfCode/2019/4/final/1"
	"testing"
)

func BenchmarkNumPasswords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.NumPasswords(382345, 843167)
	}
}

func BenchmarkNumPasswordsGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.NumPasswordsGoroutines(382345, 843167, 16)
	}
}

/*
goos: windows
goarch: amd64
pkg: aoc/2019/4/final/1
BenchmarkNumPasswords
BenchmarkNumPasswords-16              	  196543	      5907 ns/op
BenchmarkNumPasswordsGoroutines
BenchmarkNumPasswordsGoroutines-16    	  109990	     10711 ns/op
 */