package main

import (
	"aoc/2019/intCode"
	"bytes"
	"container/list"
	"fmt"
	"log"
	"os"
)

type Cell int
type Status int
type Move int

const (
	North Move = 1
	South Move = 2
	West  Move = 3
	East  Move = 4
)
const (
	Unknown Cell = iota
	Origin
	Empty
	Wall
	Goal
)
const (
	HitWall Status = iota
	Moved
	MovedAndFoundGoal
)

func move(posX, posY int, dir Move) (int, int) {
	switch dir {
	case North:
		posY -= 1
	case South:
		posY += 1
	case West:
		posX -= 1
	case East:
		posX += 1
	}
	return posX, posY
}

type cell struct {
	x, y int
}

func adjacent(grid [][]Cell, p cell) []cell {
	ret := make([]cell, 0, 4)
	for dir := North; dir <= East; dir += 1 {
		x, y := move(p.x, p.y, dir)
		if grid[y][x] == Empty || grid[y][x] == Goal {
			ret = append(ret, cell{x, y})
		}
	}
	return ret
}

func printGrid(grid [][]Cell) string {
	buf := new(bytes.Buffer)
	for y := 0; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			switch grid[y][x] {
			case Unknown:
				_, _ = fmt.Fprint(buf, "_")
			case Wall:
				_, _ = fmt.Fprint(buf, "#")
			case Empty:
				_, _ = fmt.Fprint(buf, ".")
			case Goal:
				_, _ = fmt.Fprint(buf, "O")
			case Origin:
				_, _ = fmt.Fprint(buf, "X")
			}
		}
		fmt.Fprintln(buf)
	}
	return buf.String()
}

func ShortestPathLength(grid [][]Cell, c1, c2 cell) int {
	visited := make(map[cell]bool)
	q := list.New()
	q.PushBack(c1)
	nodesLeftInLayer := 1
	nodesInNextLayer := 0
	moves := 0
	meta := make(map[cell]cell)
	depth := make(map[cell]int)
	depth[c1] = 0
	for q.Len() > 0 {
		p := q.Remove(q.Front()).(cell)
		if p == c2 {
			return moves
		} else {
			for _, cell := range adjacent(grid, p) {
				if !visited[cell] {
					depth[cell] = moves + 1
					meta[cell] = p
					q.PushBack(cell)
					visited[cell] = true
					nodesInNextLayer += 1
				}
			}
		}
		nodesLeftInLayer -= 1
		if nodesLeftInLayer == 0 {
			nodesLeftInLayer = nodesInNextLayer
			nodesInNextLayer = 0
			moves += 1
		}
	}
	return -1
}

func Day15Part1(prg *intCode.IntCodeProgram) int {
	go prg.RunProgram()
	grid, originX, originY, goalX, goalY := BuildMap(prg)
	log.Printf("grid =\n%s", printGrid(grid))
	return ShortestPathLength(grid, cell{originX, originY}, cell{goalX, goalY})
}

func doMove(prg *intCode.IntCodeProgram, grid [][]Cell, x, y int, dir Move) (posX int, posY int) {
	newX, newY := move(x, y, dir)
	if newX < 0 || SIZE < newX || newY < 0 || SIZE < newY {
		log.Printf("newX = %d, newY = %d, grid =\n%s", newX, newY, printGrid(grid))
		panic("reached edge of grid")
	}
	if grid[newX][newY] == Unknown {
		prg.GetInput() <- int(dir)
		status := Status(<-prg.GetOutput())
		switch status {
		case HitWall:
			grid[newY][newX] = Wall
		case Moved:
			posX, posY = newX, newY
			if grid[posY][posX] == Unknown {
				grid[posY][posX] = Empty
			}
		case MovedAndFoundGoal:
			posX, posY = newX, newY
			grid[posY][posX] = Goal
		default:
			panic(fmt.Errorf("unknown response: %d", status))
		}
	} else {
		posX, posY = x, y
	}
	return
}

func BuildMapFF(prg *intCode.IntCodeProgram, grid [][]Cell, x, y int, dir Move) {
	newX, newY := doMove(prg, grid, x, y, dir)
	if newX != x || newY != y {
		BuildMapFF(prg, grid, newX, newY, North)
		BuildMapFF(prg, grid, newX, newY, East)
		BuildMapFF(prg, grid, newX, newY, South)
		BuildMapFF(prg, grid, newX, newY, West)
		switch dir {
		case North:
			doMove(prg, grid, newX, newY, South)
		case South:
			doMove(prg, grid, newX, newY, North)
		case East:
			doMove(prg, grid, newX, newY, West)
		case West:
			doMove(prg, grid, newX, newY, East)
		}
	}
}

const SIZE = 50

func BuildMap(prg *intCode.IntCodeProgram) (grid [][]Cell, originX, originY, goalX, goalY int) {
	grid = make([][]Cell, SIZE)
	for y := 0; y < SIZE; y += 1 {
		grid[y] = make([]Cell, SIZE)
	}
	originX, originY = SIZE/2, SIZE/2
	grid[originY][originY] = Origin
	BuildMapFF(prg, grid, originX, originY, North)
	BuildMapFF(prg, grid, originX, originY, South)
	BuildMapFF(prg, grid, originX, originY, East)
	BuildMapFF(prg, grid, originX, originY, West)
outer:
	for y, row := range grid {
		for x, c := range row {
			if c == Goal {
				goalX, goalY = x, y
				break outer
			}
		}
	}
	return grid, originX, originY, goalX, goalY
}

func Day15Part2(prg *intCode.IntCodeProgram) int {
	return 0
}

func main() {
	prg := intCode.ReadIntCodeProgram(os.Stdin)
	log.Printf("part1 = %d", Day15Part1(prg.Copy()))
}
