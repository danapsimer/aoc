package bits

import (
	"fmt"
	"math/big"
)

type PacketType int

const (
	UnknownPacketType PacketType = iota
	LiteralPacketType
)

var LiteralTypeId = big.NewInt(4)

type Packet interface {
	Version() int8
	Type() PacketType
	TypeId() int8
	Content() interface{}
}

type packet struct {
	version int8
	typeId  int8
}

func (p *packet) Version() int8 {
	return p.version
}

func (p *packet) Type() PacketType {
	if LiteralTypeId.Cmp(big.NewInt(int64(p.typeId))) == 0 {
		return LiteralPacketType
	}
	return UnknownPacketType
}

func (p *packet) TypeId() int8 {
	return p.typeId
}

func (p *packet) Content() interface{} {
	return nil
}

type literalPacket struct {
	packet
	value *big.Int
}

func (lp *literalPacket) Content() interface{} {
	return lp.value
}

type operatorPacket struct {
	packet
	content []Packet
}

func (op *operatorPacket) Content() interface{} {
	return op.content
}

type bit = byte

func readNBits(bitStream <-chan bit, n int) (*big.Int, int, error) {
	value := big.NewInt(0)
	bits := 0
	for n > 0 {
		b, ok := <-bitStream
		if !ok {
			return nil, bits, fmt.Errorf("channel closed after reading %d bits but before reading %d bits", bits, n)
		}
		bits += 1
		value = value.Lsh(value, 1).Or(value, big.NewInt(int64(b)))
		n -= 1
	}
	return value, bits, nil
}

func readLiteralValue(bitStream <-chan bit) (*big.Int, int, error) {
	literal := big.NewInt(0)
	bits := 0
	for {
		more, ok := <-bitStream
		if !ok {
			return nil, 0, fmt.Errorf("unable to read literal, failed to get more bit")
		}
		bits += 1
		v, b, err := readNBits(bitStream, 4)
		if err != nil {
			return nil, bits + b, fmt.Errorf("unable to read literal part: %s", err.Error())
		}
		bits += b
		literal = literal.Lsh(literal, 4).Or(literal, v)
		if more == 0 {
			break
		}
	}
	return literal, bits, nil
}

func readLength(bitStream <-chan bit) (*big.Int, bool, int, error) {
	bits := 0
	i, ok := <-bitStream
	if !ok {
		return nil, false, 0, fmt.Errorf("stream closed while attempting to read size bit for operator length")
	}
	bits += 1
	n := 11
	if i == 0 {
		n = 15
	}
	v, b, err := readNBits(bitStream, n)
	if err != nil {
		return nil, false, bits + b, fmt.Errorf("error reading operation length: %s", err.Error())
	}
	bits += b
	return v, i == 0, bits, nil
}

func readOpPacketContent(bitStream <-chan bit) ([]Packet, int, error) {
	bits := 0
	length, lengthIsBits, b, err := readLength(bitStream)
	if err != nil {
		return nil, b, fmt.Errorf("error reading operator packet length: %s", err.Error())
	}
	bits += b
	if lengthIsBits {
		numBits := length.Int64()
		packets := make([]Packet, 0, 10)
		for operandBits := int64(0); operandBits < numBits; {
			packet, b, err := readPacket(bitStream)
			if err != nil {
				return nil, bits + b, fmt.Errorf("error reading sub-packet: %s", err.Error())
			}
			operandBits += int64(b)
			bits += b
			packets = append(packets, packet)
		}
		return packets, bits, nil
	} else {
		numPackets := length.Int64()
		packets := make([]Packet, numPackets)
		for p := int64(0); p < numPackets; p++ {
			packet, b, err := readPacket(bitStream)
			if err != nil {
				return nil, bits + b, fmt.Errorf("error reading the %d operaand: %s", p+1, err.Error())
			}
			bits += b
			packets[p] = packet
		}
		return packets, bits, nil
	}
}

func readPacket(bitStream <-chan bit) (Packet, int, error) {
	bits := 0

	// READ VERSION
	version, b, err := readNBits(bitStream, 3)
	if err != nil {
		return nil, bits + b, fmt.Errorf("error reading packet version: %s", err.Error())
	}
	bits += b

	// Read Type Id
	typeId, b, err := readNBits(bitStream, 3)
	if err != nil {
		return nil, bits + b, fmt.Errorf("error reading packet type id: %s", err.Error())
	}
	bits += b

	if typeId.Cmp(LiteralTypeId) == 0 {
		v, b, err := readLiteralValue(bitStream)
		if err != nil {
			return nil, bits + b, fmt.Errorf("error reading literal value: %s", err.Error())
		}
		bits += b
		return &literalPacket{packet{int8(version.Int64()), int8(typeId.Int64())}, v}, bits, nil
	} else {
		content, b, err := readOpPacketContent(bitStream)
		if err != nil {
			return nil, bits + b, fmt.Errorf("error reading op packet content: %s", err.Error())
		}
		return &operatorPacket{packet{int8(version.Int64()), int8(typeId.Int64())}, content}, b, nil
	}
}

func hexCharChannelToBitChannel(runeStream <-chan rune) <-chan bit {
	bitStream := make(chan bit)
	go func() {
		defer close(bitStream)
		for r := range runeStream {
			v := uint8(0)
			if r >= 'A' {
				v = uint8(10 + r - 'A')
			} else {
				v = uint8(r - '0')
			}
			for b := 3; b >= 0; b-- {
				bitStream <- ((uint8(1) << b) & v) >> b
			}
		}
	}()
	return bitStream
}

func stringToRuneStream(str string) <-chan rune {
	runeStream := make(chan rune)
	go func() {
		defer close(runeStream)
		for _, r := range str {
			runeStream <- r
		}
	}()
	return runeStream
}

func SumVersions(root Packet) int {
	sum := int(root.Version())
	if root.Type() != LiteralPacketType {
		content := root.Content().([]Packet)
		for _, p := range content {
			sum += SumVersions(p)
		}
	}
	return sum
}
