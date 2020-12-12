package main

import (
	"aoc/2019/utils"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type op int

const (
	MoveForward op = iota
	TurnLeft
	TurnRight
	MoveNorth
	MoveEast
	MoveWest
	MoveSouth
)

func ParseOp(s string) op {
	switch s {
	case "F":
		return MoveForward
	case "L":
		return TurnLeft
	case "R":
		return TurnRight
	case "N":
		return MoveNorth
	case "E":
		return MoveEast
	case "W":
		return MoveWest
	case "S":
		return MoveSouth
	default:
		panic(fmt.Errorf("unknown op: %s", s))
	}
}

type Instruction struct {
	Op        op
	Magnitude int
}

func ParseInstruction(s string) Instruction {
	magnitude, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(fmt.Errorf("error parsing magnitude: %s - %s", s, err.Error()))
	}
	return Instruction{ParseOp(s[:1]), magnitude}
}

func ReadInstructions(reader io.Reader) <-chan Instruction {
	instructionChan := make(chan Instruction)
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			instructionChan <- ParseInstruction(scanner.Text())
		}
		close(instructionChan)
	}()
	return instructionChan
}

func main() {
	instructions := ReadInstructions(os.Stdin)
	part1Instructions := make(chan Instruction)
	part2Instructions := make(chan Instruction)
	go func() {
		for i := range instructions {
			part1Instructions <- i
			part2Instructions <- i
		}
		close(part1Instructions)
		close(part2Instructions)
	}()
	part1Result := make(chan int)
	part2Result := make(chan int)
	go func() {
		part1Result <- Part1(part1Instructions)
	}()
	go func() {
		part2Result <- Part2(part2Instructions)
	}()
	log.Printf("part 1 manhattan distance = %d", <-part1Result)
	log.Printf("part 2 manhattan distance = %d", <-part2Result)
}

func iCos(theta int) int {
	switch utils.IAbs(theta) {
	case 90:
		return 0
	case 180:
		return -1
	case 270:
		return 0
	default:
		panic(fmt.Errorf("invalid angle %d", theta))
	}
}

func iSin(theta int) int {
	switch utils.IAbs(theta) {
	case 90:
		return 1
	case 180:
		return 0
	case 270:
		return -1
	default:
		panic(fmt.Errorf("invalid angle %d", theta))
	}
}

func Rotate(x, y, theta int) (int, int) {
	return iCos(theta)*x - iSin(theta)*y, iSin(theta)*x + iCos(theta)*y
}

func Part1(instructions <-chan Instruction) int {
	x, y, dx, dy := 0, 0, 1, 0
	for i := range instructions {
		switch i.Op {
		case MoveForward:
			x += dx * i.Magnitude
			y += dy * i.Magnitude
		case TurnLeft:
			dx, dy = Rotate(dx,dy,i.Magnitude)
		case TurnRight:
			dx, dy = Rotate(dx,dy,360-i.Magnitude)
		case MoveNorth:
			y += i.Magnitude
		case MoveEast:
			x += i.Magnitude
		case MoveSouth:
			y -= i.Magnitude
		case MoveWest:
			x -= i.Magnitude
		}
	}
	return utils.IAbs(x) + utils.IAbs(y)
}

func Part2(instructions <-chan Instruction) int {
	x, y, dx, dy := 0, 0, 10, 1
	for i := range instructions {
		switch i.Op {
		case MoveForward:
			x += dx * i.Magnitude
			y += dy * i.Magnitude
		case TurnLeft:
			dx, dy = Rotate(dx, dy, i.Magnitude)
		case TurnRight:
			dx, dy = Rotate(dx, dy, 360-i.Magnitude)
		case MoveNorth:
			dy += i.Magnitude
		case MoveEast:
			dx += i.Magnitude
		case MoveSouth:
			dy -= i.Magnitude
		case MoveWest:
			dx -= i.Magnitude
		}
	}
	return utils.IAbs(x) + utils.IAbs(y)
}
