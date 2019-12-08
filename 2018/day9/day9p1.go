package main

import (
	"container/ring"
	"fmt"
	"os"
)

func RingOfOne(value int) *ring.Ring {
	r := ring.New(1)
	r.Value = value
	return r
}

func PrintRing(ball int, base, current *ring.Ring) {
	fmt.Printf("[%5d] ", ball)
	n := base
	for {
		delim := " "
		if n.Value == current.Value {
			delim = "*"
		}
		fmt.Printf("%5d%s", n.Value, delim)
		n = n.Next()
		if n.Value == base.Value {
			break
		}
	}
	fmt.Print("\n")
}

func main() {
	var players, largestBall int
	_, err := fmt.Scanf("%d players; last marble is worth %d points", &players, &largestBall)
	if err != nil {
		fmt.Printf("ERROR: error reading input: %s\n", err.Error())
		os.Exit(-1)
	}
	largestBall *= 100
	scores := make([]int, players)
	circle := RingOfOne(0)
	base := circle
	//PrintRing(0,base,circle)
	for ball, player := 1, 1; ball <= largestBall; ball, player = ball+1, player+1 {
		if player > players {
			player = 1
		}
		if ball%23 == 0 {
			scores[player-1] += ball
			circle = circle.Move(-8)
			// adjust base if base is about to be removed
			if circle.Value == base.Value {
				base = circle.Next()
			}
			removed := circle.Unlink(1)
			circle = circle.Next()
			scores[player-1] += removed.Value.(int)
		} else {
			circle = circle.Move(1).Link(RingOfOne(ball)).Move(-1)
		}
		//PrintRing(ball,base,circle)
	}
	highScore := 0
	for _, s := range scores {
		if highScore < s {
			highScore = s
		}
	}
	fmt.Printf("high score is %d\n", highScore)
}
