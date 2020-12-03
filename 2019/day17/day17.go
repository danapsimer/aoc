package main

import (
	"aoc/2019/intCode"
	"errors"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	prg := intCode.ReadIntCodeProgram(os.Stdin)
	grid := ReadGrid(prg.Copy())
	log.Printf("part1 = %d", Day17Par1(grid))
	log.Printf("part2 = %d", Day17Part2(prg.Copy()))
}

func ReadGrid(prg *intCode.IntCodeProgram) []string {
	go prg.RunProgram()
	grid := make([]string, 0, 1000)
	grid = append(grid, "")
	for c := range prg.GetOutput() {
		if c != 10 {
			grid[len(grid)-1] = grid[len(grid)-1] + string(c)
		} else {
			grid = append(grid, "")
		}
	}
	if len(grid[len(grid)-1]) == 0 {
		grid = grid[:len(grid)-1]
	}
	PrintGrid(grid)
	return grid
}

func PrintGrid(grid []string) {
	for y := range grid {
		log.Printf("%2d: |%s|\n", y, grid[y])
	}
}

func Day17Par1(grid []string) (sum int) {
	var x, y int
	defer func() {
		if r := recover(); r != nil {
			log.Printf("error thrown at (%d,%d): %+v", x, y, r)
		}
	}()
	for y = 0; y < len(grid); y++ {
		for x = 0; x < len(grid[y]); x++ {
			if grid[y][x] == '#' &&
				(y < 1 || grid[y-1][x] == '#') &&
				(y > len(grid)-2 || grid[y+1][x] == '#') &&
				(x < 1 || grid[y][x-1] == '#') &&
				(x > len(grid[y])-2 || grid[y][x+1] == '#') {
				sum += x * y
			}
		}
	}
	return sum
}

type Dir int
type Turn int

const (
	North Dir = iota
	East
	South
	West
)
const (
	Left Turn = iota
	Right
)
const showUI = false

func findRobot(grid []string) (int, int, Dir) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			switch grid[y][x] {
			case '^':
				return x, y, North
			case 'v':
				return x, y, South
			case '>':
				return x, y, East
			case '<':
				return x, y, West
			}
		}
	}
	panic(errors.New("robot not found"))
}

func move(grid []string, x, y int, dir Dir) (int, int, bool) {
	switch dir {
	case North:
		if y > 0 && grid[y-1][x] == '#' {
			return x, y - 1, true
		}
	case South:
		if y < len(grid)-1 && grid[y+1][x] == '#' {
			return x, y + 1, true
		}
	case West:
		if x > 0 && grid[y][x-1] == '#' {
			return x - 1, y, true
		}
	case East:
		if x < len(grid[y])-1 && grid[y][x+1] == '#' {
			return x + 1, y, true
		}
	}
	return x, y, false
}

func turn(t Turn, dir Dir) Dir {
	if t == Right {
		dir += 1
		if dir > West {
			dir = North
		}
	} else {
		dir -= 1
		if dir < North {
			dir = West
		}
	}
	return dir
}

func findPath(grid []string) []string {
	path := make([]string, 0, 1000)
	rx, ry, rdir := findRobot(grid)
	d := 0
	for {
		newRX, newRY, ok := move(grid, rx, ry, rdir)
		if !ok {
			if d > 0 {
				path = append(path, strconv.Itoa(d))
				d = 0
			}
			newDir := turn(Right, rdir)
			newRX, newRY, ok = move(grid, rx, ry, newDir)
			if ok {
				rdir = newDir
				path = append(path, "R")
			} else {
				newDir = turn(Left, rdir)
				newRX, newRY, ok = move(grid, rx, ry, newDir)
				if ok {
					rdir = newDir
					path = append(path, "L")
				} else {
					// we have reached the end of the maze
					break
				}
			}
		} else {
			rx, ry = newRX, newRY
			d += 1
		}
	}
	return path
}

func Day17Part2(prg *intCode.IntCodeProgram) (sum int) {
	prg.GetProgram()[0] = 2
	go prg.RunProgram()
	grid := Day17UI(prg)
	path := findPath(grid)
	log.Printf("len(path) = %d, path = %v", len(path), path)
	pathStr := strings.Join(path,",")
	rexp := regexp.MustCompilePOSIX("^(.{1,20})\\1*(.{1,20})(?:\\1|\\2)*(.{1,20})(?:\\1|\\2|\\3)*$")
	log.Printf("pathStr = %v", rexp.FindAllStringSubmatch(pathStr, 0))

	return
}

var Height, Width = 32, 32

func ReadGridPart2(prg *intCode.IntCodeProgram) (string, bool) {
	grid := ""
	line := 0
	for c := range prg.GetOutput() {
		grid += string(c)
		if c == 10 {
			if line == 32 {
				return grid, false
			}
			line += 1
		}
	}
	return grid, true
}

func Day17UI(prg *intCode.IntCodeProgram) []string {
	var canvas *widgets.Paragraph
	if showUI {
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		canvas = widgets.NewParagraph()
		canvas.SetRect(0, 0, Height+2, Width+2)
		ui.Render(canvas)
	}
	grid, done := ReadGridPart2(prg)
	go func() {
		for !done {
			grid, done = ReadGridPart2(prg)
			if showUI {
				canvas.Text = grid
				ui.Render(canvas)
			}
		}
	}()
	gridarr := strings.Split(grid, "\n")
	if len(gridarr[len(gridarr)-1]) == 0 {
		gridarr = gridarr[:len(gridarr)-1]
	}
	return gridarr
}
