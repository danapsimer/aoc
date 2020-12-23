package deck

import (
	"bufio"
	"errors"
	"strconv"
)

type Round []int

type Deck struct {
	cards []int
}

var EndOfDeckError = errors.New("end of deck")

func ReadDeck(scanner *bufio.Scanner) *Deck {
	cards := make([]int, 0, 50)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		card, err := strconv.Atoi(line)
		if err != nil {
			panic("cannot parse line: " + line)
		}
		cards = append(cards, card)
	}
	return &Deck{cards}
}

func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}

func (d *Deck) Draw() (int, error) {
	if len(d.cards) == 0 {
		return 0, EndOfDeckError
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, nil
}

func (d *Deck) PushWin(cards Round) {
	d.cards = append(d.cards, cards...)
}

func (d *Deck) Score() int {
	sum := 0
	for idx, v := range d.cards {
		sum += (len(d.cards) - idx) * v
	}
	return sum
}
