package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type grid [128][8]bool

func main() {
	seating := readSeating(os.Stdin)
	for _, row := range seating {
		for _, col := range row {
			if col {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	foundLargest := false
	for r := 127; r >= 0; r-- {
		for c := 7; c >= 0; c-- {
			if seating[r][c] {
				if !foundLargest {
					log.Printf("largest seat id is: %d", r*8+c)
					foundLargest = true
				}
			} else {
				nr := r
				nc := c + 1
				if nc == 8 {
					nr += 1
					nc = 0
				}
				pr := r
				pc := c - 1
				if c < 0 {
					c = 7
					pr -= 1
				}
				if nr < 128 && seating[nr][nc] && seating[pr][pc] {
					log.Printf("Your seat is %d, %d, id = %d", r, c, r*8+c)
				}
			}
		}
	}
}

func binaryPosition(coord string, start, end int) int {
	size := end - start
	if size == 1 {
		return start
	}
	c := coord[0]
	switch {
	case c == 'F' || c == 'L':
		return binaryPosition(coord[1:], start, start+size/2)
	case c == 'B' || c == 'R':
		return binaryPosition(coord[1:], start+size/2, end)
	default:
		panic(fmt.Errorf("invalid seat coordinate character encountered: %v", c))
	}
}

func readSeating(r io.Reader) *grid {
	scanner := bufio.NewScanner(r)
	var g grid
	for scanner.Scan() {
		line := scanner.Text()
		row := binaryPosition(line[0:7], 0, 128)
		col := binaryPosition(line[7:10], 0, 8)
		g[row][col] = true
	}
	return &g
}
