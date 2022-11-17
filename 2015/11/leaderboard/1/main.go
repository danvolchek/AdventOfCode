package main

import (
	"github.com/danvolchek/AdventOfCode/lib"
)

func hasIncreasing(password []byte) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i+1]-password[i] == 1 && password[i+2]-password[i+1] == 1 {
			return true
		}
	}

	return false
}

func noConfusing(password []byte) bool {
	for _, char := range password {
		switch char {
		case 'i', 'o', 'l':
			return false
		}
	}

	return true
}

func twoPairs(password []byte, hasFirst bool) bool {
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			if hasFirst {
				return true
			}

			if twoPairs(password[i+2:], true) {
				return true
			}
		}
	}

	return false
}

func valid(password []byte) bool {
	return noConfusing(password) && hasIncreasing(password) && twoPairs(password, false)
}

func increment(password []byte) []byte {
	for index := len(password) - 1; index > -1; index-- {
		switch password[index] {
		case 'z':
			password[index] = 'a'
		default:
			password[index] += 1
		}
		if password[index] != 'a' {
			break
		}
	}

	// note: just for the step solver; improve this?
	return password
}

func solve(password []byte) string {
	increment(password)

	for !valid(password) {
		increment(password)
	}

	return string(password)
}

func main() {
	incrementTest := lib.Solver[[]byte, string]{
		ParseF: lib.ParseBytes,
		SolveF: lib.ToString(increment),
	}

	incrementTest.Expect("a", "b")
	incrementTest.Expect("xx", "xy")
	incrementTest.Expect("xy", "xz")
	incrementTest.Expect("xz", "ya")
	incrementTest.Expect("ya", "yb")

	validTest := lib.Solver[[]byte, bool]{
		ParseF: lib.ParseBytes,
		SolveF: valid,
	}

	validTest.Expect("hijklmmn", false)
	validTest.Expect("abbceffg", false)
	validTest.Expect("abbcegjk", false)

	solver := lib.Solver[[]byte, string]{
		ParseF: lib.ParseBytes,
		SolveF: solve,
	}

	solver.Expect("abcdefgh", "abcdffaa")
	solver.Expect("ghijklmn", "ghjaabcc")
	solver.Verify("cqjxxyzz")
}
