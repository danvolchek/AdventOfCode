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

type equation struct {
	a, n *big.Int
}

// https://en.wikipedia.org/wiki/Chinese_remainder_theorem#Existence_(constructive_proof)
func findX2(eq1, eq2 equation) *big.Int {
	//euclid := extendedEuclid(eq1.n, eq2.n)

	m1 := big.NewInt(0)
	m2 := big.NewInt(0)
	big.NewInt(0).GCD(m1, m2, eq1.n, eq2.n)

	m2.Mul(m2, eq1.a)
	m2.Mul(m2, eq2.n)

	m1.Mul(m1, eq2.a)
	m1.Mul(m1, eq1.n)

	return m2.Add(m2, m1)

	//m1, m2 := euclid.x, euclid.y

	//return eq1.a * m2 * eq2.n + eq2.a*m1 * eq1.n
}

func findX(eqs []equation) (*big.Int, *big.Int) {
	currentN := big.NewInt(0).Mul(eqs[0].n, eqs[1].n)

	sol := findX2(eqs[0], eqs[1])

	for i := 2; i < len(eqs); i++ {
		sol = findX2(equation{
			a: sol,
			n: currentN,
		}, eqs[i])

		currentN.Mul(currentN, eqs[i].n)
	}

	return sol, currentN
}

func toSmallestPositive(sol, currentN *big.Int) *big.Int {
	if sol.Cmp(big.NewInt(0)) == -1 {
		toSub := big.NewInt(0).Div(sol, currentN)
		toSub.Mul(toSub, big.NewInt(-1))

		//toSub.Add(toSub, big.NewInt(1))

		fmt.Println(sol, "+", toSub, "*", currentN, " = ")

		sol.Add(sol, toSub.Mul(toSub, currentN))
	} else {
		toSub := big.NewInt(0).Div(sol, currentN)

		//toSub.Add(toSub, big.NewInt(1))

		fmt.Println(sol, "-", toSub, "*", currentN, " = ")

		sol.Sub(sol, toSub.Mul(toSub, currentN))
	}

	//if currentN - sol > 0 && currentN - sol < sol {
	//	return currentN - sol
	//}

	return sol
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	busses := []int{}

	scanner.Scan()

	scanner.Scan()
	second := scanner.Text()
	for _, item := range strings.Split(second, ",") {
		if item == "x" {
			busses = append(busses, 0)
			continue
		}

		val, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		busses = append(busses, val)
	}

	//num := busses[0]
	//for _,  val := range busses {
	//	if val == 0 {
	//		continue
	//	}
	//	num *= val
	//}

	/*
		x=(a1 mod m1); x=(a2 mod m2);

		// most obvious
		0 = ? mod 7
		0 = ? + 1 mod 13
		0 = ? + 4 mod 59
		0 = ? + 6 mod 31
		0 = ? + 7 mod 19

		// move ? to the left to be solved for
		-? = 0 mod 7
		-? = 1 mod 13
		-? = 4 mod 59
		-? = 6 mod 31
		-? = 7 mod 19

		// make positive
		? = 0 mod 7
		? = -1 mod 13
		? = -4 mod 59
		? = -6 mod 31
		? = -7 mod 19

	*/

	var entries []equation
	for i, bus := range busses {
		if bus == 0 {
			continue
		}
		entries = append(entries, equation{
			a: big.NewInt(int64(-i)),
			n: big.NewInt(int64(bus)),
		})
	}

	negSol, n := findX(entries)

	//fmt.Println(negSol)

	actualSol := toSmallestPositive(negSol, n)

	fmt.Printf("%d %+v\n", actualSol, busses)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func main() {
	//solve(strings.NewReader("0\n3,x,x,4,5"))
	solve(strings.NewReader("0\n7,13,x,x,59,x,31,19"))
	solve(strings.NewReader("0\n17,x,13,19"))
	solve(strings.NewReader("0\n67,7,59,61"))
	solve(strings.NewReader("0\n67,x,7,59,61"))
	solve(strings.NewReader("0\n67,7,x,59,61"))
	solve(strings.NewReader("0\n1789,37,47,1889"))
	solve(input())
}
