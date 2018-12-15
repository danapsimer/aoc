package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	Generations = 20
)

type Bits struct {
	bits []bool
	zero int
}

func (s *Bits) Len() int {
	return len(s.bits)
}

func (s *Bits) First() int {
	return 0 - s.zero
}

func (s *Bits) Last() int {
	return s.Len() - s.zero - 1
}

func (s *Bits) Get(n int) bool {
	if s.First() <= n && n <= s.Last() {
		return s.bits[s.zero+n]
	} else {
		return false
	}
}

func (s *Bits) Set(n int, bit bool) {
	if n < s.First() {
		c := s.First() + n // works because n must be negative
		prepend := make([]bool,c)
		s.bits = append(prepend,s.bits...)
		s.zero += c
	} else if n > s.Last() {
		c := n - s.Last()
		app := make([]bool,c)
		s.bits = append(s.bits,app...)
	}
	s.bits[s.zero+n] = bit
}

func (s Bits) NextIndexOf(substr *Bits, start int) (int, bool) {
	substrLen := substr.Len()
	strLen := s.Len()
	for p := start; p < strLen; p++ {
		match := true
		for i, j := p, 0; j < substrLen; j++ {
			if s.Get(i) != substr.Get(j) {
				match = false
				break
			}
			i++
		}
		if match {
			return p, true
		}
	}
	return 0, false
}

func (s *Bits) Copy() *Bits {
	var newBits Bits
	newBits.bits = make([]bool,s.Len())
	copy(newBits.bits, s.bits)
	newBits.zero = s.zero
	return &newBits
}

func NewBitsFromSting(s string) *Bits {
	bits := &Bits{make([]bool,len(s)),0}
	for i, c := range s {
		bits.Set(i,c == '#')
	}
	return bits
}

func (bits *Bits) OffsetString(offset int) string {
	var b bytes.Buffer
	for i := -offset; i <= bits.Last(); i++ {
		if bits.Get(i) {
			b.WriteString("#")
		} else {
			b.WriteString(".")
		}
	}
	return string(b.Bytes())
}

func (bits *Bits) String() string {
	return bits.OffsetString(0)
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
func main() {
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
	currentState := NewBitsFromSting(state)
	fmt.Printf("initial state: %s\n", currentState.String())
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
		bits := NewBitsFromSting(pattern)
		rules[bits] = result == "#"
		fmt.Printf("%05s => %t\n", bits, rules[bits])
	}
	const ViewOff = 20
	const ViewWidth = ViewOff
	line1 := "             "
	line2 := "             "
	for i := -ViewOff; i <= 100; i++ {
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
	fmt.Printf("%11d: %s\n", 0, currentState.OffsetString(ViewWidth))
	for i := 0; i < Generations; i++ {
		nextState := currentState.Copy()

		for p, r := range rules {
			pos := currentState.First()
			for {
				pos, ok = currentState.NextIndexOf(p, pos)
				if !ok {
					break
				}
				nextState.Set(pos+2,r)
				//fmt.Printf("applying %s => %s at %2d resulting in %s\n", p,r,pos, nextState)
				pos += 1
			}
		}
		currentState = nextState
		fmt.Printf("%11d: %s\n", i+1, currentState.OffsetString(ViewWidth))
	}
	fmt.Printf("%11d: %s\n", Generations, currentState.OffsetString(ViewWidth))
	sum := 0
	for i := currentState.First(); i <= currentState.Last(); i++ {
		if currentState.Get(i) {
			sum += i
		}
	}
	fmt.Printf("result: %d\n", sum)
}
