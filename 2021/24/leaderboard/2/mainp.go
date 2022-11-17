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

// this file was an early attempt at a static anaylzer which simplified the ALU operations
// it ended up not being helpful; manually recognizing the pattern led to the solution
// it's pretty cool, though

type op int

const (
	opInp op = iota
	opAdd
	opMul
	opDiv
	opMod
	opEql
	opRaw
	opNop
)

func inputp() *os.File {
	input, err := os.Open(path.Join("2021", "24", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

type action struct {
	op         op
	arg1, arg2 string
}

type alu2Action struct {
	op         op
	arg1, arg2 *alu2Action

	value int

	debugDepth int
}

type alu2 struct {
	w, x, y, z *alu2Action

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

func (a *alu2) load(arg string, debugDepth int) *alu2Action {
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
		return makeRaw(v, debugDepth)
	}
}

func newAlu2() *alu2 {
	return &alu2{
		w: makeRaw(0, 0),
		x: makeRaw(0, 0),
		y: makeRaw(0, 0),
		z: makeRaw(0, 0),
	}
}

func makeRaw(v, debugDepth int) *alu2Action {
	return &alu2Action{
		op:         opRaw,
		value:      v,
		debugDepth: debugDepth,
	}
}

func (a *alu2) build(acts []action) {
	for i, act := range acts {
		switch act.op {
		case opInp:
			*a.store(act.arg1) = &alu2Action{
				op:         act.op,
				value:      a.inpIndex,
				debugDepth: i + 1,
			}
			a.inpIndex += 1
		case opNop:
			continue
		default:
			*a.store(act.arg1) = &alu2Action{
				op:         act.op,
				arg1:       a.load(act.arg1, i+1),
				arg2:       a.load(act.arg2, i+1),
				debugDepth: i + 1,
			}
		}
	}
}

var simplificationMap = map[*alu2Action]*alu3Action{}

type alu3Action struct {
	raw, isNotImp bool

	val      int
	inpIndex int

	val1, val2 *alu3Action
	op         op
}

func (v *alu2Action) simplify3() *alu3Action {
	if vv, ok := simplificationMap[v]; ok {
		return vv
	}

	var v1, v2 *alu3Action

	switch v.op {
	case opAdd:
		v2 = v.arg2.simplify3()
		if v2.raw && v2.isNotImp && v2.val == 0 {
			return v.arg1.simplify3()
		}

		v1 = v.arg1.simplify3()
		if v1.raw && v1.isNotImp && v1.val == 0 {
			return v2
		}

		if v1.raw && v2.raw && v1.isNotImp && v2.isNotImp {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      v1.val + v2.val,
			}
		}
	case opMul:
		v2 = v.arg2.simplify3()
		if v2.raw && v2.isNotImp && v2.val == 0 {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      0,
			}
		}

		v1 = v.arg1.simplify3()
		if v1.raw && v1.isNotImp && v1.val == 0 {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      0,
			}
		}

		if v2.raw && v2.isNotImp && v2.val == 1 {
			return v1
		} else if v1.raw && v1.isNotImp && v1.val == 1 {
			return v2
		}

		if v1.raw && v2.raw && v1.isNotImp && v2.isNotImp {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      v1.val * v2.val,
			}
		}

	case opDiv:
		v2 = v.arg2.simplify3()
		if v2.raw && v2.isNotImp && v2.val == 1 {
			return v.arg1.simplify3()
		}

		v1 = v.arg1.simplify3()

		if v1.raw && v2.raw && v1.isNotImp && v2.isNotImp {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      v1.val / v2.val,
			}
		}
	case opMod:
		v1 = v.arg1.simplify3()
		if v1.raw && v1.isNotImp && v1.val == 0 {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      0,
			}
		}

		v2 = v.arg2.simplify3()

		if v1.raw && v2.raw && v1.isNotImp && v2.isNotImp {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      v1.val % v2.val,
			}
		}

		// a valid inp is 1-9
		if v1.raw && !v1.isNotImp && v2.raw && v2.isNotImp && v2.val > 9 {
			return v1
		}
	case opEql:
		v1, v2 = v.arg1.simplify3(), v.arg2.simplify3()

		// a valid inp is 1-9
		if v1.raw && !v1.isNotImp && v2.raw && v2.isNotImp && (v2.val > 9 || v2.val < 0) {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      0,
			}
		}

		// a valid inp is 1-9
		if v2.raw && !v2.isNotImp && v1.raw && v1.isNotImp && (v1.val > 9 || v1.val < 0) {
			return &alu3Action{
				raw:      true,
				isNotImp: true,
				val:      0,
			}
		}

		if v1.raw && v2.raw && v1.isNotImp && v2.isNotImp {
			if v1.val == v2.val {
				return &alu3Action{
					raw:      true,
					isNotImp: true,
					val:      1,
				}
			} else {
				return &alu3Action{
					raw:      true,
					isNotImp: true,
					val:      0,
				}
			}
		}

		v1Min, v1Max, v1ok := v1.possibleValues()
		if v1ok {
			v2Min, v2Max, v2ok := v2.possibleValues()
			if v2ok {
				if v1Max < v2Min || v2Max < v1Min {
					return &alu3Action{
						raw:      true,
						isNotImp: true,
						val:      0,
					}
				}
			}
		}

	case opInp:
		return &alu3Action{
			raw:      true,
			isNotImp: false,
			inpIndex: v.value,
		}
	case opRaw:
		return &alu3Action{
			raw:      true,
			isNotImp: true,
			val:      v.value,
		}
	}

	ret := &alu3Action{
		val1: v1,
		val2: v2,
		op:   v.op,
	}

	simplificationMap[v] = ret

	return ret

}

