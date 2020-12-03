package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

var (
	pattern    = []int{1, 0, -1, 0}
	patternLen = len(pattern)
)

func Day16Part1(digits []byte, iterations, outputOffset int) string {
	numElements := len(digits)

	for itr := 0; itr < iterations; itr += 1 {
		for i := numElements - 2; i >= outputOffset; i-- {
			digits[i] = (digits[i] + digits[i+1]) % 10
		}
	}
	output := ""
	for i := outputOffset; i < outputOffset+8; i++ {
		output = output + string([]byte{'0' + digits[i]})
	}
	return output
}

func Day16Part2(digits []byte, iterations int) string {
	explodedDigits := make([]byte, len(digits)*10000)
	for i := 0; i < 10000; i++ {
		copy(explodedDigits[i*len(digits):], digits)
	}
	outputOffset := 0
	for i := 0; i < 7; i++ {
		outputOffset *= 10
		outputOffset += int(explodedDigits[i])
	}
	return Day16Part1(explodedDigits, iterations, outputOffset)
}

func main() {
	digits := readDigits(os.Stdin)
	digitsPart2 := make([]byte,len(digits))
	copy(digitsPart2,digits)
	log.Printf("part1 = %s", Day16Part1(digits, 100, 0))
	log.Printf("part2 = %s", Day16Part2(digitsPart2, 100))
}

func readDigits(in io.Reader) []byte {
	scanner := bufio.NewScanner(in)
	var digits []byte
	for scanner.Scan() {
		// only read the first line
		if digits != nil {
			break
		}
		line := scanner.Text()
		digits = make([]byte, len(line))
		for i, c := range line {
			digits[i] = byte(c - '0')
		}
	}
	return digits
}
