package vm

import (
	"fmt"
	"strings"
)

type VM struct {
	PC int

	RelativeBase int

	Tape []int
	In   []int
	Out  []int

	Debug bool
}

func (v *VM) Run() {
	for v.PC < len(v.Tape) {
		instruction := instruction(v.Tape[v.PC])

		code := instruction.Code()

		if v.Debug {
			fmt.Printf("PC: %v, Code: %v\n", v.PC, code)
			fmt.Printf("Tape: %v\n", v.Tape)
		}

		switch instruction.Code() {
		case 99:
			return
		case 1:
			v.three(add, instruction)
		case 2:
			v.three(mul, instruction)
		case 3:
			v.one(in, instruction)
		case 4:
			v.one(out, instruction)
		case 5:
			v.two(jit, instruction)
		case 6:
			v.two(jif, instruction)
		case 7:
			v.three(lt, instruction)
		case 8:
			v.three(eq, instruction)
		case 9:
			v.one(rel, instruction)
		default:
			panic(fmt.Sprintf("unknown code %v", instruction.Code()))
		}

	}

	panic("ran out of tape!")
}

func (v *VM) one(op func(v *VM, arg0 arg), instr instruction) {
	v.prepareArgs(v.PC, instr, 1)
	arg0 := v.parseArg(v.PC, instr, 0)

	if v.Debug {
		fmt.Printf("args: %v\n", arg0)
	}

	currPC := v.PC

	op(v, arg0)

	if v.PC == currPC {
		v.PC += 2
	}
}

func (v *VM) two(op func(v *VM, arg0, arg1 arg), instr instruction) {
	v.prepareArgs(v.PC, instr, 2)
	arg0, arg1 := v.parseArg(v.PC, instr, 0), v.parseArg(v.PC, instr, 1)

	if v.Debug {
		fmt.Printf("args: %v, %v\n", arg0, arg1)
	}

	currPC := v.PC

	op(v, arg0, arg1)

	if v.PC == currPC {
		v.PC += 3
	}
}

func (v *VM) three(op func(v *VM, arg0, arg1, arg2 arg), instr instruction) {
	v.prepareArgs(v.PC, instr, 3)
	arg0, arg1, arg2 := v.parseArg(v.PC, instr, 0), v.parseArg(v.PC, instr, 1), v.parseArg(v.PC, instr, 2)

	if v.Debug {
		fmt.Printf("args: %v, %v, %v\n", arg0, arg1, arg2)
	}

	currPC := v.PC

	op(v, arg0, arg1, arg2)

	if v.PC == currPC {
		v.PC += 4
	}
}

func (v *VM) parseArg(start int, instr instruction, i int) arg {
	raw := v.Tape[start+1+i]
	mode := instr.Mode(i)

	switch mode {
	case immediateMode:
		return immediateArg{back: raw}
	case positionMode:
		return positionArg{back: &v.Tape[raw], raw: raw}
	case relativeMode:
		return relativeArg{back: &v.Tape[raw+v.RelativeBase], raw: raw, base: v.RelativeBase}
	default:
		panic(fmt.Sprintf("unknown mode: %d", mode))
	}
}

func (v *VM) prepareArgs(start int, instr instruction, j int) {
	for i := 0; i < j; i++ {
		raw := v.Tape[start+1+i]
		mode := instr.Mode(i)

		switch mode {
		case positionMode:
			v.checkIncreaseTape(raw)
		case relativeMode:
			v.checkIncreaseTape(raw + v.RelativeBase)
		}
	}

}

func (v *VM) checkIncreaseTape(pos int) {
	if pos < len(v.Tape) {
		return
	}

	larger := make([]int, pos + 1)
	copy(larger, v.Tape)
	v.Tape = larger
}

func (v *VM) read() int {
	ret := v.In[0]

	v.In = v.In[1:]

	return ret
}

func (v *VM) write(i int) {
	v.Out = append(v.Out, i)
}

func (v VM) String() string {
	var r strings.Builder

	r.WriteString(fmt.Sprintf("PC:   %v", v.PC))

	r.WriteString(fmt.Sprintf("\nTape: %v", v.Tape))

	if len(v.In) != 0 {
		r.WriteString(fmt.Sprintf("\nIn:   %v", v.In))
	}

	if len(v.Out) != 0 {
		r.WriteString(fmt.Sprintf("\nOut:  %v", v.Out))
	}

	return r.String()
}