func (v *alu3Action) possibleValues() (int, int, bool) {
	if v.raw {
		if v.isNotImp {
			return v.val, v.val, true
		}

		return 1, 9, true
	}

	switch v.op {
	case opInp:
		panic("idk")
	case opAdd:
		v1Min, v1Max, v1ok := v.val1.possibleValues()
		v2Min, v2Max, v2ok := v.val2.possibleValues()
		return v1Min + v2Min, v1Max + v2Max, v1ok && v2ok
	case opMul:
		v1Min, v1Max, v1ok := v.val1.possibleValues()
		v2Min, v2Max, v2ok := v.val2.possibleValues()
		return v1Min * v2Min, v1Max * v2Max, v1ok && v2ok
	case opDiv:
		v1Min, v1Max, v1ok := v.val1.possibleValues()
		v2Min, v2Max, v2ok := v.val2.possibleValues()
		return v1Min / v2Max, v1Max / v2Min, v1ok && v2ok
	case opMod:
		v1Min, v1Max, v1ok := v.val1.possibleValues()
		v2Min, v2Max, v2ok := v.val2.possibleValues()
		// mod a b where a < b = a
		if v1Max < v2Min {
			return v1Min, v1Max, v1ok && v2ok
		}

		// mod a b where a >= b is between 0, b - 1
		return 0, v2Max - 1, v1ok && v2ok
	case opEql:
		return 0, 1, true
	case opRaw:
		return v.val, v.val, true
	default:
		panic(v.op)
	}
}

func (v *alu3Action) ToString(s *strings.Builder) {
	if v.raw {
		if v.isNotImp {
			s.WriteString(strconv.Itoa(v.val))
		} else {
			s.WriteString("inp_")
			s.WriteString(strconv.Itoa(v.inpIndex))
		}
		return
	}

	s.WriteByte('(')
	v.val1.ToString(s)
	s.Write([]byte{' ', fromOp(v.op), ' '})
	v.val2.ToString(s)
	s.WriteByte(')')
}

func (v *alu3Action) String() string {
	var s strings.Builder
	v.ToString(&s)
	return s.String()
}

