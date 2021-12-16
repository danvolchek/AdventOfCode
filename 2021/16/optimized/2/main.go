package main

import (
	"bytes"
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

func parse(r io.Reader) string {
	rawPacket, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	var packetBuilder strings.Builder

	for _, char := range bytes.TrimSpace(rawPacket) {
		base10Value, err := strconv.ParseInt(string(char), 16, 8)
		if err != nil {
			panic(err)
		}

		base2Value := strconv.FormatInt(base10Value, 2)
		for i := 0; i < 4-len(base2Value); i++ {
			packetBuilder.WriteByte('0')
		}

		for _, bit := range base2Value {
			packetBuilder.WriteRune(bit)
		}
	}

	return packetBuilder.String()
}

func solve(r io.Reader) {
	packetBits := parse(r)

	packet := parsePacket(&packetBits)
	if len(packetBits) != 0 && toInt(packetBits) != 0 {
		panic("bits remain after parsing outermost packet")
	}

	fmt.Println(packet.Value())
}

type Packet interface {
	Value() int
}

type literalPacket struct {
	value int
}

func (v literalPacket) Value() int {
	return v.value
}

type opPacket struct {
	op         func(subPackets []Packet) int
	subPackets []Packet
}

func (o opPacket) Value() int {
	return o.op(o.subPackets)
}

// consumeInt is like consume, but it converts the result to an int
func consumeInt(packetBits *string, numBits int) int {
	return toInt(consume(packetBits, numBits))
}

// consume returns numBits bits from packetBits, moving packetBits forward by that many bits
func consume(packetBits *string, numBits int) string {
	value := (*packetBits)[:numBits]
	*packetBits = (*packetBits)[numBits:]
	return value
}

// parsePacket parses the packet in base2 represented by packetBits, returning the packet and modifying packetBits
// to be the start of the next packet
func parsePacket(packetBits *string) Packet {
	_ = consumeInt(packetBits, 3) // skip version; not needed for part 2
	typeId := consumeInt(packetBits, 3)

	switch typeId {
	case 4:
		value := ""
		for {
			lastGroup := consume(packetBits, 1) == "0"
			group := consume(packetBits, 4)

			value += group

			if lastGroup {
				break
			}
		}

		return literalPacket{
			value: toInt(value),
		}

	default:
		op, ok := opToFunc[typeId]
		if !ok {
			panic("unknown op")
		}

		lengthTypeId := consume(packetBits, 1)

		var continueFunc func() bool

		switch lengthTypeId {
		case "0":
			packetsSize := consumeInt(packetBits, 15)

			startSize := len(*packetBits)

			continueFunc = func() bool {
				return startSize-len(*packetBits) < packetsSize
			}

		case "1":
			numPackets := consumeInt(packetBits, 11)

			i := 0
			continueFunc = func() bool {
				i++
				return i <= numPackets
			}

		default:
			panic("unknown length type id")
		}

		var subPackets []Packet
		for continueFunc() {
			subPackets = append(subPackets, parsePacket(packetBits))
		}

		return opPacket{
			op:         op,
			subPackets: subPackets,
		}
	}
}

func toInt(base2Value string) int {
	intValue, err := strconv.ParseInt(base2Value, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(intValue)
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
