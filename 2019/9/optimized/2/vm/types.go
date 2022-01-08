package vm

import (
	"fmt"
	"math"
)

const (
	positionMode  = 0
	immediateMode = 1
	relativeMode  = 2
)

type instruction int

func (i instruction) Code() int {
	return int(i) % 100
}

func (i instruction) Mode(p int) int {
	v := int(i)

	return (v / int(math.Pow(10, float64(p+2)))) % 10
}

type arg interface {
	read() int
	write(v int)
}

type immediateArg struct {
	back int
}

func (i immediateArg) read() int {
	return i.back
}

func (i immediateArg) write(v int) {
	panic("can't write immediate mode arg")
}

func (i immediateArg) String() string {
	return fmt.Sprintf("i(%v)", i.back)
}

type positionArg struct {
	raw int

	back *int
}

func (p positionArg) read() int {
	return *p.back
}

func (p positionArg) write(v int) {
	*p.back = v
}

func (p positionArg) String() string {
	return fmt.Sprintf("p(%v -> %v)", p.raw, *p.back)
}

type relativeArg struct {
	raw  int
	base int

	back *int
}

func (r relativeArg) read() int {
	return *r.back
}

func (r relativeArg) write(v int) {
	*r.back = v
}

func (r relativeArg) String() string {
	return fmt.Sprintf("r(%v + %v -> %v)", r.raw, r.base, *r.back)
}
