package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"regexp"
	"strconv"
)

type OpType int

const (
	SetMask OpType = iota
	SetMemory
)

type Operation struct {
	Op              OpType
	Address         uint64
	Value           uint64
	AndMask, OrMask uint64
}

var setMaskRegex = regexp.MustCompile("^mask\\s*=\\s*([X10]+)$")
var setMemoryRegex = regexp.MustCompile("^mem\\[(\\d+)]\\s*=\\s*(\\d+)$")

func ReadProgram(reader io.Reader) <-chan Operation {
	opChan := make(chan Operation)
	go func() {
		defer close(opChan)
		scanner := bufio.NewScanner(reader)
		lineNo := 0
		for scanner.Scan() {
			line := scanner.Text()
			match := setMaskRegex.FindStringSubmatch(line)
			if match == nil {
				match = setMemoryRegex.FindStringSubmatch(line)
				if match == nil {
					panic(fmt.Errorf("%d: unknown operation: %s", lineNo, line))
				}
				pa := func(s string) uint64 {
					address, err := strconv.ParseUint(s, 10, 36)
					if err != nil {
						panic(fmt.Errorf("%d: invalid address: %s - %s", lineNo, match[1], err.Error()))
					}
					return uint64(address)
				}
				pv := func(s string) uint64 {
					value, err := strconv.ParseUint(s, 10, 36)
					if err != nil {
						panic(fmt.Errorf("%d: invalid value: %s - %s", lineNo, match[2], err.Error()))
					}
					return uint64(value)
				}
				address := pa(match[1])
				value := pv(match[2])
				opChan <- Operation{SetMemory, address, value, 0, 0}
			} else {
				andMask := uint64(0)
				orMask := uint64(0)
				for i, c := range match[1] {
					if i > 0 {
						andMask <<= 1
						orMask <<= 1
					}
					switch c {
					case 'X':
						andMask |= 1
					case '1':
						orMask |= 1
					}
				}
				opChan <- Operation{SetMask, 0, 0, andMask, orMask}
			}
			lineNo += 1
		}
	}()
	return opChan
}

func RunProgramPart1(program <-chan Operation) map[uint64]uint64 {
	andMask := uint64(0)
	orMask := uint64(0)
	memory := make(map[uint64]uint64)
	for op := range program {
		switch op.Op {
		case SetMask:
			andMask = op.AndMask
			orMask = op.OrMask
		case SetMemory:
			memory[op.Address] = op.Value&andMask | orMask
		}
	}
	return memory
}

func RunProgramPart2(program <-chan Operation) map[uint64]uint64 {
	floatMask := uint64(0)
	orMask := uint64(0)
	memory := make(map[uint64]uint64)
	for op := range program {
		switch op.Op {
		case SetMask:
			floatMask = op.AndMask
			orMask = op.OrMask
		case SetMemory:
			bcount := bits.OnesCount64(uint64(floatMask))
			maxFloatMastCompact := uint64(0x1 << bcount)
			for floatMaskCompact := uint64(0); floatMaskCompact < maxFloatMastCompact; floatMaskCompact += 1 {
				fMask := uint64(0)
				fMaskBits := floatMaskCompact
				for b, bc := 0, bcount; b < 64 && bc > 0; b++ {
					if (floatMask >> b & 0x1) != 0 {
						if (fMaskBits & 0x1) != 0 {
							fMask |= 0x1 << b
						}
						fMaskBits >>= 1
						bc -= 1
					}
				}
				memory[op.Address&^floatMask|fMask|orMask] = op.Value
			}
		}
	}
	return memory
}

func Scatter(in <-chan Operation,out... chan<- Operation) {
	go func() {
		defer func () {
			for _, o := range out {
				close(o)
			}
		}()
		for item := range in {
			for _, o := range out {
				o <- item
			}
		}
	}()
}

func main() {
	prgChanPart1 := make(chan Operation)
	prgChanPart2 := make(chan Operation)
	Scatter(ReadProgram(os.Stdin),prgChanPart1,prgChanPart2)
	go func() {
		for op := range prgChan {
			prgChanPart1 <- op
			prgChanPart2 <- op
		}
		close(prgChanPart1)
		close(prgChanPart2)
	}()
	part1Done := make(chan bool)
	go func() {
		memory := RunProgramPart1(prgChanPart1)
		sum := uint64(0)
		for _, v := range memory {
			sum += v
		}
		log.Printf("Part1: sum of memory values = %d", sum)
		part1Done <- true
	}()
	part2Done := make(chan bool)
	go func() {
		memory := RunProgramPart2(prgChanPart2)
		sum := uint64(0)
		for _, v := range memory {
			sum += v
		}
		log.Printf("Part2: sum of memory values = %d", sum)
		part2Done <- true
	}()
	<-part1Done
	<-part2Done
}
