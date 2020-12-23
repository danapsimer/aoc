package main

import (
	"container/ring"
	"log"
	"strconv"
	"strings"
)

type Game struct {
	size  int
	r     *ring.Ring
	index []*ring.Ring
}

func (g *Game) decLabel(i int) int {
	l := i - 1
	if l < 1 {
		l = g.size
	}
	return l
}

func NewRing(size int) *Game {
	r := ring.New(size)
	index := make([]*ring.Ring, size)
	input := []int{1, 5, 6, 7, 9, 4, 8, 2, 3}
	for i := 0; i < len(input); i++ {
		r.Value = input[i]
		index[input[i]-1] = r
		r = r.Next()
	}
	for i := 10; i <= size; i++ {
		r.Value = i
		index[i-1] = r
		r = r.Next()
	}
	return &Game{size, r, index}
}

func (g *Game) RunGame(turns int) {
	curr := g.r
	for i := 0; i < turns; i++ {
		moving := curr.Unlink(3)
		var insert *ring.Ring
		startLabel := g.decLabel(curr.Value.(int))
		endLabel := g.decLabel(g.decLabel(g.decLabel(g.decLabel(startLabel))))
		for insertLabel := startLabel; insertLabel != endLabel; insertLabel = g.decLabel(insertLabel) {
			insert = g.index[insertLabel-1]
			beingMoved := false
			moving.Do(func(value interface{}) {
				if value.(int) == insertLabel {
					beingMoved = true
				}
			})
			if !beingMoved {
				break
			}
		}
		insert.Link(moving)
		curr = curr.Next()
	}
}

func Part1() {
	game := NewRing(9)
	game.RunGame(100)
	cup1 := game.index[0]
	sb := strings.Builder{}
	cup := cup1.Next()
	for {
		sb.WriteString(strconv.Itoa(cup.Value.(int)))
		cup = cup.Next()
		if cup == cup1.Next() {
			break
		}
	}
	log.Printf("labels after cup 1: %s", sb.String())
}

func Part2() {
	game := NewRing(1000000)
	game.RunGame(10000000)
	cup1 := game.index[0]
	result := cup1.Next().Value.(int) * cup1.Next().Next().Value.(int)
	log.Printf("2 labels after cup 1 multiplied together: %d", result)
}

func main() {
	Part1()
	Part2()
}

