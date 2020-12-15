package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadSeedNumbers(reader io.Reader) []int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	})
	seedNumbers := make([]int, 0, 50)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		seedNumbers = append(seedNumbers, i)
	}
	return seedNumbers
}

func main() {
	seedNumbers := ReadSeedNumbers(os.Stdin)
	lastSpoken2020, lastSpoken30000000 := SpeakNumbers(seedNumbers)
	log.Printf("The 2020th spoken number is %d", lastSpoken2020)
	log.Printf("The 30000000th spoken number is %d", lastSpoken30000000)
}

func SpeakNumbers(seedNumbers []int) (int, int) {
	lastSeen := make(map[int]int)
	lastSpoken := -1
	turn := 1
	for _, seed := range seedNumbers {
		lastSeen[seed] = turn
		lastSpoken = seed
		turn += 1
	}
	var lastTurnSeen int
	var beenSeen bool
	var spoken2020 int
	for turn <= 30000000 {
		var speak int
		if beenSeen {
			speak = turn - lastTurnSeen - 1
		} else {
			speak = 0
		}
		lastTurnSeen, beenSeen = lastSeen[speak]
		lastSeen[speak] = turn
		lastSpoken = speak
		if turn == 2020 {
			spoken2020 = speak
		}
		turn += 1
	}
	return spoken2020, lastSpoken
}
