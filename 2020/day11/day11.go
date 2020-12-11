package main

import (
	"bufio"
	"github.com/gdamore/tcell/v2"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

var defStyle tcell.Style
var screen tcell.Screen

type grid []string

func readGrid(reader io.Reader) grid {
	scanner := bufio.NewScanner(reader)
	g := make(grid, 0, 1000)
	for scanner.Scan() {
		g = append(g, scanner.Text())
	}
	return g
}

func (g grid) get(x, y int) string {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return "."
	}
	return g[y][x : x+1]
}

func (g grid) height() int {
	return len(g)
}

func (g grid) width() int {
	return len(g[0])
}

func (g grid) equals(o grid) bool {
	if g.height() == o.height() && g.width() == o.width() {
		for y := range g {
			if g[y] != o[y] {
				return false
			}
		}
		return true
	}
	return false
}

func (g grid) traverseSeats(traversFn func(dest, src grid, x, y int)) grid {
	cp := make(grid, g.height())
	copy(cp, g)
	for y := 0; y < g.height(); y++ {
		for x := 0; x < g.width(); x++ {
			c := g.get(x, y)
			if c == "." {
				continue
			}
			traversFn(cp, g, x, y)
		}
	}
	return cp
}

func (g grid) applyRules() grid {
	return g.traverseSeats(func(cp, g grid, x, y int) {
		c := g.get(x, y)
		adjOccupiedCount := 0
		for yd := -1; yd <= 1; yd++ {
			for xd := -1; xd <= 1; xd++ {
				if (xd != 0 || yd != 0) && g.get(x+xd, y+yd) == "#" {
					adjOccupiedCount += 1
				}
			}
		}
		if c == "L" && adjOccupiedCount == 0 {
			cp[y] = cp[y][:x] + "#" + cp[y][x+1:]
		} else if c == "#" && adjOccupiedCount >= 4 {
			cp[y] = cp[y][:x] + "L" + cp[y][x+1:]
		}
	})
}

func (g grid) applyRulesPart2() grid {
	return g.traverseSeats(func(cp, g grid, x, y int) {
		c := g.get(x, y)
		visibleOccupiedCount := 0
		for yd := -1; yd <= 1; yd++ {
			for xd := -1; xd <= 1; xd++ {
				if xd == 0 && yd == 0 {
					continue
				}
				for nx, ny := x+xd, y+yd; 0 <= nx && nx < g.width() && 0 <= ny && ny < g.height(); nx, ny = nx+xd, ny+yd {
					nc := g.get(nx, ny)
					if nc == "#" {
						visibleOccupiedCount += 1
						break
					} else if nc == "L" {
						break
					}
				}
			}
		}
		if c == "L" && visibleOccupiedCount == 0 {
			cp[y] = cp[y][:x] + "#" + cp[y][x+1:]
		} else if c == "#" && visibleOccupiedCount >= 5 {
			cp[y] = cp[y][:x] + "L" + cp[y][x+1:]
		}
	})
}

func (g grid) String() string {
	return strings.Join(g, "\n")
}

var occupiedEmojis = []string{
	"\U0001F64D", "\U0001F64E", "\U0001F648", "\U0001F469", "\U0001F468",
}
var chair = "\u2441"


func (g grid) Draw() {
	sw, sh := screen.Size()
	if sw > g.width()+2 {
		sw = g.width() + 2
	}
	if sh > g.height()+2 {
		sh = g.height() + 2
	}

	for col := 1; col < sw; col++ {
		screen.SetContent(col, 0, tcell.RuneHLine, nil, defStyle)
		screen.SetContent(col, sh-1, tcell.RuneHLine, nil, defStyle)
	}
	for row := 1; row < sh; row++ {
		screen.SetContent(0, row, tcell.RuneVLine, nil, defStyle)
		screen.SetContent(sw-1, row, tcell.RuneVLine, nil, defStyle)
	}
	screen.SetContent(0, 0, tcell.RuneULCorner, nil, defStyle)
	screen.SetContent(sw-1, 0, tcell.RuneURCorner, nil, defStyle)
	screen.SetContent(0, sh-1, tcell.RuneLLCorner, nil, defStyle)
	screen.SetContent(sw-1, sh-1, tcell.RuneLRCorner, nil, defStyle)
	for y := range g {
		if y < sh-2 {
			for x := range g[y] {
				if x < sw-2 {
					r := rune(g[y][x])
					switch r {
					case '#':
						r, _ = utf8.DecodeRuneInString(occupiedEmojis[rand.Intn(len(occupiedEmojis))])
					case 'L':
						r, _ = utf8.DecodeRuneInString(chair)
					case '.':
						r = ' '
					}
					screen.SetContent(x+1, y+1, r, nil, defStyle)
				}
			}
		}
	}
	screen.Show()
}

func (original grid) applyRulesUntilStable(rulesFn func(g grid) grid) <-chan grid {
	gchan := make(chan grid)
	go func() {
		g := original
		cp := g.applyRules()
		for !cp.equals(g) {
			gchan <- g
			g = cp
			cp = rulesFn(g)
		}
		close(gchan)
	}()
	return gchan
}

func (g grid) countOccupiedSeats() int {
	occupied := 0
	for y := range g {
		for x := 0; x < g.width(); x++ {
			if g.get(x, y) == "#" {
				occupied += 1
			}
		}
	}
	return occupied
}

func main() {

	s, e := tcell.NewScreen()
	if e != nil {
		log.Fatalf("%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		log.Fatalf("%v\n", e)
		os.Exit(1)
	}
	defStyle = tcell.StyleDefault.
		Background(tcell.ColorReset).
		Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	s.Clear()
	screen = s

	original := readGrid(os.Stdin)
	go func() {
		gchan := original.applyRulesUntilStable(func(g grid) grid {
			return g.applyRules()
		})
		for g := range gchan {
			g.Draw()
			time.Sleep(time.Millisecond*500)
		}

		gchan = original.applyRulesUntilStable(func(g grid) grid {
			return g.applyRulesPart2()
		})

		for g := range gchan {
			g.Draw()
			time.Sleep(time.Millisecond*500)
		}
	}()

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				os.Exit(0)
			}
		}
	}
}
