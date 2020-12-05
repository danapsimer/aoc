package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
)

var occupiedEmojis = []string{
	"\U0001F64D", "\U0001F64E", "\U0001F648", "\U0001F469", "\U0001F468",
}

func main() {
	seating := readSeating(os.Stdin)
	fmt.Println("       ╭──╮")
	fmt.Println("      ╱    ╲")
	fmt.Println("     ╱      ╲")
	fmt.Println("    ╱        ╲")
	fmt.Println("   ╱          ╲")
	fmt.Println("  ╱            ╲")
	fmt.Println(" ╱              ╲")
	fmt.Println("╱                ╲")
	fmt.Println("┣━━━━━━━━━━━━━━━━━┫")
	for r := uint8(0); r < 128; r++ {
		fmt.Print("┃")
		for c := uint8(0); c < 8; c++ {
			bp := NewBoardingPass(r, c)
			if seating[bp] {
				fmt.Print(occupiedEmojis[rand.Intn(len(occupiedEmojis))])
			} else {
				fmt.Print("\U0001F4BA")
			}
			if c == 3 {
				fmt.Print(" ")
			}
		}
		fmt.Println("┃")
	}
	fmt.Println("┗━━━━━━━━━━━━━━━━━┛")
	foundLargest := false
	for bp := boardingPass(1023); bp < 1024; bp-- {
		if seating[bp] {
			if !foundLargest {
				log.Printf("largest seat id is: %d", bp)
				foundLargest = true
			}
		} else {
			n := bp + 1
			p := bp - 1
			if (p < 1024 && seating[p]) && (n < 1024 && seating[n]) {
				log.Printf("Your seat is %d, %d, id = %d", bp.row(), bp.col(), bp)
			}
		}
	}
}

type boardingPass uint16

func (bp boardingPass) row() uint8 {
	return uint8(bp >> 3)
}

func (bp boardingPass) col() uint8 {
	return uint8(bp & 0x0007)
}

func NewBoardingPass(r, c uint8) boardingPass {
	return boardingPass(uint16(r)<<3 | uint16(c))
}

func ParseBoardingPass(bpstr string) boardingPass {
	var bp boardingPass
	for idx, c := range bpstr {
		if idx > 0 {
			bp <<= 1
		}
		if c == 'B' || c == 'R' {
			bp |= 0x1
		}
	}
	return bp
}

func readSeating(r io.Reader) []bool {
	scanner := bufio.NewScanner(r)
	g := make([]bool, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		bp := ParseBoardingPass(line)
		g[bp] = true
	}
	return g
}
