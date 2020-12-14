package main

import (
	"aoc/2019/utils"
	"bufio"
	"fmt"
	"github.com/StephaneBunel/bresenham"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
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
	part1Instructions := make(chan Instruction, 100)
	part2Instructions := make(chan Instruction, 100)
	part1Visualization := make(chan Position, 100)
	part2Visualization := make(chan Position, 100)
	go func() {
		for i := range instructions {
			part1Instructions <- i
			part2Instructions <- i
		}
		close(part1Instructions)
		close(part2Instructions)
	}()

	part1Result := processInstructions(part1Instructions, part1Visualization, Part1)
	part2Result := processInstructions(part2Instructions, part2Visualization, Part2)

	computeVisualization("part1.gif", part1Visualization)
	computeVisualization("part2.gif", part2Visualization)

	log.Printf("part 1 manhattan distance = %d", <-part1Result)
	log.Printf("part 2 manhattan distance = %d", <-part2Result)
}

func processInstructions(instructions <-chan Instruction, visualization chan<- Position, partFn func(<-chan Instruction) <-chan Position) <-chan int {
	result := make(chan int)
	go func() {
		part1Positions := partFn(instructions)
		defer close(visualization)
		lastPosition := Position{0, 0}
		for position := range part1Positions {
			visualization <- position
			lastPosition = position
		}
		result <- utils.IAbs(lastPosition.X) + utils.IAbs(lastPosition.Y)
	}()
	return result
}

func computeVisualization(outputFileName string, positions <-chan Position) {
	palette := make(color.Palette, 2)
	palette[0] = color.White
	palette[1] = color.Black
	go func() {
		g := &gif.GIF{
			make([]*image.Paletted, 0, 1000),
			make([]int, 0, 1000),
			0,
			nil,
			image.Config{
				ColorModel: nil,
				Width:      0,
				Height:     0,
			},
			0,
		}
		window := image.Rect(-10, -10, 10, 10)
		var previousImage *image.Paletted
		p := image.Pt(0, 0)
		for position := range positions {
			np := image.Pt(position.X, position.Y)

			newWindow := window
			if np.X < newWindow.Min.X {
				newWindow.Min.X = np.X - 10
			} else if np.X >= newWindow.Max.X {
				newWindow.Max.X = np.X + 10
			}
			if np.Y < newWindow.Min.Y {
				newWindow.Min.Y = np.Y - 10
			} else if np.Y >= newWindow.Max.Y {
				newWindow.Max.Y = np.Y + 10
			}

			log.Printf("creating frame# %d: drawing line from %v, to %v", len(g.Image), p, np)
			img := image.NewPaletted(newWindow, palette)
			if previousImage != nil {
				draw.Src.Draw(img, window, previousImage, window.Min)
			}
			bresenham.Bresenham(img, p.X, p.Y, np.X, np.Y, palette[1])

			g.Image = append(g.Image, img)
			g.Delay = append(g.Delay, 100)
			g.Config.Height = newWindow.Dy()
			g.Config.Width = newWindow.Dy()

			window = newWindow
			previousImage = img
			p = np
		}
		f, err := os.OpenFile(outputFileName, os.O_RDWR | os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Printf("Error writing GIF : %s", err.Error())
		}
		log.Printf("Writing GIF animation to %s", outputFileName)
		defer func() {
			f.Close()
		}()
		gif.EncodeAll(f, g)
	}()
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

type Position struct {
	X, Y int
}

func Part1(instructions <-chan Instruction) <-chan Position {
	pChan := make(chan Position)
	go func() {
		defer close(pChan)
		x, y, dx, dy := 0, 0, 1, 0
		for i := range instructions {
			switch i.Op {
			case MoveForward:
				x += dx * i.Magnitude
				y += dy * i.Magnitude
				pChan <- Position{x, y}
			case TurnLeft:
				dx, dy = Rotate(dx, dy, i.Magnitude)
			case TurnRight:
				dx, dy = Rotate(dx, dy, 360-i.Magnitude)
			case MoveNorth:
				y += i.Magnitude
				pChan <- Position{x, y}
			case MoveEast:
				x += i.Magnitude
				pChan <- Position{x, y}
			case MoveSouth:
				y -= i.Magnitude
				pChan <- Position{x, y}
			case MoveWest:
				x -= i.Magnitude
				pChan <- Position{x, y}
			}
		}
	}()
	return pChan
}

func Part2(instructions <-chan Instruction) <-chan Position {
	pChan := make(chan Position)
	go func() {
		defer close(pChan)
		x, y, dx, dy := 0, 0, 10, 1
		for i := range instructions {
			switch i.Op {
			case MoveForward:
				x += dx * i.Magnitude
				y += dy * i.Magnitude
				pChan <- Position{x, y}
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
	}()
	return pChan
}
