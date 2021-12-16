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

	fmt.Println(packet.VersionSum())
}

type Packet interface {
	VersionSum() int
}

type versioned struct {
	version int
}

type literalPacket struct {
	versioned
}

func (v literalPacket) VersionSum() int {
	return v.version
}

type opPacket struct {
	versioned

	subPackets []Packet
}

func (o opPacket) VersionSum() int {
	sum := o.version

	for _, subPacket := range o.subPackets {
		sum += subPacket.VersionSum()
	}

	return sum
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
	version := consumeInt(packetBits, 3)
	typeId := consumeInt(packetBits, 3)

	switch typeId {
	case 4:
		for {
			stop := consume(packetBits, 1) == "0"
			_ = consume(packetBits, 4) // ignore value, it's not used for part 1

			if stop {
				break
			}
		}

		return literalPacket{
			versioned: versioned{version: version},
		}

	default:
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
			versioned:  versioned{version: version},
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
	solve(strings.NewReader("D2FE28"))
	solve(strings.NewReader("38006F45291200"))
	solve(strings.NewReader("EE00D40C823060"))
	solve(strings.NewReader("8A004A801A8002F478"))
	solve(strings.NewReader("620080001611562C8802118E34"))
	solve(strings.NewReader("C0015000016115A2E0802F182340"))
	solve(strings.NewReader("A0016C880162017C3686B18A3D4780"))
	solve(input())
}