func (v *alu2Action) simplify() string {
	/*if vvv, ok := simplificationMap[v]; ok {
		return vvv
	}*/

	var ret string
	oneArg := false

	switch v.op {
	case opAdd:
		v2 := v.arg2.simplify()
		if v2 == "0" {
			ret = v.arg1.simplify()
			oneArg = true
		} else {
			v1 := v.arg1.simplify()
			if v1 == "0" {
				ret = v2
				oneArg = true
			} else {
				ret = v1 + " + " + v2
			}
		}
	case opMul:
		v2 := v.arg2.simplify()
		if v2 == "0" {
			ret = "0"
			oneArg = true
		} else if v2 == "1" {
			ret = v.arg1.simplify()
			oneArg = true
		} else {
			v1 := v.arg1.simplify()

			if v1 == "0" {
				ret = "0"
				oneArg = true
			} else {
				ret = v1 + " * " + v2
			}
		}

	case opDiv:
		v2 := v.arg2.simplify()
		if v2 == "1" {
			ret = v.arg1.simplify()
			oneArg = true
		} else {
			ret = v.arg1.simplify() + " / " + v2
		}
	case opMod:
		ret = v.arg1.simplify() + " % " + v.arg2.simplify()
	case opEql:
		ret = v.arg1.simplify() + " == " + v.arg2.simplify() + " ? 1 : 0"
	case opInp:
		oneArg = true
		ret = "inp[" + strconv.Itoa(v.value) + "]"
	case opRaw:
		oneArg = true
		ret = strconv.Itoa(v.value)
	default:
		panic(v.op)
	}

	if !oneArg {
		ret = "(" + ret + ")"
	}

	//simplificationMap[v] = ret
	return ret
}

/*
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
}*/

func solvep(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var lines []string
	var instrs []action
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			instrs = append(instrs, action{
				op: opNop,
			})
			continue
		}
		lines = append(lines, line)

		parts := strings.Split(line, " ")

		act := action{
			op:   toOp(parts[0]),
			arg1: parts[1],
		}

		if len(parts) == 3 {
			act.arg2 = parts[2]
		}

		instrs = append(instrs, act)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	b := bufio.NewReader(os.Stdin)
	start := 0
	inpNum := 1
	var oldZ *alu2Action
	for start != len(instrs) {
		end := start + 1
		for end < len(instrs) && instrs[end].op != opInp {
			end++
		}
		a := newAlu2()
		a.inpIndex = inpNum
		if oldZ != nil {
			a.z = &alu2Action{
				op:    opInp,
				value: 99 - inpNum,
			}
		}
		a.build(instrs[start:end])
		oldZ = a.z

		fmt.Printf("%d (%d-%d)\n", inpNum, start+1, end)
		fmt.Println("w:", a.w.simplify3())
		fmt.Println("x:", a.x.simplify3())
		fmt.Println("y:", a.y.simplify3())
		fmt.Println("z:", a.z.simplify3())

		b.ReadString('\n')
		start = end
		inpNum += 1
	}
	/*
		for i := 1; i <len(instrs); i++ {
			a := newAlu2()
			a.build(instrs[:i])
			fmt.Printf("%d: %s\n",i, lines[i-1])
			fmt.Println("w:", a.w.simplify3())
			fmt.Println("x:", a.x.simplify3())
			fmt.Println("y:", a.y.simplify3())
			fmt.Println("z:", a.z.simplify3())

			b.ReadString('\n')
		}*/

	//fmt.Println(a.y.simplify3())
	//fmt.Println(a.z.simplify3())

	/*a := newAlu2()
	a.build(instrs)
	fmt.Println("Simplifying....")
	ret := a.z.simplify3()

	fmt.Println("Stringing....")
	var s strings.Builder
	ret.ToString(&s)


	f, err := os.Create("24.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(s.String())
	if err != nil {
		panic(err)
	}*/

}

func toOp(v string) op {
	switch v {
	case "add":
		return opAdd
	case "inp":
		return opInp
	case "mul":
		return opMul
	case "div":
		return opDiv
	case "eql":
		return opEql
	case "mod":
		return opMod
	default:
		panic(v)
	}
}

func fromOp(o op) byte {
	switch o {
	case opAdd:
		return '+'
	case opMul:
		return '*'
	case opDiv:
		return '/'
	case opEql:
		return '='
	case opMod:
		return '%'
	default:
		panic(o)
	}
}

func mainp() {
	//solvep(strings.NewReader("inp x\nadd x 5\nmul x 10\ndiv x 2\nmod x 7\neql x 4"))
	//solvep(strings.NewReader("inp x\nmul x -1"))
	//solvep(strings.NewReader("inp z\ninp x\nmul z 3\neql z x"))
	//solvep(strings.NewReader("inp w\nadd z w\nmod z 2\ndiv w 2\nadd y w\nmod y 2\ndiv w 2\nadd x w\nmod x 2\ndiv w 2\nmod w 2"))
	solvep(inputp())
}

// 9
