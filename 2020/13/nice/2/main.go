package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"path"
	"strconv"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "13", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

// represents x = a mod n
type equation struct {
	a, n *big.Int
}

// solves the system of equations x = a1 mod n1 and x = a2 mod n2 for x
// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Case_of_two_moduli
func solveDualEquations(eq1, eq2 equation) *big.Int {
	// compute bezout coefficients
	m1 := big.NewInt(0)
	m2 := big.NewInt(0)

	big.NewInt(0).GCD(m1, m2, eq1.n, eq2.n)

	// calculate x
	m2.Mul(m2, eq1.a)
	m2.Mul(m2, eq2.n)

	m1.Mul(m1, eq2.a)
	m1.Mul(m1, eq1.n)

	return m2.Add(m2, m1)
}

// solves the system of equations x = ai mod ni for all eqs
// returns the smallest positive x which satisfies all equations
// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#General_case
func solveEquations(eqs []equation) *big.Int {
	// solve first two
	N := big.NewInt(0).Mul(eqs[0].n, eqs[1].n)
	x := solveDualEquations(eqs[0], eqs[1])

	// solve the rest without modifying the moduli of the existing equations
	for i := 2; i < len(eqs); i++ {
		x = solveDualEquations(equation{
			a: x,
			n: N,
		}, eqs[i])

		N.Mul(N, eqs[i].n)
	}

	// convert x to the smallest positive value, maintaining all moduli
	// equivalent to x -= (x / N) * N (the integer division makes this work)
	return x.Sub(x, big.NewInt(0).Mul(big.NewInt(0).Div(x, N), N))
}

func parse(r io.Reader) map[int64]int64 {
	scanner := bufio.NewScanner(r)

	chomp := func() string {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				panic(scanner.Err())
			}
		}
		return scanner.Text()
	}

	// discard earliest time
	chomp()

	busses := make(map[int64]int64)
	rawBusses := chomp()
	for index, item := range strings.Split(rawBusses, ",") {
		if item == "x" {
			continue
		}

		val, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		busses[int64(index)] = int64(val)
	}

	return busses
}

func solve(busses map[int64]int64) *big.Int {
	/*
			Each bus forms an equation:
		    	0 = ? + busIndex (mod bus)

			Where ? is the solution to this day. We need to solve the system of equations that is the above for all busses.

			This chinese remainder theorem lets us do exactly this. We can use it to efficiently solve systems of equations of the form:
				x = a (mod n)

			for x. Reordering the bus equations we have to make x = ?, we get
				? = -busIndex (mod bus)

			so a is -busIndex and n is the bus.
	*/

	var equations []equation
	for index, bus := range busses {
		equations = append(equations, equation{
			a: big.NewInt(-index),
			n: big.NewInt(bus),
		})
	}

	return solveEquations(equations)
}

func main() {
	fmt.Println(solve(parse(strings.NewReader("0\n7,13,x,x,59,x,31,19"))))
	fmt.Println(solve(parse(strings.NewReader("0\n17,x,13,19"))))
	fmt.Println(solve(parse(strings.NewReader("0\n67,7,59,61"))))
	fmt.Println(solve(parse(strings.NewReader("0\n67,x,7,59,61"))))
	fmt.Println(solve(parse(strings.NewReader("0\n67,7,x,59,61"))))
	fmt.Println(solve(parse(strings.NewReader("0\n1789,37,47,1889"))))
	fmt.Println(solve(parse(input())))
}
