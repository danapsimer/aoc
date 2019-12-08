package main

import "fmt"

type point struct {
	x, y int
}

func NewPoint(line string) (*point, error) {
	var p point
	_, err := fmt.Sscanf(line, "%d, %d", &p.x, &p.y)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(p1, p2 *point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}
