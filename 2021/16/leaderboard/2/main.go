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
	input, err := os.Open(path.Join("2021", "16", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	rawPacket := scanner.Text()
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	var bits string

	for _, char := range rawPacket {
		intNum, err := strconv.ParseInt(string(char), 16, 32)
		if err != nil {
			panic(err)
		}

		boolNum := strconv.FormatInt(intNum, 2)
		for i := 0; i < 4-len(boolNum); i++ {
			bits += "0"
		}
		//fmt.Println(string(char), intNum, boolNum)

		for _, char2 := range strconv.FormatInt(intNum, 2) {
			bits += string(char2)
		}
	}

	//fmt.Println(rawPacket, bits)

	i := 0
	var packets []packet
	for len(bits) != 0 {
		var ps []packet
		ps, bits = parsePacket(bits[i:])
		for _, p := range ps {
			packets = append(packets, p)
		}
	}

	if len(packets) != 1 {
		panic("what")
	}

	fmt.Println(calc(packets[0]))

}

func calc(p packet) int {
	switch p.typeId {
	case 4:
		return p.value
	case 0:
		return sum(p.subPackets)
	case 1:
		return product(p.subPackets)
	case 2:
		return min(p.subPackets)
	case 3:
		return max(p.subPackets)
	case 5:
		return gt(p.subPackets)
	case 6:
		return lt(p.subPackets)
	case 7:
		return eq(p.subPackets)
	default:
		panic(p.typeId)
	}
}

func sum(ps []packet) int {
	v := 0
	for _, p := range ps {
		v += calc(p)
	}

	return v
}

func product(ps []packet) int {
	v := 1
	for _, p := range ps {
		v *= calc(p)
	}

	return v
}

func min(ps []packet) int {
	v, hasMin := 0, false
	for _, p := range ps {
		val := calc(p)
		if !hasMin || val < v {
			v = val
			hasMin = true
		}
	}

	return v
}

func max(ps []packet) int {
	v, hasMin := 0, false
	for _, p := range ps {
		val := calc(p)
		if !hasMin || val > v {
			v = val
			hasMin = true
		}
	}

	return v
}

func gt(ps []packet) int {
	if calc(ps[0]) > calc(ps[1]) {
		return 1
	}

	return 0
}


func lt(ps []packet) int {
	if calc(ps[0]) < calc(ps[1]) {
		return 1
	}

	return 0
}

func eq(ps []packet) int {
	if calc(ps[0]) == calc(ps[1]) {
		return 1
	}

	return 0
}



type packet struct {
	version int
	typeId  int

	value int
	subPackets []packet
}

func consume(p *string, amnt int) string {
	ret := (*p)[:amnt]
	*p = (*p)[amnt:]
	return ret
}

func consumeInt(p *string, amnt int) int {
	return toInt(consume(p, amnt))
}

func parsePacket(pp string) ([]packet, string) {
	if pp == strings.Repeat("0", len(pp)) {
		return nil, ""
	}


	var p = &pp
	version := consumeInt(p, 3)
	id := consumeInt(p,3)

	switch id {
	case 4:
		num := ""
		for {
			first := consume(p, 1)
			rest := consume(p, 4)

			num += rest

			if first == "0" {
				break
			}
		}

		return []packet{{
			version:    version,
			typeId:     id,
			value: toInt(num),
		}}, *p
	default:
		lengthTypeId := consume(p, 1)

		switch lengthTypeId == "0" {
		case true:
			packetsSize := consumeInt(p, 15)

			startSize := len(*p)
			var subpackets []packet
			for startSize - len(*p) < packetsSize {
				var newPackets []packet
				newPackets, *p = parsePacket(*p)

				for _, ss := range newPackets {
					subpackets = append(subpackets, ss)
				}
			}

			return []packet{{
				version:    version,
				typeId:     id,
				subPackets: subpackets,
			}}, *p

		case false:
			numPackets := consumeInt(p, 11)

			var subpackets []packet
			for i := 0; i < numPackets; i++ {
				var newPackets []packet
				newPackets, *p = parsePacket(*p)

				for _, ss := range newPackets {
					subpackets = append(subpackets, ss)
				}
			}

			return []packet{{
				version:    version,
				typeId:     id,
				subPackets: subpackets,
			}}, *p
		default:
			panic("excuse me")
		}
	}
}

func toInt(v string) int {
	vv, err := strconv.ParseInt(v, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(vv)
}

func main() {
	solve(strings.NewReader("C200B40A82"))
	solve(strings.NewReader("04005AC33890"))
	solve(strings.NewReader("880086C3E88112"))
	solve(strings.NewReader("CE00C43D881120"))
	solve(strings.NewReader("D8005AC2A8F0"))
	solve(strings.NewReader("F600BC2D8F"))
	solve(strings.NewReader("9C005AC2F8F0"))
	solve(strings.NewReader("9C0141080250320F1802104A08"))
	solve(input())
}
