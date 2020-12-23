package deck

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

type Game struct {
	players []*Deck
	recursive bool
}

func ReadGame(reader io.Reader) *Game {
	scanner := bufio.NewScanner(reader)
	players := make([]*Deck, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "Player ") {
			players = append(players, ReadDeck(scanner))
		}
	}
	return &Game{players, false}
}

func (game *Game) MakeRecursive() {
	game.recursive = true
}

func (game *Game) PlayRound() (bool, int) {
	stillIn := make([]int, 0, 10)
	for pidx, player := range game.players {
		if !player.IsEmpty() {
			stillIn = append(stillIn, pidx)
		}
	}
	if len(stillIn) == 0 {
		panic("0 players with cards left!?!")
	}
	if len(stillIn) == 1 {
		return true, stillIn[0]
	}

	round := make(Round, 0, len(game.players))
	maxCard := -1
	winningPlayer := -1
	for _, pidx := range stillIn {
		card, _ := game.players[pidx].Draw()
		round = append(round, card)
		if card > maxCard {
			maxCard = card
			winningPlayer = pidx
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(round)))
	game.players[winningPlayer].PushWin(round)
	return false, winningPlayer
}

func (game *Game) Play() *Deck {
	for {
		won, winner := game.PlayRound()
		if won {
			return game.players[winner]
		}
	}
}
