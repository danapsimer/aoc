package main

import (
	"flag"
	"fmt"
	"github.com/danapsimer/aoc/2021/day04/bingo"
	"github.com/danapsimer/aoc/2021/utils"
)

var (
	callNumbersFileName = flag.String("calls", "calls.txt", "file with the called numbers for the bingo game.")
	boardsFileName      = flag.String("boards", "input.txt", "file with the boards")
)

func main() {
	boardsLines, err := utils.ReadLinesFromFile(*boardsFileName)
	if err != nil {
		panic(err)
	}
	boards, err := bingo.ReadBoards(boardsLines)
	if err != nil {
		panic(err)
	}

	calls, err := utils.ReadIntegersFromFile(*callNumbersFileName)
	if err != nil {
		panic(err)
	}

	board, score := bingo.FirstWinner(boards, calls)
	fmt.Printf("Part 1: winning score = %d, winning board:\n%v\n", score, board)

	calls, err = utils.ReadIntegersFromFile(*callNumbersFileName)
	if err != nil {
		panic(err)
	}

	board, score = bingo.LastWinner(boards, calls)
	fmt.Printf("Part 2: winning score = %d, winning board:\n%v\n", score, board)

}
