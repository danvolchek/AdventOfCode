package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

// alu is an implementation to check the z value of resulting scores, see part 2's main.go for the solver for both parts

func input() *os.File {
	input, err := os.Open(path.Join("2021", "24", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type action struct {
	op string
	arg1, arg2 string
}

type alu2Action struct {
	op string
	arg1, arg2 *alu2Action

	inpIndex int
}

type alu2 struct {
	w,x,y,z *alu2Action

	inpIndex int
}

func (a *alu2) store(arg string) **alu2Action {
	switch arg {
	case "w":
		return &a.w
	case "x":
		return &a.x
	case "y":
		return &a.y
	case "z":
		return &a.z
	default:
		panic(arg)
	}
}

func (a * alu2) load(arg string) *alu2Action {
	switch arg {
	case "w":
		return a.w
	case "x":
		return a.x
	case "y":
		return a.y
	case "z":
		return a.z
	default:
		panic(arg)
	}
}



func (a *alu2) build(acts []action) {
	for _, act := range acts {
		switch act.op {
		case "inp":
			*a.store(act.arg1) = &alu2Action{
				op:       act.op,
				inpIndex: a.inpIndex,
			}
			a.inpIndex += 1
		default:
			*a.store(act.arg1) = &alu2Action{
				op:       act.op,
				arg1: a.load(act.arg1),
				arg2: a.load(act.arg2),
			}
		}
	}
}

func (v *alu2Action) simplify() string {
	if v == nil {
		return "0"
	}

	var ret string

	switch v.op {
	case "add":
		ret = v.arg1.simplify() + " + " + v.arg2.simplify()
	case "mul":
		ret = v.arg1.simplify() + " * " + v.arg2.simplify()
	case "div":
		ret = v.arg1.simplify() + " / " + v.arg2.simplify()
	case "mod":
		ret = v.arg1.simplify() + " % " + v.arg2.simplify()
	case "eql":
		ret = v.arg1.simplify() + " == " + v.arg2.simplify() + " ? 1 : 0"
	case "inp":
		ret = "inp[" + strconv.Itoa(v.inpIndex) + "]"
	default:
		panic(v.op)
	}

	return "(" + ret + ")"
}

type alu struct {
	w,x,y,z int

	inp []int
}

func (a *alu) load(arg string) int {
	switch arg {
	case "w":
		return a.w
	case "x":
		return a.x
	case "y":
		return a.y
	case "z":
		return a.z
	default:
		v, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		return v
	}
}

func (a *alu) store(arg string) *int {
	switch arg {
	case "w":
		return &a.w
	case "x":
		return &a.x
	case "y":
		return &a.y
	case "z":
		return &a.z
	default:
		panic(arg)
	}
}

func (a *alu) String() string {
	return fmt.Sprintf("w: %d, x: %d, y: %d, z:%d", a.w, a.x, a.y, a.z)
}

func (a *alu) run(instructions []action) {
	for _, act := range instructions {
		switch act.op {
		case "inp":
			fmt.Println(a.z)
			*a.store(act.arg1) = a.inp[0]
			a.inp = a.inp[1:]
		case "add":
			*a.store(act.arg1) = a.load(act.arg1) + a.load(act.arg2)
		case "mul":
			*a.store(act.arg1) = a.load(act.arg1) * a.load(act.arg2)
		case "div":
			*a.store(act.arg1) = a.load(act.arg1) / a.load(act.arg2)
		case "mod":
			*a.store(act.arg1) = a.load(act.arg1) % a.load(act.arg2)
		case "eql":
			val := a.load(act.arg1) == a.load(act.arg2)
			if val {
				*a.store(act.arg1) = 1
			} else {
				*a.store(act.arg1) = 0
			}
		default:
			panic(act)
		}
	}
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var instrs []action
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")

		act := action{
			op:   parts[0],
			arg1: parts[1],
		}

		if len(parts) ==3 {
			act.arg2 = parts[2]
		}

		instrs = append(instrs, act)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	num := 98195993585832 //98195766853942
	sploot := splot(num)

	v := &alu{
		inp: sploot,
	}

	v.run(instrs)

	fmt.Println(v)
}

func contains(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}

	return false
}

func splot(num int) []int {
	ret := make([]int, 14)
	i := 13
	for num != 0 {
		ret[i] = num % 10
		num /= 10
		i--
	}

	return ret
}

func main() {
	//solve(strings.NewReader("inp x\nmul x -1"))
	//solve(strings.NewReader("inp z\ninp x\nmul z 3\neql z x"))
	//solve(strings.NewReader("inp w\nadd z w\nmod z 2\ndiv w 2\nadd y w\nmod y 2\ndiv w 2\nadd x w\nmod x 2\ndiv w 2\nmod w 2"))
	solve(input())
}
