package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"runtime/pprof"
)

var (
	Generations = 20
	Debug       = 0
	Zero        = big.NewInt(0)
	One         = big.NewInt(1)
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func init() {
	flag.IntVar(&Generations, "g", 20, "number of generations to run")
	flag.IntVar(&Debug, "D", 0, "Set debug output level")
}

type Bits struct {
	bits      *big.Int
	len, zero int
}

func (s *Bits) Copy() *Bits {
	var newBits Bits
	newBits.bits = big.NewInt(0)
	newBits.bits.Set(s.bits)
	newBits.zero = s.zero
	newBits.len = s.len
	return &newBits
}

func NewBitsFromString(s string) *Bits {
	bits := &Bits{big.NewInt(0), 1, 0}
	for i, c := range s {
		bits.Set(i, c == '#')
	}
	return bits
}

func NewBits(n int64) *Bits {
	return &Bits{big.NewInt(n), 64, 0}
}

func (s *Bits) Len() int {
	return s.len
}

func (s *Bits) First() int {
	return 0 - s.zero
}

func (s *Bits) Last() int {
	return s.Len() - s.zero - 1
}

func (s *Bits) ContainsAt(n int, cmp *Bits) bool {
	if n < s.First() {
		s.growNegative(n)
	}
	fmt.Printf("s.bits = %10s\n",s.bits.Text(2))
	idx := s.zero + n
	mask := big.NewInt(0)
	mask = mask.Lsh(cmp.bits,uint(idx))
	fmt.Printf("mask =   %10s\n", mask.Text(2))
	result := big.NewInt(0)
	result = result.Xor(s.bits,mask)
	fmt.Printf("result = %10s\n",result.Text(2))
	return result.Cmp(Zero) == 0
}

func (s *Bits) Get(n int) bool {
	if s.First() <= n && n <= s.Last() {
		return s.bits.Bit(s.zero+n) == 1
	} else {
		return false
	}
}

func (s *Bits) Set(n int, bit bool) {
	if n < s.First() {
		s.growNegative(n)
	}
	bitIdx := s.zero + n
	b := uint(0)
	if bit {
		b = 1
	}
	s.bits = s.bits.SetBit(s.bits, bitIdx, b)
	if bitIdx >= s.len {
		s.len = bitIdx + 1
	}
}

func (s *Bits) growNegative(n int) {
	// works because n must be negative and first is always <= 0
	c := abs(n) + s.First()
	s.zero += c
	s.len += c
	s.bits = s.bits.Lsh(s.bits, uint(c))
}

func (s *Bits) NextIndexOf(substr *Bits, start int) (int, bool) {
	strLast := s.Last()
	for p := start; p <= strLast; p++ {
		if s.ContainsAt(p,substr) {
			return p, true
		}
	}
	return 0, false
}

func (bits *Bits) OffsetString(offset, width int) string {
	var b bytes.Buffer
	for i := -offset; i <= width-offset; i++ {
		if bits.Get(i) {
			b.WriteString("#")
		} else {
			b.WriteString(".")
		}
	}
	return string(b.Bytes())
}

func (bits *Bits) String() string {
	return bits.OffsetString(bits.First(), bits.Len())
}

func (bits *Bits) Equal(other *Bits) bool {
	return bits.bits.Cmp(other.bits) == 0 && bits.len == other.len && bits.zero == other.zero
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	currentState, rules := LoadInput()
	const ViewOff = 20
	const ViewWidth = ViewOff + 150
	if Debug >= 5 {
		line1 := "             "
		line2 := "             "
		for i := -ViewOff; i <= ViewWidth-ViewOff; i++ {
			if i%10 == 0 {
				line1 = line1 + fmt.Sprintf("%d", abs(i)/10%10)
				line2 = line2 + "0"
			} else {
				line1 = line1 + " "
				line2 = line2 + " "
			}
		}
		fmt.Println(line1)
		fmt.Println(line2)
		fmt.Printf("%11d: %s\n", 0, currentState.OffsetString(ViewOff, ViewWidth))
	}
	for i := 0; i < Generations; i++ {
		nextState := currentState.Copy()

		for p, r := range rules {
			pos := currentState.First() - 10
			ok := false
			for {
				pos, ok = currentState.NextIndexOf(p, pos)
				if !ok {
					break
				}
				nextState.Set(pos+2, r)
				if Debug >= 9 {
					fmt.Printf("applying %s => %t at %2d resulting in %s\n", p, r, pos, nextState)
				}
				pos += 1
			}
		}
		currentState = nextState
		if Debug >= 5 {
			fmt.Printf("%11d: %s\n", i+1, currentState.OffsetString(ViewOff, ViewWidth))
		}
	}
	if Debug >= 1 {
		fmt.Printf("%11d: %s\n", Generations, currentState.String())
	}
	sum := 0
	for i := currentState.First(); i <= currentState.Last(); i++ {
		if currentState.Get(i) {
			sum += i
		}
	}
	fmt.Printf("result: %d\n", sum)
}

func LoadInput() (*Bits, map[*Bits]bool) {
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		fmt.Printf("ERROR: Could not read first line of input\n")
		os.Exit(-1)
	}
	line := scanner.Text()
	var state string
	_, err := fmt.Sscanf(line, "initial state: %s", &state)
	if err != nil {
		fmt.Printf("ERROR: parsing first line of input: %s\n%s\n", line, err.Error())
		os.Exit(-1)
	}
	currentState := NewBitsFromString(state)
	if Debug >= 1 {
		fmt.Printf("initial state: %s\n", currentState.String())
	}
	scanner.Scan() // skip the empty line
	rules := make(map[*Bits]bool)
	for scanner.Scan() {
		line = scanner.Text()
		var pattern, result string
		_, err = fmt.Sscanf(line, "%s => %s", &pattern, &result)
		if err != nil {
			fmt.Printf("ERROR: parsing rule: %s\n%s\n", line, err.Error())
			os.Exit(-1)
		}
		bits := NewBitsFromString(pattern)
		rules[bits] = result == "#"
		if Debug >= 1 {
			fmt.Printf("%05s => %t\n", bits, rules[bits])
		}
	}
	return currentState, rules
}
