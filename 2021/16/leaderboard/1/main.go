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

	fmt.Println(rawPacket, bits)

	i := 0
	var packets []packet
	for len(bits) != 0 {
		var ps []packet
		ps, bits = parsePacket(bits[i:])
		for _, p := range ps {
			packets = append(packets, p)
		}
	}

	vSum := 0
	for _, p := range packets {
		vSum += sum(p)
	}

	fmt.Println(vSum)

}

func sum(p packet) int {
	v := p.version
	for _, pp := range p.subPackets {
		v += sum(pp)
	}

	return v
}

type packet struct {
	version int
	typeId  int

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
	id := consumeInt(p, 3)

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

		//fmt.Println(toInt(num))
		return []packet{{
			version: version,
			typeId:  id,
		}}, *p
	default:
		lengthTypeId := consume(p, 1)

		switch lengthTypeId == "0" {
		case true:
			packetsSize := consumeInt(p, 15)

			startSize := len(*p)
			var subpackets []packet
			for startSize-len(*p) < packetsSize {
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
	vv, err := strconv.ParseInt(v, 2, 16)
	if err != nil {
		panic(err)
	}

	return int(vv)
}

func main() {
	solve(strings.NewReader("D2FE28"))
	solve(strings.NewReader("38006F45291200"))
	solve(strings.NewReader("EE00D40C823060"))
	solve(strings.NewReader("8A004A801A8002F478"))
	solve(strings.NewReader("620080001611562C8802118E34"))
	solve(strings.NewReader("C0015000016115A2E0802F182340"))
	solve(strings.NewReader("A0016C880162017C3686B18A3D4780"))
	solve(input())
}
