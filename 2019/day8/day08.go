package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	layerSize = 25 * 6
)

func main() {
	aggregateLayer := make([]byte, layerSize)
	for i := range aggregateLayer {
		aggregateLayer[i] = '2'
	}

	minZeroCount := math.MaxInt32
	var minZeroLayer []byte
	for {
		layer := make([]byte, layerSize)
		b, err := os.Stdin.Read(layer)
		if err == io.EOF {
			if b > 0 {
				panic(fmt.Sprintf("bytes are left at the end of the file: %v", layer))
			}
			break
		}
		if err != nil {
			panic(err)
		}
		if b != layerSize {
			panic(fmt.Sprintf("layersize is too small: %d", b))
		}
		zeroCount := 0
		for i, c := range layer {
			if c == '0' {
				zeroCount += 1
			}
			if aggregateLayer[i] == '2' {
				aggregateLayer[i] = layer[i]
			}
		}
		if zeroCount < minZeroCount {
			minZeroCount = zeroCount
			minZeroLayer = layer
		}
	}
	oneCount := 0
	twoCount := 0
	for _, c := range minZeroLayer {
		if c == '1' {
			oneCount += 1
		}
		if c == '2' {
			twoCount += 1
		}
	}
	fmt.Printf("part01 = %d\n", oneCount*twoCount)
	fmt.Printf("part02 = %s\n", string(aggregateLayer))
	for p := 0; p < len(aggregateLayer); p += 25 {
		line := aggregateLayer[p : p+25]
		for _, c := range line {
			if c == '1' {
				fmt.Print("X")
			} else if c == '0' {
				fmt.Print(" ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
