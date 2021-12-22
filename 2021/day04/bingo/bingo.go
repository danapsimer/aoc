package bingo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	DoneError = errors.New("Done")
)

type cell struct {
	value  int
	marked bool
}

type board [5][5]cell

func ReadBoard(in <-chan string) (*board, error) {
	b := new(board)
	for r := 0; r < 5; r++ {
		rowStr := <-in
		rowStr = strings.TrimSpace(rowStr)
		if rowStr == "" {
			return nil, fmt.Errorf("Incomplete board!")
		}
		row := strings.Split(rowStr, " ")
		c := 0
		for _, v := range row {
			if v == "" {
				continue
			}
			var err error
			b[r][c].value, err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			c += 1
		}
		if c != 5 {
			return nil, fmt.Errorf("malformed row, not 5 elements: %s", rowStr)
		}
	}
	_, more := <-in
	if !more {
		return b, DoneError
	}
	return b, nil
}

func (b *board) Mark(value int) bool {
outer:
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if b[r][c].value == value {
				b[r][c].marked = true
				break outer
			}
		}
	}
	return b.Check()
}

func (b *board) Check() bool {
	for r := 0; r < 5; r++ {
		allMarked := true
		for c := 0; c < 5 && allMarked; c++ {
			if !b[r][c].marked {
				allMarked = false
			}
		}
		if allMarked {
			return true
		}
	}
	for c := 0; c < 5; c++ {
		allMarked := true
		for r := 0; r < 5 && allMarked; r++ {
			if !b[r][c].marked {
				allMarked = false
			}
		}
		if allMarked {
			return true
		}
	}
	return false
}

func (b *board) Score(lastCall int) int {
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !b[r][c].marked {
				sum += b[r][c].value
			}
		}
	}
	return sum * lastCall
}

func ReadBoards(in <-chan string) ([]*board, error) {
	boards := make([]*board, 0, 100)
	for {
		b, err := ReadBoard(in)
		if err != nil && err != DoneError {
			return nil, err
		}
		boards = append(boards, b)
		if err == DoneError {
			break
		}
	}
	return boards, nil
}

func FirstWinner(boards []*board, calls <-chan int) (*board, int) {
	for call := range calls {
		for _, b := range boards {
			if b.Mark(call) {
				return b, b.Score(call)
			}
		}
	}
	return nil, 0
}

func LastWinner(boards []*board, calls <-chan int) (*board, int) {
	won := make([]bool, len(boards))
	var lastWinner struct {
		b     *board
		score int
	}
	for call := range calls {
		for bidx, b := range boards {
			if !won[bidx] && b.Mark(call) {
				won[bidx] = true
				lastWinner.b = b
				lastWinner.score = b.Score(call)
			}
		}
	}
	return lastWinner.b, lastWinner.score
}
