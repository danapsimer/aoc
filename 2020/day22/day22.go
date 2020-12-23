package main

import (
	"aoc/2020/day22/deck"
	"log"
	"os"
)

func main() {

	game := deck.ReadGame(os.Stdin)
	winner := game.Play()

	log.Printf("winners score = %d", sum)
}
