package main

import (
	"aoc/2020/day24/hexgrid"
	"bufio"
	"io"
	"log"
	"os"
)

func ReadInput(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0, 1000)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Part1(instructions []string) {
	grid := hexgrid.NewHexGrid()
	for _, instruction := range instructions {

		node.Flip()
	}
	log.Printf("black tiles = %d", grid.BlackNodeCount)
}

func main() {
	instructions := ReadInput(os.Stdin)
	Part1(instructions)
}
