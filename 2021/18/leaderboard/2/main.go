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

func input() *os.File {
	input, err := os.Open(path.Join("2021", "18", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

var debug = false

type number struct {
	regular bool

	value int

	isLeft, isRight bool
	left, right     *number

	parent *number
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var nums []*number
	for scanner.Scan() {
		line := scanner.Text()

		v := parseNumber(line)
		nums = append(nums, v)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	maxMag := 0
	for i, numOne := range nums {
		for j, numTwo := range nums {
			if i == j {
				continue
			}

			mag := numOne.add(numTwo).magnitude()
			if mag > maxMag {
				maxMag = mag
			}

			mag = numTwo.add(numOne).magnitude()
			if mag > maxMag {
				maxMag = mag
			}
		}
	}
	fmt.Println("ANSWER", maxMag)
}

func parseNumber(raw string) *number {
	if strings.Index(raw, "[") == -1 {
		return &number{
			regular: true,
			value:   toInt(raw),
		}
	}

	raw = raw[1 : len(raw)-1]

	var parts []string

	start, end := 0, 0
	for end < len(raw) {
		nesting := 0

		for {
			if raw[end] == '[' {
				nesting++
			} else if raw[end] == ']' {
				nesting--
			}

			end++

			if end == len(raw) || (raw[end] == ',' && nesting == 0) {
				break
			}
		}

		parts = append(parts, raw[start:end])
		end++
		start = end
	}

	if len(parts) != 2 {
		panic("confuse")
	}

	num := &number{
		regular: false,
		left:    parseNumber(parts[0]),
		right:   parseNumber(parts[1]),
	}

	num.left.parent = num
	num.left.isLeft = true
	num.right.parent = num
	num.right.isRight = true

	return num
}

func (n *number) add(other *number) *number {
	num := &number{
		regular: false,
		left:    n.clone(),
		right:   other.clone(),
	}

	num.left.parent = num
	num.left.isLeft = true
	num.right.parent = num
	num.right.isRight = true

	num.reduce()

	return num
}

func (n *number) reduce() {
	for {
		if n.explode() {
			continue
		}

		if n.split() {
			continue
		}

		break
	}
}

func (n *number) clone() *number {
	if n.regular {
		return &number{
			regular: true,
			value:   n.value,
			isLeft:  n.isLeft,
			isRight: n.isRight,
		}
	}

	newNum := &number{
		regular: false,
		value:   0,
		isLeft:  n.isLeft,
		isRight: n.isRight,
		left:    n.left.clone(),
		right:   n.right.clone(),
		parent:  nil,
	}

	newNum.left.parent = newNum
	newNum.right.parent = newNum

	return newNum
}

func (n *number) explode() bool {

	num := n.findPairToExplode(0)
	if num == nil {
		return false
	}

	if debug {
		fmt.Printf("to explode: %v\n", num)
	}
	num.doExplode()

	return true
}

func (n *number) findPairToExplode(nesting int) *number {
	if n.regular {
		return nil
	}

	if nesting < 4 {
		toExplode := n.left.findPairToExplode(nesting + 1)

		if toExplode == nil {
			toExplode = n.right.findPairToExplode(nesting + 1)
		}

		return toExplode
	}

	return n
}

func (n *number) doExplode() {

	/*
		If I am the left pair of parent:
			- The closest right is the left most child of (first left parent).right
			- The closest left is the right most child of (first right parent).left
		If I am the right pair of parent:
			- The closest left is the right most child of (first right parent).left
			- The closest right is the left most child of (first left parent).right
	*/
	a := n
	for !a.isRight {
		a = a.parent

		if !a.isRight && !a.isLeft { //root
			a = nil
			break
		}
	}
	if a != nil {
		a = a.parent.left
		for !a.regular {
			a = a.right
		}
	}

	b := n
	for !b.isLeft {
		b = b.parent

		if !b.isRight && !b.isLeft { //root
			b = nil
			break
		}
	}
	if b != nil {
		b = b.parent.right
		for !b.regular {
			b = b.left
		}
	}

	var leftRegular, rightRegular *number
	rightRegular = b
	leftRegular = a

	/*leftVal := "nil"
	if leftRegular != nil {
		leftVal = strconv.Itoa(leftRegular.value)
	}

	rightVal := "nil"
	if rightRegular != nil {
		rightVal = strconv.Itoa(rightRegular.value)
	}
	fmt.Println("left:", leftVal, "right:", rightVal)
	panic("foo")
	*/
	if leftRegular != nil {
		leftRegular.value += n.left.value
	}

	if rightRegular != nil {
		rightRegular.value += n.right.value
	}

	n.regular = true
	n.value = 0
	n.left = nil
	n.right = nil
}

func (n *number) String() string {
	if n.regular {
		return strconv.Itoa(n.value)
	}

	return "[" + n.left.String() + "," + n.right.String() + "]"
}

/// [[[1,2],[3,4]],[[5,6],[7,8]]]
//                 .
//             .       .
//           .   .   .   .
//          1 2 3 4 5 6 7 8

/*
[[6,[5,[4,[3,2]]]],1]

             .
           .   1
         6   .
           5   .
             4   .
               3   2
*/

/*
[[[[[9,8],1],2],3],4]
            .
          .   4
        .   3
      .   2
    .   1
  9   8

*/

func (n *number) split() bool {
	if n.regular {
		if n.value < 10 {
			return false
		}
		if debug {
			fmt.Println("to split: ", n, "from", n.parent)
		}

		n.regular = false

		lv, rv := n.value/2, n.value/2
		if lv+rv != n.value {
			rv += 1
		}

		n.left = &number{
			regular: true,
			value:   lv,
			isLeft:  true,
			parent:  n,
		}

		n.right = &number{
			regular: true,
			value:   rv,
			isRight: true,
			parent:  n,
		}

		n.value = 0

		return true
	}

	if n.left.split() {
		return true
	}

	return n.right.split()
}

func (n *number) magnitude() int {
	if n.regular {
		return n.value
	}

	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func toInt(v string) int {
	vv, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return vv
}

func dsolve(r io.Reader) {
	defer func() {
		if r := recover(); r != nil {
			if r != "foo" {
				panic(r)
			}
		}
	}()

	solve(r)
}

func main() {
	/*dsolve(strings.NewReader("[[[[[9,8],1],2],3],4]"))
	dsolve(strings.NewReader("[7,[6,[5,[4,[3,2]]]]]"))
	dsolve(strings.NewReader("[[6,[5,[4,[3,2]]]],1]"))
	dsolve(strings.NewReader("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"))
	dsolve(strings.NewReader("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"))*/
	solve(strings.NewReader("[[[[4,3],4],4],[7,[[8,4],9]]]\n[1,1]"))
	solve(strings.NewReader("[1,1]\n[2,2]\n[3,3]\n[4,4]"))
	solve(strings.NewReader("[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]"))
	solve(strings.NewReader("[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]\n[6,6]"))
	solve(strings.NewReader("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]\n[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]\n[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]\n[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]\n[7,[5,[[3,8],[1,4]]]]\n[[2,[2,2]],[8,[8,1]]]\n[2,9]\n[1,[[[9,3],9],[[9,0],[0,7]]]]\n[[[5,[7,4]],7],1]\n[[[[4,2],2],6],[8,7]]"))
	solve(strings.NewReader("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]\n[[[5,[2,8]],4],[5,[[9,9],0]]]\n[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]\n[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]\n[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]\n[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]\n[[[[5,4],[7,7]],8],[[8,3],8]]\n[[9,3],[[9,9],[6,[4,9]]]]\n[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]\n[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"))
	//solve(strings.NewReader("[1,2]\n[[1,2],3]\n[9,[8,7]]\n[[1,9],[8,5]]\n[[[[1,2],[3,4]],[[5,6],[7,8]]],9]\n[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]\n[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"))
	solve(input())
}
