package main

var opToFunc = map[int]func(subPackets []Packet) int{
	0: sum,
	1: product,
	2: min,
	3: max,
	5: gt,
	6: lt,
	7: eq,
}

func sum(subPackets []Packet) int {
	value := 0
	for _, subPacket := range subPackets {
		value += subPacket.Value()
	}

	return value
}

func product(subPackets []Packet) int {
	value := 1
	for _, subPacket := range subPackets {
		value *= subPacket.Value()
	}

	return value
}

func min(subPackets []Packet) int {
	value := -1
	for _, subPacket := range subPackets {
		subPacketValue := subPacket.Value()

		if value == -1 || subPacketValue < value {
			value = subPacketValue
		}
	}

	return value
}

func max(subPackets []Packet) int {
	value := -1
	for _, subPacket := range subPackets {
		subPacketValue := subPacket.Value()

		if value == -1 || subPacketValue > value {
			value = subPacketValue
		}
	}

	return value
}

func gt(subPackets []Packet) int {
	if subPackets[0].Value() > subPackets[1].Value() {
		return 1
	}

	return 0
}

func lt(subPackets []Packet) int {
	if subPackets[0].Value() < subPackets[1].Value() {
		return 1
	}

	return 0
}

func eq(subPackets []Packet) int {
	if subPackets[0].Value() == subPackets[1].Value() {
		return 1
	}

	return 0
}
