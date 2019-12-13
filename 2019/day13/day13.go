package main

import (
	"aoc/2019/intCode"
	"bytes"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"os"
	"time"
)

type Cell int

const (
	Empty Cell = iota
	Wall
	Block
	Paddle
	Ball
)

const (
	Height = 21
	Width  = 44
)

var displayChars = " +#_O"

func Respond(grid [][]Cell, out <-chan int) <-chan int {
	scoreChan := make(chan int)
	go func() {
		maxX, maxY := 0, 0
		defer func() {
			log.Printf("maxX = %d, maxY = %d", maxX, maxY)
			close(scoreChan)
		}()
		state := 1
		x, y := 0, 0
		for g := range out {
			switch state {
			case 1:
				x = g
				state = 2
			case 2:
				y = g
				state = 3
			case 3:
				if x == -1 && y == 0 {
					scoreChan <- g
				} else {
					if Cell(g) < Empty || Ball < Cell(g) {
						panic(fmt.Errorf("unexpected cell type %d", g))
					}
					if y > maxY {
						maxY = y
					}
					if x > maxX {
						maxX = x
					}
					grid[y][x] = Cell(g)
				}
				state = 1
			default:
				panic(fmt.Errorf("NOT POSSIBLE: unexpected state: %d", state))
			}
		}
	}()
	return scoreChan
}

func NewGrid() [][]Cell {
	grid := make([][]Cell, Height)
	for i := range grid {
		grid[i] = make([]Cell, Width)
	}
	return grid
}

func FindBallAndPaddleX(grid [][]Cell) (ball, paddle int) {
	ball = -1
	paddle = -1
	for _, row := range grid {
		for x, c := range row {
			if c == Ball {
				ball = x
			} else if c == Paddle {
				paddle = x
			}
			if ball >= 0 && paddle >= 0 {
				return
			}
		}
	}
	return
}

func DrawGrid(grid [][]Cell) string {
	buf := new(bytes.Buffer)
	for _, row := range grid {
		for _, c := range row {
			if _, err := fmt.Fprint(buf, displayChars[c:c+1]); err != nil {
				panic(fmt.Errorf("error writing to buffer! %s", displayChars[c:c+1]))
			}
		}
		fmt.Fprintln(buf)
	}
	return buf.String()
}

func Day13Part2(prg *intCode.IntCodeProgram, grid [][]Cell) int {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	canvas := widgets.NewParagraph()
	canvas.SetRect(0, 0, Width+2, Height+2)
	ui.Render(canvas)

	scoreWidget := widgets.NewParagraph()
	scoreWidget.SetRect(0, Height+2, 20, Height+5)

	prg.GetProgram()[0] = 2
	go prg.RunProgram()
	scoreChan := Respond(grid, prg.GetOutput())
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok || err.Error() != "send on closed channel" {
					panic(r)
				}
			}
		}()
		for {
			time.Sleep(time.Millisecond * 10)
			canvas.Text = DrawGrid(grid)
			ui.Render(canvas)
			ball, paddle := FindBallAndPaddleX(grid)
			if ball < 0 || paddle < 0 {
				panic(fmt.Errorf("could not find ball and paddle: %d, %d", ball, paddle))
			}
			dir := 0
			if paddle < ball {
				dir = 1
			} else if paddle > ball {
				dir = -1
			}
			prg.GetInput() <- dir
		}
	}()
	var score int
	for score = range scoreChan {
		scoreWidget.Text = fmt.Sprintf("Score: %10d", score)
		ui.Render(scoreWidget)
	}
	return score
}

func Day13Part1(prg *intCode.IntCodeProgram, grid [][]Cell) int {
	go prg.RunProgram()
	doneChan := Respond(grid, prg.GetOutput())
	for score := range doneChan {
		log.Printf("score = %d", score)
	}
	blocks := 0
	for _, row := range grid {
		for _, c := range row {
			if c == Block {
				blocks += 1
			}
		}
	}
	return blocks
}

func main() {
	prg := intCode.ReadIntCodeProgram(os.Stdin)
	log.Printf("part1 = %d", Day13Part1(prg.Copy(), NewGrid()))
	log.Printf("part2 = %d", Day13Part2(prg.Copy(), NewGrid()))
}
