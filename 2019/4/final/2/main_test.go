package main_test

import (
	main "github.com/danvolchek/AdventOfCode/2019/4/final/2"
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
