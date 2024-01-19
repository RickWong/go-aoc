package day16

import (
	_ "embed"
	"encoding/hex"
	"fmt"
	"github.com/samber/lo"
	"github.com/zyedidia/generic/stack"
	"strconv"
	"strings"
	"testing"
)

//go:embed example.txt
var Example string

//go:embed input.txt
var Input string

var data = Input

type Packet struct {
	version, typeId int
	literal         *int
	lengthType      *int
	subpackets      []*Packet
}

func consumeBits(binary string, offset *int, bits int) string {
	if bits == 0 {
		bits = len(binary)
	}
	read := binary[*offset : *offset+bits]
	*offset += bits
	return read
}

func consumeBool(binary string, offset *int) bool {
	return consumeBits(binary, offset, 1) == "1"
}

func consumeInt(binary string, offset *int, bits int) int {
	return bin2dec(consumeBits(binary, offset, bits))
}

func bin2dec(binary string) int {
	v, _ := strconv.ParseInt(binary, 2, 0)
	return int(v)
}

func part1() int {
	lines := strings.Split(data, "\n")
	bytes, _ := hex.DecodeString(lines[0])
	binary := strings.Join(
		lo.Map(bytes, func(byte byte, _ int) string {
			return fmt.Sprintf("%08b", byte)
		}),
		"")

	offset := 0
	packet := decodePacket(binary, &offset)

	sum := 0
	packets := stack.New[*Packet]()
	packets.Push(packet)
	for packets.Size() > 0 {
		current := packets.Pop()
		sum += current.version
		for _, subpacket := range current.subpackets {
			packets.Push(subpacket)
		}
	}

	return sum
}

func decodePacket(packet string, offset *int) *Packet {
	version := consumeInt(packet, offset, 3)
	typeId := consumeInt(packet, offset, 3)

	switch typeId {
	case 4:
		literalStr := ""
		for i := 0; ; i += 5 {
			more := consumeBool(packet, offset)
			nibble := consumeBits(packet, offset, 4)
			literalStr += nibble
			if !more {
				break
			}
		}

		literal := bin2dec(literalStr)

		return &Packet{
			version,
			typeId,
			&literal,
			nil,
			nil,
		}
	default:
		subpackets := make([]*Packet, 0)
		lengthType := consumeInt(packet, offset, 1)
		if lengthType == 0 {
			numBitsSubpackets := consumeInt(packet, offset, 15)
			subpacketsBits := consumeBits(packet, offset, numBitsSubpackets)
			subpacketsOffset := 0
			for subpacketsOffset < numBitsSubpackets {
				subpackets = append(subpackets, decodePacket(subpacketsBits, &subpacketsOffset))
			}
		} else {
			numSubpackets := consumeInt(packet, offset, 11)
			for i := 0; i < numSubpackets; i++ {
				subpackets = append(subpackets, decodePacket(packet, offset))
			}
		}

		return &Packet{
			version,
			typeId,
			nil,
			&lengthType,
			subpackets,
		}
	}
}

func TestPart1(t *testing.T) {
	result := part1()
	expect := 20
	if data == Input {
		expect = 886
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}

func part2() int {
	lines := strings.Split(data, "\n")
	bytes, _ := hex.DecodeString(lines[0])
	binary := strings.Join(
		lo.Map(bytes, func(byte byte, _ int) string {
			return fmt.Sprintf("%08b", byte)
		}),
		"")

	offset := 0
	packet := decodePacket(binary, &offset)
	value := evaluatePacket(packet)

	return value
}

func evaluatePacket(packet *Packet) int {
	value := 0

	subvalues := make([]int, 0, len(packet.subpackets))
	for _, sp := range packet.subpackets {
		subvalues = append(subvalues, evaluatePacket(sp))
	}

	switch packet.typeId {
	case 0:
		value = lo.Sum(subvalues)
	case 1:
		value = 1
		for _, subvalue := range subvalues {
			value *= subvalue
		}
	case 2:
		value = lo.Min(subvalues)
	case 3:
		value = lo.Max(subvalues)
	case 4:
		value = *packet.literal
	case 5:
		if subvalues[0] > subvalues[1] {
			value = 1
		}
	case 6:
		if subvalues[0] < subvalues[1] {
			value = 1
		}
	case 7:
		if subvalues[0] == subvalues[1] {
			value = 1
		}
	}

	return value
}

func TestPart2(t *testing.T) {
	result := part2()
	expect := 1
	if data == Input {
		expect = 184487454837
	}

	if result != expect {
		t.Errorf("Result was incorrect, got: %d, expect: %d.", result, expect)
	}
}
